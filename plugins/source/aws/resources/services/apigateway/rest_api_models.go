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

func restApiModels() *schema.Table {
	tableName := "aws_apigateway_rest_api_models"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html`,
		Resolver:    fetchApigatewayRestApiModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.Model{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApigatewayRestAPIModelArn,
				PrimaryKey: true,
			},
			{
				Name:     "model_template",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveApigatewayRestAPIModelModelTemplate,
			},
		},
	}
}

func fetchApigatewayRestApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway
	config := apigateway.GetModelsInput{RestApiId: r.Id, Limit: aws.Int32(500)}
	for p := apigateway.NewGetModelsPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIModelArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	m := resource.Item.(types.Model)
	rapi := resource.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/models/%s", aws.ToString(rapi.Id), aws.ToString(m.Name)),
	}.String())
}

func resolveApigatewayRestAPIModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Model)
	api := resource.Parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway

	if api.Id == nil || r.Name == nil {
		return nil
	}

	config := apigateway.GetModelTemplateInput{
		RestApiId: api.Id,
		ModelName: r.Name,
	}

	response, err := svc.GetModelTemplate(ctx, &config, func(options *apigateway.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if client.IsAWSError(err, "BadRequestException") {
			// This is an application level error and the user has nothing to do with that.
			// https://github.com/cloudquery/cq-provider-aws/pull/567#discussion_r827095787
			// The suer will be able to find incorrect configured models via
			// select * from aws_apigateway_rest_api_models where model_template is nil
			return nil
		}
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, response.Value)
}
