package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Apigatewayv2VpcLinks() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_vpc_links",
		Description:  "Represents a VPC link.",
		Resolver:     fetchApigatewayv2VpcLinks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpclinks", *resource.Item.(types.VpcLink).VpcLinkId}, nil
				}),
			},
			{
				Name:        "name",
				Description: "The name of the VPC link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_group_ids",
				Description: "A list of security group IDs for the VPC link.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "subnet_ids",
				Description: "A list of subnet IDs to include in the VPC link.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "The ID of the VPC link.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VpcLinkId"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the VPC link was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "tags",
				Description: "Tags for the VPC link.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "vpc_link_status",
				Description: "The status of the VPC link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_link_status_message",
				Description: "A message summarizing the cause of the status of the VPC link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_link_version",
				Description: "The version of the VPC link.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayv2VpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigatewayv2.GetVpcLinksInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetVpcLinks(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
