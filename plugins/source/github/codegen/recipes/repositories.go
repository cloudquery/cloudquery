package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Repositories() []*Resource {
	repositoryID := codegen.ColumnDefinition{
		Name:     "repository_id",
		Type:     schema.TypeInt,
		Resolver: `client.ResolveParentColumn("ID")`,
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}

	alert := dependabotAlert()
	alert.Service = "repositories"
	alert.TableName = "repository_dependabot_alerts"
	alert.SkipFields = []string{"RepositoryID"}
	alert.ExtraColumns = append(alert.ExtraColumns, repositoryID)

	sec := dependabotSecret()
	sec.Service = "repositories"
	sec.TableName = "repository_dependabot_secrets"
	sec.SkipFields = []string{"RepositoryID"}
	sec.ExtraColumns = append(sec.ExtraColumns, repositoryID)

	repo := repository()
	repo.Service = "repositories"
	repo.TableName = "repositories"
	repo.Multiplex = orgMultiplex
	repo.Relations = []string{"Alerts()", "Secrets()", "Releases()"}

	return []*Resource{
		repo, alert, sec,
		{
			Service:      "repositories",
			TableName:    "releases",
			Multiplex:    orgMultiplex,
			SubService:   "releases",
			Struct:       new(github.RepositoryRelease),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn, repositoryID},
			Relations:    []string{"Assets()"},
		},
		{
			Service:    "repositories",
			TableName:  "release_assets",
			Multiplex:  orgMultiplex,
			SubService: "assets",
			Struct:     new(github.ReleaseAsset),
			PKColumns:  []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn, codegen.ColumnDefinition{
				Name:     "repository_id",
				Type:     schema.TypeInt,
				Resolver: `client.ResolveGrandParentColumn("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			}},
		},
	}
}

func repository() *Resource {
	return &Resource{
		SubService:   "repositories",
		Struct:       new(github.Repository),
		PKColumns:    []string{"id"},
		ExtraColumns: codegen.ColumnDefinitions{orgColumn},
	}
}
