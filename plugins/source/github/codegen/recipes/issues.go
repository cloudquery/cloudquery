package recipes

import (
	"github.com/google/go-github/v48/github"
)

func Issues() []*Resource {
	return []*Resource{
		{
			Service:      "issues",
			SubService:   "issues",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Issue),
			TableName:    "issues",
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
		},
	}
}
