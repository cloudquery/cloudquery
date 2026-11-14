package auth

import (
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/stretchr/testify/require"
)

func TestTokenNeeded(t *testing.T) {
	tests := []struct {
		name     string
		registry specs.Registry
		path     string
		want     bool
	}{
		{name: "cloudquery registry", registry: specs.RegistryCloudQuery, path: "cloudquery/aws", want: true},
		{name: "cloudquery docker registry", registry: specs.RegistryDocker, path: "docker.cloudquery.io/cloudquery/source-aws:v1.0.0", want: true},
		{name: "look-alike host suffix", registry: specs.RegistryDocker, path: "docker.cloudquery.io.evil.com/cloudquery/source-aws:v1.0.0", want: false},
		{name: "look-alike host without separator", registry: specs.RegistryDocker, path: "docker.cloudquery.iox/cloudquery/source-aws", want: false},
		{name: "other docker registry", registry: specs.RegistryDocker, path: "ghcr.io/cloudquery/source-aws:v1.0.0", want: false},
		{name: "github registry", registry: specs.RegistryGitHub, path: "docker.cloudquery.io/cloudquery/source-aws", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tokenNeeded(tt.registry, tt.path))
		})
	}
}
