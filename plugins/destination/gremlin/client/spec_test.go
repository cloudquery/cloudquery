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
			Name: "spec with endpoint",
			Spec: `{"endpoint": "ws://localhost:8182"}`,
		},
		{
			Name: "spec with wss endpoint, without port",
			Spec: `{"endpoint": "wss://localhost"}`,
		},
		{
			Name: "spec with endpoint, auth_mode basic, username/password",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "basic", "username": "abc", "password": "def"}`,
		},
		{
			Name: "spec with endpoint, auth_mode basic, username but no password",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "basic", "username": "abc"}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, username",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"reg", "username": "abc"}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, password",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"reg", "password": "abc"}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, username/password",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"reg", "username": "abc", "password": "def"}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, no region",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws"}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, empty region",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":""}`,
			Err:  true,
		},
		{
			Name: "spec with endpoint, auth_mode aws, region",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"reg"}`,
		},
		{
			Name: "spec with bool endpoint",
			Spec: `{"endpoint": true}`,
			Err:  true,
		},
		{
			Name: "spec with null endpoint",
			Spec: `{"endpoint": null}`,
			Err:  true,
		},
		{
			Name: "spec with int endpoint",
			Spec: `{"endpoint": 123}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"endpoint": "ws://localhost:8182", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"endpoint": "ws://localhost:8182", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"endpoint": "ws://localhost:8182", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"endpoint": "ws://localhost:8182", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"endpoint": "ws://localhost:8182", "unknown": "test"}`,
			Err:  true,
		},
	})
}
