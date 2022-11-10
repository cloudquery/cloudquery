// Auto generated code - DO NOT EDIT.

package security

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Contacts() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_contacts",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security#Contact`,
		Resolver:    fetchSecurityContacts,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "phone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Phone"),
			},
			{
				Name:     "alert_notifications",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlertNotifications"),
			},
			{
				Name:     "alerts_to_admins",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlertsToAdmins"),
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
		},
	}
}

func fetchSecurityContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.Contacts

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
