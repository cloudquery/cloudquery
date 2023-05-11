package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Apps() *schema.Table {
	tableName := "aws_amplify_apps"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amplify/latest/APIReference/API_ListApps.html`,
		Resolver:    fetchApps,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "amplify"),
		Transform:   transformers.TransformWithStruct(&types.App{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := meta.(*client.Client).Services().Amplify

	config := amplify.ListAppsInput{
		MaxResults: int32(100),
	}
	// No paginator available
	for {
		output, err := svc.ListApps(ctx, &config, func(options *amplify.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.Apps
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
