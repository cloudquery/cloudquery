package trafficmanager

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Profiles() *schema.Table {
	return &schema.Table{
		Name:                 "azure_trafficmanager_profiles",
		Resolver:             fetchProfiles,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/trafficmanager/profiles/list-by-subscription?tabs=HTTP#profile",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_trafficmanager_profiles", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armtrafficmanager.Profile{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armtrafficmanager.NewProfilesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
