package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func init() {
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
			Relations:       []string{"Users()"},
		},
		{
			SubService: "users",
			Struct:     &sqladmin.User{},
			SkipMock:   true,
			SkipFetch:  true,
			ChildTable: true,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "project_id",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: "client.ResolveProject",
				},
				{
					Name:     "instance",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Instance")`,
				},
				{
					Name:     "name",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Name")`,
				},
			},
		},
	}

	for _, resource := range resources {
		resource.Service = "sql"
		resource.Template = "newapi_list"
	}

	Resources = append(Resources, resources...)

}
