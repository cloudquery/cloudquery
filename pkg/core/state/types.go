package state

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type tableList map[string][]string

func (a tableList) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *tableList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type stringMap map[string]string

func (a stringMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *stringMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
