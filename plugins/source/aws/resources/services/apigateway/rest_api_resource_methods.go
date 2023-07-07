package apigateway

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func restApiResourceMethods() *schema.Table {
	tableName := "aws_apigateway_rest_api_resource_methods"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html`,
		Resolver:    fetchApigatewayRestApiResourceMethods,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&apigateway.GetMethodOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("rest_api_arn"),
			},
			{
				Name:     "resource_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApigatewayRestAPIResourceMethodArn,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			restApiResourceMethodIntegrations(),
		},
	}
}

func fetchApigatewayRestApiResourceMethods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	api := parent.Parent.Item.(types.RestApi)
	resource := parent.Item.(types.Resource)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway
	for method := range resource.ResourceMethods {
		config := apigateway.GetMethodInput{RestApiId: api.Id, ResourceId: resource.Id, HttpMethod: aws.String(method)}
		resp, err := svc.GetMethod(ctx, &config, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp
	}
	return nil
}

func resolveApigatewayRestAPIResourceMethodArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Parent.Item.(types.Resource)
	method := resource.Item.(*apigateway.GetMethodOutput)
	rapi := resource.Parent.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/resources/%s/methods/%s", aws.ToString(rapi.Id), aws.ToString(r.Id), aws.ToString(method.HttpMethod)),
	}.String())
}
