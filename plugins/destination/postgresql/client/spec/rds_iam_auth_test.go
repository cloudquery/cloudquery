package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRDSIAMAuth_Omitted_OK(t *testing.T) {
	s := baseSpec()
	require.NoError(t, s.Validate())
	require.False(t, s.HasRDSIAMAuthConfig())
}

func TestRDSIAMAuth_Empty_OK(t *testing.T) {
	// Every field is optional: the AWS SDK resolves the region and credentials from
	// the standard AWS configuration sources.
	s := baseSpec()
	s.RDSIAMAuth = &RDSIAMAuthSpec{}
	require.NoError(t, s.Validate())
	require.True(t, s.HasRDSIAMAuthConfig())
}

func TestRDSIAMAuth_Valid_OK(t *testing.T) {
	s := baseSpec()
	s.RDSIAMAuth = &RDSIAMAuthSpec{
		Region:          "us-east-1",
		Endpoint:        "mydb.123456789012.us-east-1.rds.amazonaws.com:5432",
		LocalProfile:    "my_profile",
		RoleARN:         "arn:aws:iam::123456789012:role/my-role",
		RoleSessionName: "cloudquery",
		ExternalID:      "external-id",
	}
	require.NoError(t, s.Validate())
	require.True(t, s.HasRDSIAMAuthConfig())
}

func TestRDSIAMAuth_WithLakebase_Error(t *testing.T) {
	s := baseSpec()
	s.RDSIAMAuth = &RDSIAMAuthSpec{Region: "us-east-1"}
	s.Lakebase = &LakebaseSpec{Endpoint: "projects/p/branches/b/endpoints/e"}
	require.Error(t, s.Validate())
}
