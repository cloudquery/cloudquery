package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ElastiCacheResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "clusters",
			Struct:      &types.CacheCluster{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "engine_versions",
			Struct:      &types.CacheEngineVersion{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheEngineVersion.html",
			SkipFields:  []string{"Engine", "EngineVersion"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:        "account_id",
					Description: "The AWS Account ID of the resource.",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveAWSAccount`,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:        "region",
					Description: "The AWS Region of the resource.",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveAWSRegion`,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
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
			SubService:  "global_replication_groups",
			Struct:      &types.GlobalReplicationGroup{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_GlobalReplicationGroup.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "parameter_groups",
			Struct:      &types.CacheParameterGroup{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheParameterGroup.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "replication_groups",
			Struct:      &types.ReplicationGroup{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "reserved_cache_nodes_offerings",
			Struct:      &types.ReservedCacheNodesOffering{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNodesOffering.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveCacheNodesOfferingArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "reserved_cache_nodes",
			Struct:      &types.ReservedCacheNode{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNode.html",
			SkipFields:  []string{"ReservationARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ReservationARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "service_updates",
			Struct:      &types.ServiceUpdate{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ServiceUpdate.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveServiceUpdateArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "snapshots",
			Struct:      &types.Snapshot{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "user_groups",
			Struct:      &types.UserGroup{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UserGroup.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "users",
			Struct:      &types.User{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_User.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		}, {
			SubService:  "subnet_groups",
			Struct:      &types.CacheSubnetGroup{},
			Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheSubnetGroup.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "elasticache"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticache")`
	}
	return resources
}
