package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClientCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_client_certificates",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_ClientCertificate.html`,
		Resolver:    fetchApigatewayClientCertificates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.ClientCertificate{}),
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
				Resolver: resolveApigatewayClientCertificateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
