package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func Billing() []*Resource {
	return []*Resource{
		{
			Service:      "billing",
			SubService:   "action",
			Struct:       new(github.ActionBilling),
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
		{
			Service:      "billing",
			SubService:   "package",
			Struct:       new(github.PackageBilling),
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
		{
			Service:      "billing",
			SubService:   "storage",
			Struct:       new(github.StorageBilling),
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
		},
	}
}
