package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2RouteTables() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_route_tables",
		Description: "Describes a route table.",
		Resolver:    fetchEc2RouteTables,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
					return []string{"route-table", *resource.Item.(types.RouteTable).RouteTableId}, nil
				}),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the route table.",
				Type:        schema.TypeString,
			},
			{
				Name:            "id",
				Description:     "The ID of the route table.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("RouteTableId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the route table.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "associations",
				Description: "Describes an association between a route table and a subnet or gateway.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Associations"),
			},
			{
				Name:        "propagating_vgws",
				Description: "Describes a virtual private gateway propagating route.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("PropagatingVgws"),
			},
			{
				Name:        "routes",
				Description: "Describes a route in a route table.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Routes"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2RouteTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeRouteTablesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeRouteTables(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.RouteTables
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
