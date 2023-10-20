package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestOAuthSpec_JSONSchemaExtend(t *testing.T) {
	sc, err := jsonschema.Generate(OAuthSpec{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(sc), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // either access_token, or client_id/client_secret
			Spec: `{}`,
		},
		{
			Name: "empty access_token",
			Err:  true,
			Spec: `{"access_token":""}`,
		},
		{
			Name: "access_token",
			Spec: `{"access_token":"abc"}`,
		},
		{
			Name: "no access_token with empty client_id & no client_secret",
			Err:  true,
			Spec: `{"client_id":""}`,
		},
		{
			Name: "no access_token with no client_id & empty client_secret",
			Err:  true,
			Spec: `{"client_secret":""}`,
		},
		{
			Name: "no access_token with empty client_id & empty client_secret",
			Err:  true,
			Spec: `{"client_id":"","client_secret":""}`,
		},
		{
			Name: "no access_token with client_id & empty client_secret",
			Err:  true,
			Spec: `{"client_id":"id","client_secret":""}`,
		},
		{
			Name: "no access_token with empty client_id & client_secret",
			Err:  true,
			Spec: `{"client_id":"","client_secret":"secret"}`,
		},
		{
			Name: "no access_token with client_id & client_secret",
			Spec: `{"client_id":"id","client_secret":"secret"}`,
		},
		{
			Name: "empty access_token with empty client_id & no client_secret",
			Err:  true,
			Spec: `{"access_token":"","client_id":""}`,
		},
		{
			Name: "empty access_token with no client_id & empty client_secret",
			Err:  true,
			Spec: `{"access_token":"","client_secret":""}`,
		},
		{
			Name: "empty access_token with empty client_id & empty client_secret",
			Err:  true,
			Spec: `{"access_token":"","client_id":"","client_secret":""}`,
		},
		{
			Name: "empty access_token with client_id & empty client_secret",
			Err:  true,
			Spec: `{"access_token":"","client_id":"id","client_secret":""}`,
		},
		{
			Name: "empty access_token with empty client_id & client_secret",
			Err:  true,
			Spec: `{"access_token":"","client_id":"","client_secret":"secret"}`,
		},
		{
			Name: "empty access_token with client_id & client_secret",
			Spec: `{"access_token":"","client_id":"id","client_secret":"secret"}`,
		},
		{
			Name: "access_token with empty client_id & no client_secret",
			Spec: `{"access_token":"token","client_id":""}`,
		},
		{
			Name: "access_token with no client_id & empty client_secret",
			Spec: `{"access_token":"token","client_secret":""}`,
		},
		{
			Name: "access_token with empty client_id & empty client_secret",
			Spec: `{"access_token":"token","client_id":"","client_secret":""}`,
		},
		{
			Name: "access_token with client_id & empty client_secret",
			Spec: `{"access_token":"token","client_id":"id","client_secret":""}`,
		},
		{
			Name: "access_token with empty client_id & client_secret",
			Spec: `{"access_token":"token","client_id":"","client_secret":"secret"}`,
		},
		{
			Name: "access_token with client_id & client_secret",
			Spec: `{"access_token":"token","client_id":"id","client_secret":"secret"}`,
		},
	})
}
