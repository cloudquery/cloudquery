package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func WorkerRouteResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:        "client.ZoneMultiplex",
			CFStruct:         &cloudflare.WorkerRoute{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_worker_routes",
			TableFuncName:    "WorkerRoutes",
			Filename:         "worker_routes.go",
			Package:          "worker_routes",
			ResolverFuncName: "fetchWorkerRoutes",
		},
	}
}
