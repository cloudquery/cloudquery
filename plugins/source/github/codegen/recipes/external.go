package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func External() []*Resource {
	return []*Resource{
		{
			Service:      "external",
			SubService:   "groups",
			Struct:       new(github.ExternalGroup),
			PKColumns:    []string{"group_id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
	}
}
