package recipies

import (
	"github.com/google/go-github/v45/github"
)

func External() []*Resource {
	return []*Resource{
		{
			Service:      "external",
			SubService:   "group",
			Struct:       new(github.ExternalGroup),
			SkipFields:   []string{"UpdatedAt"},
			ExtraColumns: append(orgColumns, timestampField("updated_at", "UpdatedAt")),
		},
	}
}
