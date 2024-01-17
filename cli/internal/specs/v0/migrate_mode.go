package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type MigrateMode int

const (
	MigrateModeSafe MigrateMode = iota
	MigrateModeForced
)

var (
	AllMigrateModes = [...]string{
		MigrateModeSafe:   "safe",
		MigrateModeForced: "forced",
	}
)

func (m MigrateMode) String() string {
	return AllMigrateModes[m]
}

func (m MigrateMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m *MigrateMode) UnmarshalJSON(data []byte) (err error) {
	var migrateMode string
	if err := json.Unmarshal(data, &migrateMode); err != nil {
		return err
	}
	if *m, err = MigrateModeFromString(migrateMode); err != nil {
		return err
	}
	return nil
}

func (MigrateMode) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllMigrateModes))
	for i, k := range AllMigrateModes {
		sc.Enum[i] = k
	}
}

func MigrateModeFromString(s string) (MigrateMode, error) {
	for m, str := range AllMigrateModes {
		if s == str {
			return MigrateMode(m), nil
		}
	}
	return MigrateModeSafe, fmt.Errorf("invalid migrate mode: %s", s)
}
