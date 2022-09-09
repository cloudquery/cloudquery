package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2InstanceStatuses() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instance_statuses",
		Description: "Describes the status of an instance.",
		Resolver:    fetchEc2InstanceStatuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"instance", *resource.Item.(types.InstanceStatus).InstanceId}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_state_code",
				Description: "The state of the instance as a 16-bit unsigned integer.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("InstanceState.Code"),
			},
			{
				Name:     "instance_state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceState"),
			},
			{
				Name:     "instance_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceStatus"),
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:     "system_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemStatus"),
			},
			{
				Name:        "events",
				Description: "Any scheduled events associated with the instance.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Events"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2InstanceStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeInstanceStatusInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeInstanceStatus(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.InstanceStatuses
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
