package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayVpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetVpcLinksInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	paginator := apigateway.NewGetVpcLinksPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayVpcLinkArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	link := resource.Item.(types.VpcLink)
	arn := cl.RegionGlobalARN(client.ApigatewayService, "/vpclinks", *link.Id)
	return resource.Set(c.Name, arn)
}
