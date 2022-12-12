package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/domain/model"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DomainResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &model.Domain{},
			Service:    "domain",
			PKColumns:  []string{"id"},
			Relations:  []string{"DomainRecords()"},
		},
		{
			DataStruct: &model.DomainRecord{},
			Service:    "domain",
			PKColumns:  []string{"domain_name", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "domain_name",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("name")`,
				},
			},
		},
	}
}
