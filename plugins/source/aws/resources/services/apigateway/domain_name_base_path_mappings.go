package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DomainNameBasePathMappings() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_domain_name_base_path_mappings",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html`,
		Resolver:    fetchApigatewayDomainNameBasePathMappings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.BasePathMapping{}),
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
				Name:     "domain_name_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayDomainNameBasePathMappingArn,
			},
		},
	}
}
