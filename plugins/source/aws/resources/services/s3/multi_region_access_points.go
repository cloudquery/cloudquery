package s3

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func MultiRegionAccessPoints() *schema.Table {
	tableName := "aws_s3_multi_region_access_points"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_MultiRegionAccessPointReport.html`,
		Resolver:    fetchMultiRegionAccessPoints,
		Transform:   transformers.TransformWithStruct(&types.MultiRegionAccessPointReport{}),
		Multiplex:   client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Description:         `The Amazon Resource Name (ARN) of the multi-Region access point.`,
				Resolver:            resolveARN,
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchMultiRegionAccessPoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3control).S3control

	paginator := s3control.NewListMultiRegionAccessPointsPaginator(svc, &s3control.ListMultiRegionAccessPointsInput{
		AccountId: aws.String(cl.AccountID),
	})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *s3control.Options) {
			// According to the docs: This action will always be routed to the US West (Oregon) Region
			o.Region = "us-west-2"
		})
		if err != nil {
			return err
		}
		res <- output.AccessPoints
	}

	return nil
}

func resolveARN(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	mrap := r.Item.(types.MultiRegionAccessPointReport)
	return r.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.S3Service),
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("accesspoint/%s", aws.ToString(mrap.Alias)),
	}.String())
}
