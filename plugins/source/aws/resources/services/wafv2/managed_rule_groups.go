package wafv2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ManagedRuleGroups() *schema.Table {
	tableName := "aws_wafv2_managed_rule_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html`,
		Resolver:    fetchWafv2ManagedRuleGroups,
		Transform:   transformers.TransformWithStruct(&types.ManagedRuleGroupSummary{}, transformers.WithPrimaryKeys("Name", "VendorName")),
		Multiplex:   client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "scope",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveWAFScope,
				PrimaryKey: true,
			},
			{
				Name:     "properties",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveManageRuleGroupProperties,
			},
		},
	}
}

func fetchWafv2ManagedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	config := wafv2.ListAvailableManagedRuleGroupsInput{Scope: cl.WAFScope}
	for {
		output, err := service.ListAvailableManagedRuleGroups(ctx, &config, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
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

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	// Resolve managed rule group via describe managed rule group
	output, err := service.DescribeManagedRuleGroup(ctx, &wafv2.DescribeManagedRuleGroupInput{
		Name:       managedRuleGroupSum.Name,
		VendorName: managedRuleGroupSum.VendorName,
		Scope:      cl.WAFScope,
	}, func(o *wafv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, output)
}
