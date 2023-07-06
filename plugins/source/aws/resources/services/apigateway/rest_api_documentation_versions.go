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

func restApiDocumentationVersions() *schema.Table {
	tableName := "aws_apigateway_rest_api_documentation_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationVersion.html`,
		Resolver:    fetchApigatewayRestApiDocumentationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.DocumentationVersion{}),
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
				Resolver:   resolveApigatewayRestAPIDocumentationVersionArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchApigatewayRestApiDocumentationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway
	config := apigateway.GetDocumentationVersionsInput{RestApiId: r.Id, Limit: aws.Int32(500)}
	// No paginator available
	for {
		response, err := svc.GetDocumentationVersions(ctx, &config, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
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
func resolveApigatewayRestAPIDocumentationVersionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	v := resource.Item.(types.DocumentationVersion)
	rapi := resource.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/documentation/versions/%s", aws.ToString(rapi.Id), aws.ToString(v.Version)),
	}.String())
}
