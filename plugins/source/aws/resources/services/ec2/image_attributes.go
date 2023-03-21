package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func imageAttributesLaunchPermissions() *schema.Table {
	const tableName = "aws_ec2_image_launch_permissions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html`,
		Resolver:    fetchEc2ImageAttributeLaunchPermissions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.LaunchPermission{}),
		Columns: []schema.Column{
			{
				Name:     "image_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchEc2ImageAttributeLaunchPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(types.Image)

	svc := c.Services().Ec2
	output, err := svc.DescribeImageAttribute(ctx, &ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLaunchPermission,
		ImageId:   p.ImageId,
	})
	if err != nil {
		return err
	}
	res <- output.LaunchPermissions
	return nil
}
