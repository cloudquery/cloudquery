package helpers

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	"github.com/iancoleman/strcase"
)

func TableAndFetcherNames(r *recipes.Resource) (string, string) {
	cqSubservice := Coalesce(r.CQSubserviceOverride, r.AWSSubService)

	tableNameFromSubService := strcase.ToSnake(cqSubservice)
	fetcherNameFromSubService := strcase.ToCamel(cqSubservice)
	{
		// Generate table and fetcher names using parent info

		prev := tableNameFromSubService
		var (
			preTableNames   []string
			preFetcherNames []string
		)
		rp := r.Parent
		for rp != nil {
			if rp.CQSubserviceOverride != "" {
				preTableNames = append(preTableNames, rp.CQSubserviceOverride)
				preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.CQSubserviceOverride))
			} else {
				ins := strcase.ToSnake(rp.ItemName)
				if !strings.HasPrefix(prev, ins) {
					preTableNames = append(preTableNames, ins)
					preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.ItemName))
					prev = ins
				}
			}
			rp = rp.Parent
		}
		if len(preTableNames) > 0 {
			tableNameFromSubService = strings.Join(ReverseStringSlice(preTableNames), "_") + "_" + tableNameFromSubService
			fetcherNameFromSubService = strings.Join(ReverseStringSlice(preFetcherNames), "") + fetcherNameFromSubService
		}
	}

	return tableNameFromSubService, fetcherNameFromSubService
}
