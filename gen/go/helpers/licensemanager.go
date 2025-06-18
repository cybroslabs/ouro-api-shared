package helpers

import (
	"context"
	"errors"

	"github.com/cybroslabs/ouro-api-shared/gen/go/system"
)

// LicenseManager is an interface that defines methods for license operations in the system.
type LicenseManager interface {
	// GetLicense retrieves the current license key.
	GetLicense(ctx context.Context) (*system.License, error)
}

// The LicenseManagerOpts struct contains options for creating a LicenseManager instance.
type LicenseManagerOpts struct {
	Connectors Connectors
}

type licenseManager struct {
	connectors Connectors
}

// NewLicenseManager creates a new instance of LicenseManager with the provided options.
func NewLicenseManager(opts *LicenseManagerOpts) LicenseManager {
	if opts == nil {
		return nil
	}
	return &licenseManager{
		connectors: opts.Connectors,
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
