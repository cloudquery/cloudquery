package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func apiIntegrationResponses() *schema.Table {
	tableName := "aws_apigatewayv2_api_integration_responses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid.html`,
		Resolver:    fetchApigatewayv2ApiIntegrationResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.IntegrationResponse{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "api_integration_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "integration_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("integration_id"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApiIntegrationResponseArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchApigatewayv2ApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Integration)
	p := parent.Parent.Item.(types.Api)
	config := apigatewayv2.GetIntegrationResponsesInput{
		ApiId:         p.ApiId,
		IntegrationId: r.IntegrationId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetIntegrationResponses(ctx, &config, func(options *apigatewayv2.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveApiIntegrationResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.IntegrationResponse)
	i := resource.Parent.Item.(types.Integration)
	api := resource.Parent.Parent.Item.(types.Api)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/apis/%s/integrations/%s/integrationresponses/%s", aws.ToString(api.ApiId), aws.ToString(i.IntegrationId), aws.ToString(r.IntegrationResponseId)),
	}.String())
}
