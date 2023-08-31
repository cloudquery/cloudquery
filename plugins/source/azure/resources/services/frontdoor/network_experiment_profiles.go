package frontdoor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func NetworkExperimentProfiles() *schema.Table {
	return &schema.Table{
		Name:                 "azure_frontdoor_network_experiment_profiles",
		Resolver:             fetchNetworkExperimentProfiles,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/internetanalyzer/network-experiment-profiles/list?tabs=HTTP#profile",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_frontdoor_network_experiment_profiles", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armfrontdoor.Profile{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchNetworkExperimentProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armfrontdoor.NewNetworkExperimentProfilesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
