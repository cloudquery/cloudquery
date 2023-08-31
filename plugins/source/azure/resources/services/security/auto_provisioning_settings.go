package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AutoProvisioningSettings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_auto_provisioning_settings",
		Resolver:             fetchAutoProvisioningSettings,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/auto-provisioning-settings/list?tabs=HTTP#autoprovisioningsetting",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_auto_provisioning_settings", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.AutoProvisioningSetting{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAutoProvisioningSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewAutoProvisioningSettingsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
