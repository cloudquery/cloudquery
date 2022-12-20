package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafv2ManagedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	service := c.Services().Wafv2

	config := wafv2.ListAvailableManagedRuleGroupsInput{Scope: c.WAFScope}
	for {
		output, err := service.ListAvailableManagedRuleGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ManagedRuleGroups

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveManageRuleGroupProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	managedRuleGroupSum := resource.Item.(types.ManagedRuleGroupSummary)

	c := meta.(*client.Client)
	service := c.Services().Wafv2

	// Resolve managed rule group via describe managed rule group
	output, err := service.DescribeManagedRuleGroup(ctx, &wafv2.DescribeManagedRuleGroupInput{
		Name:       managedRuleGroupSum.Name,
		VendorName: managedRuleGroupSum.VendorName,
		Scope:      c.WAFScope,
	}, func(options *wafv2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, output)
}
