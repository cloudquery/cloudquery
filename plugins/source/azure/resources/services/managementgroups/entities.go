package managementgroups

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Entities() *schema.Table {
	return &schema.Table{
		Name:                 "azure_managementgroups_entities",
		Resolver:             fetchEntities,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/managementgroups/entities/list?tabs=HTTP",
		Transform:            transformers.TransformWithStruct(&armmanagementgroups.EntityInfo{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchEntities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmanagementgroups.NewEntitiesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
