package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)



func init(){
	resources := []*Resource{
		{
			SubService: "instances",
			Struct:     &sqladmin.DatabaseInstance{},
			SkipMock:   true,
			SkipFetch:  true,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "self_link",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("SelfLink")`,
				},
			},
			NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6"}),
		},
	}

	for _, resource := range resources {
		resource.Service = "sql"
		resource.Template = "newapi_list"
	}

	Resources = append(Resources, resources...)

}
