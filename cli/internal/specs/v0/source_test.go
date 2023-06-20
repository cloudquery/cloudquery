package specs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var sourceUnmarshalSpecTestCases = []struct {
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

func TestSourceUnmarshalSpec(t *testing.T) {
	for _, tc := range sourceUnmarshalSpecTestCases {
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

var sourceUnmarshalSpecValidateTestCases = []struct {
	name   string
	spec   string
	err    string
	source *Source
}{
	{
		"required_name",
		`kind: source
spec:`,
		"name is required",
		nil,
	},
	{
		"required_path",
		`kind: source
spec:
  name: test
`,
		"path is required",
		nil,
	},
	{
		"required_version",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  tables: ["test"]
`,
		"version is required",
		nil,
	},
	{
		"required_version_format",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: 1.1.0
  tables: ["test"]
`,
		"version must start with v",
		nil,
	},
	{
		"tables_required",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
`,
		"tables configuration is required. Hint: set the tables you want to sync by adding `tables: [...]` or use `cloudquery tables` to list available tables",
		nil,
	},
	{
		"destination_required",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
  tables: ["test"]
`,
		"at least one destination is required",
		nil,
	},
	{
		"valid_scheduler",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
  destinations: ["test"]
  scheduler: round-robin
  tables: ["test"]
`,
		"",
		&Source{
			Name:         "test",
			Registry:     RegistryGithub,
			Path:         "cloudquery/test",
			Concurrency:  defaultConcurrency,
			Version:      "v1.1.0",
			Destinations: []string{"test"},
			Scheduler:    SchedulerRoundRobin,
			Tables:       []string{"test"},
		},
	},
	{
		"success",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
  destinations: ["test"]
  tables: ["test"]
`,
		"",
		&Source{
			Name:         "test",
			Registry:     RegistryGithub,
			Path:         "cloudquery/test",
			Concurrency:  defaultConcurrency,
			Version:      "v1.1.0",
			Destinations: []string{"test"},
			Scheduler:    SchedulerDFS,
			Tables:       []string{"test"},
		},
	},
}

func TestSourceUnmarshalSpecValidate(t *testing.T) {
	for _, tc := range sourceUnmarshalSpecValidateTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			var spec Spec
			err = SpecUnmarshalYamlStrict([]byte(tc.spec), &spec)
			if err != nil {
				t.Fatal(err)
			}
			source := spec.Spec.(*Source)
			source.SetDefaults()
			err = source.Validate()
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("expected:%s got:%s", tc.err, err.Error())
				}
				return
			}

			if cmp.Diff(source, tc.source) != "" {
				t.Fatalf("expected:%v got:%v", tc.source, source)
			}
		})
	}
}

func TestSpec_VersionString(t *testing.T) {
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
			s := Source{
				Name:     tt.fields.Name,
				Version:  tt.fields.Version,
				Path:     tt.fields.Path,
				Registry: tt.fields.Registry,
			}
			if got := s.VersionString(); got != tt.want {
				t.Errorf("Source.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
