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

func fetchEc2EbsSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	config.OwnerIds = []string{c.AccountID}
	for {
		output, err := svc.DescribeSnapshots(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Snapshots
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveEbsSnapshotAttribute(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Snapshot)
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
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
