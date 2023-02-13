package servicequotas

import (
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicequotas_services",
		Description: `https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html`,
		Resolver:    fetchServicequotasServices,
		Transform:   transformers.TransformWithStruct(&types.ServiceInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicequotas"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "service_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceCode"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Quotas(),
		},
	}
}
