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

func Ec2VpcPeeringConnections() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_vpc_peering_connections",
		Description:   "Describes a VPC peering connection.",
		Resolver:      fetchEc2VpcPeeringConnections,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
		IgnoreInTests: true,
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
					return []string{"vpc-peering-connection", *resource.Item.(types.VpcPeeringConnection).VpcPeeringConnectionId}, nil
				}),
			},
			{
				Name:        "accepter_cidr_block",
				Description: "The IPv4 CIDR block for the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccepterVpcInfo.CidrBlock"),
			},
			{
				Name:        "accepter_cidr_block_set",
				Description: "Information about the IPv4 CIDR blocks for the VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("AccepterVpcInfo.CidrBlockSet.CidrBlock"),
			},
			{
				Name:        "accepter_ipv6_cidr_block_set",
				Description: "The IPv6 CIDR block for the VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("AccepterVpcInfo.Ipv6CidrBlockSet.Ipv6CidrBlock"),
			},
			{
				Name:        "accepter_owner_id",
				Description: "The AWS account ID of the VPC owner.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccepterVpcInfo.OwnerId"),
			},
			{
				Name:        "accepter_allow_dns_resolution_from_remote_vpc",
				Description: "Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccepterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc"),
			},
			{
				Name:        "accepter_allow_egress_local_classic_link_to_remote_vpc",
				Description: "Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc"),
			},
			{
				Name:        "accepter_allow_egress_local_vpc_to_remote_classic_link",
				Description: "Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccepterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink"),
			},
			{
				Name:        "accepter_vpc_region",
				Description: "The Region in which the VPC is located.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccepterVpcInfo.Region"),
			},
			{
				Name:        "accepter_vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccepterVpcInfo.VpcId"),
			},
			{
				Name:        "expiration_time",
				Description: "The time that an unaccepted VPC peering connection will expire.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "requester_cidr_block",
				Description: "The IPv4 CIDR block for the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RequesterVpcInfo.CidrBlock"),
			},
			{
				Name:        "requester_cidr_block_set",
				Description: "Information about the IPv4 CIDR blocks for the VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("RequesterVpcInfo.CidrBlockSet.CidrBlock"),
			},
			{
				Name:        "requester_ipv6_cidr_block_set",
				Description: "The IPv6 CIDR block for the VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("RequesterVpcInfo.Ipv6CidrBlockSet.Ipv6CidrBlock"),
			},
			{
				Name:        "requester_owner_id",
				Description: "The AWS account ID of the VPC owner.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RequesterVpcInfo.OwnerId"),
			},
			{
				Name:        "requester_allow_dns_resolution_from_remote_vpc",
				Description: "Indicates whether a local VPC can resolve public DNS hostnames to private IP addresses when queried from instances in a peer VPC.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RequesterVpcInfo.PeeringOptions.AllowDnsResolutionFromRemoteVpc"),
			},
			{
				Name:        "requester_allow_egress_local_classic_link_to_remote_vpc",
				Description: "Indicates whether a local ClassicLink connection can communicate with the peer VPC over the VPC peering connection.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalClassicLinkToRemoteVpc"),
			},
			{
				Name:        "requester_allow_egress_local_vpc_to_remote_classic_link",
				Description: "Indicates whether a local VPC can communicate with a ClassicLink connection in the peer VPC over the VPC peering connection.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RequesterVpcInfo.PeeringOptions.AllowEgressFromLocalVpcToRemoteClassicLink"),
			},
			{
				Name:        "requester_vpc_region",
				Description: "The Region in which the VPC is located.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RequesterVpcInfo.Region"),
			},
			{
				Name:        "requester_vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RequesterVpcInfo.VpcId"),
			},
			{
				Name:        "status_code",
				Description: "The status of the VPC peering connection.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Code"),
			},
			{
				Name:        "status_message",
				Description: "A message that provides more information about the status, if applicable.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Message"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the resource.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "id",
				Description: "The ID of the VPC peering connection.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VpcPeeringConnectionId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
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
			return diag.WrapError(err)
		}
		res <- output.VpcPeeringConnections
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
