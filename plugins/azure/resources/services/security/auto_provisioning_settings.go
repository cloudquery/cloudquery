package security

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecurityAutoProvisioningSettings() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_auto_provisioning_settings",
		Description:  "AutoProvisioningSetting auto provisioning setting",
		Resolver:     fetchSecurityAutoProvisioningSettings,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "auto_provision",
				Description: "Describes what kind of security agent provisioning action to take Possible values include: 'AutoProvisionOn', 'AutoProvisionOff'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoProvisioningSettingProperties.AutoProvision"),
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "Resource type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Type"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSecurityAutoProvisioningSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.AutoProvisioningSettings
	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}

	}
	return nil
}
