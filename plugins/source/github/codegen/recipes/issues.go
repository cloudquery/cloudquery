package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func Issues() []*Resource {
	return []*Resource{
		{
			TableName:    "issues",
			Service:      "issues",
			SubService:   "issues",
			Struct:       new(github.Issue),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
	}
}
