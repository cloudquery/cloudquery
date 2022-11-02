package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchServicequotasQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config servicequotas.ListServicesInput
	svc := meta.(*client.Client).Services().ServiceQuotas
	servicePaginator := servicequotas.NewListServicesPaginator(svc, &config)
	for servicePaginator.HasMorePages() {
		output, err := servicePaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, service := range output.Services {
			var config servicequotas.ListServiceQuotasInput
			config.ServiceCode = service.ServiceCode
			quotasPaginator := servicequotas.NewListServiceQuotasPaginator(svc, &config)
			for quotasPaginator.HasMorePages() {
				output, err := quotasPaginator.NextPage(ctx)
				if err != nil {
					return err
				}
				res <- output.Quotas
			}
		}
	}
	return nil
}
