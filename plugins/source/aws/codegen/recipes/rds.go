package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RDSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "certificates",
			Struct:      &types.Certificate{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Certificate.html",
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
			SkipFields:  []string{"Engine", "EngineVersion"},
			ExtraColumns: append(defaultRegionalColumnsPK, []codegen.ColumnDefinition{
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
			}...),
			Relations: []string{
				"ClusterParameters()",
			},
		},
		{
			SubService:   "cluster_parameters",
			Struct:       &types.Parameter{},
			SkipFields:   []string{"DBClusterParameterGroupArn"},
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService:  "cluster_parameter_groups",
			Struct:      &types.DBClusterParameterGroup{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterParameterGroup.html",
			SkipFields:  []string{"DBClusterParameterGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterParameterGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsClusterParameterGroupTags`,
					},
				}...),
			Relations: []string{"ClusterParameterGroupParameters()"},
		},
		{
			SubService:  "cluster_parameter_group_parameters",
			Struct:      &types.Parameter{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "cluster_parameter_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "cluster_snapshots",
			Struct:      &types.DBClusterSnapshot{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html",
			SkipFields:  []string{"DBClusterSnapshotArn", "TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterSnapshotArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRDSClusterSnapshotTags`,
					},
					{
						Name:     "attributes",
						Type:     schema.TypeJSON,
						Resolver: `resolveRDSClusterSnapshotAttributes`,
					},
				}...),
		},
		{
			SubService:  "clusters",
			Struct:      &types.DBCluster{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html",
			SkipFields:  []string{"DBClusterArn", "TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsClusterTags`,
					},
				}...),
		},
		{
			SubService:  "db_parameter_groups",
			Struct:      &types.DBParameterGroup{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html",
			SkipFields:  []string{"DBParameterGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBParameterGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsDbParameterGroupTags`,
					},
				}...),
			Relations: []string{"DbParameterGroupDbParameters()"},
		},
		{
			SubService:  "db_parameter_group_db_parameters",
			Struct:      &types.Parameter{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "db_parameter_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "db_security_groups",
			Struct:      &types.DBSecurityGroup{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html",
			SkipFields:  []string{"DBSecurityGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBSecurityGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsDbSecurityGroupTags`,
					},
				}...),
		},
		{
			SubService:  "db_snapshots",
			Struct:      &types.DBSnapshot{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSnapshot.html",
			SkipFields:  []string{"DBSnapshotArn", "TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBSnapshotArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRDSDBSnapshotTags`,
					},
					{
						Name:     "attributes",
						Type:     schema.TypeJSON,
						Resolver: `resolveRDSDBSnapshotAttributes`,
					},
				}...),
		},
		{
			SubService:  "event_subscriptions",
			Struct:      &types.EventSubscription{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_EventSubscription.html",
			SkipFields:  []string{"EventSubscriptionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("EventSubscriptionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRDSEventSubscriptionTags`,
					},
				}...),
		},
		{
			SubService:  "instances",
			Struct:      &types.DBInstance{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html",
			SkipFields:  []string{"DBInstanceArn", "ProcessorFeatures", "TagList"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DBInstanceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "processor_features",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsInstanceProcessorFeatures`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRdsInstanceTags`,
					},
				}...),
		},
		{
			SubService:  "subnet_groups",
			Struct:      &types.DBSubnetGroup{},
			Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSubnetGroup.html",
			SkipFields:  []string{"DBSubnetGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
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
		r.Service = "rds"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("rds")`
	}
	return resources
}
