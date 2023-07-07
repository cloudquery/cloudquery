package worker_meta_data

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func WorkersSecrets() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_workers_secrets",
		Resolver:  fetchWorkersSecrets,
		Transform: client.TransformWithStruct(&cloudflare.WorkersSecret{}),
		Columns: []schema.Column{
			{
				Name:     "worker_meta_data_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
