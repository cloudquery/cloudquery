package appstream

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DirectoryConfigs() *schema.Table {
	tableName := "aws_appstream_directory_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html`,
		Resolver:    fetchAppstreamDirectoryConfigs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.DirectoryConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "directory_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DirectoryName"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAppstreamDirectoryConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeDirectoryConfigsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Appstream
	// No paginator available
	for {
		response, err := svc.DescribeDirectoryConfigs(ctx, &input, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DirectoryConfigs

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
