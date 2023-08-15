package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func imageAttributesLastLaunchTime() *schema.Table {
	return &schema.Table{
		Name: "aws_ec2_image_last_launched_times",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImageAttribute.html. 
The date and time, in ISO 8601 date-time format, when the AMI was last used to launch an EC2 instance. When the AMI is used to launch an instance, there is a 24-hour delay before that usage is reported.`,
		Resolver: fetchEc2ImageAttributeLastLaunchTime,
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
	cl := meta.(*client.Client)
	p := parent.Item.(types.Image)
	if aws.ToString(p.OwnerId) != cl.AccountID {
		return nil
	}
	svc := cl.Services().Ec2
	output, err := svc.DescribeImageAttribute(ctx, &ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLastLaunchedTime,
		ImageId:   p.ImageId,
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.LastLaunchedTime
	return nil
}
