package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EbsSnapshots() *schema.Table {
	tableName := "aws_ec2_ebs_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchEc2EbsSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   client.TransformWithStruct(&types.Snapshot{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveEbsSnapshotArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attribute",
				Type:     schema.TypeJSON,
				Resolver: resolveEbsSnapshotAttribute,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2EbsSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	paginator := ec2.NewDescribeSnapshotsPaginator(svc, &ec2.DescribeSnapshotsInput{
		OwnerIds:   []string{c.AccountID},
		MaxResults: aws.Int32(1000),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Snapshots
	}
	return nil
}

func resolveEbsSnapshotAttribute(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Snapshot)
	svc := meta.(*client.Client).Services().Ec2
	output, err := svc.DescribeSnapshotAttribute(ctx, &ec2.DescribeSnapshotAttributeInput{
		Attribute:  types.SnapshotAttributeNameCreateVolumePermission,
		SnapshotId: r.SnapshotId,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output)
}

func resolveEbsSnapshotArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "snapshot/" + aws.ToString(resource.Item.(types.Snapshot).SnapshotId),
	}
	return resource.Set(c.Name, a.String())
}
