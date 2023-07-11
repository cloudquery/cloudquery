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

func VpcLinks() *schema.Table {
	tableName := "aws_apigatewayv2_vpc_links"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/vpclinks.html#vpclinks-prop-vpclink`,
		Resolver:    fetchApigatewayv2VpcLinks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.VpcLink{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveVpcLinkArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchApigatewayv2VpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apigatewayv2.GetVpcLinksInput
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetVpcLinks(ctx, &config, func(options *apigatewayv2.Options) {
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

func resolveVpcLinkArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/vpclinks/%s", aws.ToString(resource.Item.(types.VpcLink).VpcLinkId)),
	}.String())
}
