package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestRegionMetadata(t *testing.T) {
	data, err := jsonschema.Generate(RegionMetadata{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{}`,
		},
		{
			Name: "bad type",
			Err:  true,
			Spec: `123`,
		},
		{
			Name: "proper",
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "missing realmKey",
			Err:  true,
			Spec: `{"realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "empty realmKey",
			Err:  true,
			Spec: `{"realmKey":"","realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "null realmKey",
			Err:  true,
			Spec: `{"realmKey":null,"realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "bad realmKey",
			Err:  true,
			Spec: `{"realmKey":123,"realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "missing realmDomainComponent",
			Err:  true,
			Spec: `{"realmKey":"OC1","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "empty realmDomainComponent",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"","regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "null realmDomainComponent",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":null,"regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "bad realmDomainComponent",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":123,"regionKey":"SYD","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "missing regionKey",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "empty regionKey",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"","regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "null regionKey",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":null,"regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "bad regionKey",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":123,"regionIdentifier":"ap-sydney-1"}`,
		},
		{
			Name: "missing regionIdentifier",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"SYD"}`,
		},
		{
			Name: "empty regionIdentifier",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":""}`,
		},
		{
			Name: "null regionIdentifier",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":null}`,
		},
		{
			Name: "bad regionIdentifier",
			Err:  true,
			Spec: `{"realmKey":"OC1","realmDomainComponent":"oraclecloud.com","regionKey":"SYD","regionIdentifier":123}`,
		},
	})
}
