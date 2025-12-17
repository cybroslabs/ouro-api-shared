package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcapi"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svccrypto"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcdataproxy"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcdeviceregistry"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcourooperator"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svctaskmaster"
	"github.com/cybroslabs/ouro-api-shared/gen/go/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Mock Connectors for testing
type mockConnectors struct {
	license *system.License
	err     error
}

func (m *mockConnectors) OpenApiServiceClient() (svcapi.ApiInternalServiceClient, context.CancelFunc, error) {
	return nil, func() {}, errors.New("not implemented")
}

func (m *mockConnectors) OpenTaskmasterServiceClient() (svctaskmaster.TaskmasterServiceClient, context.CancelFunc, error) {
	return nil, func() {}, errors.New("not implemented")
}

func (m *mockConnectors) OpenDataproxyServiceClient() (svcdataproxy.DataproxyServiceClient, context.CancelFunc, error) {
	return nil, func() {}, errors.New("not implemented")
}

func (m *mockConnectors) OpenDeviceRegistryServiceClient() (svcdeviceregistry.DeviceRegistryServiceClient, context.CancelFunc, error) {
	return nil, func() {}, errors.New("not implemented")
}

func (m *mockConnectors) OpenOuroOperatorServiceClient() (svcourooperator.OuroOperatorServiceClient, context.CancelFunc, error) {
	if m.err != nil {
		return nil, func() {}, m.err
	}
	return &mockOuroOperatorServiceClient{license: m.license}, func() {}, nil
}

func (m *mockConnectors) OpenCryptoServiceClient() (svccrypto.CryptoServiceClient, context.CancelFunc, error) {
	return nil, func() {}, errors.New("not implemented")
}

type mockOuroOperatorServiceClient struct {
	svcourooperator.OuroOperatorServiceClient
	license *system.License
}

func (m *mockOuroOperatorServiceClient) GetLicense(ctx context.Context, req *emptypb.Empty, opts ...grpc.CallOption) (*system.License, error) {
	if m.license == nil {
		return nil, errors.New("no license available")
	}
	return m.license, nil
}

// Helper function to create a license with version constraints
func createLicenseWithVersions(serviceExp time.Time, versions ...string) *system.License {
	versionJSON, _ := json.Marshal(versions)
	return system.License_builder{
		ServiceExpiration: timestamppb.New(serviceExp),
		Options: map[string]string{
			string(system.LicensedItemComponentVersions): string(versionJSON),
		},
	}.Build()
}

func TestGetSemVerParts_Success(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected []int
	}{
		{"basic version", "1.2.3", []int{1, 2, 3}},
		{"with v prefix", "v1.2.3", []int{1, 2, 3}},
		{"with prerelease", "1.2.3-alpha.1", []int{1, 2, 3}},
		{"with build metadata", "1.2.3+build.123", []int{1, 2, 3}},
		{"full semver", "v2.10.15-beta.2+sha.abc123", []int{2, 10, 15}},
		{"large numbers", "999.999.999", []int{999, 999, 999}},
		{"zeros", "0.0.0", []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getSemVerParts(tt.version)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetSemVerParts_Invalid(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{"missing patch", "1.2"},
		{"missing minor and patch", "1"},
		{"letters in version", "a.b.c"},
		{"negative numbers", "-1.0.0"},
		{"empty string", ""},
		{"just v", "v"},
		{"spaces", "1 . 2 . 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getSemVerParts(tt.version)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Contains(t, err.Error(), "valid semver")
		})
	}
}

func TestGetSemVerOpParts_Success(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		expectedOp  string
		expectedVer []int
	}{
		{"no operator defaults to =", "1.2.3", "=", []int{1, 2, 3}},
		{"equals operator", "=1.2.3", "=", []int{1, 2, 3}},
		{"greater than", ">1.2.3", ">", []int{1, 2, 3}},
		{"greater than or equal", ">=1.2.3", ">=", []int{1, 2, 3}},
		{"less than", "<1.2.3", "<", []int{1, 2, 3}},
		{"less than or equal", "<=1.2.3", "<=", []int{1, 2, 3}},
		{"with v prefix", ">=v2.0.0", ">=", []int{2, 0, 0}},
		{"with prerelease", ">1.0.0-alpha", ">", []int{1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op, ver, err := getSemVerOpParts(tt.version)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedOp, op)
			assert.Equal(t, tt.expectedVer, ver)
		})
	}
}

func TestGetSemVerOpParts_Invalid(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{"invalid operator", "~1.2.3"},
		{"missing version", ">="},
		{"invalid version format", ">=1.2"},
		{"empty string", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op, ver, err := getSemVerOpParts(tt.version)
			assert.Error(t, err)
			assert.Empty(t, op)
			assert.Nil(t, ver)
		})
	}
}

