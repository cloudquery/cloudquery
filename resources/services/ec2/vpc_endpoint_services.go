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

func Ec2VpcEndpointServices() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpc_endpoint_services",
		Description:  "Describes a VPC endpoint service.",
		Resolver:     fetchEc2VpcEndpointServices,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
					return []string{"vpc-endpoint-service", *resource.Item.(types.ServiceDetail).ServiceId}, nil
				}),
			},
			{
				Name:        "acceptance_required",
				Description: "Indicates whether VPC endpoint connection requests to the service must be accepted by the service owner.",
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
				Name:        "manages_vpc_endpoints",
				Description: "Indicates whether the service manages its VPC endpoints.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "owner",
				Description: "The Amazon Web Services account ID of the service owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "payer_responsibility",
				Description: "The payer responsibility.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_name",
				Description: "The private DNS name for the service.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_name_verification_state",
				Description: "The verification state of the VPC endpoint service.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_names",
				Description: "The private DNS names assigned to the VPC endpoint service.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveEc2VpcEndpointServicePrivateDnsNames,
			},
			{
				Name:        "id",
				Description: "The ID of the endpoint service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceId"),
			},
			{
				Name:        "service_name",
				Description: "The Amazon Resource Name (ARN) of the service.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_type",
				Description: "The type of service.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveEc2VpcEndpointServiceServiceType,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the service.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2VpcEndpointServiceTags,
			},
			{
				Name:        "vpc_endpoint_policy_supported",
				Description: "Indicates whether the service supports endpoint policies.",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2VpcEndpointServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpcEndpointServicesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeVpcEndpointServices(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.ServiceDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2VpcEndpointServicePrivateDnsNames(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ServiceDetail)
	pdn := make([]string, 0, len(r.PrivateDnsNames))
	for _, n := range r.PrivateDnsNames {
		pdn = append(pdn, *n.PrivateDnsName)
	}
	return diag.WrapError(resource.Set(c.Name, pdn))
}
func resolveEc2VpcEndpointServiceServiceType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ServiceDetail)
	st := make([]string, 0, len(r.ServiceType))
	for _, std := range r.ServiceType {
		st = append(st, string(std.ServiceType))
	}
	return diag.WrapError(resource.Set(c.Name, st))
}
func resolveEc2VpcEndpointServiceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ServiceDetail)
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(r.Tags)))
}
