package recipes

import (
	"github.com/google/go-github/v45/github"
)

func Actions() []*Resource {
	return []*Resource{
		{
			Service:      "actions",
			SubService:   "workflows",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Workflow),
			TableName:    "workflows",
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
		},
	}
}
