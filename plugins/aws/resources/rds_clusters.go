package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_clusters",
		Resolver:     fetchRdsClusters,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "activity_stream_kinesis_stream_name",
				Type: schema.TypeString,
			},
			{
				Name: "activity_stream_kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "activity_stream_mode",
				Type: schema.TypeString,
			},
			{
				Name: "activity_stream_status",
				Type: schema.TypeString,
			},
			{
				Name: "allocated_storage",
				Type: schema.TypeInt,
			},
			{
				Name: "availability_zones",
				Type: schema.TypeStringArray,
			},
			{
				Name: "backtrack_consumed_change_records",
				Type: schema.TypeBigInt,
			},
			{
				Name: "backtrack_window",
				Type: schema.TypeBigInt,
			},
			{
				Name: "backup_retention_period",
				Type: schema.TypeInt,
			},
			{
				Name: "capacity",
				Type: schema.TypeInt,
			},
			{
				Name: "character_set_name",
				Type: schema.TypeString,
			},
			{
				Name: "clone_group_id",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "copy_tags_to_snapshot",
				Type: schema.TypeBool,
			},
			{
				Name: "cross_account_clone",
				Type: schema.TypeBool,
			},
			{
				Name: "custom_endpoints",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "db_cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterArn"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_parameter_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroup"),
			},
			{
				Name:     "db_subnet_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name: "database_name",
				Type: schema.TypeString,
			},
			{
				Name: "db_cluster_resource_id",
				Type: schema.TypeString,
			},
			{
				Name: "deletion_protection",
				Type: schema.TypeBool,
			},
			{
				Name: "earliest_backtrack_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "earliest_restorable_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "enabled_cloudwatch_logs_exports",
				Type: schema.TypeStringArray,
			},
			{
				Name: "endpoint",
				Type: schema.TypeString,
			},
			{
				Name: "engine",
				Type: schema.TypeString,
			},
			{
				Name: "engine_mode",
				Type: schema.TypeString,
			},
			{
				Name: "engine_version",
				Type: schema.TypeString,
			},
			{
				Name: "global_write_forwarding_requested",
				Type: schema.TypeBool,
			},
			{
				Name: "global_write_forwarding_status",
				Type: schema.TypeString,
			},
			{
				Name: "hosted_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "http_endpoint_enabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "latest_restorable_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "master_username",
				Type: schema.TypeString,
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "pending_modified_values_db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.DBClusterIdentifier"),
			},
			{
				Name:     "pending_modified_values_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.EngineVersion"),
			},
			{
				Name:     "pending_modified_values_iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PendingModifiedValues.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:     "pending_modified_values_master_user_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
			},
			{
				Name:     "pending_cloudwatch_logs_types_to_disable",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable"),
			},
			{
				Name:     "pending_cloudwatch_logs_types_to_enable",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable"),
			},
			{
				Name: "percent_progress",
				Type: schema.TypeString,
			},
			{
				Name: "port",
				Type: schema.TypeInt,
			},
			{
				Name: "preferred_backup_window",
				Type: schema.TypeString,
			},
			{
				Name: "preferred_maintenance_window",
				Type: schema.TypeString,
			},
			{
				Name: "read_replica_identifiers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "reader_endpoint",
				Type: schema.TypeString,
			},
			{
				Name: "replication_source_identifier",
				Type: schema.TypeString,
			},
			{
				Name:     "scaling_configuration_info_auto_pause",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ScalingConfigurationInfo.AutoPause"),
			},
			{
				Name:     "scaling_configuration_info_max_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScalingConfigurationInfo.MaxCapacity"),
			},
			{
				Name:     "scaling_configuration_info_min_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScalingConfigurationInfo.MinCapacity"),
			},
			{
				Name:     "scaling_configuration_info_seconds_until_auto_pause",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScalingConfigurationInfo.SecondsUntilAutoPause"),
			},
			{
				Name:     "scaling_configuration_info_timeout_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScalingConfigurationInfo.TimeoutAction"),
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "storage_encrypted",
				Type: schema.TypeBool,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsClusterTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_rds_cluster_associated_roles",
				Resolver: fetchRdsClusterAssociatedRoles,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "feature_name",
						Type: schema.TypeString,
					},
					{
						Name: "role_arn",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_cluster_db_cluster_members",
				Resolver: fetchRdsClusterDbClusterMembers,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "db_cluster_parameter_group_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBClusterParameterGroupStatus"),
					},
					{
						Name:     "db_instance_identifier",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBInstanceIdentifier"),
					},
					{
						Name: "is_cluster_writer",
						Type: schema.TypeBool,
					},
					{
						Name: "promotion_tier",
						Type: schema.TypeInt,
					},
				},
			},
			{
				Name:     "aws_rds_cluster_db_cluster_option_group_memberships",
				Resolver: fetchRdsClusterDbClusterOptionGroupMemberships,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "db_cluster_option_group_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBClusterOptionGroupName"),
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_cluster_domain_memberships",
				Resolver: fetchRdsClusterDomainMemberships,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "domain",
						Type: schema.TypeString,
					},
					{
						Name: "f_q_d_n",
						Type: schema.TypeString,
					},
					{
						Name:     "iam_role_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("IAMRoleName"),
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_cluster_vpc_security_groups",
				Resolver: fetchRdsClusterVpcSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_security_group_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config rds.DescribeDBClustersInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBClusters(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBClusters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRdsClusterTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBCluster)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchRdsClusterAssociatedRoles(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.DBCluster)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- cluster.AssociatedRoles
	return nil
}
func fetchRdsClusterDbClusterMembers(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.DBCluster)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- cluster.DBClusterMembers
	return nil
}
func fetchRdsClusterDbClusterOptionGroupMemberships(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.DBCluster)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- cluster.DBClusterOptionGroupMemberships
	return nil
}
func fetchRdsClusterDomainMemberships(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.DBCluster)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- cluster.DomainMemberships
	return nil
}
func fetchRdsClusterVpcSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.DBCluster)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- cluster.VpcSecurityGroups
	return nil
}
