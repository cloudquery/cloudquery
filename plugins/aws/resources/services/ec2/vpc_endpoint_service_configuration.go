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

func Ec2VpcEndpointServiceConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpc_endpoint_service_configurations",
		Description:  "Describes a service configuration for a VPC endpoint service.",
		Resolver:     fetchEc2VpcEndpointServiceConfigurations,
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpc-endpoint-service-configuration", *resource.Item.(types.ServiceConfiguration).ServiceId}, nil
				}),
			},
			{
				Name:        "acceptance_required",
				Description: "Indicates whether requests from other AWS accounts to create an endpoint to the service must first be accepted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "availability_zones",
				Description: "The Availability Zones in which the service is available.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "base_endpoint_dns_names",
				Description: "The DNS names for the service.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "gateway_load_balancer_arns",
				Description:   "The Amazon Resource Names (ARNs) of the Gateway Load Balancers for the service.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "manages_vpc_endpoints",
				Description: "Indicates whether the service manages its VPC endpoints.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "network_load_balancer_arns",
				Description: "The Amazon Resource Names (ARNs) of the Network Load Balancers for the service.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "payer_responsibility",
				Description: "The payer responsibility.",
				Type:        schema.TypeString,
			},
			{
				Name:          "private_dns_name",
				Description:   "The private DNS name for the service.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "private_dns_name_configuration_name",
				Description:   "The name of the record subdomain the service provider needs to create.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PrivateDnsNameConfiguration.Name"),
				IgnoreInTests: true,
			},
			{
				Name:        "private_dns_name_configuration_state",
				Description: "The verification state of the VPC endpoint service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateDnsNameConfiguration.State"),
			},
			{
				Name:          "private_dns_name_configuration_type",
				Description:   "The endpoint service verification type, for example TXT.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PrivateDnsNameConfiguration.Type"),
				IgnoreInTests: true,
			},
			{
				Name:          "private_dns_name_configuration_value",
				Description:   "The value the service provider adds to the private DNS name domain record before verification.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PrivateDnsNameConfiguration.Value"),
				IgnoreInTests: true,
			},
			{
				Name:        "service_id",
				Description: "The ID of the service.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_name",
				Description: "The name of the service.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_state",
				Description: "The service state.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_type",
				Description: "The type of service.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ServiceType.ServiceType"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the service.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2VpcEndpointServiceConfigurations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpcEndpointServiceConfigurationsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeVpcEndpointServiceConfigurations(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.ServiceConfigurations
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
