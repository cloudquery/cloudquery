package queries

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	MergeTree = "MergeTree"
)

type Engine struct {
	Name       string `json:"name,omitempty"`
	Parameters []any  `json:"parameters,omitempty"`
}

func (e *Engine) String() string {
	if len(e.Parameters) == 0 {
		return e.Name
	}

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
			panic(fmt.Errorf("unsupported engine option type %T", t))
		}
	}

	return res
}
func (e *Engine) Validate() error {
	if !strings.HasSuffix(e.Name, MergeTree) {
		return fmt.Errorf("only *MergeTree table engine family is supported at the moment, got %q", e.Name)
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
