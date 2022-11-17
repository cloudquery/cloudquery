package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func PostgreSQL() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armpostgresql.Server),
			Resolver: postgresql.ServersClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armpostgresql.Configuration),
					Resolver: postgresql.ConfigurationsClient.NewListByServerPager,
				},
				{
					Struct:   new(armpostgresql.FirewallRule),
					Resolver: postgresql.FirewallRulesClient.NewListByServerPager,
				},
			},
		},
	}
}
