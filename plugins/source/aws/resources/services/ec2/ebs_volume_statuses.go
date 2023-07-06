package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EbsVolumesStatuses() *schema.Table {
	tableName := "aws_ec2_ebs_volume_statuses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusItem.html`,
		Resolver:    fetchEc2EbsVolumeStatuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.VolumeStatusItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "volume_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveEbsVolumeStatusArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchEc2EbsVolumeStatuses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	config := ec2.DescribeVolumeStatusInput{MaxResults: aws.Int32(1000)}
	paginator := ec2.NewDescribeVolumeStatusPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.VolumeStatuses
	}
	return nil
}

func resolveEbsVolumeStatusArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	volume := resource.Item.(types.VolumeStatusItem)
	a := resolveVolumeARN(cl.Partition, cl.Region, cl.AccountID, aws.ToString(volume.VolumeId))
	return resource.Set(c.Name, a)
}
