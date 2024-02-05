package spec

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/stretchr/testify/require"
)

func TestLogLevel_JSONSchema(t *testing.T) {
	sc := LogLevel(0).JSONSchema()
	data, err := json.Marshal(sc)
	require.NoError(t, err)

	cases := make([]jsonschema.TestCase, len(sc.Enum))
	for i, e := range sc.Enum {
		val := strconv.Quote(e.(string))
		cases[i] = jsonschema.TestCase{Name: val, Spec: val}
	}

	jsonschema.TestJSONSchema(t, string(data),
		append(cases,
			jsonschema.TestCase{Name: "null", Spec: `null`, Err: true},
			jsonschema.TestCase{Name: "integer", Spec: `123`, Err: true},
			jsonschema.TestCase{Name: "bad value", Spec: `"some_extra_value"`, Err: true},
		),
	)
}

func TestLogLevel_MarshalJSON(t *testing.T) {
	type spec struct {
		LogLevel LogLevel `json:"log_level"`
	}
	for l := tracelog.LogLevelNone; l <= tracelog.LogLevelTrace; l++ {
		level := LogLevel(l)
		t.Run(level.String(), func(t *testing.T) {
			data, err := json.Marshal(spec{LogLevel: level})
			require.NoError(t, err)
			require.Exactly(t, `{"log_level":"`+level.String()+`"}`, string(data))
		})
	}
}

func TestLogLevel_UnmarshalJSON(t *testing.T) {
	type spec struct {
		LogLevel LogLevel `json:"log_level"`
	}
	for l := tracelog.LogLevelNone; l <= tracelog.LogLevelTrace; l++ {
		level := LogLevel(l)
		t.Run(level.String(), func(t *testing.T) {
			var s spec
			err := json.Unmarshal([]byte(`{"log_level":"`+level.String()+`"}`), &s)
			require.NoError(t, err)
			require.Exactly(t, level, s.LogLevel)
		})
	}
}
