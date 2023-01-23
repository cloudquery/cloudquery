package appsync

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GraphqlApis() *schema.Table {
	return &schema.Table{
		Name:        "aws_appsync_graphql_apis",
		Description: `https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html`,
		Resolver:    fetchAppsyncGraphqlApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appsync"),
		Transform:   transformers.TransformWithStruct(&types.GraphqlApi{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
