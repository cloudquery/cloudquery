// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_instances",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html`,
		Resolver:    fetchRdsInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "processor_features",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceProcessorFeatures,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceTags,
			},
			{
				Name:     "activity_stream_engine_native_audit_fields_included",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ActivityStreamEngineNativeAuditFieldsIncluded"),
			},
			{
				Name:     "activity_stream_kinesis_stream_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamKinesisStreamName"),
			},
			{
				Name:     "activity_stream_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamKmsKeyId"),
			},
			{
				Name:     "activity_stream_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamMode"),
			},
			{
				Name:     "activity_stream_policy_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamPolicyStatus"),
			},
			{
				Name:     "activity_stream_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamStatus"),
			},
			{
				Name:     "allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedStorage"),
			},
			{
				Name:     "associated_roles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AssociatedRoles"),
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "automatic_restart_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AutomaticRestartTime"),
			},
			{
				Name:     "automation_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutomationMode"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "aws_backup_recovery_point_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsBackupRecoveryPointArn"),
			},
			{
				Name:     "backup_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BackupRetentionPeriod"),
			},
			{
				Name:     "backup_target",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupTarget"),
			},
			{
				Name:     "ca_certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CACertificateIdentifier"),
			},
			{
				Name:     "character_set_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CharacterSetName"),
			},
			{
				Name:     "copy_tags_to_snapshot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CopyTagsToSnapshot"),
			},
			{
				Name:     "custom_iam_instance_profile",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomIamInstanceProfile"),
			},
			{
				Name:     "customer_owned_ip_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CustomerOwnedIpEnabled"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_instance_automated_backups_replications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBInstanceAutomatedBackupsReplications"),
			},
			{
				Name:     "db_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceClass"),
			},
			{
				Name:     "db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceIdentifier"),
			},
			{
				Name:     "db_instance_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceStatus"),
			},
			{
				Name:     "db_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBName"),
			},
			{
				Name:     "db_parameter_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBParameterGroups"),
			},
			{
				Name:     "db_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBSecurityGroups"),
			},
			{
				Name:     "db_subnet_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name:     "db_instance_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DbInstancePort"),
			},
			{
				Name:     "dbi_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DbiResourceId"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "domain_memberships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainMemberships"),
			},
			{
				Name:     "enabled_cloudwatch_logs_exports",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledCloudwatchLogsExports"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "enhanced_monitoring_resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnhancedMonitoringResourceArn"),
			},
			{
				Name:     "iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:     "instance_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InstanceCreateTime"),
			},
			{
				Name:     "iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Iops"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "latest_restorable_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LatestRestorableTime"),
			},
			{
				Name:     "license_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseModel"),
			},
			{
				Name:     "listener_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ListenerEndpoint"),
			},
			{
				Name:     "master_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterUsername"),
			},
			{
				Name:     "max_allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxAllocatedStorage"),
			},
			{
				Name:     "monitoring_interval",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MonitoringInterval"),
			},
			{
				Name:     "monitoring_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonitoringRoleArn"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "nchar_character_set_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NcharCharacterSetName"),
			},
			{
				Name:     "network_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkType"),
			},
			{
				Name:     "option_group_memberships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OptionGroupMemberships"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:     "performance_insights_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PerformanceInsightsEnabled"),
			},
			{
				Name:     "performance_insights_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PerformanceInsightsKMSKeyId"),
			},
			{
				Name:     "performance_insights_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PerformanceInsightsRetentionPeriod"),
			},
			{
				Name:     "preferred_backup_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredBackupWindow"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "promotion_tier",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PromotionTier"),
			},
			{
				Name:     "publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PubliclyAccessible"),
			},
			{
				Name:     "read_replica_db_cluster_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadReplicaDBClusterIdentifiers"),
			},
			{
				Name:     "read_replica_db_instance_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadReplicaDBInstanceIdentifiers"),
			},
			{
				Name:     "read_replica_source_db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadReplicaSourceDBInstanceIdentifier"),
			},
			{
				Name:     "replica_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicaMode"),
			},
			{
				Name:     "resume_full_automation_mode_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ResumeFullAutomationModeTime"),
			},
			{
				Name:     "secondary_availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryAvailabilityZone"),
			},
			{
				Name:     "status_infos",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatusInfos"),
			},
			{
				Name:     "storage_encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("StorageEncrypted"),
			},
			{
				Name:     "storage_throughput",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageThroughput"),
			},
			{
				Name:     "storage_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageType"),
			},
			{
				Name:     "tde_credential_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TdeCredentialArn"),
			},
			{
				Name:     "timezone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timezone"),
			},
			{
				Name:     "vpc_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcSecurityGroups"),
			},
		},
	}
}
