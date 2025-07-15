package spec

import (
	"fmt"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestSpec_Validate(t *testing.T) {
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database"}, WantErr: false},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "", Database: "database"}, WantErr: true},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: ""}, WantErr: true},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{Default: true}}, WantErr: false},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{RoleARN: "arn:aws:iam::123456789012:role/role_name"}}, WantErr: false},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{LocalProfile: "test_profile"}}, WantErr: false},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{Default: true, RoleARN: "arn:aws:iam::123456789012:role/role_name"}}, WantErr: true},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{Default: true, LocalProfile: "test_profile"}}, WantErr: true},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{Default: true, RoleSessionName: "test-session"}}, WantErr: true},
		{Give: Spec{BatchSize: int64(0), BatchSizeBytes: int64(0), ConnectionString: "test-connection-string", Database: "database", AWSCredentials: &Credentials{}}, WantErr: true},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			err := tc.Give.Validate()
			if tc.WantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string",
			Spec: `{"connection_string": "conn"}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string and database",
			Spec: `{"connection_string": "conn", "database":"foo"}`,
		},
		{
			Name: "spec with bool connection_string",
			Spec: `{"connection_string": true, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with null connection_string",
			Spec: `{"connection_string": null, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with int connection_string",
			Spec: `{"connection_string": 123, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "database":"foo", "unknown": "test"}`,
			Err:  true,
		},
		{
			Name: "spec with valid default aws_credentials",
			Spec: `{"connection_string": "abc", "database":"foo", "aws_credentials": {"default": true}}`,
			Err:  false,
		},
		{
			Name: "spec with valid assume_role in aws_credentials",
			Spec: `{"connection_string": "abc", "database":"foo", "aws_credentials": {"role_arn": "arn:aws:iam::123456789012:role/role_name"}}`,
			Err:  false,
		},
		{
			Name: "spec with valid local_profile in aws_credentials",
			Spec: `{"connection_string": "abc", "database":"foo", "aws_credentials": {"local_profile": "test_profile"}}`,
			Err:  false,
		},
		{
			Name: "invalid spec with both valid local_profile and default in aws_credentials",
			Spec: `{"connection_string": "abc", "database":"foo", "aws_credentials": {"default": true,"local_profile": "test_profile"}}`,
			Err:  false,
		},
	})
}
