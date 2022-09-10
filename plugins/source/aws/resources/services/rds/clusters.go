package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RdsClusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_clusters",
		Description: "Contains the details of an Amazon Aurora DB cluster",
		Resolver:    fetchRdsClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "activity_stream_kinesis_stream_name",
				Description: "The name of the Amazon Kinesis data stream used for the database activity stream.",
				Type:        schema.TypeString,
			},
			{
				Name:        "activity_stream_kms_key_id",
				Description: "The AWS KMS key identifier used for encrypting messages in the database activity stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "activity_stream_mode",
				Description: "The mode of the database activity stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "activity_stream_status",
				Description: "The status of the database activity stream.",
				Type:        schema.TypeString,
			},
			{
				Name:        "allocated_storage",
				Description: "For all database engines except Amazon Aurora, AllocatedStorage specifies the allocated storage size in gibibytes (GiB)",
				Type:        schema.TypeInt,
			},
			{
				Name:        "availability_zones",
				Description: "Provides the list of Availability Zones (AZs) where instances in the DB cluster can be created.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "backtrack_consumed_change_records",
				Description: "The number of change records stored for Backtrack.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "backtrack_window",
				Description: "The target backtrack window, in seconds",
				Type:        schema.TypeInt,
			},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic DB snapshots are retained.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "capacity",
				Description: "The current capacity of an Aurora Serverless DB cluster",
				Type:        schema.TypeInt,
			},
			{
				Name:        "character_set_name",
				Description: "If present, specifies the name of the character set that this cluster is associated with.",
				Type:        schema.TypeString,
			},
			{
				Name:        "clone_group_id",
				Description: "Identifies the clone group to which the DB cluster is associated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "copy_tags_to_snapshot",
				Description: "Specifies whether tags are copied from the DB cluster to snapshots of the DB cluster.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "cross_account_clone",
				Description: "Specifies whether the DB cluster is a clone of a DB cluster owned by a different AWS account.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "custom_endpoints",
				Description: "Identifies all custom endpoints associated with the cluster.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "db_cluster_identifier",
				Description: "Contains a user-supplied DB cluster identifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:        "db_cluster_parameter_group",
				Description: "Specifies the name of the DB cluster parameter group for the DB cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterParameterGroup"),
			},
			{
				Name:        "db_cluster_option_group_memberships",
				Description: "Provides the map of option group memberships for this DB cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DBClusterOptionGroupMemberships"),
			},
			{
				Name:        "db_subnet_group",
				Description: "Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name:        "database_name",
				Description: "Contains the name of the initial database of this DB cluster that was provided at create time, if one was specified when the DB cluster was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The AWS Region-unique, immutable identifier for the DB cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DbClusterResourceId"),
			},
			{
				Name:        "deletion_protection",
				Description: "Indicates if the DB cluster has deletion protection enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "earliest_backtrack_time",
				Description: "The earliest time to which a DB cluster can be backtracked.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "earliest_restorable_time",
				Description: "The earliest time to which a database can be restored with point-in-time restore.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this DB cluster is configured to export to CloudWatch Logs",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "endpoint",
				Description: "Specifies the connection endpoint for the primary instance of the DB cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine",
				Description: "The name of the database engine to be used for this DB cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_mode",
				Description: "The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery, global, or multimaster",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "global_write_forwarding_requested",
				Description: "Specifies whether you have requested to enable write forwarding for a secondary cluster in an Aurora global database",
				Type:        schema.TypeBool,
			},
			{
				Name:        "global_write_forwarding_status",
				Description: "Specifies whether a secondary cluster in an Aurora global database has write forwarding enabled, not enabled, or is in the process of enabling it.",
				Type:        schema.TypeString,
			},
			{
				Name:        "hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        schema.TypeString,
			},
			{
				Name:        "http_endpoint_enabled",
				Description: "A value that indicates whether the HTTP endpoint for an Aurora Serverless DB cluster is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "A value that indicates whether the mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "kms_key_id",
				Description: "If StorageEncrypted is enabled, the AWS KMS key identifier for the encrypted DB cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "master_username",
				Description: "Contains the master username for the DB cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "multi_az",
				Description: "Specifies whether the DB cluster has instances in multiple Availability Zones.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MultiAZ"),
			},
			{
				Name:        "pending_modified_values",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "percent_progress",
				Description: "Specifies the progress of the operation as a percentage.",
				Type:        schema.TypeString,
			},
			{
				Name:        "port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "preferred_backup_window",
				Description: "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.",
				Type:        schema.TypeString,
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
				Type:        schema.TypeString,
			},
			{
				Name:        "read_replica_identifiers",
				Description: "Contains one or more identifiers of the read replicas associated with this DB cluster.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "reader_endpoint",
				Description: "The reader endpoint for the DB cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "replication_source_identifier",
				Description: "Contains the identifier of the source DB cluster if this DB cluster is a read replica.",
				Type:        schema.TypeString,
			},
			{
				Name:        "scaling_configuration_info",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "status",
				Description: "Specifies the current state of this DB cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "tags",
				Description: "A list of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsClusterTags,
			},
			{
				Name: "associated_roles",
				Type: schema.TypeJSON,
			},
			{
				Name: "db_cluster_memebers",
				Resolver: schema.PathResolver("DBClusterMembers"),
				Type: schema.TypeJSON,
			},
			{
				Name: "domain_memberships",
				Type: schema.TypeJSON,
			},
			{
				Name: "vpc_security_groups",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRdsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeDBClustersInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBClusters(ctx, &config)
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
func resolveRdsClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBCluster)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
