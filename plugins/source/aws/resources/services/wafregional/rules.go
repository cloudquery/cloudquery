package wafregional

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafregional_rules",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_Rule.html`,
		Resolver:    fetchWafregionalRules,
		Transform:   transformers.TransformWithStruct(&types.Rule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafregionalRuleArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRuleTags,
				Description: `Rule tags.`,
			},
		},
	}
}
