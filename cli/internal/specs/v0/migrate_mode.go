package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MigrateMode int

const (
	MigrateModeSafe MigrateMode = iota
	MigrateModeForced
)

var (
	migrateModeStrings = []string{"safe", "forced"}
)

func (m MigrateMode) String() string {
	return migrateModeStrings[m]
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

func MigrateModeFromString(s string) (MigrateMode, error) {
	switch s {
	case "safe":
		return MigrateModeSafe, nil
	case "forced":
		return MigrateModeForced, nil
	}
	return 0, fmt.Errorf("invalid migrate mode: %s", s)
}
