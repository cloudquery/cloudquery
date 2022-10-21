package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DocumentDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "clusters",
			Struct:      &types.DBCluster{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			SkipFields:  []string{"DBClusterArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBClusterTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"ClusterSnapshots()",
			},
		},
		{
			SubService:  "cluster_snapshots",
			Struct:      &types.DBClusterSnapshot{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshot.html",
			SkipFields:  []string{"DBClusterSnapshotArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBClusterSnapshotTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterSnapshotArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "attributes",
						Type:     schema.TypeJSON,
						Resolver: `resolveDocdbClusterSnapshotAttributes`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "docdb"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("docdb")`
	}
	return resources
}
