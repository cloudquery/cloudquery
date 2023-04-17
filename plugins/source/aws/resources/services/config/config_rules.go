package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ConfigRules() *schema.Table {
	tableName := "aws_config_config_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeConfigRules.html`,
		Resolver:    fetchConfigConfigRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
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
			configRuleCompliances(),
		},
	}
}

func fetchConfigConfigRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Configservice

	input := &configservice.DescribeConfigRulesInput{}
	p := configservice.NewDescribeConfigRulesPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ConfigRules
	}
	return nil
}
