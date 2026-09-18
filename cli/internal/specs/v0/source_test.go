package specs

import (
	"testing"
	"time"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
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
			require.Equal(t, tc.source, source)
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
			Metadata: Metadata{
				Name:             "test",
				Registry:         RegistryCloudQuery,
				Path:             "cloudquery/test",
				Version:          "v1.1.0",
				registryInferred: true,
			},
			Destinations:        []string{"test"},
			Tables:              []string{"test"},
			Spec:                map[string]any{},
			SkipDependentTables: &boolTrue,
			SyncMode:            SyncModeFull,
		},
	},
	{
		"success github",
		`kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
  registry: github
  destinations: ["test"]
  tables: ["test"]
`,
		"",
		&Source{
			Metadata: Metadata{
				Name:     "test",
				Registry: RegistryGitHub,
				Path:     "cloudquery/test",
				Version:  "v1.1.0",
			},
			Destinations:        []string{"test"},
			Tables:              []string{"test"},
			Spec:                map[string]any{},
			SkipDependentTables: &boolTrue,
			SyncMode:            SyncModeFull,
		},
	}}

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
			err = source.Validate(false)
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("expected:%s got:%s", tc.err, err.Error())
				}
				return
			}

			require.Equal(t, tc.source, source)
		})
	}
}

func TestSpec_VersionString(t *testing.T) {
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
			name: "should handle non CloudQuery registry",
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
			s := Source{Metadata: tt.meta}
			if got := s.VersionString(); got != tt.want {
				t.Errorf("Source.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackendOptions_JSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(BackendOptions))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{}`,
		},
		{
			Name: "missing table_name",
			Err:  true,
			Spec: `{"connection":"a"}`,
		},
		{
			Name: "empty table_name",
			Err:  true,
			Spec: `{"table_name":"","connection":"a"}`,
		},
		{
			Name: "null table_name",
			Err:  true,
			Spec: `{"table_name":null,"connection":"a"}`,
		},
		{
			Name: "bad table_name type",
			Err:  true,
			Spec: `{"table_name":123,"connection":"a"}`,
		},
		{
			Name: "missing connection",
			Err:  true,
			Spec: `{"table_name":"a"}`,
		},
		{
			Name: "empty connection",
			Err:  true,
			Spec: `{"table_name":"a","connection":""}`,
		},
		{
			Name: "null connection",
			Err:  true,
			Spec: `{"table_name":"a","connection":null}`,
		},
		{
			Name: "bad connection type",
			Err:  true,
			Spec: `{"table_name":"a","connection":123}`,
		},
		{
			Name: "proper",
			Spec: `{"table_name":"a","connection":"b"}`,
		},
	})
}

func TestSource_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(Source))
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
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"]}`,
		},
		{
			Name: "empty spec",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],"spec":{}}`,
		},
		{
			Name: "null spec",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],"spec":null}`,
		},
		{
			Name: "bad spec type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],"spec":[]}`,
		},
		{
			Name: "missing tables",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"]}`,
		},
		{
			Name: "empty tables",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":[]
}`,
		},
		{
			Name: "null tables",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":null
}`,
		},
		{
			Name: "bad tables type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":123
}`,
		},
		{
			Name: "empty tables entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":[""]
}`,
		},
		{
			Name: "null tables entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":[null]
}`,
		},
		{
			Name: "bad tables entry type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":[123]
}`,
		},
		{
			Name: "proper tables",
			Spec: `{"name":"a","path":"b","registry":"local","destinations":["a"],
"tables":["*"]
}`,
		},
		{
			Name: "missing destinations",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"]}`,
		},
		{
			Name: "empty destinations",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":[]
}`,
		},
		{
			Name: "null destinations",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":null
}`,
		},
		{
			Name: "bad destinations type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":123
}`,
		},
		{
			Name: "empty destinations entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":[""]
}`,
		},
		{
			Name: "null destinations entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":[null]
}`,
		},
		{
			Name: "bad destinations entry type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":[123]
}`,
		},
		{
			Name: "proper destinations",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],
