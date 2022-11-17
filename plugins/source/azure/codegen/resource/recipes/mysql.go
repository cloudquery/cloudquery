package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func MySQL() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armmysql.Server),
			Resolver: mysql.ServersClient.NewListPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armmysql.Configuration),
					Resolver: mysql.ConfigurationsClient.NewListByServerPager,
				},
			},
		},
	}
}