func TestDrainChannel(t *testing.T) {
	t.Run("empty channel", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		drainChannel(ch)
		// Should not block
	})

	t.Run("channel with one item", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		ch <- struct{}{}
		drainChannel(ch)
		// Channel should be empty
		select {
		case <-ch:
			t.Error("channel should be empty")
		default:
			// Expected
		}
	})

	t.Run("channel with multiple items", func(t *testing.T) {
		ch := make(chan struct{}, 3)
		ch <- struct{}{}
		ch <- struct{}{}
		ch <- struct{}{}
		drainChannel(ch)
		// Channel should be empty
		select {
		case <-ch:
			t.Error("channel should be empty")
		default:
			// Expected
		}
	})

	t.Run("closed channel", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		close(ch)
		drainChannel(ch)
		// Should not panic
	})
}

func TestNewLicenseManager_Success(t *testing.T) {
	cancelCalled := false
	opts := &LicenseManagerOpts{
		AppVersion:   "1.0.0",
		AppName:      "test-app",
		AppBuiltDate: time.Now().Format(time.RFC3339),
		Connectors:   &mockConnectors{},
		Cancel: func() {
			cancelCalled = true
		},
	}

	lm, err := NewLicenseManager(opts)
	require.NoError(t, err)
	assert.NotNil(t, lm)

	// Clean up
	lm.Stop()
	assert.True(t, cancelCalled, "Cancel function should have been called")
}

func TestNewLicenseManager_DevVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{"dev lowercase", "dev"},
		{"dev uppercase", "Dev"},
		{"development lowercase", "development"},
		{"development mixed case", "Development"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &LicenseManagerOpts{
				AppVersion:   tt.version,
				AppName:      "test-app",
				AppBuiltDate: "", // Not used for dev version
				Connectors:   &mockConnectors{},
				Cancel:       func() {}, // Provide cancel function to avoid nil panic
			}

			lm, err := NewLicenseManager(opts)
			require.NoError(t, err)
			assert.NotNil(t, lm)

			// Dev version should have far-future semver
			lmImpl := lm.(*licenseManager)
			assert.Equal(t, []int{999, 999, 999}, lmImpl.appSemVer)
			assert.Equal(t, time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), lmImpl.appBuiltDate)

			lm.Stop()
		})
	}
}

func TestNewLicenseManager_NilOptions(t *testing.T) {
	lm, err := NewLicenseManager(nil)
	assert.Error(t, err)
	assert.Nil(t, lm)
	assert.Contains(t, err.Error(), "options cannot be nil")
}

func TestNewLicenseManager_InvalidVersion(t *testing.T) {
	opts := &LicenseManagerOpts{
		AppVersion:   "invalid",
		AppName:      "test-app",
		AppBuiltDate: time.Now().Format(time.RFC3339),
		Connectors:   &mockConnectors{},
		Cancel:       func() {},
	}

	lm, err := NewLicenseManager(opts)
	assert.Error(t, err)
	assert.Nil(t, lm)
	assert.Contains(t, err.Error(), "valid semver")
}

func TestNewLicenseManager_InvalidAppName(t *testing.T) {
	tests := []struct {
		name    string
		appName string
	}{
		{"uppercase letters", "TestApp"},
		{"starting with hyphen", "-testapp"},
		{"ending with hyphen", "testapp-"},
		{"numbers", "test123"},
		{"spaces", "test app"},
		{"empty", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &LicenseManagerOpts{
				AppVersion:   "1.0.0",
				AppName:      tt.appName,
				AppBuiltDate: time.Now().Format(time.RFC3339),
				Connectors:   &mockConnectors{},
				Cancel:       func() {},
			}

			lm, err := NewLicenseManager(opts)
			assert.Error(t, err)
			assert.Nil(t, lm)
			assert.Contains(t, err.Error(), "lowercase")
		})
	}
}

func TestNewLicenseManager_ValidAppNames(t *testing.T) {
	tests := []struct {
		name    string
		appName string
	}{
		{"simple", "test"},
		{"with hyphen", "test-app"},
		{"multiple hyphens", "test-app-name"},
		{"single letter", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &LicenseManagerOpts{
				AppVersion:   "1.0.0",
				AppName:      tt.appName,
				AppBuiltDate: time.Now().Format(time.RFC3339),
				Connectors:   &mockConnectors{},
				Cancel:       func() {},
			}

			lm, err := NewLicenseManager(opts)
			require.NoError(t, err)
			assert.NotNil(t, lm)
			lm.Stop()
		})
	}
}