"destinations":["a"]
}`,
		},
		{
			Name: "empty skip_tables",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":[]
}`,
		},
		{
			Name: "null skip_tables",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":null
}`,
		},
		{
			Name: "bad skip_tables type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":123
}`,
		},
		{
			Name: "empty skip_tables entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":[""]
}`,
		},
		{
			Name: "null skip_tables entry",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":[null]
}`,
		},
		{
			Name: "bad skip_tables entry type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":[123]
}`,
		},
		{
			Name: "proper skip_tables",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_tables":["a"]
}`,
		},
		{
			Name: "null skip_dependent_tables",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_dependent_tables":null
}`,
		},
		{
			Name: "bad skip_dependent_tables type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_dependent_tables":123
}`,
		},
		{
			Name: "skip_dependent_tables:true",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_dependent_tables":true
}`,
		},
		{
			Name: "skip_dependent_tables:false",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"skip_dependent_tables":false
}`,
		},
		// backend_options is tested in depth separately
		{
			Name: "null backend_options",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"backend_options":null
}`,
		},
		{
			Name: "bad backend_options type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"backend_options":123
}`,
		},
		{
			Name: "bad sync_mode",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"sync_mode":"streaming"
}`,
		},
		{
			Name: "sync_mode:full",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"sync_mode":"full"
}`,
		},
		{
			Name: "sync_mode:events",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"sync_mode":"events"
}`,
		},
		{
			Name: "null events",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":null
}`,
		},
		{
			Name: "empty events",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":{}
}`,
		},
		{
			Name: "unknown events option",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":{"followw":true}
}`,
		},
		{
			Name: "bad events.full_sync",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":{"full_sync":"sometimes"}
}`,
		},
		{
			Name: "bad events.batch_max_wait",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":{"batch_max_wait":5}
}`,
		},
		{
			Name: "events.batch_max_changes below minimum",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"events":{"batch_max_changes":0}
}`,
		},
		{
			Name: "proper events",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"sync_mode":"events",
"events":{"follow":true,"full_sync":"always","max_checkpoint_age":"20h","batch_max_changes":10,"batch_max_wait":"1s","summary_interval":"30s"}
}`,
		},
		{
			Name: "empty spec",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"spec":{}
}`,
		},
		{
			Name: "null spec",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"spec":null
}`,
		},
		{
			Name: "bad spec type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"spec":123
}`,
		},
		{
			Name: "null deterministic_cq_id",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"deterministic_cq_id":null
}`,
		},
		{
			Name: "bad deterministic_cq_id type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"deterministic_cq_id":123
}`,
		},
		{
			Name: "deterministic_cq_id:true",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"deterministic_cq_id":true
}`,
		},
		{
			Name: "deterministic_cq_id:false",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"deterministic_cq_id":false
}`,
		},
		{
			Name: "empty otel_endpoint",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint":""
}`,
		},
		{
			Name: "null otel_endpoint",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint":null
}`,
		},
		{
			Name: "bad otel_endpoint type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint":123
}`,
		},
		{
			Name: "proper otel_endpoint",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint":"a"
}`,
		},
		{
			Name: "null otel_endpoint_insecure",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint_insecure":null
}`,
		},
		{
			Name: "bad otel_endpoint_insecure type",
			Err:  true,
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint_insecure":123
}`,
		},
		{
			Name: "otel_endpoint_insecure:true",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint_insecure":true
}`,
		},
		{
			Name: "otel_endpoint_insecure:false",
			Spec: `{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a"],
"otel_endpoint_insecure":false
}`,
		},
	})
}

func TestBackendOptionsPluginName(t *testing.T) {
	tests := []struct {
		name     string
		options  *BackendOptions
		expected string
	}{
		{
			name:     "nil doesn't blow up",
			options:  nil,
			expected: "",
		},
		{
			name: "No interpolation results in empty plugin name",
			options: &BackendOptions{
				TableName:  "test_table",
				Connection: "localhost:7777",
			},
			expected: "",
		},
		{
			name: "Proper variable name results in correct plugin name",
			options: &BackendOptions{
				TableName:  "test_table",
				Connection: "@@plugins.aws.connection",
			},
			expected: "aws",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.options.PluginName()
			if actual != tt.expected {
				t.Errorf("unexpected plugin name, got: %s, want: %s", actual, tt.expected)
			}
		})
	}
}

