package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func MariaDB() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armmariadb.Server),
			Resolver: mariadb.ServersClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armmariadb.Configuration),
					Resolver: mariadb.ConfigurationsClient.NewListByServerPager,
				},
			},
		},
	}
}
