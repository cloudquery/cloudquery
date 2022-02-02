package registry

import (
	"fmt"
	"strings"
)

const (
	defaultOrganization = "cloudquery"
)

// ParseProviderName parses a name of a provider which can be just a name or a name + organization
// For example aws <-> cloudquery/aws will download the cq-provider-aws in cloudquery organization
// The organization defaults to cloudquery, if you want to download from a different repo set the name <your_org_name>/<provider_name>
func ParseProviderName(name string) (string, string, error) {
	names := strings.Split(name, "/")
	if len(names) == 2 {
		return strings.ToLower(names[0]), names[1], nil
	}
	if len(names) == 1 {
		return defaultOrganization, name, nil
	}
	return "", "", fmt.Errorf("invalid provider name %s", name)
}

// ProviderRepoName returns a repository name for a given provider name.
func ProviderRepoName(name string) string {
	return fmt.Sprintf("cq-provider-%s", name)
}
