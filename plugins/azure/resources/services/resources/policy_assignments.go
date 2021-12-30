package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResourcesPolicyAssignments() *schema.Table {
	return &schema.Table{
		Name:         "azure_resources_policy_assignments",
		Description:  "Azure network watcher",
		Resolver:     fetchResourcesPolicyAssignments,
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
				Name:        "display_name",
				Description: "The display name of the policy assignment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssignmentProperties.DisplayName"),
			},
			{
				Name:        "policy_definition_id",
				Description: "The ID of the policy definition or policy set definition being assigned",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssignmentProperties.PolicyDefinitionID"),
			},
			{
				Name:        "scope",
				Description: "The scope for the policy assignment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssignmentProperties.Scope"),
			},
			{
				Name:        "not_scopes",
				Description: "The policy's excluded scopes",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("AssignmentProperties.NotScopes"),
			},
			{
				Name:        "parameters",
				Description: "The parameter values for the assigned policy rule",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AssignmentProperties.Parameters"),
			},
			{
				Name:        "description",
				Description: "This message will be part of response in case of policy violation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssignmentProperties.Description"),
			},
			{
				Name:        "metadata",
				Description: "The policy assignment metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveResourcesPolicyAssignmentMetadata,
			},
			{
				Name:        "enforcement_mode",
				Description: "The policy assignment enforcement mode",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AssignmentProperties.EnforcementMode"),
			},
			{
				Name:        "id",
				Description: "The ID of the policy assignment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "type",
				Description: "The type of the policy assignment",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the policy assignment",
				Type:        schema.TypeString,
			},
			{
				Name:        "sku_name",
				Description: "The name of the policy sku",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The policy sku tier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "location",
				Description: "The location of the policy assignment",
				Type:        schema.TypeString,
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal ID of the resource identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant ID of the resource identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The identity type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchResourcesPolicyAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Resources.Assignments
	response, err := svc.List(ctx, meta.(*client.Client).SubscriptionId, "", nil)
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
func resolveResourcesPolicyAssignmentMetadata(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	a, ok := resource.Item.(policy.Assignment)
	if !ok {
		return fmt.Errorf("expected policy.Assignment but got %T", resource.Item)
	}
	if a.Metadata == nil {
		return nil
	}

	out, err := json.Marshal(a.Metadata)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
