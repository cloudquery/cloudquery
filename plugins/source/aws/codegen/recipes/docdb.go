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
			SubService: "cluster_snapshots",
			Struct:     &types.DBClusterSnapshot{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBClusterSnapshotTags`,
					},
				}...),
			Relations: []string{"ClusterSnapshotAttributes()"},
		},
		{
			SubService: "cluster_snapshot_attributes",
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			Struct: &types.DBClusterSnapshotAttributesResult{},
		},
		{
			SubService: "cluster_parameter_groups",
			Struct:     &types.DBClusterParameterGroup{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			SkipFields: []string{"DBClusterParameterGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBClusterParameterGroupTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterParameterGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"ClusterParameters()",
			},
		},
		{
			SubService: "cluster_parameters",
			Struct:     &types.Parameter{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
		},
		{
			SubService: "certificates",
			Struct:     &types.Certificate{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService: "engine_versions",
			Struct:     &types.DBEngineVersion{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService: "instances",
			Struct:     &types.DBInstance{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			SkipFields: []string{"DBInstanceArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBInstanceTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBInstanceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "subnet_groups",
			Struct:     &types.DBSubnetGroup{},
			//Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html",
			SkipFields: []string{"DBSubnetGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveDBSubnetGroupTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBSubnetGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
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
