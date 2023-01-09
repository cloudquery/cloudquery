package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DomainNames() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_domain_names",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html`,
		Resolver:    fetchApigatewayv2DomainNames,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.DomainName{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveDomainNameArn(),
			},
		},
		Relations: []*schema.Table{
			DomainNameRestApiMappings(),
		},
	}
}
