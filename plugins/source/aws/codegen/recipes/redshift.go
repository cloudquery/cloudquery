package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RedshiftResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "clusters",
			Struct:     &types.Cluster{},
			SkipFields: []string{"ClusterParameterGroups", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver:    `resolveClusterArn()`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The list of tags for the cluster.",
						Type:        schema.TypeJSON,
						Resolver:    `client.ResolveTags`,
					},
					{
						Name:        "logging_status",
						Description: "Describes the status of logging for a cluster.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveRedshiftClusterLoggingStatus`,
					},
				}...),
			Relations: []string{
				"Snapshots()",
				"ClusterParameterGroups()",
			},
		},
		{
			SubService: "cluster_parameter_groups",
			Struct:     &types.ClusterParameterGroupStatus{},
			SkipFields: []string{"ParameterGroupName"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "cluster_arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver:    `schema.ParentColumnResolver("arn")`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "parameter_group_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ParameterGroupName")`,
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
			SkipFields: []string{"ParameterName"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "cluster_arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver:    `schema.ParentColumnResolver("cluster_arn")`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "parameter_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ParameterName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},

		{
			SubService: "event_subscriptions",
			Struct:     &types.EventSubscription{},
			SkipFields: []string{"Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "ARN of the event subscription.",
						Type:        schema.TypeString,
						Resolver:    `resolveEventSubscriptionARN`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Tags",
						Type:        schema.TypeJSON,
						Resolver:    `client.ResolveTags`,
					},
				}...),
		},

		{
			SubService: "snapshots",
			Struct:     &types.Snapshot{},
			SkipFields: []string{"Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "ARN of the snapshot.",
						Type:        schema.TypeString,
						Resolver:    `resolveSnapshotARN`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Tags consisting of a name/value pair for a resource.",
						Type:        schema.TypeJSON,
						Resolver:    `client.ResolveTags`,
					},
				}...),
		},

		{
			SubService: "subnet_groups",
			Struct:     &types.ClusterSubnetGroup{},
			SkipFields: []string{"Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver:    `resolveSubnetGroupArn()`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The list of tags for the cluster subnet group.",
						Type:        schema.TypeJSON,
						Resolver:    `client.ResolveTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "redshift"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("redshift")`
	}
	return resources
}
