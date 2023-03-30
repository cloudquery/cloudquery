package wafregional

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Rules() *schema.Table {
	tableName := "aws_wafregional_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_Rule.html`,
		Resolver:    fetchWafregionalRules,
		Transform:   client.TransformWithStruct(&types.Rule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf-regional"),
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
