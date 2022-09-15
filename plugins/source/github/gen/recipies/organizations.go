package recipies

import (
	"github.com/google/go-github/v45/github"
)

func Organizations() []*Resource {
	return []*Resource{
		{
			Service:      "organizations",
			SubService:   "organizations",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Organization),
			TableName:    "organizations",
			SkipFields:   append(skipID, htmlURL),
			ExtraColumns: append(orgColumns, idColumn, htmlURLCol),
			Relations:    []string{"Members()"},
		},
		{
			Service:      "organizations",
			SubService:   "members",
			Multiplex:    "", // we skip multiplexing here as it's a relation
			Struct:       new(github.User),
			TableName:    "organization_members",
			SkipFields:   append(skipID, htmlURL),
			ExtraColumns: append(orgColumns, idColumn, htmlURLCol), // we can use orgColumns here
		},
	}
}
