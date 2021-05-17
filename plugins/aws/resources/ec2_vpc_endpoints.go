package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpc_endpoints",
		Resolver:     fetchEc2VpcEndpoints,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "creation_timestamp",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "last_error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastError.Code"),
			},
			{
				Name:     "last_error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastError.Message"),
			},
			{
				Name: "network_interface_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name: "policy_document",
				Type: schema.TypeString,
			},
			{
				Name: "private_dns_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "requester_managed",
				Type: schema.TypeBool,
			},
			{
				Name: "route_table_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "service_name",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2vpcEndpointTags,
			},
			{
				Name: "vpc_endpoint_id",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_endpoint_type",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_vpc_endpoint_dns_entries",
				Resolver: fetchEc2VpcEndpointDnsEntries,
				Columns: []schema.Column{
					{
						Name:     "vpc_endpoint_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "dns_name",
						Type: schema.TypeString,
					},
					{
						Name: "hosted_zone_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_vpc_endpoint_groups",
				Resolver: fetchEc2VpcEndpointGroups,
				Columns: []schema.Column{
					{
						Name:     "vpc_endpoint_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "group_id",
						Type: schema.TypeString,
					},
					{
						Name: "group_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2VpcEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveEc2vpcEndpointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VpcEndpoint)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2VpcEndpointDnsEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	endpoint, ok := parent.Item.(types.VpcEndpoint)
	if !ok {
		return fmt.Errorf("not vpc endpoint")
	}
	res <- endpoint.DnsEntries

	return nil
}
func fetchEc2VpcEndpointGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	endpoint, ok := parent.Item.(types.VpcEndpoint)
	if !ok {
		return fmt.Errorf("not vpc endpoint")
	}
	res <- endpoint.Groups

	return nil
}
