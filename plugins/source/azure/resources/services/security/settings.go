// Auto generated code - DO NOT EDIT.

package security

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func Settings() *schema.Table {
	return &schema.Table{
		Name:      "azure_security_settings",
		Resolver:  fetchSecuritySettings,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: resolveEnabled,
			},
		},
	}
}

func resolveEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(security.BasicSetting)
	if v, ok := item.AsDataExportSettings(); ok {
		return errors.WithStack(resource.Set(c.Name, v.Enabled))
	}
	if v, ok := item.AsAlertSyncSettings(); ok {
		return errors.WithStack(resource.Set(c.Name, v.Enabled))
	}
	return errors.WithStack(resource.Set(c.Name, true))
}

func fetchSecuritySettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.Settings

	response, err := svc.List(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	for response.NotDone() {
		for _, item := range response.Values() {
			if v, ok := item.AsSetting(); ok {
				res <- v
			} else if v, ok := item.AsDataExportSettings(); ok {
				res <- v
			} else if v, ok := item.AsAlertSyncSettings(); ok {
				res <- v
			} else {
				return errors.WithStack(fmt.Errorf("unexpected BasicSetting: %#v", item))
			}
		}
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
