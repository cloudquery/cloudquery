package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AccessPoints() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_access_points",
		Resolver:  fetchAccessPoints,
		Transform: transformers.TransformWithStruct(&types.AccessPoint{}),
		Multiplex: client.ServiceAccountRegionMultiplexer("s3-control"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessPointArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAccessPoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().S3control
	paginator := s3control.NewListAccessPointsPaginator(svc, &s3control.ListAccessPointsInput{
		AccountId: aws.String(c.AccountID),
	})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.AccessPointList
	}

	return nil
}
