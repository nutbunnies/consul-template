package dependency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	VaultDefaultLeaseDuration = 0
}

func TestVaultRenewDuration(t *testing.T) {
	//VaultDefaultLeaseDuration = 5 * time.Minute
	testCases := []struct {
		desc        string
		secret      *Secret
		expectedMin float64
		expectedMax float64
	}{
		//{
		//	desc:        "base renewable lease - no duration",
		//	secret:      Secret{LeaseDuration: 0, Renewable: true, Auth: &SecretAuth{LeaseDuration: 0, Renewable: true}},
		//	expectedMin: 50,
		//	expectedMax: 100,
		//},
		//{
		//	desc:        "base non-renewable lease - no duration",
		//	secret:      Secret{LeaseDuration: 0, Renewable: false, Auth: &SecretAuth{LeaseDuration: 0, Renewable: false}},
		//	expectedMin: 255, // 85% of 5 minutes
		//	expectedMax: 594, // 99% of 5 minutes
		//},
		{
			desc:        "base renewable lease",
			secret:      &Secret{LeaseDuration: 10000, Renewable: true},
			expectedMin: 1665,
			expectedMax: 3334,
		},
		{
			desc:        "base non-renewable lease",
			secret:      &Secret{LeaseDuration: 10000, Renewable: false},
			expectedMin: 8500,
			expectedMax: 9900,
		},
		{
			desc:        "auth renewable lease",
			secret:      &Secret{LeaseDuration: 0, Renewable: true, Auth: &SecretAuth{LeaseDuration: 10000, Renewable: true}},
			expectedMin: 1665,
			expectedMax: 3334,
		},
		{
			desc:        "auth non-renewable lease",
			secret:      &Secret{LeaseDuration: 0, Renewable: false, Auth: &SecretAuth{LeaseDuration: 10000, Renewable: false}},
			expectedMin: 8500,
			expectedMax: 9900,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dur := vaultRenewDuration(tC.secret)
			assert.True(t, dur.Seconds() >= tC.expectedMin, fmt.Sprintf("expected: %f to be >= %f", dur.Seconds(), tC.expectedMin))
			assert.True(t, dur.Seconds() <= tC.expectedMax, fmt.Sprintf("expected: %f to be <= %f", dur.Seconds(), tC.expectedMax))
		})
	}
	//VaultDefaultLeaseDuration = 0
}
