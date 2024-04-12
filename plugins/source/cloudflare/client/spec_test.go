package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "spec with token",
			Spec: `{"api_token": "secret"}`,
		},
		{
			Name: "spec with integer token",
			Spec: `{"api_token": 123}`,
			Err:  true,
		},
		{
			Name: "spec with api key and email",
			Spec: `{"api_key": "key", "api_email": "email"}`,
		},
		{
			Name: "spec with integer api key",
			Spec: `{"api_key": 1234, "api_email": "email"}`,
			Err:  true,
		},
		{
			Name: "spec with integer api email",
			Spec: `{"api_key": "email", "api_email": 1234}`,
			Err:  true,
		},
		{
			Name: "spec with accounts",
			Spec: `{"api_token": "secret", "accounts": ["account1", "account2"]}`,
		},
		{
			Name: "spec with empty accounts",
			Spec: `{"api_token": "secret", "accounts": []}`,
		},
		{
			Name: "spec with null accounts",
			Spec: `{"api_token": "secret", "accounts": null}`,
		},
		{
			Name: "spec with integer accounts",
			Spec: `{"accounts": [123, 456]}`,
			Err:  true,
		},
		{
			Name: "spec with zones",
			Spec: `{"api_token": "secret", "zones": ["zone1", "zone2"]}`,
		},
		{
			Name: "spec with empty zones",
			Spec: `{"api_token": "secret", "zones": []}`,
		},
		{
			Name: "spec with null zones",
			Spec: `{"api_token": "secret", "zones": null}`,
		},
		{
			Name: "spec with concurrency",
			Spec: `{"api_token": "secret", "concurrency": 100}`,
		},
		{
			Name: "spect with unknown field",
			Spec: `{"bad_configuration": "bad"}`,
			Err:  true,
		},
	})
}
