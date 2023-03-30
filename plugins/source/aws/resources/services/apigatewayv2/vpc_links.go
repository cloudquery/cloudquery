package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcLinks() *schema.Table {
	tableName := "aws_apigatewayv2_vpc_links"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html`,
		Resolver:    fetchApigatewayv2VpcLinks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&types.VpcLink{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVpcLinkArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
