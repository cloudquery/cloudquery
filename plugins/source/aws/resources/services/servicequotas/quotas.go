package servicequotas

import (
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Quotas() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicequotas_quotas",
		Description: `https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceQuota.html`,
		Resolver:    fetchServicequotasQuotas,
		Transform:   transformers.TransformWithStruct(&types.ServiceQuota{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QuotaArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
