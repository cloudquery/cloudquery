// Code generated by codegen; DO NOT EDIT.

package postgresql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	postgresql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Postgresql

	server := parent.Item.(*postgresql.Server)
	id, err := arm.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}

	pager := svc.ConfigurationsClient.NewListByServerPager(id.ResourceGroupName, *server.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
