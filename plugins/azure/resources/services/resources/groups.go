package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResourcesGroups() *schema.Table {
	return &schema.Table{
		Name:         "azure_resources_groups",
		Description:  "Azure resource group",
		Resolver:     fetchResourcesGroups,
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
				Name:        "id",
				Description: "The ID of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource group",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource group",
				Type:        schema.TypeString,
			},
			{
				Name:        "properties_provisioning_state",
				Description: "The provisioning state",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "location",
				Description: "The location of the resource group It cannot be changed after the resource group has been created It must be one of the supported Azure locations",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by",
				Description: "The ID of the resource that manages this resource group",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags attached to the resource group",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchResourcesGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Resources.Groups
	response, err := svc.List(ctx, "", nil)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
