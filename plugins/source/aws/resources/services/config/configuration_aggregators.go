package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ConfigurationAggregators() *schema.Table {
	tableName := "aws_config_configuration_aggregators"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationAggregator.html`,
		Resolver:    fetchConfigurationAggregators,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&types.ConfigurationAggregator{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigurationAggregatorArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{},
	}
}

func fetchConfigurationAggregators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Configservice

	p := configservice.NewDescribeConfigurationAggregatorsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *configservice.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.ConfigurationAggregators
	}
	return nil
}
