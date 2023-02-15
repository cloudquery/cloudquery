package config

import (
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_config_config_rules",
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeConfigRules.html`,
		Resolver:    fetchConfigConfigRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer("config"),
		Transform:   transformers.TransformWithStruct(&types.ConfigRule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigRuleArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			ConfigRuleCompliances(),
		},
	}
}
