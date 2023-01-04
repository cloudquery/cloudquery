package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Teams() []*Resource {
	teamID := codegen.ColumnDefinition{
		Name:     "team_id",
		Type:     schema.TypeInt,
		Resolver: `client.ResolveParentColumn("ID")`,
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}

	repos := repository()
	repos.Service = "teams"
	repos.TableName = "team_repositories"
	repos.SkipFields = append(repos.SkipFields, "TeamID")
	repos.ExtraColumns = append(repos.ExtraColumns, teamID)

	return []*Resource{
		{
			TableName:    "teams",
			Service:      "teams",
			SubService:   "teams",
			Struct:       new(github.Team),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
			Relations:    []string{"Members()", "Repositories()"},
		},
		{
			TableName:  "team_members",
			Service:    "teams",
			SubService: "members",
			Struct:     new(github.User),
			PKColumns:  []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{
				orgColumn,
				teamID,
				codegen.ColumnDefinition{
					Name:     "membership",
					Type:     schema.TypeJSON,
					Resolver: "resolveMembership",
				},
			},
			Multiplex: "", // we skip multiplexing here as it's a relation
		},
		repos,
	}
}
