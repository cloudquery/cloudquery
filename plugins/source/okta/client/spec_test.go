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
			Spec: `{"token": "tok"}`,
			Err:  true,
		},
		{
			Name: "spec with token and domain",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com"}`,
		},
		{
			Name: "spec with token and invalid domain",
			Spec: `{"token": "tok", "domain": "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"}`,
			Err:  true,
		},
		{
			Name: "spec with token and domain and empty rate limit",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "rate_limit": {}}`,
		},
		{
			Name: "spec with token and domain and null rate limit",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "rate_limit": null}`,
		},
		{
			Name: "spec with token and domain and rate limit",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "rate_limit": {"max_backoff": 60}}`,
		},
		{
			Name: "spec with token and domain and invalid rate limit",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "rate_limit": {"max_backoff": true}}`,
			Err:  true,
		},
		{
			Name: "spec with token and domain and zero rate limit",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "rate_limit": {"max_backoff": 0}}`,
			Err:  true,
		},
		{
			Name: "spec with bool concurrency",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "concurrency":false}`,
			Err:  true,
		},
		{
			Name: "spec with null concurrency",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "concurrency":null}`,
			Err:  true,
		},
		{
			Name: "spec with string concurrency",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "concurrency":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with proper concurrency",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "concurrency": 7}`,
		},
		{
			Name: "spec with array concurrency",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "concurrency":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"token": "tok", "domain": "https://domain.okta.com", "unknown": "test"}`,
			Err:  true,
		},
	})
}