func TestNewLicenseManager_InvalidBuiltDate(t *testing.T) {
	opts := &LicenseManagerOpts{
		AppVersion:   "1.0.0",
		AppName:      "test-app",
		AppBuiltDate: "not-a-date",
		Connectors:   &mockConnectors{},
		Cancel:       func() {},
	}

	lm, err := NewLicenseManager(opts)
	assert.Error(t, err)
	assert.Nil(t, lm)
	assert.Contains(t, err.Error(), "RFC3339")
}

func TestLicenseManager_GetLicense(t *testing.T) {
	lm := &licenseManager{}

	t.Run("no license", func(t *testing.T) {
		license := lm.GetLicense()
		assert.Nil(t, license)
	})

	t.Run("with license", func(t *testing.T) {
		expectedLicense := &system.License{}
		lm.license = expectedLicense

		license := lm.GetLicense()
		assert.Equal(t, expectedLicense, license)
	})
}

func TestLicenseManager_HasLicense(t *testing.T) {
	lm := &licenseManager{}

	t.Run("no license", func(t *testing.T) {
		has := lm.HasLicense()
		assert.False(t, has)
	})

	t.Run("with license", func(t *testing.T) {
		lm.license = &system.License{}
		has := lm.HasLicense()
		assert.True(t, has)
	})
}

func TestLicenseManager_WaitLicense_NilManager(t *testing.T) {
	var lm *licenseManager
	ctx := context.Background()

	license, err := lm.WaitLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.Contains(t, err.Error(), "not initialized")
}

func TestLicenseManager_WaitLicense_AlreadyHasLicense(t *testing.T) {
	expectedLicense := &system.License{}
	lm := &licenseManager{
		license: expectedLicense,
		event:   make(chan struct{}, 1),
	}

	ctx := context.Background()
	license, err := lm.WaitLicense(ctx)
	require.NoError(t, err)
	assert.Equal(t, expectedLicense, license)
}

