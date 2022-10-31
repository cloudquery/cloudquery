package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func Actions() []*Resource {
	return []*Resource{
		{
			Service:    "actions",
			SubService: "workflows",
			Multiplex:  orgMultiplex,
			Struct:     new(github.Workflow),
			TableName:  "workflows",
			SkipFields: skipID,
			ExtraColumns: append(orgColumns, idColumn,
				codegen.ColumnDefinition{
					Name:     "contents",
					Type:     schema.TypeJSON,
					Resolver: `resolveContents`,
				},
			),
		},
	}
}
