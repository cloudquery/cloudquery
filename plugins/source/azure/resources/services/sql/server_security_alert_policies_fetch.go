// Code generated by codegen; DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	sql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchServerSecurityAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sql

	server := parent.Item.(*sql.Server)
	id, err := arm.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}

	pager := svc.ServerSecurityAlertPoliciesClient.NewListByServerPager(id.ResourceGroupName, *server.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
