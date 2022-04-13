package ec2

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource egress_only_internet_gateways --config gen.hcl --output .
func EgressOnlyInternetGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_egress_only_internet_gateways",
		Description:  "Describes an egress-only internet gateway.",
		Resolver:     fetchEc2EgressOnlyInternetGateways,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Description: "The Amazon Resource Name (ARN) for the egress-only internet gateway.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"egress-only-internet-gateway", *resource.Item.(types.EgressOnlyInternetGateway).EgressOnlyInternetGatewayId}, nil
				}),
			},
			{
				Name:        "attachments",
				Description: "Information about the attachment of the egress-only internet gateway.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEgressOnlyInternetGatewaysAttachments,
			},
			{
				Name:        "id",
				Description: "The ID of the egress-only internet gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EgressOnlyInternetGatewayId"),
			},
			{
				Name:        "tags",
				Description: "The tags assigned to the egress-only internet gateway.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2EgressOnlyInternetGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	input := ec2.DescribeEgressOnlyInternetGatewaysInput{}
	for {
		output, err := svc.DescribeEgressOnlyInternetGateways(ctx, &input, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.EgressOnlyInternetGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveEgressOnlyInternetGatewaysAttachments(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	egress := resource.Item.(types.EgressOnlyInternetGateway)
	b, err := json.Marshal(egress.Attachments)
	if err != nil {
		return diag.WrapError(err)
	}

	return resource.Set(c.Name, b)
}
