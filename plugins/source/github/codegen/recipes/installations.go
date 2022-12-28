package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func Installations() []*Resource {
	return []*Resource{
		{
			TableName:    "installations",
			Service:      "installations",
			SubService:   "installations",
			Struct:       new(github.Installation),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
	}
}
