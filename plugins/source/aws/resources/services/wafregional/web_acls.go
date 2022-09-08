package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)


func WebAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_web_acls",
		Description:  "Contains the Rules that identify the requests that you want to allow, block, or count.",
		Resolver:     fetchWafregionalWebAcls,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
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
				Name:        "tags",
				Description: "Web ACL tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalWebACLTags,
			},
			{
				Name:        "default_action",
				Description: "Specifies how you want AWS WAF to respond to requests that match the settings in a Rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultAction.Type"),
			},
			{
				Name:        "id",
				Description: "A unique identifier for a WebACL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WebACLId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this WebACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "A friendly name or description of the WebACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "Tha Amazon Resource Name (ARN) of the web ACL.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "rules",
				Description: "The action for each Rule in a WebACL",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafregionalWebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	var params wafregional.ListWebACLsInput
	for {
		result, err := svc.ListWebACLs(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, item := range result.WebACLs {
			detail, err := svc.GetWebACL(
				ctx,
				&wafregional.GetWebACLInput{WebACLId: item.WebACLId},
				func(o *wafregional.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if detail.WebACL == nil {
				continue
			}
			res <- *detail.WebACL
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}
func resolveWafregionalWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	params := wafregional.ListTagsForResourceInput{ResourceARN: resource.Item.(types.WebACL).WebACLArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return err
		}
		if result.TagInfoForResource != nil {
			client.TagsIntoMap(result.TagInfoForResource.TagList, tags)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
func resolveWebACLRulesExcludedRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.ActivatedRule)
	ids := make([]string, len(rule.ExcludedRules))
	for _, item := range rule.ExcludedRules {
		if item.RuleId != nil {
			ids = append(ids, *item.RuleId)
		}
	}
	return diag.WrapError(resource.Set(c.Name, ids))
}
