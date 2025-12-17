package helpers

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cybroslabs/ouro-api-shared/gen/go/system"
)

const (
	_periodNoLicense  = 15 * time.Second
	_periodRevalidate = 30 * time.Minute
)

var (
	_semver    = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(?:-([\da-z\-]+(?:\.[\da-z\-]+)*))?(?:\+([\da-z\-]+(?:\.[\da-z\-]+)*))?$`)
	_semverOp  = regexp.MustCompile(`^(<=|<|=|>|>=|)v?(\d+)\.(\d+)\.(\d+)(?:-([\da-z\-]+(?:\.[\da-z\-]+)*))?(?:\+([\da-z\-]+(?:\.[\da-z\-]+)*))?$`)
	_lowercase = regexp.MustCompile(`^[a-z](?:[a-z-]*[a-z])?$`)
)

// LicenseManager is an interface that defines methods for license operations in the system.
type LicenseManager interface {
	// GetLicense retrieves the current license key. If no license is set, it returns nil.
	GetLicense() *system.License
	// HasLicense checks if a license is currently set.
	HasLicense() bool
	// WaitLicense blocks until a license is available or the context is done.
	WaitLicense(ctx context.Context) (*system.License, error)
	// Stop stops the license manager and releases any resources it holds.
	Stop()
}

// The LicenseManagerOpts struct contains options for creating a LicenseManager instance.
type LicenseManagerOpts struct {
	// This must be valid semver string, e.g. "1.0.0", "2.1.3-beta.1" with optional "v" prefix.
	AppVersion string
	// This must be valid application name.
	AppName string
	// The date time when the application was built, in RFC3339 format.
	AppBuiltDate string
	// Connectors provides access to external services required by the LicenseManager.
	Connectors Connectors
	// Cancel is an optional context cancel function that will be called when the LicenseManager is stopped.
	Cancel context.CancelFunc
	// OnLicenseGranted is an optional callback function that will be called when a license is successfully fetched.
	OnLicenseGranted func(license *system.License)
	// OnLicenseRevoked is an optional callback function that will be called when the license is revoked or becomes invalid.
	OnLicenseRevoked func()
}

type licenseManager struct {
	appVersion       string
	appName          string
	appBuiltDate     time.Time
	appSemVer        []int
	connectors       Connectors
	license          *system.License
	onLicenseGranted func(license *system.License)
	onLicenseRevoked func()

	runCancel    context.CancelFunc
	outterCancel context.CancelFunc

	event chan struct{}
}

// NewLicenseManager creates a new instance of LicenseManager with the provided options.
func NewLicenseManager(opts *LicenseManagerOpts) (LicenseManager, error) {
	if opts == nil {
		return nil, errors.New("options cannot be nil")
	}

	lm := &licenseManager{
		appVersion:       opts.AppVersion,
		appName:          opts.AppName,
		connectors:       opts.Connectors,
		outterCancel:     opts.Cancel,
		onLicenseGranted: opts.OnLicenseGranted,
		onLicenseRevoked: opts.OnLicenseRevoked,
		event:            make(chan struct{}, 1),
	}

	is_dev := strings.EqualFold(opts.AppVersion, "dev") || strings.EqualFold(opts.AppVersion, "development")
	if is_dev {
		lm.appSemVer = []int{999, 999, 999} // Far future version for 'dev' version
	} else if sv, err := getSemVerParts(opts.AppVersion); err != nil {
		return nil, err
	} else {
		lm.appSemVer = sv
	}
	if !_lowercase.MatchString(opts.AppName) {
		return nil, errors.New("AppName must be lowercase letters and hyphens only, starting and ending with a letter")
	}
	if is_dev {
		lm.appBuiltDate = time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC) // Far future date for 'dev' version
	} else if t, err := time.Parse(time.RFC3339, opts.AppBuiltDate); err != nil {
		return nil, errors.New("AppBuiltDate must be in RFC3339 format")
	} else {
		lm.appBuiltDate = t
	}

	ctx, ctx_cancel := context.WithCancel(context.Background())
	lm.runCancel = ctx_cancel

	go lm.run(ctx)

	return lm, nil
}

// Stop stops the license manager and releases any resources it holds.
// It cancels the running context and the outer cancel function if provided.
func (lm *licenseManager) Stop() {
	if lm == nil {
		return
	}
	if lm.runCancel != nil {
		lm.runCancel()
	}
	if lm.outterCancel != nil {
		lm.outterCancel()
	}
}

func getSemVerParts(version string) ([]int, error) {
	if sv := _semver.FindStringSubmatch(version); sv == nil {
		return nil, errors.New("version must be a valid semver string")
	} else {
		mustAtoi := func(s string) int {
			v, _ := strconv.Atoi(s)
			return v
		}
		return []int{
			mustAtoi(sv[1]),
			mustAtoi(sv[2]),
			mustAtoi(sv[3]),
		}, nil
	}
}

func getSemVerOpParts(version string) (string, []int, error) {
	if sv := _semverOp.FindStringSubmatch(version); sv == nil {
		return "", nil, errors.New("version must be a valid semver string with optional operator")
	} else {
		mustAtoi := func(s string) int {
			v, _ := strconv.Atoi(s)
			return v
		}
		op := sv[1]
		if op == "" {
			op = "="
		}
		return op, []int{
			mustAtoi(sv[2]),
			mustAtoi(sv[3]),
			mustAtoi(sv[4]),
		}, nil
	}
}

func drainChannel(ch chan struct{}) {
	// Paranoidly drain the channel
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}

func (lm *licenseManager) run(ctx context.Context) {
	check_period := _periodNoLicense
	check_failed_count := 0
	granted := false
	for {
		license, err := lm.fetchLicense(ctx)
		if err != nil {
			if granted {
				check_failed_count++
				if check_failed_count == 3 {
					// If we failed to get the license three times in a row, we will revoke current license.
					lm.license = nil
					drainChannel(lm.event)
					if f := lm.onLicenseRevoked; f != nil {
						f()
					}
				}
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(check_period):
				continue
			}
		} else {
			check_failed_count = 0
			granted = true
		}

		lm.license = license
		select {
		case lm.event <- struct{}{}:
		default:
		}
		if f := lm.onLicenseGranted; f != nil {
			f(license)
		}

		// We've got a license, let's revalidate it periodically.
		select {
		case <-ctx.Done():
			return
		case <-time.After(_periodRevalidate):
			continue
		}
	}
}

// Returns the current license.
func (lm *licenseManager) fetchLicense(ctx context.Context) (*system.License, error) {
	if lm == nil {
		return nil, errors.New("license manager is not initialized")
	}

	cli, closer, err := lm.connectors.OpenOuroOperatorServiceClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	resp, err := cli.GetLicense(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Check that the build date is not after the service expiry date.
	if service_expiration := resp.GetServiceExpiration(); service_expiration != nil {
		ts := service_expiration.AsTime()
		if !lm.appBuiltDate.IsZero() && lm.appBuiltDate.After(ts) {
			return nil, errors.New("the service has expired, unable to run version " + lm.appVersion)
		}
	}

	// Validate that the license is valid for this application version.
	// If no component versions are specified, the license is valid for all versions.
	component_versions := resp.GetLicensedStringArray(system.LicensedItemComponentVersions)
	component_prefix := lm.appName + ":"
	for _, v := range component_versions {
		ver_info, ok := strings.CutPrefix(v, component_prefix)
		if !ok {
			continue
		}

		parts := strings.Fields(ver_info)
		prev_sv := []int{0, 0, 0}
		version_granted := false
		min_ok := len(parts) > 0

	GRANTED:
		for _, p := range parts {
			var op string
			var sv []int
			if op, sv, err = getSemVerOpParts(p); err != nil {
				return nil, errors.New("invalid component version format in license")
			}
			if sv[0] < prev_sv[0] || (sv[0] == prev_sv[0] && sv[1] < prev_sv[1]) || (sv[0] == prev_sv[0] && sv[1] == prev_sv[1] && sv[2] < prev_sv[2]) {
				return nil, errors.New("invalid component version format in license: versions must be in ascending order")
			}
			prev_sv = sv

			switch op {
			case "=":
				if sv[0] == lm.appSemVer[0] && sv[1] == lm.appSemVer[1] && sv[2] == lm.appSemVer[2] {
					version_granted = true
					break GRANTED
				}
				min_ok = false
			case ">":
				min_ok = lm.appSemVer[0] > sv[0] || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] > sv[1]) || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] == sv[1] && lm.appSemVer[2] > sv[2])
			case ">=":
				min_ok = lm.appSemVer[0] > sv[0] || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] > sv[1]) || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] == sv[1] && lm.appSemVer[2] >= sv[2])
			case "<":
				if min_ok && (lm.appSemVer[0] < sv[0] || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] < sv[1]) || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] == sv[1] && lm.appSemVer[2] < sv[2])) {
					version_granted = true
					break GRANTED
				}
				min_ok = false
			case "<=":
				if min_ok && (lm.appSemVer[0] < sv[0] || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] < sv[1]) || (lm.appSemVer[0] == sv[0] && lm.appSemVer[1] == sv[1] && lm.appSemVer[2] <= sv[2])) {
					version_granted = true
					break GRANTED
				}
				min_ok = false
			default:
				return nil, errors.New("invalid component version format in license: unknown operator " + op)
			}
		}

		if !version_granted && !min_ok {
			return nil, errors.New("license is not valid for this application version")
		}

		break
	}

	return resp, nil
}

// GetLicense retrieves the current license key. If no license is set, it returns nil.
func (lm *licenseManager) GetLicense() *system.License {
	return lm.license
}

// HasLicense checks if a license is currently loaded.
func (lm *licenseManager) HasLicense() bool {
	return lm.license != nil
}

// WaitLicense blocks until a license is available or the context is done.
// It returns an error if the context is canceled, otherwise it returns nil when a license is available.
func (lm *licenseManager) WaitLicense(ctx context.Context) (*system.License, error) {
	if lm == nil {
		return nil, errors.New("license manager is not initialized")
	}
	for {
		if l := lm.license; l != nil {
			return l, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-lm.event:
		}
	}
}
