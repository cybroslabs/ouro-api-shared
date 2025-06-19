package helpers

import (
	"context"
	"errors"
	"time"

	"github.com/cybroslabs/ouro-api-shared/gen/go/system"
)

const (
	_periodNoLicense  = 15 * time.Second
	_periodRevalidate = 30 * time.Minute
)

// LicenseManager is an interface that defines methods for license operations in the system.
type LicenseManager interface {
	// GetLicense retrieves the current license key.
	GetLicense(ctx context.Context) (*system.License, error)
	// HasLicense checks if a license is currently set.
	HasLicense() bool
	// Stop stops the license manager and releases any resources it holds.
	Stop()
}

// The LicenseManagerOpts struct contains options for creating a LicenseManager instance.
type LicenseManagerOpts struct {
	Connectors Connectors
	Cancel     context.CancelFunc
	OnLicense  func(license *system.License)
}

type licenseManager struct {
	connectors Connectors
	license    *system.License
	onLicense  func(license *system.License)

	runCancel    context.CancelFunc
	outterCancel context.CancelFunc
}

// NewLicenseManager creates a new instance of LicenseManager with the provided options.
func NewLicenseManager(opts *LicenseManagerOpts) LicenseManager {
	if opts == nil {
		return nil
	}
	ctx, ctx_cancel := context.WithCancel(context.Background())
	lm := &licenseManager{
		connectors:   opts.Connectors,
		runCancel:    ctx_cancel,
		outterCancel: opts.Cancel,
		onLicense:    opts.OnLicense,
	}
	go lm.run(ctx)
	return lm
}

func (lm *licenseManager) Stop() {
	lm.runCancel()
	lm.outterCancel()
}

func (lm *licenseManager) run(ctx context.Context) {
	check_period := _periodNoLicense
	check_failed_count := 0
	for {
		license, err := lm.GetLicense(ctx)
		if err != nil {
			check_failed_count++
			if check_failed_count == 3 {
				// If we failed to get the license three times in a row, we will revoke current license.
				lm.license = nil
				if f := lm.onLicense; f != nil {
					f(nil)
				}
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(check_period):
				continue
			}
		}

		lm.license = license
		if f := lm.onLicense; f != nil {
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
func (lm *licenseManager) GetLicense(ctx context.Context) (*system.License, error) {
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
	return resp, nil
}

func (lm *licenseManager) HasLicense() bool {
	return lm.license != nil
}
