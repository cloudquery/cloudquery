package client

import (
	"os"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/require"
)

func TestOCIConfigurationProvider(t *testing.T) {
	configProvider, err := common.ComposingConfigurationProvider(
		[]common.ConfigurationProvider{
			common.ConfigurationProviderEnvironmentVariables("OCI_CLI", ""),
			common.DefaultConfigProvider(),
		},
	)
	require.NoError(t, err)

	prefixes := []string{"TF_VAR", "OCI_CLI"}
	for _, pfx := range prefixes {
		t.Run(pfx, func(t *testing.T) {
			envRegion, ok := os.LookupEnv(pfx + "_region")
			require.False(t, ok)
			require.Empty(t, envRegion)

			sdkRegion, err := configProvider.Region()
			require.Error(t, err)
			require.Empty(t, sdkRegion)

			require.NoError(t, os.Setenv(pfx+"_region", "region_"+pfx))

			envRegion, ok = os.LookupEnv(pfx + "_region")
			require.True(t, ok)
			require.Equal(t, "region_"+pfx, envRegion)

			sdkRegion, err = configProvider.Region()
			require.NoError(t, err)
			require.Equal(t, "region_"+pfx, sdkRegion)

			require.NoError(t, os.Unsetenv(pfx+"_region"))
		})
	}
}
