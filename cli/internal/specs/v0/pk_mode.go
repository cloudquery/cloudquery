package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type PKMode int

const (
	PKModeDefaultKeys PKMode = iota
	PKModeCQID
)

var (
	AllPKModes = [...]string{
		PKModeDefaultKeys: "default",
		PKModeCQID:        "cq-id-only",
	}
)

func (m PKMode) String() string {
	return AllPKModes[m]
}

func (m PKMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m *PKMode) UnmarshalJSON(data []byte) (err error) {
	var pkMode string
	if err := json.Unmarshal(data, &pkMode); err != nil {
		return err
	}
	if *m, err = PKModeFromString(pkMode); err != nil {
		return err
	}
	return nil
}

func (PKMode) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllPKModes))
	for i, k := range AllPKModes {
		sc.Enum[i] = k
	}
}

func PKModeFromString(s string) (PKMode, error) {
	for m, str := range AllPKModes {
		if s == str {
			return PKMode(m), nil
		}
	}
	return PKModeDefaultKeys, fmt.Errorf("invalid pk mode: %s", s)
}
