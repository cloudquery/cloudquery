package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func imageAttributesLastLaunchTime() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_image_last_launched_times",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html`,
		Resolver:    fetchEc2ImageAttributeLastLaunchTime,
		Columns: []schema.Column{
			{
				Name:       "image_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "last_launched_time",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("Value"),
			},
		},
	}
}

func fetchEc2ImageAttributeLastLaunchTime(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(types.Image)
	if aws.ToString(p.OwnerId) != c.AccountID {
		return nil
	}
	svc := c.Services().Ec2
	output, err := svc.DescribeImageAttribute(ctx, &ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLastLaunchedTime,
		ImageId:   p.ImageId,
	}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.LastLaunchedTime
	return nil
}
