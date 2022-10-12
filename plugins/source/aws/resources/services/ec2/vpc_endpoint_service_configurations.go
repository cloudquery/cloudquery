// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcEndpointServiceConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_vpc_endpoint_service_configurations",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceConfiguration.html",
		Resolver:    fetchEc2VpcEndpointServiceConfigurations,
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
				Resolver: resolveVpcEndpointServiceConfigurationArn,
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
				Name:     "gateway_load_balancer_arns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("GatewayLoadBalancerArns"),
			},
			{
				Name:     "manages_vpc_endpoints",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ManagesVpcEndpoints"),
			},
			{
				Name:     "network_load_balancer_arns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NetworkLoadBalancerArns"),
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
				Name:     "private_dns_name_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateDnsNameConfiguration"),
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
				Name:     "service_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceState"),
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
		},
	}
}
