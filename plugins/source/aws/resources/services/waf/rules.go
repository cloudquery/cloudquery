package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Rules() *schema.Table {
	tableName := "aws_waf_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleSummary.html`,
		Resolver:    fetchWafRules,
		Transform:   transformers.TransformWithStruct(&types.Rule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafRuleArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafRuleTags,
			},
		},
	}
}

func fetchWafRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRulesInput{}
	for {
		output, err := service.ListRules(ctx, &config)
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
	cl := meta.(*client.Client)
	rule := resource.Item.(*types.Rule)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "waf",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("rule/%s", aws.ToString(rule.RuleId)),
	}.String())
}

func resolveWafRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(*types.Rule)

	// Resolve tags for resource
	cl := meta.(*client.Client)
	service := cl.Services().Waf

	// Generate arn
	arnStr := arn.ARN{
		Partition: cl.Partition,
		Service:   "waf",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("rule/%s", aws.ToString(rule.RuleId)),
	}.String()

	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: aws.String(arnStr)}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig)
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
