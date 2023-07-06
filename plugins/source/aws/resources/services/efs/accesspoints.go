package efs

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func AccessPoints() *schema.Table {
	tableName := "aws_efs_access_points"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/efs/latest/ug/API_AccessPointDescription.html`,
		Resolver:    fetchAccessPoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticfilesystem"),
		Transform:   transformers.TransformWithStruct(&types.AccessPointDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AccessPointArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     cqtypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchAccessPoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Efs
	paginator := efs.NewDescribeAccessPointsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *efs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AccessPoints
	}
	return nil
}
