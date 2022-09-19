package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayUsagePlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetUsagePlansInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetUsagePlansPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayUsagePlanArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	up := resource.Item.(types.UsagePlan)
	arn := cl.RegionGlobalARN(client.ApigatewayService, usagePlanIDPart, *up.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayUsagePlanKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.UsagePlan)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetUsagePlanKeysInput{UsagePlanId: r.Id}
	for p := apigateway.NewGetUsagePlanKeysPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayUsagePlanKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	up := resource.Parent.Item.(types.UsagePlan)
	key := resource.Item.(types.UsagePlanKey)
	arn := cl.RegionGlobalARN(client.ApigatewayService, usagePlanIDPart, *up.Id, "keys", *key.Id)
	return resource.Set(c.Name, arn)
}
