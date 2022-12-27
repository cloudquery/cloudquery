package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Organizations() []*Resource {
	return []*Resource{
		{
			Service:      "organizations",
			SubService:   "organizations",
			Struct:       new(github.Organization),
			TableName:    "organizations",
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
			Relations:    []string{"Members()"},
		},
		{
			Service:    "organizations",
			SubService: "members",
			Struct:     new(github.User),
			TableName:  "organization_members",
			PKColumns:  []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{
				orgColumn, // we can use orgColumn here
				{
					Name:     "membership",
					Type:     schema.TypeJSON,
					Resolver: "resolveMembership",
				},
			},
			Multiplex: "", // we skip multiplexing here as it's a relation
		},
	}
}
