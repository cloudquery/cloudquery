package recipes

import (
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func init() {
	resources := []*Resource{
		{
			SubService:      "instances",
			Struct:          &sqladmin.DatabaseInstance{},
			SkipMock:        true,
			SkipFetch:       true,
			PrimaryKeys:     []string{"self_link"},
			NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_6": "ipv6"}),
			Relations:       []string{"Users()"},
		},
		{
			SubService:  "users",
			Struct:      &sqladmin.User{},
			SkipMock:    true,
			SkipFetch:   true,
			ChildTable:  true,
			PrimaryKeys: []string{ProjectIdColumn.Name, "instance", "name"},
		},
	}

	for _, resource := range resources {
		resource.Service = "sql"
		resource.Template = "newapi_list"
	}

	Resources = append(Resources, resources...)

}
