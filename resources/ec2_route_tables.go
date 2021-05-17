package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2RouteTables() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_route_tables",
		Resolver:     fetchEc2RouteTables,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteTableId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2routeTableTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_route_table_associations",
				Resolver: fetchEc2RouteTableAssociations,
				Columns: []schema.Column{
					{
						Name:     "route_table_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "association_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AssociationState.State"),
					},
					{
						Name:     "association_state_status_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AssociationState.StatusMessage"),
					},
					{
						Name: "gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "main",
						Type: schema.TypeBool,
					},
					{
						Name: "route_table_association_id",
						Type: schema.TypeString,
					},
					{
						Name: "subnet_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_route_table_propagating_vgws",
				Resolver: fetchEc2RouteTablePropagatingVgws,
				Columns: []schema.Column{
					{
						Name:     "route_table_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "gateway_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_route_table_routes",
				Resolver: fetchEc2RouteTableRoutes,
				Columns: []schema.Column{
					{
						Name:     "route_table_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "carrier_gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "destination_cidr_block",
						Type: schema.TypeString,
					},
					{
						Name: "destination_ipv6_cidr_block",
						Type: schema.TypeString,
					},
					{
						Name: "destination_prefix_list_id",
						Type: schema.TypeString,
					},
					{
						Name: "egress_only_internet_gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "instance_id",
						Type: schema.TypeString,
					},
					{
						Name: "instance_owner_id",
						Type: schema.TypeString,
					},
					{
						Name: "local_gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "nat_gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "network_interface_id",
						Type: schema.TypeString,
					},
					{
						Name: "origin",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "transit_gateway_id",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_peering_connection_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2RouteTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeRouteTablesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeRouteTables(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
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
func resolveEc2routeTableTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.RouteTable)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2RouteTableAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.Associations
	return nil
}
func fetchEc2RouteTablePropagatingVgws(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.PropagatingVgws
	return nil
}
func fetchEc2RouteTableRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.RouteTable)
	res <- r.Routes
	return nil
}
