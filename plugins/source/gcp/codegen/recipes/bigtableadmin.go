package recipes

import (
	bigtable "cloud.google.com/go/bigtable"
	btapb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "instances",
			Struct:      &bigtable.InstanceInfo{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances#Instance",
			NewFunction: bigtable.NewInstanceAdminClient,
			SkipMock:    true,
			SkipFetch:   true,
			Relations:   []string{"AppProfiles(), Clusters(), Tables()"},
		},
		{
			SubService:  "app_profiles",
			Struct:      &btapb.AppProfile{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles#AppProfile",
			ChildTable:  true,
			SkipMock:    true,
			SkipFetch:   true,
		},
		{
			SubService:  "clusters",
			Struct:      &bigtable.ClusterInfo{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters#Cluster",
			ChildTable:  true,
			SkipMock:    true,
			SkipFetch:   true,
			Relations:   []string{"Backups()"},
		},
		{
			SubService:          "tables",
			Struct:              &bigtable.TableInfo{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			Description:         "https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.tables#Table",
			ChildTable:          true,
			SkipMock:            true,
			SkipFetch:           true,
			PreResourceResolver: "getTableInfo",
		},
		{
			SubService:  "backups",
			Struct:      &bigtable.BackupInfo{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters.backups#Backup",
			ChildTable:  true,
			SkipMock:    true,
			SkipFetch:   true,
		},
	}

	for _, resource := range resources {
		resource.Service = "bigtableadmin"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
