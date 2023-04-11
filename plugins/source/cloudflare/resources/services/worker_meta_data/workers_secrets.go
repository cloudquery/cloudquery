package worker_meta_data

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func WorkersSecrets() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_workers_secrets",
		Resolver:  fetchWorkersSecrets,
		Transform: transformers.TransformWithStruct(&cloudflare.WorkersSecret{}),
		Columns: []schema.Column{
			{
				Name:     "worker_meta_data_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
