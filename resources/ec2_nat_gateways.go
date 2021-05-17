package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2NatGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_nat_gateways",
		Resolver:     fetchEc2NatGateways,
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
				Name: "create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "delete_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "failure_code",
				Type: schema.TypeString,
			},
			{
				Name: "failure_message",
				Type: schema.TypeString,
			},
			{
				Name: "nat_gateway_id",
				Type: schema.TypeString,
			},
			{
				Name:     "provisioned_bandwidth_provision_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ProvisionedBandwidth.ProvisionTime"),
			},
			{
				Name:     "provisioned_bandwidth_provisioned",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisionedBandwidth.Provisioned"),
			},
			{
				Name:     "provisioned_bandwidth_request_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ProvisionedBandwidth.RequestTime"),
			},
			{
				Name:     "provisioned_bandwidth_requested",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisionedBandwidth.Requested"),
			},
			{
				Name:     "provisioned_bandwidth_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisionedBandwidth.Status"),
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2natGatewayTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_nat_gateway_addresses",
				Resolver: fetchEc2NatGatewayAddresses,
				Columns: []schema.Column{
					{
						Name:     "nat_gateway_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "allocation_id",
						Type: schema.TypeString,
					},
					{
						Name: "network_interface_id",
						Type: schema.TypeString,
					},
					{
						Name: "private_ip",
						Type: schema.TypeString,
					},
					{
						Name: "public_ip",
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
func fetchEc2NatGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeNatGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeNatGateways(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.NatGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil

}
func resolveEc2natGatewayTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.NatGateway)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2NatGatewayAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.NatGateway)
	res <- r.NatGatewayAddresses
	return nil
}
