package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func NeptuneResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "cluster_parameter_groups",
			Struct:      &types.DBClusterParameterGroup{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameters",
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
						Resolver: `resolveNeptuneClusterParameterGroupTags`,
					},
				}...),
			Relations: []string{"ClusterParameterGroupParameters()"},
		},
		{
			SubService:  "cluster_parameter_group_parameters",
			Struct:      &types.Parameter{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameterGroups",
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
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-snapshots.html#DescribeDBClusterSnapshots",
			SkipFields:  []string{"DBClusterSnapshotArn"},
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
						Name:     "attributes",
						Type:     schema.TypeJSON,
						Resolver: `resolveNeptuneClusterSnapshotAttributes`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveNeptuneClusterSnapshotTags`,
					},
				}...),
		},
		{
			SubService:  "clusters",
			Struct:      &types.DBCluster{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters",
			SkipFields:  []string{"DBClusterArn"},
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
						Resolver: `resolveNeptuneClusterTags`,
					},
				}...),
		},
		{
			SubService:  "db_parameter_groups",
			Struct:      &types.DBParameterGroup{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameterGroups",
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
						Resolver: `resolveNeptuneDbParameterGroupTags`,
					},
				}...),
			Relations: []string{"DbParameterGroupDbParameters()"},
		},
		{
			SubService:  "db_parameter_group_db_parameters",
			Struct:      &types.Parameter{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameters",
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
			SubService:  "global_clusters",
			Struct:      &types.GlobalCluster{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#DescribeDBInstances",
			SkipFields:  []string{"GlobalClusterArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("GlobalClusterArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveNeptuneGlobalClusterTags`,
					},
				}...),
		},
		{
			SubService:  "event_subscriptions",
			Struct:      &types.EventSubscription{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-events.html#DescribeEventSubscriptions",
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
						Resolver: `resolveNeptuneEventSubscriptionTags`,
					},
				}...),
		},
		{
			SubService:  "instances",
			Struct:      &types.DBInstance{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#DescribeDBInstances",
			SkipFields:  []string{"DBInstanceArn"},
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
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveNeptuneInstanceTags`,
					},
				}...),
		},
		{
			SubService:  "subnet_groups",
			Struct:      &types.DBSubnetGroup{},
			Description: "https://docs.aws.amazon.com/neptune/latest/userguide/api-subnets.html#DescribeDBSubnetGroups",
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
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveNeptuneSubnetGroupTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "neptune"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("neptune")`
	}
	return resources
}
