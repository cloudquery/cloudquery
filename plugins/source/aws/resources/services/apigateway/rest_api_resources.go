package apigateway

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func restApiResources() *schema.Table {
	tableName := "aws_apigateway_rest_api_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Resource.html`,
		Resolver:    fetchApigatewayRestApiResources,
		Transform:   transformers.TransformWithStruct(&types.Resource{}),
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
				Resolver:   resolveApigatewayRestAPIResourceArn,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			restApiResourceMethods(),
		},
	}
}

func fetchApigatewayRestApiResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApigateway).Apigateway
	config := apigateway.GetResourcesInput{RestApiId: r.Id, Limit: aws.Int32(500)}
	for p := apigateway.NewGetResourcesPaginator(svc, &config); p.HasMorePages(); {
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
func resolveApigatewayRestAPIResourceArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Resource)
	rapi := resource.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/resources/%s", aws.ToString(rapi.Id), aws.ToString(r.Id)),
	}.String())
}
