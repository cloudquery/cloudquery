package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLakebase_Omitted_OK(t *testing.T) {
	s := baseSpec()
	require.NoError(t, s.Validate())
	require.False(t, s.HasLakebaseConfig())
}

func TestLakebase_Valid_OK(t *testing.T) {
	s := baseSpec()
	s.Lakebase = &LakebaseSpec{
		Endpoint:     "projects/p/branches/b/endpoints/e",
		Host:         "https://example.cloud.databricks.com",
		ClientID:     "id",
		ClientSecret: "secret",
	}
	require.NoError(t, s.Validate())
	require.True(t, s.HasLakebaseConfig())
}

func TestLakebase_MissingEndpoint_Error(t *testing.T) {
	s := baseSpec()
	s.Lakebase = &LakebaseSpec{
		Host: "https://example.cloud.databricks.com",
	}
	require.Error(t, s.Validate())
}
