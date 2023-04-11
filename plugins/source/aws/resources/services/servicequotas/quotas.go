package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func quotas() *schema.Table {
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

func fetchServicequotasQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Servicequotas
	service := parent.Item.(types.ServiceInfo)
	config := servicequotas.ListServiceQuotasInput{
		ServiceCode: service.ServiceCode,
		MaxResults:  aws.Int32(100),
	}
	quotasPaginator := servicequotas.NewListServiceQuotasPaginator(svc, &config)
	for quotasPaginator.HasMorePages() {
		output, err := quotasPaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Quotas
	}
	return nil
}
