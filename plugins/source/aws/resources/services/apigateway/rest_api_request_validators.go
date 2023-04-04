package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func restApiRequestValidators() *schema.Table {
	tableName := "aws_apigateway_rest_api_request_validators"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RequestValidator.html`,
		Resolver:    fetchApigatewayRestApiRequestValidators,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.RequestValidator{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIRequestValidatorArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchApigatewayRestApiRequestValidators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetRequestValidatorsInput{RestApiId: r.Id, Limit: aws.Int32(500)}
	for {
		response, err := svc.GetRequestValidators(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
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
func resolveApigatewayRestAPIRequestValidatorArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.RequestValidator)
	rapi := resource.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/requestvalidators/%s", aws.ToString(rapi.Id), aws.ToString(r.Id)),
	}.String())
}