func TestSourceSyncModeValidate(t *testing.T) {
	const header = `kind: source
spec:
  name: test
  path: cloudquery/test
  version: v1.1.0
  destinations: ["test"]
  tables: ["test"]
`
	cases := []struct {
		name string
		spec string
		err  string
	}{
		{
			name: "full is the default",
			spec: header,
		},
		{
			name: "events requires backend_options",
			spec: header + `  sync_mode: events
`,
			err: "`sync_mode: events` requires `backend_options` with a `table_name`. Hint: point `backend_options.connection` at one of your destinations",
		},
		{
			name: "events requires a backend table_name",
			spec: header + `  sync_mode: events
  backend_options:
    connection: "@@plugins.test.connection"
`,
			err: "`sync_mode: events` requires `backend_options` with a `table_name`. Hint: point `backend_options.connection` at one of your destinations",
		},
		{
			name: "events options without events sync mode",
			spec: header + `  events:
    follow: true
`,
			err: "`events` options require `sync_mode: events`",
		},
		{
			name: "events options with an explicit full sync mode",
			spec: header + `  sync_mode: full
  events:
    follow: true
`,
			err: "`events` options require `sync_mode: events`",
		},
		{
			name: "unknown sync mode",
			spec: header + `  sync_mode: streaming
`,
			err: `invalid sync_mode "streaming", must be one of [full events]`,
		},
		{
			name: "unknown full_sync policy",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    full_sync: sometimes
`,
			err: `invalid events.full_sync "sometimes", must be one of [always on_first_run never]`,
		},
		{
			name: "non-positive batch_max_changes",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    batch_max_changes: -1
`,
			err: "events.batch_max_changes must be greater than 0",
		},
		{
			name: "non-positive batch_max_wait",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    batch_max_wait: 0s
`,
			err: "events.batch_max_wait must be greater than 0",
		},
		{
			name: "non-positive summary_interval",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    summary_interval: -1m
`,
			err: "events.summary_interval must be greater than 0",
		},
		{
			name: "non-positive max_checkpoint_age",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    max_checkpoint_age: 0h
`,
			err: "events.max_checkpoint_age must be greater than 0",
		},
		{
			name: "valid events source",
			spec: header + `  sync_mode: events
  backend_options:
    table_name: cq_state
    connection: "@@plugins.test.connection"
  events:
    follow: true
    full_sync: always
    max_checkpoint_age: 20h
    batch_max_changes: 10
    batch_max_wait: 1s
    summary_interval: 30s
`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var spec Spec
			require.NoError(t, SpecUnmarshalYamlStrict([]byte(tc.spec), &spec))
			source := spec.Spec.(*Source)
			source.SetDefaults()
			err := source.Validate(false)
			if tc.err == "" {
				require.NoError(t, err)
				return
			}
			require.EqualError(t, err, tc.err)
		})
	}
}

func TestSourceEventsSetDefaults(t *testing.T) {
	t.Run("full mode leaves events unset", func(t *testing.T) {
		s := &Source{}
		s.SetDefaults()
		require.Equal(t, SyncModeFull, s.SyncMode)
		require.Nil(t, s.Events)
	})

	t.Run("events mode fills every default", func(t *testing.T) {
		s := &Source{SyncMode: SyncModeEvents}
		s.SetDefaults()
		require.NotNil(t, s.Events)
		require.False(t, s.Events.Follow)
		require.Equal(t, FullSyncOnFirstRun, s.Events.FullSync)
		require.Nil(t, s.Events.MaxCheckpointAge)
		require.Equal(t, 500, s.Events.BatchMaxChanges)
		require.Equal(t, 5*time.Second, s.Events.BatchMaxWait.Duration())
		require.Equal(t, 5*time.Minute, s.Events.SummaryInterval.Duration())
	})

	t.Run("events mode keeps what the user set", func(t *testing.T) {
		s := &Source{
			SyncMode: SyncModeEvents,
			Events: &EventsOptions{
				Follow:           true,
				FullSync:         FullSyncNever,
				MaxCheckpointAge: ptrDuration(20 * time.Hour),
				BatchMaxChanges:  7,
				BatchMaxWait:     ptrDuration(time.Second),
				SummaryInterval:  ptrDuration(time.Minute),
			},
		}
		s.SetDefaults()
		require.True(t, s.Events.Follow)
		require.Equal(t, FullSyncNever, s.Events.FullSync)
		require.Equal(t, 20*time.Hour, s.Events.MaxCheckpointAge.Duration())
		require.Equal(t, 7, s.Events.BatchMaxChanges)
		require.Equal(t, time.Second, s.Events.BatchMaxWait.Duration())
		require.Equal(t, time.Minute, s.Events.SummaryInterval.Duration())
	})
}
