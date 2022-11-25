package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkerMetaDataResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns:     []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			DataStruct:       &cloudflare.WorkerMetaData{},
			PKColumns:        []string{"id"},
			TableName:        "cloudflare_worker_meta_data",
			TableFuncName:    "WorkerMetaData",
			Filename:         "worker_meta_data.go",
			Service:          "worker_meta_data",
			Relations:        []string{"workerCronTriggers()", "workersSecrets()"},
			ResolverFuncName: "fetchWorkerMetaData",
		},
		{
			DataStruct: &cloudflare.WorkerCronTrigger{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			TableName:        "cloudflare_worker_cron_triggers",
			TableFuncName:    "workerCronTriggers",
			Filename:         "worker_cron_triggers.go",
			Service:          "worker_meta_data",
			ResolverFuncName: "fetchWorkerCronTriggers",
		},
		{
			DataStruct: &cloudflare.WorkersSecret{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			TableName:        "cloudflare_workers_secrets",
			TableFuncName:    "workersSecrets",
			Filename:         "workers_secrets.go",
			Service:          "worker_meta_data",
			ResolverFuncName: "fetchWorkersSecrets",
		},
	}
}
