package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AdGroups() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_groups",
		Description:  "ADGroup active Directory group information",
		Resolver:     fetchAdGroups,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "display_name",
				Description: "The display name of the group",
				Type:        schema.TypeString,
			},
			{
				Name:        "mail_enabled",
				Description: "Whether the group is mail-enabled Must be false This is because only pure security groups can be created using the Graph API",
				Type:        schema.TypeBool,
			},
			{
				Name:        "mail_nickname",
				Description: "The mail alias for the group",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_enabled",
				Description: "Whether the group is security-enable",
				Type:        schema.TypeBool,
			},
			{
				Name:        "mail",
				Description: "The primary email address of the group",
				Type:        schema.TypeString,
			},
			{
				Name:        "additional_properties",
				Description: "Unmatched properties from the message are deserialized this collection",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "object_id",
				Description: "The object ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectID"),
			},
			{
				Name:        "deletion_timestamp_time",
				Description: "The time at which the directory object was deleted.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DeletionTimestamp.Time"),
			},
			{
				Name:        "object_type",
				Description: "Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAdGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().AD.Groups
	response, err := svc.List(ctx, "")
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
