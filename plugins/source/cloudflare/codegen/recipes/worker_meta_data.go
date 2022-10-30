package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkerMetaDataResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			CFStruct:         &cloudflare.WorkerMetaData{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_worker_meta_data",
			TableFuncName:    "WorkerMetaData",
			Filename:         "worker_meta_data.go",
			Package:          "worker_meta_data",
			Relations:        []string{"workerCronTriggers()", "workersSecrets()"},
			ResolverFuncName: "fetchWorkerMetaData",
		},
		{
			CFStruct: &cloudflare.WorkerCronTrigger{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_worker_cron_triggers",
			TableFuncName:    "workerCronTriggers",
			Filename:         "worker_cron_triggers.go",
			Package:          "worker_meta_data",
			ResolverFuncName: "fetchWorkerCronTriggers",
		},
		{
			CFStruct: &cloudflare.WorkersSecret{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_workers_secrets",
			TableFuncName:    "workersSecrets",
			Filename:         "workers_secrets.go",
			Package:          "worker_meta_data",
			ResolverFuncName: "fetchWorkersSecrets",
		},
	}
}