func TestLicenseManager_WaitLicense_ContextCanceled(t *testing.T) {
	lm := &licenseManager{
		event: make(chan struct{}, 1),
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	license, err := lm.WaitLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestLicenseManager_Stop(t *testing.T) {
	runCancelCalled := false
	outterCancelCalled := false

	lm := &licenseManager{
		runCancel: func() {
			runCancelCalled = true
		},
		outterCancel: func() {
			outterCancelCalled = true
		},
	}

	lm.Stop()

	assert.True(t, runCancelCalled)
	assert.True(t, outterCancelCalled)
}

func TestLicenseManager_FetchLicense_NilManager(t *testing.T) {
	var lm *licenseManager
	ctx := context.Background()

	license, err := lm.fetchLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.Contains(t, err.Error(), "not initialized")
}

func TestLicenseManager_FetchLicense_ConnectorError(t *testing.T) {
	lm := &licenseManager{
		connectors: &mockConnectors{
			err: errors.New("connection failed"),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
}

func TestLicenseManager_FetchLicense_ServiceExpired(t *testing.T) {
	// License expired 1 day ago
	expiredTime := time.Now().Add(-24 * time.Hour)

	lm := &licenseManager{
		appVersion:   "1.0.0",
		appSemVer:    []int{1, 0, 0},
		appBuiltDate: time.Now(),
		connectors: &mockConnectors{
			license: system.License_builder{
				ServiceExpiration: timestamppb.New(expiredTime),
			}.Build(),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.Contains(t, err.Error(), "expired")
}

func TestLicenseManager_FetchLicense_Success_NoVersionConstraint(t *testing.T) {
	// License with no component versions should accept any version
	lm := &licenseManager{
		appName:      "test-app",
		appVersion:   "1.0.0",
		appSemVer:    []int{1, 0, 0},
		appBuiltDate: time.Now().Add(-24 * time.Hour),
		connectors: &mockConnectors{
			license: system.License_builder{
				ServiceExpiration: timestamppb.New(time.Now().Add(24 * time.Hour)),
			}.Build(),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	require.NoError(t, err)
	assert.NotNil(t, license)
}

func TestLicenseManager_FetchLicense_VersionCheck_Equals(t *testing.T) {
	tests := []struct {
		name       string
		appVer     []int
		licenseVer string
		shouldPass bool
	}{
		{"exact match", []int{1, 2, 3}, "=1.2.3", true},
		{"major mismatch", []int{2, 2, 3}, "=1.2.3", false},
		{"minor mismatch", []int{1, 3, 3}, "=1.2.3", false},
		{"patch mismatch", []int{1, 2, 4}, "=1.2.3", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm := &licenseManager{
				appName:      "test-app",
				appVersion:   "1.2.3",
				appSemVer:    tt.appVer,
				appBuiltDate: time.Now().Add(-24 * time.Hour),
				connectors: &mockConnectors{
					license: createLicenseWithVersions(
						time.Now().Add(24*time.Hour),
						"test-app:"+tt.licenseVer,
					),
				},
			}

			ctx := context.Background()
			license, err := lm.fetchLicense(ctx)
			if tt.shouldPass {
				require.NoError(t, err)
				assert.NotNil(t, license)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "not valid for this application version")
			}
		})
	}
}

func TestLicenseManager_FetchLicense_VersionCheck_GreaterThan(t *testing.T) {
	tests := []struct {
		name       string
		appVer     []int
		licenseVer string
		shouldPass bool
	}{
		{"app greater major", []int{2, 0, 0}, ">1.0.0", true},
		{"app greater minor", []int{1, 1, 0}, ">1.0.0", true},
		{"app greater patch", []int{1, 0, 1}, ">1.0.0", true},
		{"app equal", []int{1, 0, 0}, ">1.0.0", false},
		{"app less", []int{0, 9, 9}, ">1.0.0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm := &licenseManager{
				appName:      "test-app",
				appVersion:   "1.0.0",
				appSemVer:    tt.appVer,
				appBuiltDate: time.Now().Add(-24 * time.Hour),
				connectors: &mockConnectors{
					license: createLicenseWithVersions(
						time.Now().Add(24*time.Hour),
						"test-app:"+tt.licenseVer,
					),
				},
			}

			ctx := context.Background()
			license, err := lm.fetchLicense(ctx)
			if tt.shouldPass {
				require.NoError(t, err)
				assert.NotNil(t, license)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLicenseManager_FetchLicense_VersionCheck_Range(t *testing.T) {
	// Test range: >=1.0.0 <2.0.0
	tests := []struct {
		name       string
		appVer     []int
		shouldPass bool
	}{
		{"below range", []int{0, 9, 9}, false},
		{"at lower bound", []int{1, 0, 0}, true},
		{"within range", []int{1, 5, 0}, true},
		{"at upper bound", []int{2, 0, 0}, false},
		{"above range", []int{2, 1, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm := &licenseManager{
				appName:      "test-app",
				appVersion:   "1.0.0",
				appSemVer:    tt.appVer,
				appBuiltDate: time.Now().Add(-24 * time.Hour),
				connectors: &mockConnectors{
					license: createLicenseWithVersions(
						time.Now().Add(24*time.Hour),
						"test-app:>=1.0.0 <2.0.0",
					),
				},
			}

			ctx := context.Background()
			license, err := lm.fetchLicense(ctx)
			if tt.shouldPass {
				require.NoError(t, err)
				assert.NotNil(t, license)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLicenseManager_FetchLicense_VersionCheck_InvalidFormat(t *testing.T) {
	lm := &licenseManager{
		appName:      "test-app",
		appVersion:   "1.0.0",
		appSemVer:    []int{1, 0, 0},
		appBuiltDate: time.Now().Add(-24 * time.Hour),
		connectors: &mockConnectors{
			license: createLicenseWithVersions(
				time.Now().Add(24*time.Hour),
				"test-app:invalid",
			),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.Contains(t, err.Error(), "invalid component version format")
}

func TestLicenseManager_FetchLicense_VersionCheck_NotAscending(t *testing.T) {
	lm := &licenseManager{
		appName:      "test-app",
		appVersion:   "1.0.0",
		appSemVer:    []int{1, 5, 0},
		appBuiltDate: time.Now().Add(-24 * time.Hour),
		connectors: &mockConnectors{
			license: createLicenseWithVersions(
				time.Now().Add(24*time.Hour),
				"test-app:>=2.0.0 <1.0.0", // Versions not in ascending order
			),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	assert.Error(t, err)
	assert.Nil(t, license)
	assert.Contains(t, err.Error(), "ascending order")
}

func TestLicenseManager_FetchLicense_VersionCheck_DifferentApp(t *testing.T) {
	// License for different app should be ignored
	lm := &licenseManager{
		appName:      "test-app",
		appVersion:   "1.0.0",
		appSemVer:    []int{1, 0, 0},
		appBuiltDate: time.Now().Add(-24 * time.Hour),
		connectors: &mockConnectors{
			license: createLicenseWithVersions(
				time.Now().Add(24*time.Hour),
				"other-app:>=1.0.0",
			),
		},
	}

	ctx := context.Background()
	license, err := lm.fetchLicense(ctx)
	require.NoError(t, err)
	assert.NotNil(t, license)
}

// Benchmark tests
func BenchmarkGetSemVerParts(b *testing.B) {
	for b.Loop() {
		_, _ = getSemVerParts("v1.2.3-beta.1+build.123")
	}
}

func BenchmarkGetSemVerOpParts(b *testing.B) {
	for b.Loop() {
		_, _, _ = getSemVerOpParts(">=v1.2.3")
	}
}
