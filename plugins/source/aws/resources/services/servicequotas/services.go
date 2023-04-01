package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	tableName := "aws_servicequotas_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html`,
		Resolver:    fetchServicequotasServices,
		Transform:   transformers.TransformWithStruct(&types.ServiceInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicequotas"),
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
			quotas(),
		},
	}
}

func fetchServicequotasServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := servicequotas.ListServicesInput{
		MaxResults: aws.Int32(100),
	}

	svc := meta.(*client.Client).Services().Servicequotas
	servicePaginator := servicequotas.NewListServicesPaginator(svc, &config)
	for servicePaginator.HasMorePages() {
		output, err := servicePaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Services
	}
	return nil
}
