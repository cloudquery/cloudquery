package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Registry int

const (
	RegistryUnset Registry = iota
	RegistryGithub
	RegistryLocal
	RegistryGrpc
	RegistryDocker
	RegistryCloudQuery
)

func (r Registry) String() string {
	return [...]string{"", "github", "local", "grpc", "docker", "cloudquery"}[r]
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

func (r Registry) NeedVersion() bool {
	return r == RegistryGithub || r == RegistryCloudQuery
}

func RegistryFromString(s string) (Registry, error) {
	switch s {
	case "":
		return RegistryUnset, nil
	case "github":
		return RegistryGithub, nil
	case "local":
		return RegistryLocal, nil
	case "grpc":
		return RegistryGrpc, nil
	case "docker":
		return RegistryDocker, nil
	case "cloudquery":
		return RegistryCloudQuery, nil
	default:
		return RegistryGithub, fmt.Errorf("unknown registry %q", s)
	}
}
