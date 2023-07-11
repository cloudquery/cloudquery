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

func imageAttributesLaunchPermissions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_image_launch_permissions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html`,
		Resolver:    fetchEc2ImageAttributeLaunchPermissions,
		Transform:   transformers.TransformWithStruct(&types.LaunchPermission{}),
		Columns: []schema.Column{
			{
				Name:     "image_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchEc2ImageAttributeLaunchPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(types.Image)
	if aws.ToString(p.OwnerId) != cl.AccountID {
		return nil
	}
	svc := cl.Services().Ec2
	output, err := svc.DescribeImageAttribute(ctx, &ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLaunchPermission,
		ImageId:   p.ImageId,
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.LaunchPermissions
	return nil
}
