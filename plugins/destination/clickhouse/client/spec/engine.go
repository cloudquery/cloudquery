package spec

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	MergeTree = "MergeTree"
)

// Engine allows to specify a custom table engine to be used.
type Engine struct {
	// Name of the table engine.
	// Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.
	Name string `json:"name,omitempty" jsonschema:"pattern=^.*MergeTree$,default=MergeTree"`

	// Engine parameters.
	// Currently, no restrictions are imposed on the parameter types.
	Parameters []any `json:"parameters,omitempty"`
}

func (e *Engine) String() string {
	return e.Name + "(" + strings.Join(e.params(), ", ") + ")"
}

func (e *Engine) params() []string {
	res := make([]string, len(e.Parameters))

	for i, p := range e.Parameters {
		switch t := p.(type) {
		case string:
			res[i] = "'" + t + "'"
		case int:
			res[i] = strconv.Itoa(t)
		case int32:
			res[i] = strconv.FormatInt(int64(t), 10)
		case int64:
			res[i] = strconv.FormatInt(t, 10)
		case float32:
			res[i] = strconv.FormatFloat(float64(t), 'f', -1, 32)
		case float64:
			res[i] = strconv.FormatFloat(t, 'f', -1, 64)
		case json.Number:
			res[i] = t.String()
		case bool:
			if t {
				res[i] = "true"
			} else {
				res[i] = "false"
			}
		default:
			res[i] = fmt.Sprint(p)
		}
	}

	return res
}

func (e *Engine) Validate() error {
	if !strings.HasSuffix(e.Name, MergeTree) {
		return fmt.Errorf("only *MergeTree table engine family is supported at the moment")
	}

	for _, p := range e.Parameters {
		switch t := p.(type) {
		case string, int, int32, int64, float32, float64, json.Number, bool: // supported types
		default:
			return fmt.Errorf("unsupported engine option type %T", t)
		}
	}

	return nil
}

func DefaultEngine() *Engine {
	return &Engine{Name: MergeTree}
}
