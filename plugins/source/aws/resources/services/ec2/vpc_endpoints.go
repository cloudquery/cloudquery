package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_vpc_endpoints",
		Description:   "Describes a VPC endpoint.",
		Resolver:      fetchEc2VpcEndpoints,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:     "last_error",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastError"),
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:            "id",
				Description:     "The ID of the VPC endpoint.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("VpcEndpointId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
			{
				Name:        "dns_entries",
				Description: "Describes DNS entries.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DnsEntries"),
			},
			{
				Name:        "groups",
				Description: "Describes security groups.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Groups"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
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
			return err
		}
		res <- output.VpcEndpoints
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
