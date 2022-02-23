package registry

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

const (
	DefaultOrganization = "cloudquery"
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
		return DefaultOrganization, name, nil
	}
	return "", "", fmt.Errorf("invalid provider name %q", name)
}

// ProviderRepoName returns a repository name for a given provider name.
func ProviderRepoName(name string) string {
	return fmt.Sprintf("cq-provider-%s", name)
}

// ParseProviderNameWithVersion parses the given string as "repository/provider@version" or "provider@version" (or @version omitted, defaulting to "latest")
func ParseProviderNameWithVersion(nameWithVersion string) (string, string, string, error) {
	versionParts := strings.Split(nameWithVersion, "@")

	if l := len(versionParts); l == 1 || (l == 2 && versionParts[1] == "latest") {
		org, name, err := ParseProviderName(versionParts[0])
		return org, name, "latest", err
	} else if l != 2 {
		return "", "", "", fmt.Errorf("invalid provider name@version %q", nameWithVersion)
	}

	org, name, err := ParseProviderName(versionParts[0])
	if err != nil {
		return "", "", "", err
	}

	ver, err := semver.NewVersion(versionParts[1])
	if err != nil {
		return "", "", "", fmt.Errorf("invalid version %q: %w", versionParts[1], err)
	}

	return org, name, "v" + ver.String(), err
}
