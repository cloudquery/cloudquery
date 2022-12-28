package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Actions() []*Resource {
	return []*Resource{
		{
			TableName:  "workflows",
			Service:    "actions",
			SubService: "workflows",
			Struct:     new(github.Workflow),
			PKColumns:  []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{
				orgColumn,
				{
					Name:     "contents",
					Type:     schema.TypeString,
					Resolver: `resolveContents`,
				},
			},
			Multiplex: orgMultiplex,
		},
	}
}
