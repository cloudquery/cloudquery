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

func ebsSnapshotAttributes() *schema.Table {
	tableName := "aws_ec2_ebs_snapshot_attributes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshotAttribute.html`,
		Resolver:    fetchEbsSnapshotAttributes,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&ec2.DescribeSnapshotAttributeOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "snapshot_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchEbsSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Snapshot)
	cl := meta.(*client.Client)
	if aws.ToString(r.OwnerId) != cl.AccountID {
		return nil
	}
	svc := cl.Services().Ec2
	permissions, err := svc.DescribeSnapshotAttribute(ctx, &ec2.DescribeSnapshotAttributeInput{
		Attribute:  types.SnapshotAttributeNameCreateVolumePermission,
		SnapshotId: r.SnapshotId,
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	productCodes, err := svc.DescribeSnapshotAttribute(ctx, &ec2.DescribeSnapshotAttributeInput{
		Attribute:  types.SnapshotAttributeNameProductCodes,
		SnapshotId: r.SnapshotId,
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		// If the call for `ProductCodes` fails, we still want to return the `CreateVolumePermission` data
		res <- permissions
		return err
	}
	permissions.ProductCodes = productCodes.ProductCodes
	res <- permissions
	return nil
}
