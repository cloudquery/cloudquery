// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func rules() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_rules",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn#Rule`,
		Resolver:    fetchCDNRules,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cdn_rule_set_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "order",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Order"),
			},
			{
				Name:     "conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Conditions"),
			},
			{
				Name:     "actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Actions"),
			},
			{
				Name:     "match_processing_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MatchProcessingBehavior"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "deployment_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentStatus"),
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
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
		},
	}
}

func fetchCDNRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CDN.Rules

	profile := parent.Parent.Item.(cdn.Profile)
	resource, err := client.ParseResourceID(*profile.ID)
	if err != nil {
		return err
	}
	ruleSet := parent.Item.(cdn.RuleSet)
	response, err := svc.ListByRuleSet(ctx, resource.ResourceGroup, *profile.Name, *ruleSet.Name)

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
