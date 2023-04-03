package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Volumes() *schema.Table {
	tableName := "aws_fsx_volumes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_Volume.html`,
		Resolver:    fetchFsxVolumes,
		Transform:   transformers.TransformWithStruct(&types.Volume{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "administrative_actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdministrativeActions"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchFsxVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	input := fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeVolumesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.Volumes
	}
	return nil
}
