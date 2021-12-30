package resources

import (
	"embed"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-template/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// CHANGEME: Change to your provider name
const providerName = "github"

var (
	//go:embed migrations/*.sql
	providerMigrations embed.FS
	Version            = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      providerName,
		Version:   Version,
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			// CHANGEME: place here all supported resources
			"demo_resource": DemoResource(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
