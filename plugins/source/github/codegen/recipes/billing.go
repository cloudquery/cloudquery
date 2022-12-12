package recipes

import (
	"github.com/google/go-github/v48/github"
)

func Billing() []*Resource {
	return []*Resource{
		{
			Service:      "billing",
			SubService:   "action",
			Multiplex:    orgMultiplex,
			Struct:       new(github.ActionBilling),
			ExtraColumns: orgColumns,
		},
		{
			Service:      "billing",
			SubService:   "package",
			Multiplex:    orgMultiplex,
			Struct:       new(github.PackageBilling),
			ExtraColumns: orgColumns,
		},
		{
			Service:      "billing",
			SubService:   "storage",
			Multiplex:    orgMultiplex,
			Struct:       new(github.StorageBilling),
			ExtraColumns: orgColumns,
		},
	}
}
