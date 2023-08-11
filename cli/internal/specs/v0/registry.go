package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Registry int

const (
	RegistryGithub Registry = iota
	RegistryLocal
	RegistryGo
	RegistryGrpc
	RegistryDocker
)

func (r Registry) String() string {
	return [...]string{"github", "local", "go", "grpc", "docker"}[r]
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

func RegistryFromString(s string) (Registry, error) {
	switch s {
	case "github":
		return RegistryGithub, nil
	case "local":
		return RegistryLocal, nil
	case "go":
		return RegistryGo, nil
	case "grpc":
		return RegistryGrpc, nil
	case "docker":
		return RegistryDocker, nil
	default:
		return RegistryGithub, fmt.Errorf("unknown registry %s", s)
	}
}
