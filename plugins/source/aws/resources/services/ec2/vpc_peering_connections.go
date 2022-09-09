package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2VpcPeeringConnections() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_vpc_peering_connections",
		Description:   "Describes a VPC peering connection.",
		Resolver:      fetchEc2VpcPeeringConnections,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpc-peering-connection", *resource.Item.(types.VpcPeeringConnection).VpcPeeringConnectionId}, nil
				}),
			},
			{
				Name:     "accepter_vpc_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccepterVpcInfo"),
			},
			{
				Name:        "expiration_time",
				Description: "The time that an unaccepted VPC peering connection will expire.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "requester_vpc_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequesterVpcInfo"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the resource.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:            "id",
				Description:     "The ID of the VPC peering connection.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("VpcPeeringConnectionId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2VpcPeeringConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpcPeeringConnectionsInput
	c := meta.(*client.Client)
	svc := meta.(*client.Client).Services().EC2
	for {
		output, err := svc.DescribeVpcPeeringConnections(ctx, &config, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.VpcPeeringConnections
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
