package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayVpcLinks() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_vpc_links",
		Description:  "An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC).",
		Resolver:     fetchApigatewayVpcLinks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
					return []string{"/vpclinks", *resource.Item.(types.VpcLink).Id}, nil
				}),
			},
			{
				Name:        "description",
				Description: "The description of the VPC link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier of the VpcLink. It is used in an Integration to reference this VpcLink.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name used to label and identify the VPC link.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING, or FAILED. Deploying an API will wait if the status is PENDING and will fail if the status is DELETING.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "A description about the VPC link status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "target_arns",
				Description: "The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayVpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetVpcLinksInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetVpcLinks(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
