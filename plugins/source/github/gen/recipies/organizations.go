package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
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
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
			Relations:    []string{"Members()"},
		},
		{
			Service:    "organizations",
			SubService: "members",
			Multiplex:  "", // we skip multiplexing here as it's a relation
			Struct:     new(github.User),
			TableName:  "organization_members",
			SkipFields: skipID,
			ExtraColumns: append(orgColumns, idColumn, // we can use orgColumns here
				codegen.ColumnDefinition{
					Name:     "membership",
					Type:     schema.TypeJSON,
					Resolver: "resolveMembership",
				},
			),
		},
	}
}
