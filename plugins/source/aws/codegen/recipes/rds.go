package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RDSResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "certificates",
			Struct:     &types.Certificate{},
			SkipFields: []string{"CertificateArn"},
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
			SubService: "cluster_parameter_groups",
			Struct:     &types.DBClusterParameterGroup{},
			SkipFields: []string{"DBClusterParameterGroupArn"},
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
			SubService: "cluster_parameter_group_parameters",
			Struct:     &types.Parameter{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "cluster_parameter_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "cluster_snapshots",
			Struct:     &types.DBClusterSnapshot{},
			SkipFields: []string{"DBClusterSnapshotArn"},
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
			SubService: "clusters",
			Struct:     &types.DBCluster{},
			SkipFields: []string{"DBClusterArn"},
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
			SubService: "db_parameter_groups",
			Struct:     &types.DBParameterGroup{},
			SkipFields: []string{"DBParameterGroupArn"},
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
			SubService: "db_parameter_group_db_parameters",
			Struct:     &types.Parameter{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "db_parameter_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "db_security_groups",
			Struct:     &types.DBSecurityGroup{},
			SkipFields: []string{"DBSecurityGroupArn"},
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
			SubService: "db_snapshots",
			Struct:     &types.DBSnapshot{},
			SkipFields: []string{"DBSnapshotArn"},
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
			SubService: "event_subscriptions",
			Struct:     &types.EventSubscription{},
			SkipFields: []string{"EventSubscriptionArn"},
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
			SubService: "instances",
			Struct:     &types.DBInstance{},
			SkipFields: []string{"DBInstanceArn", "ProcessorFeatures", "Tags"},
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
			SubService: "subnet_groups",
			Struct:     &types.DBSubnetGroup{},
			SkipFields: []string{"DBSubnetGroupArn"},
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
