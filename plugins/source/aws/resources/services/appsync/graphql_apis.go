package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func GraphqlApis() *schema.Table {
	tableName := "aws_appsync_graphql_apis"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html`,
		Resolver:    fetchAppsyncGraphqlApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appsync"),
		Transform:   transformers.TransformWithStruct(&types.GraphqlApi{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAppsyncGraphqlApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config appsync.ListGraphqlApisInput
	c := meta.(*client.Client)
	svc := c.Services().Appsync
	// No paginator available
	for {
		output, err := svc.ListGraphqlApis(ctx, &config, func(options *appsync.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.GraphqlApis
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
