package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecurityContacts() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_contacts",
		Description:  "Contact contact details for security issues",
		Resolver:     fetchSecurityContacts,
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
				Name:        "email",
				Description: "The email of this security contact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactProperties.Email"),
			},
			{
				Name:        "phone",
				Description: "The phone number of this security contact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactProperties.Phone"),
			},
			{
				Name:        "alert_notifications",
				Description: "Whether to send security alerts notifications to the security contact Possible values include: 'On', 'Off'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactProperties.AlertNotifications"),
			},
			{
				Name:        "alerts_to_admins",
				Description: "Whether to send security alerts notifications to subscription admins Possible values include: 'AlertsToAdminsOn', 'AlertsToAdminsOff'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ContactProperties.AlertsToAdmins"),
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
func fetchSecurityContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
