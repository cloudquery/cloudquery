package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func Repositories() []*Resource {
	repo := repository()
	repo.Service = "repositories"
	repo.TableName = "repositories"
	repo.Multiplex = orgMultiplex

	return []*Resource{repo}
}

func repository() *Resource {
	return &Resource{
		SubService:   "repositories",
		Struct:       new(github.Repository),
		PKColumns:    []string{"id"},
		ExtraColumns: codegen.ColumnDefinitions{orgColumn},
	}
}
