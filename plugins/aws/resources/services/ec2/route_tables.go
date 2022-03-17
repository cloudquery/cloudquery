package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2RouteTables() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_route_tables",
		Description:  "Describes a route table.",
		Resolver:     fetchEc2RouteTables,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
					return []string{"route-table", *resource.Item.(types.RouteTable).RouteTableId}, nil
				}),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the route table.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the route table.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RouteTableId"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the route table.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2routeTableTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_route_table_associations",
				Description: "Describes an association between a route table and a subnet or gateway.",
				Resolver:    fetchEc2RouteTableAssociations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"route_table_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "route_table_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_route_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The ID of the association.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RouteTableAssociationId"),
					},
					{
						Name:        "association_state",
						Description: "The state of the association.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociationState.State"),
					},
					{
						Name:          "association_state_status_message",
						Description:   "The status message, if applicable.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("AssociationState.StatusMessage"),
						IgnoreInTests: true,
					},
					{
						Name:          "gateway_id",
						Description:   "The ID of the internet gateway or virtual private gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "main",
						Description: "Indicates whether this is the main route table.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_route_table_propagating_vgws",
				Description:   "Describes a virtual private gateway propagating route.",
				Resolver:      fetchEc2RouteTablePropagatingVgws,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"route_table_cq_id", "gateway_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "route_table_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_route_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "gateway_id",
						Description: "The ID of the virtual private gateway.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_ec2_route_table_routes",
				Description: "Describes a route in a route table.",
				Resolver:    fetchEc2RouteTableRoutes,
				Columns: []schema.Column{
					{
						Name:        "route_table_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_route_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "carrier_gateway_id",
						Description:   "The ID of the carrier gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "destination_cidr_block",
						Description: "The IPv4 CIDR block used for the destination match.",
						Type:        schema.TypeString,
					},
					{
						Name:        "destination_ipv6_cidr_block",
						Description: "The IPv6 CIDR block used for the destination match.",
						Type:        schema.TypeString,
					},
					{
						Name:          "destination_prefix_list_id",
						Description:   "The prefix of the AWS service.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "egress_only_internet_gateway_id",
						Description:   "The ID of the egress-only internet gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "gateway_id",
						Description: "The ID of a gateway attached to your VPC.",
						Type:        schema.TypeString,
					},
					{
						Name:          "instance_id",
						Description:   "The ID of a NAT instance in your VPC.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "instance_owner_id",
						Description:   "The AWS account ID of the owner of the instance.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "local_gateway_id",
						Description:   "The ID of the local gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "nat_gateway_id",
						Description:   "The ID of a NAT gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "network_interface_id",
						Description:   "The ID of the network interface.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "origin",
						Description: "Describes how the route was created.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "The state of the route.",
						Type:        schema.TypeString,
					},
					{
						Name:          "transit_gateway_id",
						Description:   "The ID of a transit gateway.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "vpc_peering_connection_id",
						Description:   "The ID of a VPC peering connection.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2RouteTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeRouteTablesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeRouteTables(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.RouteTables
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2routeTableTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.RouteTable)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2RouteTableAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.Associations
	return nil
}
func fetchEc2RouteTablePropagatingVgws(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.PropagatingVgws
	return nil
}
func fetchEc2RouteTableRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.Routes
	return nil
}
