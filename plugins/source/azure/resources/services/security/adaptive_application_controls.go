package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AdaptiveApplicationControls() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_adaptive_application_controls",
		Resolver:             fetchAdapterApplicationControls,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/adaptive-application-controls/list?view=rest-defenderforcloud-2020-01-01&tabs=HTTP",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_adaptive_application_controls", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.AdaptiveApplicationControlGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAdapterApplicationControls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewAdaptiveApplicationControlsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	r, err := svc.List(ctx, nil)
	if err != nil {
		return err
	}
	res <- r.Value
	return nil
}
