package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/google/go-github/v48/github"
)

func dependabotAlert() *Resource {
	return &Resource{
		SubService:   "alerts",
		Struct:       new(github.DependabotAlert),
		PKColumns:    []string{"number"},
		ExtraColumns: codegen.ColumnDefinitions{orgColumn},
	}
}

func dependabotSecret() *Resource {
	return &Resource{
		SubService:   "secrets",
		Struct:       new(github.Secret),
		PKColumns:    []string{"name"},
		ExtraColumns: codegen.ColumnDefinitions{orgColumn},
	}
}
