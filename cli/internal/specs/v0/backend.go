package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Backend int

const (
	BackendNone Backend = iota
	BackendLocal
)

var AllBackendNames = [...]string{
	BackendNone:  "none",
	BackendLocal: "local",
}

type Backends []Backend

func (s Backends) String() string {
	var buffer bytes.Buffer
	for i, backend := range s {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(backend.String())
	}
	return buffer.String()
}

func (s Backend) String() string {
	return AllBackendNames[s]
}
func (s Backend) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *Backend) UnmarshalJSON(data []byte) (err error) {
	var backend string
	if err := json.Unmarshal(data, &backend); err != nil {
		return err
	}
	if *s, err = BackendFromString(backend); err != nil {
		return err
	}
	return nil
}

func (Backend) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllBackendNames))
	for i, name := range AllBackendNames {
		sc.Enum[i] = name
	}
}

func BackendFromString(s string) (Backend, error) {
	for i, str := range AllBackendNames {
		if s == str {
			return Backend(i), nil
		}
	}
	return BackendNone, fmt.Errorf("unknown backend %s", s)
}
