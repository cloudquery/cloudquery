package worker_meta_data

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkerMetaData() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_meta_data",
		Resolver:  fetchWorkerMetaData,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&cloudflare.WorkerMetaData{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			WorkerCronTriggers(),
			WorkersSecrets(),
		},
	}
}
