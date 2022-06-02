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

func Ec2InternetGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_internet_gateways",
		Description:  "Describes an internet gateway.",
		Resolver:     fetchEc2InternetGateways,
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
					return []string{"internet-gateway", *resource.Item.(types.InternetGateway).InternetGatewayId}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The ID of the internet gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InternetGatewayId"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the internet gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the internet gateway.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2internetGatewayTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_internet_gateway_attachments",
				Description: "Describes the attachment of a VPC to an internet gateway or an egress-only internet gateway.",
				Resolver:    fetchEc2InternetGatewayAttachments,
				Columns: []schema.Column{
					{
						Name:        "internet_gateway_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_internet_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "state",
						Description: "The current state of the attachment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_id",
						Description: "The ID of the VPC.",
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
func fetchEc2InternetGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeInternetGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeInternetGateways(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.InternetGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2internetGatewayTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.InternetGateway)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return diag.WrapError(resource.Set("tags", tags))
}
func fetchEc2InternetGatewayAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.InternetGateway)
	res <- r.Attachments
	return nil
}
