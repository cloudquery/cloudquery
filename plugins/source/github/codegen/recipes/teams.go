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
			Service:      "teams",
			SubService:   "teams",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Team),
			TableName:    "teams",
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
			Relations:    []string{"Members()", "Repositories()"},
		},
		{
			Service:    "teams",
			SubService: "members",
			Multiplex:  "", // we skip multiplexing here as it's a relation
			Struct:     new(github.User),
			TableName:  "team_members",
			SkipFields: skipID,
			ExtraColumns: append(orgColumns, idColumn, teamID,
				codegen.ColumnDefinition{
					Name:     "membership",
					Type:     schema.TypeJSON,
					Resolver: "resolveMembership",
				},
			),
		},
		repos,
	}
}
