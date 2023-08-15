package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Profiles() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cdn_profiles",
		Resolver:             fetchProfiles,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cdn/profiles/list?tabs=HTTP#profile",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cdn_profiles", client.Namespacemicrosoft_cdn),
		Transform:            transformers.TransformWithStruct(&armcdn.Profile{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},

		Relations: []*schema.Table{
			endpoints(),
			rule_sets(),
			security_policies(),
		},
	}
}

func fetchProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcdn.NewProfilesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
