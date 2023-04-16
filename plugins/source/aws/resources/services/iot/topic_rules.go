package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func TopicRules() *schema.Table {
	tableName := "aws_iot_topic_rules"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_GetTopicRule.html`,
		Resolver:            fetchIotTopicRules,
		PreResourceResolver: getTopicRule,
		Transform:           transformers.TransformWithStruct(&iot.GetTopicRuleOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIotTopicRuleTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotTopicRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTopicRulesInput{
		MaxResults: aws.Int32(250),
	}
	paginator := iot.NewListTopicRulesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Rules
	}
	return nil
}

func getTopicRule(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.GetTopicRule(ctx, &iot.GetTopicRuleInput{
		RuleName: resource.Item.(types.TopicRuleListItem).RuleName,
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

func resolveIotTopicRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetTopicRuleOutput)
	svc := meta.(*client.Client).Services().Iot
	return resolveIotTags(ctx, svc, resource, c, i.RuleArn)
}
