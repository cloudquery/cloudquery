package resources

import (
	"embed"

	"github.com/cloudquery/cq-provider-okta/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*/*.sql
	migrationFiles embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "okta",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"users": Users(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
		ErrorClassifier: client.ErrorClassifier,
		Migrations:      migrationFiles,
	}
}
