package policy

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DataPolicyManifests() *schema.Table {
	return &schema.Table{
		Name:                 "azure_policy_data_policy_manifests",
		Resolver:             fetchDataPolicyManifests,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy@v0.6.0#DataPolicyManifest",
		Transform:            transformers.TransformWithStruct(&armpolicy.DataPolicyManifest{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchDataPolicyManifests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpolicy.NewDataPolicyManifestsClient(cl.Creds, cl.Options)
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
