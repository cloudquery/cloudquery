package worker_meta_data

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func WorkerMetaData() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_meta_data",
		Resolver:  fetchWorkerMetaData,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&cloudflare.WorkerMetaData{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			WorkerCronTriggers(),
			WorkersSecrets(),
		},
	}
}
