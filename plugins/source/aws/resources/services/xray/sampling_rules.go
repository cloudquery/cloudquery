package xray

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SamplingRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_sampling_rules",
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html`,
		Resolver:    fetchXraySamplingRules,
		Transform:   transformers.TransformWithStruct(&types.SamplingRuleRecord{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SamplingRule.RuleARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXraySamplingRuleTags,
			},
		},
	}
}
