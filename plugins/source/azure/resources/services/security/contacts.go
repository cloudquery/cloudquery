// Code generated by codegen; DO NOT EDIT.

package security

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Contacts() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_contacts",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity#Contact`,
		Resolver:    fetchContacts,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "alert_notifications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.AlertNotifications"),
			},
			{
				Name:     "emails",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Emails"),
			},
			{
				Name:     "notifications_by_role",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NotificationsByRole"),
			},
			{
				Name:     "phone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Phone"),
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
