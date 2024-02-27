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
			Name: "spec with bad endpoint 1",
			Spec: `{"endpoint": "http://localhost:8182"}`,
			Err:  true,
		},
		{
			Name: "spec with bad endpoint 2",
			Spec: `{"endpoint": "wss://"}`,
			Err:  true,
		},
		{
			Name: "spec with bad endpoint 3",
			Spec: `{"endpoint": "ws://"}`,
			Err:  true,
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
			Name: "spec with endpoint, none auth_mode",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "none"}`,
		},
		{
			Name: "spec with endpoint, invalid auth_mode",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "invalid"}`,
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
			Name: "spec with endpoint, auth_mode aws, region, without host",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"us-east-1"}`,
		},
		{
			Name: "spec with endpoint, auth_mode aws, region, host",
			Spec: `{"endpoint": "ws://localhost:8182", "auth_mode": "aws", "aws_region":"us-east-1", "aws_neptune_host":"my-neptune.cluster-testtesttest.us-east-1.neptune.amazonaws.com"}`,
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
