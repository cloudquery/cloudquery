package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func NetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_network_interfaces",
		Description: "Describes a network interface.",
		Resolver:    fetchEc2NetworkInterfaces,
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
				Description: "The Amazon Resource Name (ARN) for the egress-only internet gateway.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"network-interface", *resource.Item.(types.NetworkInterface).NetworkInterfaceId}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the network interface.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTagField("TagSet"),
			},
			{
				Name:     "association",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Association"),
			},
			{
				Name:     "attachment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attachment"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone.",
				Type:        schema.TypeString,
			},
			{
				Name:          "deny_all_igw_traffic",
				Description:   "Indicates whether a network interface with an IPv6 address is unreachable from the public internet",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:        "description",
				Description: "A description.",
				Type:        schema.TypeString,
			},
			{
				Name:        "groups",
				Description: "The tags assigned to the egress-only internet gateway.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Groups"),
			},
			{
				Name:        "interface_type",
				Description: "The type of network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:     "ipv_4_prefixes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ipv4Prefixes"),
			},
			{
				Name:          "ipv6_address",
				Description:   "The IPv6 globally unique address associated with the network interface.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:     "ipv_6_addresses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ipv6Addresses"),
			},
			{
				Name:          "ipv6_native",
				Description:   "Indicates whether this is an IPv6 only network interface.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:     "ipv_6_prefixes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ipv6Prefixes"),
			},
			{
				Name:        "mac_address",
				Description: "The MAC address.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the network interface.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkInterfaceId"),
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "owner_id",
				Description: "The Amazon Web Services account ID of the owner of the network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_name",
				Description: "The private DNS name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_address",
				Description: "The IPv4 address of the network interface within the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "requester_id",
				Description: "The alias or Amazon Web Services account ID of the principal or service that created the network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "requester_managed",
				Description: "Indicates whether the network interface is being managed by Amazon Web Services.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "source_dest_check",
				Description: "Indicates whether source/destination checking is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "status",
				Description: "The status of the network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_addresses",
				Description: "Describes the private IPv4 address of a network interface.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("PrivateIpAddresses"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2NetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	input := ec2.DescribeNetworkInterfacesInput{}
	for {
		output, err := svc.DescribeNetworkInterfaces(ctx, &input, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.NetworkInterfaces
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
