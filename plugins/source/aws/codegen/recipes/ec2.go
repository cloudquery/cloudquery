package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EC2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "byoip_cidrs",
			Struct:      &types.ByoipCidr{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html",
			SkipFields:  []string{"Cidr"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:    "cidr",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "customer_gateways",
			Struct:      &types.CustomerGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CustomerGateway.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveCustomerGatewayArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "ebs_snapshots",
			Struct:      &types.Snapshot{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Snapshot.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveEbsSnapshotArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "attribute",
						Type:     schema.TypeJSON,
						Resolver: "resolveEbsSnapshotAttribute",
					},
				}...),
		},
		{
			SubService:  "ebs_volumes",
			Struct:      &types.Volume{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveEbsVolumeArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "egress_only_internet_gateways",
			Struct:      &types.EgressOnlyInternetGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EgressOnlyInternetGateway.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveEgressOnlyInternetGatewaysArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:   "eips",
			Struct:       &types.Address{},
			Description:  "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html",
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService:  "flow_logs",
			Struct:      &types.FlowLog{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FlowLog.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveFlowLogArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "hosts",
			Struct:      &types.Host{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Host.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveHostArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "images",
			Struct:      &types.Image{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveImageArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "instance_statuses",
			Struct:      &types.InstanceStatus{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatus.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveInstanceStatusArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "instances",
			Struct:      &types.Instance{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveInstanceArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:          "state_transition_reason_time",
						Type:          schema.TypeTimestamp,
						Resolver:      "resolveEc2InstanceStateTransitionReasonTime",
						IgnoreInTests: true,
					},
				}...),
		},
		{
			SubService:  "instance_types",
			Struct:      &types.InstanceTypeInfo{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceTypeInfo.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveInstanceTypeArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "internet_gateways",
			Struct:      &types.InternetGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InternetGateway.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveInternetGatewayArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "key_pairs",
			Struct:      &types.KeyPairInfo{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_KeyPairInfo.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveKeyPairArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "nat_gateways",
			Struct:      &types.NatGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NatGateway.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveNatGatewayArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "network_acls",
			Struct:      &types.NetworkAcl{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveNetworkAclArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "network_interfaces",
			Struct:      &types.NetworkInterface{},
			SkipFields:  []string{"TagSet"},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterface.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveNetworkInterfaceArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTagField("TagSet")`,
					},
				}...),
		},
		{
			Name:        "aws_regions", // rename table for backwards-compatibility
			SubService:  "regions",
			Struct:      &types.Region{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html",
			SkipFields:  []string{"RegionName"},
			Multiplex:   `client.AccountMultiplex`,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
				},
				{
					Name:     "enabled",
					Type:     schema.TypeBool,
					Resolver: `resolveRegionEnabled`,
				},
				{
					Name:     "partition",
					Type:     schema.TypeString,
					Resolver: `resolveRegionPartition`,
				},
				// for backwards-compatibility: renamed "region_name" to "region"
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("RegionName")`,
				},
			},
		},
		{
			SubService: "regional_configs",
			Struct:     &models.RegionalConfig{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "reserved_instances",
			Struct:      &types.ReservedInstances{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReservedInstances.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveReservedInstanceArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "route_tables",
			Struct:      &types.RouteTable{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTable.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveRouteTableArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "security_groups",
			Struct:      &types.SecurityGroup{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: "resolveSecurityGroupArn",
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "subnets",
			Struct:      &types.Subnet{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Subnet.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SubnetArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "transit_gateways",
			Struct:      &types.TransitGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGateway.html",
			SkipFields:  []string{"TransitGatewayId", "TransitGatewayArn"},
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TransitGatewayId")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TransitGatewayArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"TransitGatewayAttachments()",
				"TransitGatewayRouteTables()",
				"TransitGatewayVpcAttachments()",
				"TransitGatewayPeeringAttachments()",
				"TransitGatewayMulticastDomains()",
			},
		},
		{
			SubService:  "transit_gateway_attachments",
			Struct:      &types.TransitGatewayAttachment{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachment.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "transit_gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "transit_gateway_route_tables",
			Struct:      &types.TransitGatewayRouteTable{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteTable.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "transit_gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "transit_gateway_vpc_attachments",
			Struct:      &types.TransitGatewayVpcAttachment{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayVpcAttachment.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "transit_gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "transit_gateway_peering_attachments",
			Struct:      &types.TransitGatewayPeeringAttachment{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "transit_gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "transit_gateway_multicast_domains",
			Struct:      &types.TransitGatewayMulticastDomain{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomain.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "transit_gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "vpc_endpoint_service_configurations",
			Struct:      &types.ServiceConfiguration{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceConfiguration.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcEndpointServiceConfigurationArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "vpc_endpoint_services",
			Struct:      &types.ServiceDetail{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcEndpointServiceArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "vpc_endpoints",
			Struct:      &types.VpcEndpoint{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpoint.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcEndpointArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "vpc_peering_connections",
			Struct:      &types.VpcPeeringConnection{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnection.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcPeeringConnectionArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "vpcs",
			Struct:      &types.Vpc{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Vpc.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "vpn_gateways",
			Struct:      &types.VpnGateway{},
			Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnGateway.html",
			ExtraColumns: append(defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpnGatewayArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}
	for _, r := range resources {
		r.Service = "ec2"
		if r.Multiplex == "" {
			r.Multiplex = `client.ServiceAccountRegionMultiplexer("ec2")`
		}
	}
	return resources
}
