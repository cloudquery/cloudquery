package ec2

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource network_interfaces --config gen.hcl --output .
func NetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_network_interfaces",
		Description:  "Describes a network interface.",
		Resolver:     fetchEc2NetworkInterfaces,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the network interface.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTagField("TagSet"),
			},
			{
				Name:        "association_allocation_id",
				Description: "The allocation ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Association.AllocationId"),
			},
			{
				Name:        "association_id",
				Description: "The association ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Association.AssociationId"),
			},
			{
				Name:          "association_carrier_ip",
				Description:   "The carrier IP address associated with the network interface",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Association.CarrierIp"),
				IgnoreInTests: true,
			},
			{
				Name:          "association_customer_owned_ip",
				Description:   "The customer-owned IP address associated with the network interface.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Association.CustomerOwnedIp"),
				IgnoreInTests: true,
			},
			{
				Name:        "association_ip_owner_id",
				Description: "The ID of the Elastic IP address owner.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Association.IpOwnerId"),
			},
			{
				Name:        "association_public_dns_name",
				Description: "The public DNS name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Association.PublicDnsName"),
			},
			{
				Name:        "association_public_ip",
				Description: "The address of the Elastic IP address bound to the network interface.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Association.PublicIp"),
			},
			{
				Name:        "attachment_attach_time",
				Description: "The timestamp indicating when the attachment initiated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attachment.AttachTime"),
			},
			{
				Name:        "attachment_id",
				Description: "The ID of the network interface attachment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attachment.AttachmentId"),
			},
			{
				Name:        "attachment_delete_on_termination",
				Description: "Indicates whether the network interface is deleted when the instance is terminated.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attachment.DeleteOnTermination"),
			},
			{
				Name:        "attachment_device_index",
				Description: "The device index of the network interface attachment on the instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Attachment.DeviceIndex"),
			},
			{
				Name:        "attachment_instance_id",
				Description: "The ID of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attachment.InstanceId"),
			},
			{
				Name:        "attachment_instance_owner_id",
				Description: "The Amazon Web Services account ID of the owner of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attachment.InstanceOwnerId"),
			},
			{
				Name:        "attachment_network_card_index",
				Description: "The index of the network card.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Attachment.NetworkCardIndex"),
			},
			{
				Name:        "attachment_status",
				Description: "The attachment state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attachment.Status"),
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
				Resolver:    resolveNetworkInterfacesGroups,
			},
			{
				Name:        "interface_type",
				Description: "The type of network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv4_prefixes",
				Description: "Describes an IPv4 prefix.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ipv4Prefixes.Ipv4Prefix"),
			},
			{
				Name:          "ipv6_address",
				Description:   "The IPv6 globally unique address associated with the network interface.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "ipv6_addresses",
				Description: "Describes an IPv6 address associated with a network interface.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ipv6Addresses.Ipv6Address"),
			},
			{
				Name:          "ipv6_native",
				Description:   "Indicates whether this is an IPv6 only network interface.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:        "ipv6_prefixes",
				Description: "Describes the IPv6 prefix.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ipv6Prefixes.Ipv6Prefix"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_network_interface_private_ip_addresses",
				Description: "Describes the private IPv4 address of a network interface.",
				Resolver:    schema.PathTableResolver("PrivateIpAddresses"),
				Columns: []schema.Column{
					{
						Name:        "network_interface_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_network_interfaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "association_allocation_id",
						Description: "The allocation ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.AllocationId"),
					},
					{
						Name:        "association_id",
						Description: "The association ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.AssociationId"),
					},
					{
						Name:          "association_carrier_ip",
						Description:   "The carrier IP address associated with the network interface",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.CarrierIp"),
						IgnoreInTests: true,
					},
					{
						Name:          "association_customer_owned_ip",
						Description:   "The customer-owned IP address associated with the network interface.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.CustomerOwnedIp"),
						IgnoreInTests: true,
					},
					{
						Name:        "association_ip_owner_id",
						Description: "The ID of the Elastic IP address owner.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.IpOwnerId"),
					},
					{
						Name:        "association_public_dns_name",
						Description: "The public DNS name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.PublicDnsName"),
					},
					{
						Name:        "association_public_ip",
						Description: "The address of the Elastic IP address bound to the network interface.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Association.PublicIp"),
					},
					{
						Name:        "primary",
						Description: "Indicates whether this IPv4 address is the primary private IPv4 address of the network interface.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "private_dns_name",
						Description: "The private DNS name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_ip_address",
						Description: "The private IPv4 address.",
						Type:        schema.TypeString,
					},
				},
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
			return diag.WrapError(err)
		}
		res <- output.NetworkInterfaces
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveNetworkInterfacesGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ni := resource.Item.(types.NetworkInterface)
	b, err := json.Marshal(ni.Groups)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, b))
}
