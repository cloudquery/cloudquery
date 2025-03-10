package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

type testDestinationSpec struct {
	ConnectionString string `json:"connection_string"`
}

func TestDestinationSpecUnmarshalSpec(t *testing.T) {
	destination := Destination{
		Spec: map[string]any{
			"connection_string": "postgres://user:pass@host:port/db",
		},
	}
	var spec testDestinationSpec
	if err := destination.UnmarshalSpec(&spec); err != nil {
		t.Fatal(err)
	}
	if spec.ConnectionString != "postgres://user:pass@host:port/db" {
		t.Fatalf("expected postgres://user:pass@host:port/db, got %s", spec.ConnectionString)
	}
}

var destinationUnmarshalSpecTestCases = []struct {
	name   string
	spec   string
	err    string
	source *Source
}{
	{
		"invalid_kind",
		`kind: nice`,
		"failed to decode spec: unknown kind nice",
		nil,
	},
	{
		"invalid_type",
		`kind: source
spec:
  name: 3
`,
		"failed to decode spec: json: cannot unmarshal number into Go struct field Source.Metadata.name of type string",
		&Source{
			Metadata: Metadata{Name: "test"},
			Tables:   []string{"*"},
		},
	},
	{
		"unknown_field",
		`kind: source
spec:
  namea: 3
`,
		`failed to decode spec: json: unknown field "namea"`,
		&Source{
			Metadata: Metadata{Name: "test"},
			Tables:   []string{"*"},
		},
	},
}

func TestDestinationUnmarshalSpec(t *testing.T) {
	for _, tc := range destinationUnmarshalSpecTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			var spec Spec
			err = SpecUnmarshalYamlStrict([]byte(tc.spec), &spec)
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("expected:%s got:%s", tc.err, err.Error())
				}
				return
			}

			source := spec.Spec.(*Source)
			require.Equal(t, tc.source, source)
		})
	}
}

var destinationUnmarshalSpecValidateTestCases = []struct {
	name        string
	spec        string
	err         string
	destination *Destination
}{
	{
		"required_name",
		`kind: destination
spec:`,
		"name is required",
		nil,
	},
	{
		"required_version",
		`kind: destination
spec:
  name: test
  path: cloudquery/test
`,
		"version is required",
		nil,
	},
	{
		"required_version_format",
		`kind: destination
spec:
  name: test
  path: cloudquery/test
  version: 1.1.0
`,
		"version must start with v",
		nil,
	},
	{
		"version_is_not_required_for_grpc_registry",
		`kind: destination
spec:
  name: test
  registry: grpc
  path: "localhost:9999"
`,
		"",
		&Destination{
			Metadata: Metadata{
				Name:     "test",
				Registry: RegistryGRPC,
				Path:     "localhost:9999",
			},
			Spec: map[string]any{},
		},
	},
	{
		"version_is_not_required_for_local_registry",
		`kind: destination
spec:
  name: test
  registry: local
  path: "/home/user/some_executable"
`,
		"",
		&Destination{
			Metadata: Metadata{
				Name:     "test",
				Registry: RegistryLocal,
				Path:     "/home/user/some_executable",
			},
			Spec: map[string]any{},
		},
	},
	{
		"success",
		`kind: destination
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
`,
		"",
		&Destination{
			Metadata: Metadata{
				Name:             "test",
				Registry:         RegistryCloudQuery,
				Path:             "cloudquery/test",
				Version:          "v1.1.0",
				registryInferred: true,
			},
			Spec: map[string]any{},
		},
	},
	{
		"success github",
		`kind: destination
spec:
  name: test
  path: cloudquery/test
  registry: cloudquery
  version: v1.1.0
`,
		"",
		&Destination{
			Metadata: Metadata{
				Name:     "test",
				Registry: RegistryCloudQuery,
				Path:     "cloudquery/test",
				Version:  "v1.1.0",
			},
			Spec: map[string]any{},
		},
	},
}

func TestDestinationUnmarshalSpecValidate(t *testing.T) {
	for _, tc := range destinationUnmarshalSpecValidateTestCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var err error
			var spec Spec
			err = SpecUnmarshalYamlStrict([]byte(tc.spec), &spec)
			if err != nil {
				t.Fatal(err)
			}
			destination := spec.Spec.(*Destination)
			destination.SetDefaults()
			err = destination.Validate()
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("expected:\n%s\ngot:\n%s", tc.err, err.Error())
				}
				return
			}

			require.EqualValues(t, tc.destination, destination)
		})
	}
}

func TestDestination_VersionString(t *testing.T) {
	tests := []struct {
		name string
		meta Metadata
		want string
	}{
		{
			name: "should use short version without name part in path when those are the same",
			meta: Metadata{
				Name:     "aws",
				Version:  "v10.0.0",
				Path:     "cloudquery/aws",
				Registry: RegistryCloudQuery,
			},
			want: "aws (cloudquery/aws@v10.0.0)",
		},
		{
			name: "should use long version with path when name doesn't match path",
			meta: Metadata{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "cloudquery/aws",
				Registry: RegistryCloudQuery,
			},
			want: "my-aws-spec (cloudquery/aws@v10.0.0)",
		},
		{
			name: "should handle non CloudQuery Hub registry",
			meta: Metadata{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "localhost:7777",
				Registry: RegistryGRPC,
			},
			want: "my-aws-spec (grpc@localhost:7777)",
		},
		{
			name: "should handle malformed path",
			meta: Metadata{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "aws",
				Registry: RegistryCloudQuery,
			},
			want: "my-aws-spec (aws@v10.0.0)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Destination{Metadata: tt.meta}
			if got := d.VersionString(); got != tt.want {
				t.Errorf("Destination.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDestination_JSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(Destination))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{}`,
		},
		{
			Name: "null",
			Err:  true,
			Spec: `null`,
		},
		{
			Name: "bad type",
			Err:  true,
			Spec: `[]`,
		},
		{
			Name: "missing spec",
			Spec: `{"name":"a","path":"b","registry":"local"}`,
		},
		{
			Name: "empty spec",
			Spec: `{"name":"a","path":"b","registry":"local","spec":{}}`,
		},
		{
			Name: "null spec",
			Spec: `{"name":"a","path":"b","registry":"local","spec":null}`,
		},
		{
			Name: "bad spec type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","spec":[]}`,
		},
		// write_mode, migrate_mode & pk_mode are tested separately
	})
}
