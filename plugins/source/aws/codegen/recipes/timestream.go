package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TimestreamResources() []*Resource {
	return []*Resource{
		{
			Service:    "timestream",
			SubService: "databases",
			Struct:     new(types.Database),
			Multiplex:  `client.ServiceAccountRegionMultiplexer("ingest.timestream")`,
			PKColumns:  []string{"arn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: "fetchDatabaseTags",
				},
			),
			Relations: []string{"Tables()"},
		},
		{
			Service:      "timestream",
			SubService:   "tables",
			Struct:       new(types.Table),
			Multiplex:    "", // skip multiplexing as it's a relation for databases
			PKColumns:    []string{"arn"},
			ExtraColumns: defaultRegionalColumns,
		},
	}
}
