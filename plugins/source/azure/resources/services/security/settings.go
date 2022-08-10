package security

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecuritySettings() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_settings",
		Description:  "Setting the kind of the security setting",
		Resolver:     fetchSecuritySettings,
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
				Name:        "kind",
				Description: "Possible values include: 'KindSetting', 'KindDataExportSettings'",
				Type:        schema.TypeString,
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
			{
				Name:        "enabled",
				Description: "Export setting enabled flag",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DataExportSettingProperties.Enabled"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSecuritySettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.Settings
	response, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		for _, item := range response.Values() {
			if v, ok := item.AsSetting(); ok {
				res <- v
			} else if v, ok := item.AsDataExportSettings(); ok {
				res <- v
			} else if v, ok := item.AsAlertSyncSettings(); ok { // nolint:revive
				res <- v
			} else {
				return diag.WrapError(fmt.Errorf("unexpected BasicSetting: %#v", item))
			}
		}
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
