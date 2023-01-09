package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InstanceStatuses() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instance_statuses",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatus.html`,
		Resolver:    fetchEc2InstanceStatuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.InstanceStatus{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveInstanceStatusArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
