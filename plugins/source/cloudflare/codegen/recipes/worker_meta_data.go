package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkerMetaDataResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns:           []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:              "client.AccountMultiplex",
			DataStruct:             &cloudflare.WorkerMetaData{},
			PKColumns:              []string{"id"},
			Service:                "worker_meta_data",
			Relations:              []string{"WorkerCronTriggers()", "WorkersSecrets()"},
			SkipServiceInTableName: true,
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
			Service:                "worker_meta_data",
			SkipServiceInTableName: true,
			SkipParentInTableName:  true,
			//TableName: "worker_cron_triggers",
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
			Service:                "worker_meta_data",
			SkipServiceInTableName: true,
			SkipParentInTableName:  true,
			//TableName: "workers_secrets",
		},
	}
}
