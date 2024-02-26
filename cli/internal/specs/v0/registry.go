package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Registry int

const (
	RegistryUnset  Registry = iota
	RegistryGitHub          // deprecated
	RegistryLocal
	RegistryGRPC
	RegistryDocker
	RegistryCloudQuery
)

var (
	AllRegistries = [...]string{
		RegistryUnset:      "",
		RegistryGitHub:     "github",
		RegistryLocal:      "local",
		RegistryGRPC:       "grpc",
		RegistryDocker:     "docker",
		RegistryCloudQuery: "cloudquery",
	}
)

func (r Registry) String() string {
	return AllRegistries[r]
}

func (r Registry) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (r *Registry) UnmarshalJSON(data []byte) (err error) {
	var registry string
	if err := json.Unmarshal(data, &registry); err != nil {
		return err
	}
	if *r, err = RegistryFromString(registry); err != nil {
		return err
	}
	return nil
}

func (Registry) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllRegistries))
	for i, k := range AllRegistries {
		sc.Enum[i] = k
	}
}

// NeedVersion has to be in sync with Metadata.JSONSchemaExtend
func (r Registry) NeedVersion() bool {
	return r == RegistryGitHub || r == RegistryCloudQuery
}

func RegistryFromString(s string) (Registry, error) {
	for r, str := range AllRegistries {
		if s == str {
			return Registry(r), nil
		}
	}
	return RegistryUnset, fmt.Errorf("unknown registry %q", s)
}
