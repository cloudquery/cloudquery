package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=web_acls.hcl -domain=wafregional -resource=web_acls
func WebAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_web_acls",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafregionalWebAcls,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
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
				Resolver:    resolveWebAclTags,
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
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_wafregional_web_acl_rules",
				Description: "The action for each Rule in a WebACL",
				Resolver:    fetchWafregionalWebAclRules,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_wafregional_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "priority",
						Description: "Specifies the order in which the Rules in a WebACL are evaluated",
						Type:        schema.TypeInt,
					},
					{
						Name:        "rule_id",
						Description: "The RuleId for a Rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "action",
						Description: "Specifies how you want AWS WAF to respond to requests that match the settings in a Rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Action.Type"),
					},
					{
						Name:        "excluded_rules",
						Description: "An array of rules to exclude from a rule group",
						Type:        schema.TypeStringArray,
						Resolver:    resolveWebACLRulesExcludedRules,
					},
					{
						Name:          "override_action",
						Description:   "Describes an override action for the rule.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("OverrideAction.Type"),
						IgnoreInTests: true,
					},
					{
						Name:        "type",
						Description: "The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup",
						Type:        schema.TypeString,
					},
				},
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
		result, err := svc.ListWebACLs(ctx, &params, func(o *wafregional.Options) { o.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range result.WebACLs {
			detail, err := svc.GetWebACL(
				ctx,
				&wafregional.GetWebACLInput{WebACLId: item.WebACLId},
				func(o *wafregional.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
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
func fetchWafregionalWebAclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	acl := parent.Item.(types.WebACL)
	res <- acl.Rules
	return nil
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
func resolveWebAclTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	params := wafregional.ListTagsForResourceInput{ResourceARN: resource.Item.(types.WebACL).WebACLArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return diag.WrapError(err)
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
