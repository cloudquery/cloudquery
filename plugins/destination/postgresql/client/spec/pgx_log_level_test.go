package spec

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
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
