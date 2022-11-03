// Code generated by codegen; DO NOT EDIT.

package servicequotas

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicequotas_services",
		Description: "https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html",
		Resolver:    fetchServicequotasServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicequotas"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
