package wafv2

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
)

func Wafv2ManagedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:                 "aws_wafv2_managed_rule_groups",
		Description:          "High-level information about a managed rule group, returned by ListAvailableManagedRuleGroups",
		Resolver:             fetchWafv2ManagedRuleGroups,
		Multiplex:            client.ServiceAccountRegionMultiplexer("waf-regional"),
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolveDescribeManagedRuleGroup,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "vendor_name", "name"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
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
				Name:        "description",
				Description: "The description of the managed rule group, provided by AWS Managed Rules or the AWS Marketplace seller who manages it.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the managed rule group",
				Type:        schema.TypeString,
			},
			{
				Name:        "vendor_name",
				Description: "The name of the managed rule group vendor",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafv2ManagedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2

	// Dependent on the region select the right scope
	scope := types.ScopeRegional
	region := c.Region
	if region == strings.ToLower("global") {
		region = "us-east-1"
		scope = types.ScopeCloudfront
	}
	config := wafv2.ListAvailableManagedRuleGroupsInput{Scope: scope}
	for {
		output, err := service.ListAvailableManagedRuleGroups(ctx, &config, func(options *wafv2.Options) {
			options.Region = region
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

	// Dependent on the region select the right scope
	scope := types.ScopeRegional
	region := c.Region
	if region == strings.ToLower("global") {
		region = "us-east-1"
		scope = types.ScopeCloudfront
	}

	// Resolve managed rule group via describe managed rule group
	descrManagedRuleGroup, err := service.DescribeManagedRuleGroup(ctx, &wafv2.DescribeManagedRuleGroupInput{
		Name:       managedRuleGroupSum.Name,
		VendorName: managedRuleGroupSum.VendorName,
		Scope:      scope,
	}, func(options *wafv2.Options) {
		options.Region = region
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
