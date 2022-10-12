// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcEndpointServices() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_vpc_endpoint_services",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html",
		Resolver:    fetchEc2VpcEndpointServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVpcEndpointServiceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "acceptance_required",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AcceptanceRequired"),
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "base_endpoint_dns_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("BaseEndpointDnsNames"),
			},
			{
				Name:     "manages_vpc_endpoints",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ManagesVpcEndpoints"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "payer_responsibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PayerResponsibility"),
			},
			{
				Name:     "private_dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateDnsName"),
			},
			{
				Name:     "private_dns_name_verification_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateDnsNameVerificationState"),
			},
			{
				Name:     "private_dns_names",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateDnsNames"),
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceId"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "service_type",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceType"),
			},
			{
				Name:     "supported_ip_address_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedIpAddressTypes"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_endpoint_policy_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VpcEndpointPolicySupported"),
			},
		},
	}
}
