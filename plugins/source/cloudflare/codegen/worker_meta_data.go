// Code generated by codegen using template resource_manual.go.tpl; DO NOT EDIT.

package codegen

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/services"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkerMetaData() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_worker_meta_data",
		Resolver: services.FetchWorkerMetaData,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ETAG"),
			},
			{
				Name:     "size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "created_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ModifiedOn"),
			},
		},

		Relations: []*schema.Table{
			workerCronTriggers(),
			workersSecrets(),
		},
	}
}
