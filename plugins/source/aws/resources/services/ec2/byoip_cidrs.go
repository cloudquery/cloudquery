package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ByoipCidrs() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_byoip_cidrs",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html`,
		Resolver:    fetchEc2ByoipCidrs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.ByoipCidr{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "cidr",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
