package worker_meta_data

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
