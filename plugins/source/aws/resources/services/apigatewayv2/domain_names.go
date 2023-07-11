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

func DomainNames() *schema.Table {
	tableName := "aws_apigatewayv2_domain_names"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames.html`,
		Resolver:    fetchApigatewayv2DomainNames,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.DomainName{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveDomainNameArn,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			domainNameRestApiMappings(),
		},
	}
}

func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config apigatewayv2.GetDomainNamesInput
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(options *apigatewayv2.Options) {
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

func resolveDomainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/domainnames/%s", aws.ToString(resource.Item.(types.DomainName).DomainName)),
	}.String())
}
