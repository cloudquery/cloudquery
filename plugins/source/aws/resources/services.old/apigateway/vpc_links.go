package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource vpc_links --config vpc_links.hcl --output .
func VpcLinks() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_vpc_links",
		Description:  "An API Gateway VPC link for a RestApi to access resources in an Amazon Virtual Private Cloud (VPC)",
		Resolver:     fetchApigatewayVpcLinks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayVpcLinkArn,
			},
			{
				Name:        "description",
				Description: "The description of the VPC link",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier of the VpcLink",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name used to label and identify the VPC link",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the VPC link",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "A description about the VPC link status",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "target_arns",
				Description: "The ARN of the network load balancer of the VPC targeted by the VPC link",
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
	paginator := apigateway.NewGetVpcLinksPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayVpcLinkArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	link := resource.Item.(types.VpcLink)
	arn := cl.RegionGlobalARN(client.ApigatewayService, "/vpclinks", *link.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
}
