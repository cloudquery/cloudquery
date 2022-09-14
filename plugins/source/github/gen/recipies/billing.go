package recipies

import (
	"github.com/google/go-github/v45/github"
)

func Billing() []*Resource {
	return []*Resource{
		{
			Service:      "billing",
			SubService:   "action",
			Struct:       new(github.ActionBilling),
			ExtraColumns: defaultOrgColumns,
		},
		{
			Service:      "billing",
			SubService:   "package",
			Struct:       new(github.PackageBilling),
			ExtraColumns: defaultOrgColumns,
		},
		{
			Service:      "billing",
			SubService:   "storage",
			Struct:       new(github.StorageBilling),
			ExtraColumns: defaultOrgColumns,
		},
	}
}
