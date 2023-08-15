package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RetentionConfigurations() *schema.Table {
	tableName := "aws_config_retention_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_RetentionConfiguration.html`,
		Resolver:    fetchRetentionConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform: transformers.TransformWithStruct(&types.RetentionConfiguration{},
			transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
		Relations: []*schema.Table{},
	}
}

func fetchRetentionConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Configservice

	p := configservice.NewDescribeRetentionConfigurationsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *configservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.RetentionConfigurations
	}
	return nil
}
