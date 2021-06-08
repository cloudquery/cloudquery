package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Wafv2ManagedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:                 "aws_wafv2_managed_rule_groups",
		Resolver:             fetchWafv2ManagedRuleGroups,
		Multiplex:            client.AccountRegionMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveDescribeManagedRuleGroup,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "available_labels",
				Type: schema.TypeStringArray,
			},
			{
				Name: "consumed_labels",
				Type: schema.TypeStringArray,
			},
			{
				Name: "capacity",
				Type: schema.TypeBigInt,
			},
			{
				Name: "label_namespace",
				Type: schema.TypeString,
			},
			{
				Name: "rules",
				Type: schema.TypeJSON,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "vendor_name",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafv2ManagedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2
	config := wafv2.ListAvailableManagedRuleGroupsInput{}
	for {
		output, err := service.ListAvailableManagedRuleGroups(ctx, &config, func(options *wafv2.Options) {
			options.Region = c.Region
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

func resolveDescribeManagedRuleGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	managedRuleGroupSum, ok := resource.Item.(types.ManagedRuleGroupSummary)
	if !ok {
		return fmt.Errorf("not a ManagedRuleGroupSummary instance: %#v", resource.Item)
	}

	c := meta.(*client.Client)
	service := c.Services().WafV2

	// Resolve managed rule group via describe managed rule group
	descrManagedRuleGroup, err := service.DescribeManagedRuleGroup(ctx, &wafv2.DescribeManagedRuleGroupInput{
		Name:       managedRuleGroupSum.Name,
		VendorName: managedRuleGroupSum.VendorName,
	}, func(options *wafv2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	// Available labels
	labels := make([]string, len(descrManagedRuleGroup.AvailableLabels))
	for id, label := range descrManagedRuleGroup.AvailableLabels {
		labels[id] = aws.ToString(label.Name)
	}
	if err := resource.Set("available_labels", labels); err != nil {
		return err
	}
	// Consumed labels
	labels = make([]string, len(descrManagedRuleGroup.ConsumedLabels))
	for id, label := range descrManagedRuleGroup.ConsumedLabels {
		labels[id] = aws.ToString(label.Name)
	}
	if err := resource.Set("consumed_labels", labels); err != nil {
		return err
	}
	// Capacity
	if err := resource.Set("capacity", cast.ToInt(descrManagedRuleGroup.Capacity)); err != nil {
		return err
	}
	// Label namespace
	if err := resource.Set("label_namespace", aws.ToString(descrManagedRuleGroup.LabelNamespace)); err != nil {
		return err
	}
	// Rules
	if len(descrManagedRuleGroup.Rules) > 0 {
		data, err := json.Marshal(descrManagedRuleGroup.Rules)
		if err != nil {
			return err
		}
		if err := resource.Set("rules", data); err != nil {
			return err
		}
	}
	return nil
}
