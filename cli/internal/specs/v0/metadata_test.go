package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestMetadata_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(Metadata))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{}`, // name & path are always required
		},
		{
			Name: "null",
			Err:  true,
			Spec: `null`,
		},
		{
			Name: "bad type",
			Err:  true,
			Spec: `123`,
		},
		{
			Name: "minimal",
			Spec: `{"name":"a","path":"b","registry":"docker"}`,
		},
		{
			Name: "empty name",
			Err:  true,
			Spec: `{"name":"","path":"b","registry":"docker"}`,
		},
		{
			Name: "null name",
			Err:  true,
			Spec: `{"name":null,"path":"b","registry":"docker"}`,
		},
		{
			Name: "bad name type",
			Err:  true,
			Spec: `{"name":123,"path":"b","registry":"docker"}`,
		},
		{
			Name: "empty path",
			Err:  true,
			Spec: `{"name":"a","path":"","registry":"docker"}`,
		},
		{
			Name: "null path",
			Err:  true,
			Spec: `{"name":"a","path":null,"registry":"docker"}`,
		},
		{
			Name: "bad path type",
			Err:  true,
			Spec: `{"name":"a","path":123,"registry":"docker"}`,
		},
		{
			Name: "null registry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":null}`,
		},
		{
			Name: "bad registry type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":123}`,
		},
		{
			Name: "registry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":123}`,
		},
		{
			Name: "empty registry without version",
			Err:  true, // this will imply CQ, but CQ requires version
			Spec: `{"name":"a","path":"b","registry":""}`,
		},
		{
			Name: "empty registry with version",
			Spec: `{"name":"a","path":"b","registry":"","version":"v0"}`,
		},
		{
			Name: "github registry without version",
			Err:  true, // this requires version
			Spec: `{"name":"a","path":"b","registry":"github"}`,
		},
		{
			Name: "github registry with version",
			Spec: `{"name":"a","path":"b","registry":"github","version":"v0"}`,
		},
		{
			Name: "local registry without version",
			Spec: `{"name":"a","path":"b","registry":"local"}`,
		},
		{
			Name: "local registry with version", // we just ignore version
			Spec: `{"name":"a","path":"b","registry":"local","version":"v0"}`,
		},
		{
			Name: "grpc registry without version",
			Spec: `{"name":"a","path":"b","registry":"grpc"}`,
		},
		{
			Name: "grpc registry with version", // we just ignore version
			Spec: `{"name":"a","path":"b","registry":"grpc","version":"v0"}`,
		},
		{
			Name: "docker registry without version",
			Spec: `{"name":"a","path":"b","registry":"docker"}`,
		},
		{
			Name: "docker registry with version", // we just ignore version
			Spec: `{"name":"a","path":"b","registry":"docker","version":"v0"}`,
		},
		{
			Name: "cloudquery registry without version",
			Err:  true, // this requires version
			Spec: `{"name":"a","path":"b","registry":"cloudquery"}`,
		},
		{
			Name: "cloudquery registry with version",
			Spec: `{"name":"a","path":"b","registry":"cloudquery","version":"v0"}`,
		},
		{
			Name: "empty version",
			Spec: `{"name":"a","path":"b","registry":"local","version":""}`,
		},
		{
			Name: "null version",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","version":null}`,
		},
		{
			Name: "bad version type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","version":123}`,
		},
		{
			Name: "bad version format when ignored", // we ignore version for registry: local
			Spec: `{"name":"a","path":"b","registry":"local","version":"a"}`,
		},
		{
			Name: "proper version format when ignored", // we ignore version for registry: local
			Spec: `{"name":"a","path":"b","registry":"local","version":"v0"}`,
		},
		{
			Name: "bad version format when accounted for",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"","version":"a"}`,
		},
		{
			Name: "proper version format when accounted for",
			Spec: `{"name":"a","path":"b","registry":"","version":"v0"}`,
		},
	})
}
