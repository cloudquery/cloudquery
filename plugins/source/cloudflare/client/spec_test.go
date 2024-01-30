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
			Name: "spec with api key",
			Spec: `{"api_key": "key"}`,
		},
		{
			Name: "spec with integer api key",
			Spec: `{"api_key": 1234}`,
			Err:  true,
		},
		{
			Name: "spec with api email",
			Spec: `{"api_email": "email"}`,
		},
		{
			Name: "spec with accounts",
			Spec: `{"accounts": ["account1", "account2"]}`,
		},
		{
			Name: "spec with empty accounts",
			Spec: `{"accounts": []}`,
		},
		{
			Name: "spec with null accounts",
			Spec: `{"accounts": null}`,
		},
		{
			Name: "spec with integer accounts",
			Spec: `{"accounts": [123, 456]}`,
			Err:  true,
		},
		{
			Name: "spec with zones",
			Spec: `{"zones": ["zone1", "zone2"]}`,
		},
		{
			Name: "spec with empty zones",
			Spec: `{"zones": []}`,
		},
		{
			Name: "spec with null zones",
			Spec: `{"zones": null}`,
		},
		{
			Name: "spec with concurrency",
			Spec: `{"concurrency": 100}`,
		},
		{
			Name: "spect with unknown field",
			Spec: `{"bad_configuration": "bad"}`,
			Err:  true,
		},
	})
}
