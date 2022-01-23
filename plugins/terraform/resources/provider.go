package resources

import (
	"embed"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-terraform/client"
)

var (
	//go:embed migrations/*/*.sql
	migrationFiles embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "terraform",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"tf.data": TFData(),
		},
		Migrations: migrationFiles,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
