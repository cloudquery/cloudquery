// pluginmanager take care of lifecycle of plugins.
// Including: downloading, upgrading, spawning, closing
// Currently we use github releases as our plugin store. We might change in the future
// to our own hosted one.
package plugin

import (
	"os/exec"
	"strings"

	"context"

	"github.com/cloudquery/cloudquery/internal/destinations"
	"github.com/cloudquery/cq-provider-sdk/clients"
	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/spec"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func hasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// get unnormalized plugin name and returns normalized/
// aws -> cloudquery, aws, latest
// cloudquery/aws -> cloudquery, aws, latest
// aws@v1.0.0 -> cloudquery, aws, v1.0.0
func parsePluginName(name string) (string, string, string, error) {
	if len(name) == 0 {
		return "", "", "", errors.New("plugin name is empty")
	}

	organization := "cloudquery"
	pluginName := name
	version := "latest"

	pluginPart := strings.Split(name, "@")
	if len(pluginPart) > 2 {
		return "", "", "", errors.Errorf("invalid plugin name: %s. only one @ is allowed", name)
	}
	if len(pluginPart) == 2 {
		version = pluginPart[1]
	}

	pluginPart = strings.Split(pluginPart[0], "/")
	if len(pluginPart) > 2 {
		return "", "", "", errors.Errorf("invalid plugin name: %s. only one / is allowed", name)
	}

	if len(pluginPart) == 2 {
		organization = pluginPart[0]
	}

	return organization, pluginName, version, nil
}

type PluginManager struct {
	plugins []exec.Cmd
}

func NewPluginManager() *PluginManager {
	return &PluginManager{}
}

func (p *PluginManager) Download(ctx context.Context, name string) error {
	// org, name, version, err := parsePluginName(name)
	// if err != nil {
	// 	return errors.Wrap("failed to parse plugin name", err)
	// }

	return nil
}

func (p *PluginManager) GetDestinationClient(ctx context.Context, spec spec.DestinationSpec, opts plugins.DestinationPluginOptions) (*clients.DestinationClient, error) {
	switch spec.Name {
	case "postgresql":
		return clients.NewLocalDestinationClient(&destinations.PostgreSqlPlugin{}), nil
	default:
		return nil, errors.Errorf("unknown destination type: %s", spec.Name)
	}
}

func (p *PluginManager) GetSourcePluginClient(ctx context.Context, spec spec.SourceSpec) (*clients.SourceClient, error) {
	if spec.Registry == "grpc" {
		cc, err := grpc.Dial(spec.Path, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		return clients.NewSourceClient(cc), nil
	}
	// org, name, version, err := parsePluginName(name)
	return nil, nil
}

func (p *PluginManager) ClosePlugin(ctx context.Context, name string) error {
	return nil
}

func (p *PluginManager) CloseAll(ctx context.Context) error {
	return nil
}
