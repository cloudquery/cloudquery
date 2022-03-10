package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Eips() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_eips",
		Description:  "Describes an Elastic IP address, or a carrier IP address.",
		Resolver:     fetchEc2Eips,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "allocation_id"}},
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
				Name:        "allocation_id",
				Description: "The ID representing the allocation of the address for use with EC2-VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "association_id",
				Description: "The ID representing the association of the address with an instance in a VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:          "carrier_ip",
				Description:   "The carrier IP address associated",
				Type:          schema.TypeInet,
				Resolver:      schema.IPAddressResolver("CarrierIp"),
				IgnoreInTests: true,
			},
			{
				Name:          "customer_owned_ip",
				Description:   "The customer-owned IP address.",
				Type:          schema.TypeInet,
				Resolver:      schema.IPAddressResolver("CustomerOwnedIp"),
				IgnoreInTests: true,
			},
			{
				Name:          "customer_owned_ipv4_pool",
				Description:   "The ID of the customer-owned address pool.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "domain",
				Description: "Indicates whether this Elastic IP address is for use with instances in EC2-Classic (standard) or instances in a VPC (vpc).",
				Type:        schema.TypeString,
			},
			{
				Name:          "instance_id",
				Description:   "The ID of the instance that the address is associated with (if any).",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "network_border_group",
				Description: "The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_interface_id",
				Description: "The ID of the network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_interface_owner_id",
				Description: "The ID of the AWS account that owns the network interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address associated with the Elastic IP address.",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("PrivateIpAddress"),
			},
			{
				Name:        "public_ip",
				Description: "The Elastic IP address.",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("PublicIp"),
			},
			{
				Name:        "public_ipv4_pool",
				Description: "The ID of an address pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the Elastic IP address.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2eipTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2Eips(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	output, err := svc.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{}, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.Addresses
	return nil
}
func resolveEc2eipTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Address)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
