package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type WriteMode int

const (
	WriteModeOverwriteDeleteStale WriteMode = iota
	WriteModeOverwrite
	WriteModeAppend
)

var (
	AllWriteModes = [...]string{
		WriteModeOverwriteDeleteStale: "overwrite-delete-stale",
		WriteModeOverwrite:            "overwrite",
		WriteModeAppend:               "append",
	}
)

func (m WriteMode) String() string {
	return AllWriteModes[m]
}

func (m WriteMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m *WriteMode) UnmarshalJSON(data []byte) (err error) {
	var writeMode string
	if err := json.Unmarshal(data, &writeMode); err != nil {
		return err
	}
	if *m, err = WriteModeFromString(writeMode); err != nil {
		return err
	}
	return nil
}

func (WriteMode) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllWriteModes))
	for i, k := range AllWriteModes {
		sc.Enum[i] = k
	}
}

func WriteModeFromString(s string) (WriteMode, error) {
	for m, str := range AllWriteModes {
		if s == str {
			return WriteMode(m), nil
		}
	}
	return WriteModeOverwriteDeleteStale, fmt.Errorf("invalid write mode: %s", s)
}
