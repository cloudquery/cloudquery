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

func apiModels() *schema.Table {
	tableName := "aws_apigatewayv2_api_models"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-models.html`,
		Resolver:    fetchApigatewayv2ApiModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.Model{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "api_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "api_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApiModelArn,
				PrimaryKey: true,
			},
			{
				Name:     "model_template",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveApigatewayv2apiModelModelTemplate,
			},
		},
	}
}
func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetModelsInput{
		ApiId: r.ApiId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetModels(ctx, &config, func(options *apigatewayv2.Options) {
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

func resolveApigatewayv2apiModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Model)
	p := resource.Parent.Item.(types.Api)
	config := apigatewayv2.GetModelTemplateInput{
		ApiId:   p.ApiId,
		ModelId: r.ModelId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	response, err := svc.GetModelTemplate(ctx, &config, func(options *apigatewayv2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Value)
}

func resolveApiModelArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Model)
	p := resource.Parent.Item.(types.Api)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/apis/%s/models/%s", aws.ToString(p.ApiId), aws.ToString(r.ModelId)),
	}.String())
}
