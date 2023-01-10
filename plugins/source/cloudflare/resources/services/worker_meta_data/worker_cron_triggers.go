package worker_meta_data

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkerCronTriggers() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_cron_triggers",
		Resolver:  fetchWorkerCronTriggers,
		Transform: transformers.TransformWithStruct(&cloudflare.WorkerCronTrigger{}),
		Columns: []schema.Column{
			{
				Name:     "worker_meta_data_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
