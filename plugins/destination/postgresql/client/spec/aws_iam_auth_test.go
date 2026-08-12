package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAWSIAMAuth_Omitted_OK(t *testing.T) {
	s := baseSpec()
	require.NoError(t, s.Validate())
	require.False(t, s.HasAWSIAMAuthConfig())
}

func TestAWSIAMAuth_Empty_OK(t *testing.T) {
	s := baseSpec()
	s.AWSIAMAuth = &AWSIAMAuthSpec{}
	require.NoError(t, s.Validate())
	require.True(t, s.HasAWSIAMAuthConfig())
	require.Equal(t, AWSIAMAuthServiceRDS, s.AWSIAMAuth.ServiceOrDefault())
}

func TestAWSIAMAuth_Valid_OK(t *testing.T) {
	s := baseSpec()
	s.AWSIAMAuth = &AWSIAMAuthSpec{
		Service:         AWSIAMAuthServiceRDS,
		Region:          "us-east-1",
		Endpoint:        "mydb.123456789012.us-east-1.rds.amazonaws.com:5432",
		LocalProfile:    "my_profile",
		RoleARN:         "arn:aws:iam::123456789012:role/my-role",
		RoleSessionName: "cloudquery",
		ExternalID:      "external-id",
	}
	require.NoError(t, s.Validate())
	require.True(t, s.HasAWSIAMAuthConfig())
}

func TestAWSIAMAuth_UnknownService_Error(t *testing.T) {
	s := baseSpec()
	s.AWSIAMAuth = &AWSIAMAuthSpec{Service: "dsql"}
	require.ErrorContains(t, s.Validate(), "`aws_iam_auth.service` must be one of: rds")
}

func TestAWSIAMAuth_WithLakebase_Error(t *testing.T) {
	s := baseSpec()
	s.AWSIAMAuth = &AWSIAMAuthSpec{Region: "us-east-1"}
	s.Lakebase = &LakebaseSpec{Endpoint: "projects/p/branches/b/endpoints/e"}
	require.Error(t, s.Validate())
}
