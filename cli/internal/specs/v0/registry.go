package specs

import (
	"encoding/json"
	"fmt"
)

type Registry string

const (
	RegistryGithub     = Registry("github")
	RegistryLocal      = Registry("local")
	RegistryGrpc       = Registry("grpc")
	RegistryDocker     = Registry("docker")
	RegistryCloudQuery = Registry("cloudquery")
)

func (r *Registry) UnmarshalJSON(data []byte) (err error) {
	var registry string
	if err := json.Unmarshal(data, &registry); err != nil {
		return err
	}
	*r = Registry(registry)
	return r.IsValid()
}

func (r Registry) IsValid() error {
	switch r {
	case RegistryGithub, RegistryLocal, RegistryGrpc, RegistryDocker, RegistryCloudQuery:
		return nil
	default:
		return fmt.Errorf("unknown registry %q", r)
	}
}
