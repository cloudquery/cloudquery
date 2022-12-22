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
			Description:     "https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances#DatabaseInstance",
		},
		{
			SubService:  "users",
			Struct:      &sqladmin.User{},
			SkipMock:    true,
			SkipFetch:   true,
			ChildTable:  true,
			PrimaryKeys: []string{ProjectIdColumn.Name, "instance", "name"},
			Description: "https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/users#User",
		},
	}

	for _, resource := range resources {
		resource.Service = "sql"
		resource.Template = "newapi_list"
		resource.ServiceDNS = "sqladmin.googleapis.com"
	}

	Resources = append(Resources, resources...)

}
