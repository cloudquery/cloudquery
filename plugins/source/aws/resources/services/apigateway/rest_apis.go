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

func RestApis() *schema.Table {
	tableName := "aws_apigateway_rest_apis"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html`,
		Resolver:    fetchApigatewayRestApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.RestApi{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApigatewayRestAPIArn,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			restApiAuthorizers(),
			restApiDeployments(),
			restApiDocumentationParts(),
			restApiDocumentationVersions(),
			restApiGatewayResponses(),
			restApiModels(),
			restApiRequestValidators(),
			restApiResources(),
			restApiStages(),
		},
	}
}

func fetchApigatewayRestApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := apigateway.GetRestApisInput{
		Limit: aws.Int32(500),
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway
	for p := apigateway.NewGetRestApisPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	rapi := resource.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s", aws.ToString(rapi.Id)),
	}.String())
}
