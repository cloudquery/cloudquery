package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var ec2Resoruces []*Resource = []*Resource{
	{
		SubService: "byoip_cidrs",
		Struct: &types.ByoipCidr{},
		SkipFields: []string{"Cidr"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name: "account_id",
				Type: schema.TypeString,
				Resolver: "client.ResolveAWSAccount",
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "region",
				Type: schema.TypeString,
				Resolver: "client.ResolveAWSRegion",
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "cidr",
				Type: schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
	{
		SubService: "customer_gateways",
		Struct: &types.CustomerGateway{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveCustomerGatewayArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "ebs_snapshots",
		Struct: &types.Snapshot{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveEbsSnapshotArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name: "attribute",
					Type: schema.TypeString,
					Resolver: "resolveEbsSnapshotAttribute",
				},
			}...),
	},
	{
		SubService: "ebs_volumes",
		Struct: &types.Volume{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveEbsVolumeArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "egress_only_internet_gateways",
		Struct: &types.Volume{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveEgressOnlyInternetGatewaysArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "eips",
		Struct: &types.Address{},
		ExtraColumns: defaultRegionalColumns,
	},
	{
		SubService: "flow_logs",
		Struct: &types.FlowLog{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveFlowLogArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "hosts",
		Struct: &types.Host{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveHostArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "images",
		Struct: &types.Image{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveImageArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "instance_statuses",
		Struct: &types.InstanceStatus{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveInstanceStatusArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "instances",
		Struct: &types.Instance{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveInstanceArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "instance_types",
		Struct: &types.InstanceTypeInfo{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveInstanceTypeArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "internet_gateways",
		Struct: &types.InternetGateway{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveInternetGatewayArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "key_pairs",
		Struct: &types.KeyPairInfo{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveKeyPairArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "nat_gateways",
		Struct: &types.NatGateway{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveNatGatewayArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "nat_gateways",
		Struct: &types.NatGateway{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveNatGatewayArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "network_acls",
		Struct: &types.NatGateway{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveNetworkAclArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "network_interfaces",
		Struct: &types.NetworkInterface{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveNetworkInterfaceArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "route_tables",
		Struct: &types.RouteTable{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveRouteTableArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "security_groups",
		Struct: &types.SecurityGroup{},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: "resolveSecurityGroupArn",
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "subnets",
		Struct: &types.Subnet{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `schema.PathResolver("SubnetArn")`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpc_endpoint_service_configurations",
		Struct: &types.ServiceConfiguration{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpcEndpointServiceConfigurationArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpc_endpoint_services",
		Struct: &types.ServiceDetail{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpcEndpointServiceArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpc_endpoints",
		Struct: &types.VpcEndpoint{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpcEndpointArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpc_peering_connections",
		Struct: &types.VpcPeeringConnection{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpcPeeringConnectionArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpcs",
		Struct: &types.Vpc{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpcArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
	{
		SubService: "vpn_gateways",
		Struct: &types.VpnGateway{},
		SkipFields: []string{"arn"},
		ExtraColumns: append(defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveVpnGatewayArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
	},
}

func Ec2Resources() []*Resource {
	for _, r := range ec2Resoruces {
		r.Service = "ec2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("ec2")`
		// resource.
	}
	return ec2Resoruces
}