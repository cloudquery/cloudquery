package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_rules",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafRules,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafRuleArn,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafRuleTags,
			},
			{
				Name:        "id",
				Description: "A unique identifier for a Rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this Rule",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The friendly name or description for the Rule",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_waf_rule_predicates",
				Description: "This is AWS WAF Classic documentation",
				Resolver:    fetchWafRulePredicates,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"rule_cq_id", "data_id"}},
				Columns: []schema.Column{
					{
						Name:        "rule_cq_id",
						Description: "Unique CloudQuery ID of aws_waf_rules table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "data_id",
						Description: "A unique identifier for a predicate in a Rule, such as ByteMatchSetId or IPSetId",
						Type:        schema.TypeString,
					},
					{
						Name:        "negated",
						Description: "Set Negated to False if you want AWS WAF to allow, block, or count requests based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, RegexMatchSet, GeoMatchSet, or SizeConstraintSet",
						Type:        schema.TypeBool,
					},
					{
						Name:        "type",
						Description: "The type of predicate in a Rule, such as ByteMatch or IPSet.",
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
func fetchWafRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRulesInput{}
	for {
		output, err := service.ListRules(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, ruleSum := range output.Rules {
			rule, err := service.GetRule(ctx, &waf.GetRuleInput{RuleId: ruleSum.RuleId}, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- rule.Rule
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafRuleArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(*types.Rule)
	if !ok {
		return fmt.Errorf("not a Rule instance: %#v", resource.Item)
	}
	usedClient := meta.(*client.Client)

	// Generate arn
	arnStr := client.GenerateResourceARN(
		"waf",
		"rule",
		aws.ToString(rule.RuleId),
		usedClient.Region,
		usedClient.AccountID)

	return resource.Set(c.Name, arnStr)
}
func resolveWafRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(*types.Rule)
	if !ok {
		return fmt.Errorf("not a Rule instance: %#v", resource.Item)
	}

	// Resolve tags for resource
	usedClient := meta.(*client.Client)
	service := usedClient.Services().Waf

	// Generate arn
	arnStr := client.GenerateResourceARN(
		"waf",
		"rule",
		aws.ToString(rule.RuleId),
		"",
		usedClient.AccountID)

	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: aws.String(arnStr)}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *waf.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
		})
		if err != nil {
			return err
		}
		for _, t := range tags.TagInfoForResource.TagList {
			outputTags[*t.Key] = t.Value
		}
		if aws.ToString(tags.NextMarker) == "" {
			break
		}
		tagsConfig.NextMarker = tags.NextMarker
	}
	return resource.Set("tags", outputTags)
}
func fetchWafRulePredicates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	rule, ok := parent.Item.(*types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", rule)
	}
	res <- rule.Predicates
	return nil
}
