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
				"Instances()",
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
		{
			SubService:  "cluster_parameter_groups",
			Struct:      &types.DBClusterParameterGroup{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html",
			SkipFields:  []string{"DBClusterParameterGroupArn"},
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
					{
						Name:     "parameters",
						Type:     schema.TypeJSON,
						Resolver: `resolveDocdbClusterParameterGroupParameters`,
					},
				}...),
		},
		{
			SubService:  "certificates",
			Struct:      &types.Certificate{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html",
			SkipFields:  []string{"CertificateArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CertificateArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "engine_versions",
			Struct:      &types.DBEngineVersion{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html",
			Multiplex:   `client.AccountMultiplex`,
			SkipFields:  []string{"Engine", "EngineVersion"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "engine",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Engine")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "engine_version",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("EngineVersion")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "instances",
			Struct:      &types.DBInstance{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html",
			SkipFields:  []string{"DBInstanceArn"},
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
			SubService:  "subnet_groups",
			Struct:      &types.DBSubnetGroup{},
			Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBSubnetGroup.html",
			SkipFields:  []string{"DBSubnetGroupArn"},
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
