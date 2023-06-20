package specs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
		"failed to decode spec: json: cannot unmarshal number into Go struct field Source.name of type string",
		&Source{
			Name:   "test",
			Tables: []string{"*"},
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
			Name:   "test",
			Tables: []string{"*"},
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
			if cmp.Diff(source, tc.source) != "" {
				t.Fatalf("expected:%v got:%v", tc.source, source)
			}
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
			Name:           "test",
			Registry:       RegistryGrpc,
			Path:           "localhost:9999",
			BatchSize:      10000,
			BatchSizeBytes: 10000000,
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
			Name:           "test",
			Registry:       RegistryLocal,
			Path:           "/home/user/some_executable",
			BatchSize:      10000,
			BatchSizeBytes: 10000000,
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
			Name:           "test",
			Registry:       RegistryGithub,
			Path:           "cloudquery/test",
			Version:        "v1.1.0",
			BatchSize:      10000,
			BatchSizeBytes: 10000000,
		},
	},
}

func TestDestinationUnmarshalSpecValidate(t *testing.T) {
	for _, tc := range destinationUnmarshalSpecValidateTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			var spec Spec
			err = SpecUnmarshalYamlStrict([]byte(tc.spec), &spec)
			if err != nil {
				t.Fatal(err)
			}
			destination := spec.Spec.(*Destination)
			destination.SetDefaults(10000, 10000000)
			err = destination.Validate()
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("expected:\n%s\ngot:\n%s", tc.err, err.Error())
				}
				return
			}

			if cmp.Diff(destination, tc.destination) != "" {
				t.Fatalf("expected:\n%v\ngot:\n%v\n", tc.destination, destination)
			}
		})
	}
}

func TestDestination_VersionString(t *testing.T) {
	type fields struct {
		Name     string
		Version  string
		Path     string
		Registry Registry
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should use short version without name part in path when those are the same",
			fields: fields{
				Name:     "aws",
				Version:  "v10.0.0",
				Path:     "cloudquery/aws",
				Registry: RegistryGithub,
			},
			want: "aws (v10.0.0)",
		},
		{
			name: "should use long version with path when name doesn't match path",
			fields: fields{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "cloudquery/aws",
				Registry: RegistryGithub,
			},
			want: "my-aws-spec (aws@v10.0.0)",
		},
		{
			name: "should handle non GitHub registry",
			fields: fields{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "localhost:7777",
				Registry: RegistryGrpc,
			},
			want: "my-aws-spec (grpc@localhost:7777)",
		},
		{
			name: "should handle malformed path",
			fields: fields{
				Name:     "my-aws-spec",
				Version:  "v10.0.0",
				Path:     "aws",
				Registry: RegistryGithub,
			},
			want: "my-aws-spec (aws@v10.0.0)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Destination{
				Name:     tt.fields.Name,
				Version:  tt.fields.Version,
				Path:     tt.fields.Path,
				Registry: tt.fields.Registry,
			}
			if got := d.VersionString(); got != tt.want {
				t.Errorf("Destination.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
