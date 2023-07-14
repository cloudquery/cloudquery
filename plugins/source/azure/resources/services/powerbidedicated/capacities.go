package powerbidedicated

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Capacities() *schema.Table {
	return &schema.Table{
		Name:                 "azure_powerbidedicated_capacities",
		Resolver:             fetchCapacities,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/power-bi-embedded/capacities/list?tabs=HTTP#dedicatedcapacity",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_powerbidedicated_capacities", client.Namespacemicrosoft_powerbidedicated),
		Transform:            transformers.TransformWithStruct(&armpowerbidedicated.DedicatedCapacity{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchCapacities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpowerbidedicated.NewCapacitiesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
