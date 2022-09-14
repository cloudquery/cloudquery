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
			Relations:        []string{"workerCronTriggers()", "workersSecrets()"},
			ResolverFuncName: "services.FetchWorkerMetaData",
		},
		{
			CFStruct: &cloudflare.WorkerCronTrigger{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_cq_id",
					Type:     schema.TypeUUID,
					Resolver: "schema.ParentIDResolver",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_worker_cron_triggers",
			TableFuncName:    "workerCronTriggers",
			Filename:         "worker_cron_triggers.go",
			ResolverFuncName: "services.FetchWorkerCronTriggers",
		},
		{
			CFStruct: &cloudflare.WorkersSecret{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "worker_meta_data_cq_id",
					Type:     schema.TypeUUID,
					Resolver: "schema.ParentIDResolver",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_workers_secrets",
			TableFuncName:    "workersSecrets",
			Filename:         "workers_secrets.go",
			ResolverFuncName: "services.FetchWorkersSecrets",
		},
	}
}
