package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func domainNameRestApiMappings() *schema.Table {
	tableName := "aws_apigatewayv2_domain_name_rest_api_mappings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_ApiMapping.html`,
		Resolver:    fetchApigatewayv2DomainNameRestApiMappings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.ApiMapping{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "domain_name_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveDomainNameRestApiMappingArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchApigatewayv2DomainNameRestApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.DomainName)
	config := apigatewayv2.GetApiMappingsInput{
		DomainName: r.DomainName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApiMappings(ctx, &config)

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

func resolveDomainNameRestApiMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/domainnames/%s/apimappings/%s", aws.ToString(resource.Parent.Item.(types.DomainName).DomainName), aws.ToString(resource.Item.(types.ApiMapping).ApiMappingId)),
	}.String())
}
