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

func Ec2VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_vpc_endpoints",
		Description:   "Describes a VPC endpoint.",
		Resolver:      fetchEc2VpcEndpoints,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
					return []string{"vpc-endpoint", *resource.Item.(types.VpcEndpoint).VpcEndpointId}, nil
				}),
			},
			{
				Name:        "creation_timestamp",
				Description: "The date and time that the VPC endpoint was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_error_code",
				Description: "The error code for the VPC endpoint error.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastError.Code"),
			},
			{
				Name:        "last_error_message",
				Description: "The error message for the VPC endpoint error.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastError.Message"),
			},
			{
				Name:        "network_interface_ids",
				Description: "(Interface endpoint) One or more network interfaces for the endpoint.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the VPC endpoint.",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy_document",
				Description: "The policy document associated with the endpoint, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_enabled",
				Description: "(Interface endpoint) Indicates whether the VPC is associated with a private hosted zone.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "requester_managed",
				Description: "Indicates whether the VPC endpoint is being managed by its service.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "route_table_ids",
				Description: "(Gateway endpoint) One or more route tables associated with the endpoint.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "service_name",
				Description: "The name of the service to which the endpoint is associated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the VPC endpoint.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "(Interface endpoint) One or more subnets in which the endpoint is located.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the VPC endpoint.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2vpcEndpointTags,
			},
			{
				Name:        "id",
				Description: "The ID of the VPC endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VpcEndpointId"),
			},
			{
				Name:        "vpc_endpoint_type",
				Description: "The type of endpoint.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC to which the endpoint is associated.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ec2_vpc_endpoint_dns_entries",
				Description:   "Describes a DNS entry.",
				Resolver:      fetchEc2VpcEndpointDnsEntries,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"vpc_endpoint_cq_id", "dns_name"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "vpc_endpoint_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_vpc_endpoints table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "dns_name",
						Description: "The DNS name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "hosted_zone_id",
						Description: "The ID of the private hosted zone.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_vpc_endpoint_groups",
				Description:   "Describes a security group.",
				Resolver:      fetchEc2VpcEndpointGroups,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"vpc_endpoint_cq_id", "group_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "vpc_endpoint_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_vpc_endpoints table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "group_id",
						Description: "The ID of the security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_name",
						Description: "The name of the security group.",
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
func fetchEc2VpcEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpcEndpointsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeVpcEndpoints(ctx, &config, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.VpcEndpoints
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2vpcEndpointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VpcEndpoint)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2VpcEndpointDnsEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpoint := parent.Item.(types.VpcEndpoint)
	res <- endpoint.DnsEntries

	return nil
}
func fetchEc2VpcEndpointGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	endpoint := parent.Item.(types.VpcEndpoint)
	res <- endpoint.Groups

	return nil
}
