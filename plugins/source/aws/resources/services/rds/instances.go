package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RdsInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_instances",
		Description: "Contains the details of an Amazon RDS DB instance",
		Resolver:    fetchRdsInstances,
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
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size specified in gibibytes.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "A value that indicates that minor version patches are applied automatically.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "availability_zone",
				Description: "Specifies the name of the Availability Zone the DB instance is located in.",
				Type:        schema.TypeString,
			},
			{
				Name:        "aws_backup_recovery_point_arn",
				Description: "The Amazon Resource Name (ARN) of the recovery point in AWS Backup.",
				Type:        schema.TypeString,
			},
			{
				Name:        "backup_retention_period",
				Description: "Specifies the number of days for which automatic DB snapshots are retained.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "ca_certificate_identifier",
				Description: "The identifier of the CA certificate for this DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CACertificateIdentifier"),
			},
			{
				Name:        "character_set_name",
				Description: "If present, specifies the name of the character set that this instance is associated with.",
				Type:        schema.TypeString,
			},
			{
				Name:        "copy_tags_to_snapshot",
				Description: "Specifies whether tags are copied from the DB instance to snapshots of the DB instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "customer_owned_ip_enabled",
				Description: "Specifies whether a customer-owned IP address (CoIP) is enabled for an RDS on Outposts DB instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "cluster_identifier",
				Description: "If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "db_instance_class",
				Description: "Contains the name of the compute and memory capacity class of the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBInstanceClass"),
			},
			{
				Name:        "user_instance_id",
				Description: "Contains a user-supplied database identifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBInstanceIdentifier"),
			},
			{
				Name:        "db_instance_status",
				Description: "Specifies the current state of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBInstanceStatus"),
			},
			{
				Name:        "db_name",
				Description: "The meaning of this parameter differs according to the database engine you use. MySQL, MariaDB, SQL Server, PostgreSQL Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBName"),
			},
			{
				Name:        "instance_port",
				Description: "Specifies the port that the DB instance listens on",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DbInstancePort"),
			},
			{
				Name:        "id",
				Description: "The AWS Region-unique, immutable identifier for the DB instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DbiResourceId"),
			},
			{
				Name:        "deletion_protection",
				Description: "Indicates if the DB instance has deletion protection enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enabled_cloudwatch_logs_exports",
				Description: "A list of log types that this DB instance is configured to export to CloudWatch Logs",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "endpoint",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "engine",
				Description: "The name of the database engine to be used for this DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "Indicates the database engine version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enhanced_monitoring_resource_arn",
				Description: "The Amazon Resource Name (ARN) of the Amazon CloudWatch Logs log stream that receives the Enhanced Monitoring metrics data for the DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "instance_create_time",
				Description: "Provides the date and time the DB instance was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "iops",
				Description: "Specifies the Provisioned IOPS (I/O operations per second) value.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "kms_key_id",
				Description: "If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_restorable_time",
				Description: "Specifies the latest time to which a database can be restored with point-in-time restore.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "license_model",
				Description: "License model information for this DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "listener_endpoint",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "master_username",
				Description: "Contains the master username for the DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_allocated_storage",
				Description: "The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "monitoring_interval",
				Description: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "monitoring_role_arn",
				Description: "The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "multi_az",
				Description: "Specifies if the DB instance is a Multi-AZ deployment.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MultiAZ"),
			},
			{
				Name:        "nchar_character_set_name",
				Description: "The name of the NCHAR character set for the Oracle DB instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "pending_modified_values",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "performance_insights_enabled",
				Description: "True if Performance Insights is enabled for the DB instance, and otherwise false.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "performance_insights_kms_key_id",
				Description: "The AWS KMS key identifier for encryption of Performance Insights data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PerformanceInsightsKMSKeyId"),
			},
			{
				Name:        "performance_insights_retention_period",
				Description: "The amount of time, in days, to retain Performance Insights data",
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
				Name:        "processor_features",
				Description: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsInstanceProcessorFeatures,
			},
			{
				Name:        "promotion_tier",
				Description: "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance",
				Type:        schema.TypeInt,
			},
			{
				Name:        "publicly_accessible",
				Description: "Specifies the accessibility options for the DB instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "read_replica_db_cluster_identifiers",
				Description: "Contains one or more identifiers of Aurora DB clusters to which the RDS DB instance is replicated as a read replica",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ReadReplicaDBClusterIdentifiers"),
			},
			{
				Name:        "read_replica_db_instance_identifiers",
				Description: "Contains one or more identifiers of the read replicas associated with this DB instance.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ReadReplicaDBInstanceIdentifiers"),
			},
			{
				Name:        "read_replica_source_db_instance_identifier",
				Description: "Contains the identifier of the source DB instance if this DB instance is a read replica.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReadReplicaSourceDBInstanceIdentifier"),
			},
			{
				Name:        "replica_mode",
				Description: "The open mode of an Oracle read replica",
				Type:        schema.TypeString,
			},
			{
				Name:        "secondary_availability_zone",
				Description: "If present, specifies the name of the secondary Availability Zone for a DB instance with multi-AZ support.",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB instance is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "storage_type",
				Description: "Specifies the storage type associated with DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A list of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsInstanceTags,
			},
			{
				Name:        "tde_credential_arn",
				Description: "The ARN from the key store with which the instance is associated for TDE encryption.",
				Type:        schema.TypeString,
			},
			{
				Name:        "timezone",
				Description: "The time zone of the DB instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_infos",
				Description: "The status of a read replica. If the instance isn't a read replica, this is  blank.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("StatusInfos"),
			},
			{
				Name:        "associated_roles",
				Description: "The AWS Identity and Access Management (IAM) roles associated with the DB instance",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "db_instance_automated_backups_replications",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("DBInstanceAutomatedBackupsReplications"),
			},
			{
				Name:        "db_parameter_groups",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("DBParameterGroups"),
			},
			{
				Name:        "db_security_groups",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("DBSecurityGroups"),
			},
			{
				Name:        "db_subnet_group",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name:        "domain_memberships",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("DomainMemberships"),
			},
			{
				Name:        "option_group_memberships",
				Type: 			schema.TypeJSON,
				Resolver: 		schema.PathResolver("OptionGroupMemberships"),
			},
			{
				Name:        "vpc_security_groups",
				Type: 			schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRdsInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeDBInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBInstances(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DBInstances
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRdsInstancePendingModifiedValuesProcessorFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	pendingProcessorFeatures := map[string]*string{}
	for _, t := range r.PendingModifiedValues.ProcessorFeatures {
		pendingProcessorFeatures[*t.Name] = t.Value
	}
	return resource.Set(c.Name, pendingProcessorFeatures)
}
func resolveRdsInstanceProcessorFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	processorFeatures := map[string]*string{}
	for _, t := range r.ProcessorFeatures {
		processorFeatures[*t.Name] = t.Value
	}
	return resource.Set(c.Name, processorFeatures)
}
func resolveRdsInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
