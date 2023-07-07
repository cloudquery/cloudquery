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

func domainNameBasePathMappings() *schema.Table {
	tableName := "aws_apigateway_domain_name_base_path_mappings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html`,
		Resolver:    fetchApigatewayDomainNameBasePathMappings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.BasePathMapping{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "domain_name_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApigatewayDomainNameBasePathMappingArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchApigatewayDomainNameBasePathMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.DomainName)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway
	config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName, Limit: aws.Int32(500)}
	for p := apigateway.NewGetBasePathMappingsPaginator(svc, &config); p.HasMorePages(); {
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

func resolveApigatewayDomainNameBasePathMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	domain := resource.Parent.Item.(types.DomainName)
	mapping := resource.Item.(types.BasePathMapping)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/domainnames/%s/basepathmappings/%s", aws.ToString(domain.DomainName), aws.ToString(mapping.BasePath)),
	}.String())
}
