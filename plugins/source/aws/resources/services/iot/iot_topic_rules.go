package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotTopicRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_topic_rules",
		Description: "The output from the GetTopicRule operation.",
		Resolver:    fetchIotTopicRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotTopicRuleTags,
			},
			{
				Name:     "rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rule"),
			},
			{
				Name:            "arn",
				Description:     "The rule ARN.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("RuleArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotTopicRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTopicRulesInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListTopicRules(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.Rules {
			rule, err := svc.GetTopicRule(ctx, &iot.GetTopicRuleInput{
				RuleName: s.RuleName,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- rule
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotTopicRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetTopicRuleOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.RuleArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input)

		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
func resolveIotTopicRulesErrorActionHttpHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetTopicRuleOutput)
	if i.Rule == nil || i.Rule.ErrorAction == nil || i.Rule.ErrorAction.Http == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Rule.ErrorAction.Http.Headers {
		j[*h.Key] = *h.Value
	}
	return resource.Set(c.Name, j)
}
func resolveIotTopicRulesErrorActionTimestreamDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetTopicRuleOutput)
	if i.Rule == nil || i.Rule.ErrorAction == nil || i.Rule.ErrorAction.Timestream == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, h := range i.Rule.ErrorAction.Timestream.Dimensions {
		j[*h.Name] = *h.Value
	}
	return resource.Set(c.Name, j)
}
