package ec2

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
		Description:  "Describes a NAT gateway.",
		Resolver:     fetchEc2NatGateways,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
					return []string{"natgateway", *resource.Item.(types.NatGateway).NatGatewayId}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The ID of the NAT gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NatGatewayId"),
			},
			{
				Name:        "create_time",
				Description: "The date and time the NAT gateway was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "delete_time",
				Description:   "The date and time the NAT gateway was deleted, if applicable.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "failure_code",
				Description:   "If the NAT gateway could not be created, specifies the error code for the failure.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "failure_message",
				Description:   "If the NAT gateway could not be created, specifies the error message for the failure, that corresponds to the error code.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth_provision_time",
				Description:   "Reserved.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("ProvisionedBandwidth.ProvisionTime"),
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth_provisioned",
				Description:   "Reserved.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ProvisionedBandwidth.Provisioned"),
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth_request_time",
				Description:   "Reserved.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("ProvisionedBandwidth.RequestTime"),
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth_requested",
				Description:   "Reserved.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ProvisionedBandwidth.Requested"),
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth_status",
				Description:   "Reserved.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ProvisionedBandwidth.Status"),
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "The state of the NAT gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet in which the NAT gateway is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags for the NAT gateway.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2natGatewayTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the NAT gateway is located.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_nat_gateway_addresses",
				Description: "Describes the IP addresses and network interface associated with a NAT gateway.",
				Resolver:    fetchEc2NatGatewayAddresses,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"nat_gateway_cq_id", "network_interface_id"}},
				Columns: []schema.Column{
					{
						Name:        "nat_gateway_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_nat_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allocation_id",
						Description: "The allocation ID of the Elastic IP address that's associated with the NAT gateway.",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_interface_id",
						Description: "The ID of the network interface associated with the NAT gateway.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_ip",
						Description: "The private IP address associated with the Elastic IP address.",
						Type:        schema.TypeString,
					},
					{
						Name:        "public_ip",
						Description: "The Elastic IP address associated with the NAT gateway.",
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
func fetchEc2NatGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchEc2NatGatewayAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.NatGateway)
	res <- r.NatGatewayAddresses
	return nil
}
