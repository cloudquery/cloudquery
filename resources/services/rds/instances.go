package rds

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsInstances() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_instances",
		Description:  "Contains the details of an Amazon RDS DB instance",
		Resolver:     fetchRdsInstances,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "subnet_group_arn",
				Description: "The Amazon Resource Name (ARN) for the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup.DBSubnetGroupArn"),
			},
			{
				Name:        "subnet_group_description",
				Description: "Provides the description of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup.DBSubnetGroupDescription"),
			},
			{
				Name:        "subnet_group_name",
				Description: "The name of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup.DBSubnetGroupName"),
			},
			{
				Name:        "subnet_group_subnet_group_status",
				Description: "Provides the status of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup.SubnetGroupStatus"),
			},
			{
				Name:        "subnet_group_vpc_id",
				Description: "Provides the VpcId of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroup.VpcId"),
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
				Name:        "endpoint_address",
				Description: "Specifies the DNS address of the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoint.Address"),
			},
			{
				Name:        "endpoint_hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoint.HostedZoneId"),
			},
			{
				Name:        "endpoint_port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Endpoint.Port"),
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
				Name:        "listener_endpoint_address",
				Description: "Specifies the DNS address of the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ListenerEndpoint.Address"),
			},
			{
				Name:        "listener_endpoint_hosted_zone_id",
				Description: "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ListenerEndpoint.HostedZoneId"),
			},
			{
				Name:        "listener_endpoint_port",
				Description: "Specifies the port that the database engine is listening on.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ListenerEndpoint.Port"),
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
				Name:        "pending_modified_values_allocated_storage",
				Description: "The allocated storage size for the DB instance specified in gigabytes .",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.AllocatedStorage"),
			},
			{
				Name:        "pending_modified_values_backup_retention_period",
				Description: "The number of days for which automated backups are retained.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.BackupRetentionPeriod"),
			},
			{
				Name:        "pending_modified_values_ca_certificate_identifier",
				Description: "The identifier of the CA certificate for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.CACertificateIdentifier"),
			},
			{
				Name:        "pending_modified_values_db_instance_class",
				Description: "The name of the compute and memory capacity class for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.DBInstanceClass"),
			},
			{
				Name:        "pending_modified_values_db_instance_identifier",
				Description: "The database identifier for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.DBInstanceIdentifier"),
			},
			{
				Name:        "pending_modified_values_db_subnet_group_name",
				Description: "The DB subnet group for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.DBSubnetGroupName"),
			},
			{
				Name:        "pending_modified_values_engine_version",
				Description: "The database engine version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.EngineVersion"),
			},
			{
				Name:        "pending_modified_values_iam_database_authentication_enabled",
				Description: "Whether mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PendingModifiedValues.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "pending_modified_values_iops",
				Description: "The Provisioned IOPS value for the DB instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.Iops"),
			},
			{
				Name:        "pending_modified_values_license_model",
				Description: "The license model for the DB instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.LicenseModel"),
			},
			{
				Name:        "pending_modified_values_master_user_password",
				Description: "The master credentials for the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
			},
			{
				Name:        "pending_modified_values_multi_az",
				Description: "A value that indicates that the Single-AZ DB instance will change to a Multi-AZ deployment.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PendingModifiedValues.MultiAZ"),
			},
			{
				Name:        "pending_cloudwatch_logs_types_to_disable",
				Description: "Log types that are in the process of being enabled",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable"),
			},
			{
				Name:        "pending_cloudwatch_logs_types_to_enable",
				Description: "Log types that are in the process of being deactivated",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable"),
			},
			{
				Name:        "pending_modified_values_port",
				Description: "The port for the DB instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.Port"),
			},
			{
				Name:        "pending_modified_values_processor_features",
				Description: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsInstancePendingModifiedValuesProcessorFeatures,
			},
			{
				Name:        "pending_modified_values_storage_type",
				Description: "The storage type of the DB instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.StorageType"),
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
				Resolver:    resolveRdsInstanceStatusInfos,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_rds_instance_associated_roles",
				Description: "Describes an AWS Identity and Access Management (IAM) role that is associated with a DB instance. ",
				Resolver:    fetchRdsInstanceAssociatedRoles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "role_arn"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "feature_name",
						Description: "The name of the feature associated with the AWS Identity and Access Management (IAM) role",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Describes the state of association between the IAM role and the DB instance",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_db_instance_automated_backups_replications",
				Description: "Automated backups of a DB instance replicated to another AWS Region",
				Resolver:    fetchRdsInstanceDbInstanceAutomatedBackupsReplications,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "db_instance_automated_backups_arn"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "db_instance_automated_backups_arn",
						Description: "The Amazon Resource Name (ARN) of the replicated automated backups.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DBInstanceAutomatedBackupsArn"),
					},
				},
			},
			{
				Name:        "aws_rds_instance_db_parameter_groups",
				Description: "The status of the DB parameter group",
				Resolver:    fetchRdsInstanceDbParameterGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "db_parameter_group_name"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "db_parameter_group_name",
						Description: "The name of the DB parameter group.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DBParameterGroupName"),
					},
					{
						Name:        "parameter_apply_status",
						Description: "The status of parameter updates.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_db_security_groups",
				Description: "This data type is used as a response element in the following actions:  * ModifyDBInstance  * RebootDBInstance  * RestoreDBInstanceFromDBSnapshot  * RestoreDBInstanceToPointInTime ",
				Resolver:    fetchRdsInstanceDbSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "db_security_group_name"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "db_security_group_name",
						Description: "The name of the DB security group.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DBSecurityGroupName"),
					},
					{
						Name:        "status",
						Description: "The status of the DB security group.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_db_subnet_group_subnets",
				Description: "This data type is used as a response element for the DescribeDBSubnetGroups operation. ",
				Resolver:    fetchRdsInstanceDbSubnetGroupSubnets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "subnet_identifier"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "subnet_availability_zone_name",
						Description: "The name of the Availability Zone.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:        "subnet_identifier",
						Description: "The identifier of the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_outpost_arn",
						Description: "The Amazon Resource Name (ARN) of the Outpost.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetOutpost.Arn"),
					},
					{
						Name:        "subnet_status",
						Description: "The status of the subnet.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_domain_memberships",
				Description: "An Active Directory Domain membership record associated with the DB instance or cluster. ",
				Resolver:    fetchRdsInstanceDomainMemberships,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "domain"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "domain",
						Description: "The identifier of the Active Directory Domain.",
						Type:        schema.TypeString,
					},
					{
						Name:        "fqdn",
						Description: "The fully qualified domain name of the Active Directory Domain.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FQDN"),
					},
					{
						Name:        "iam_role_name",
						Description: "The name of the IAM role to be used when making API calls to the Directory Service.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IAMRoleName"),
					},
					{
						Name:        "status",
						Description: "The status of the Active Directory Domain membership for the DB instance or cluster",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_option_group_memberships",
				Description: "Provides information on the option groups the DB instance is a member of. ",
				Resolver:    fetchRdsInstanceOptionGroupMemberships,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "option_group_name"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "option_group_name",
						Description: "The name of the option group that the instance belongs to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the DB instance's option group membership",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_rds_instance_vpc_security_groups",
				Description: "This data type is used as a response element for queries on VPC security group membership. ",
				Resolver:    fetchRdsInstanceVpcSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "vpc_security_group_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The AWS Region-unique, immutable identifier for the DB instance",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "status",
						Description: "The status of the VPC security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_security_group_id",
						Description: "The name of the VPC security group.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeDBInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBInstances(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
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
func fetchRdsInstanceAssociatedRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.AssociatedRoles
	return nil
}
func fetchRdsInstanceDbInstanceAutomatedBackupsReplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBInstanceAutomatedBackupsReplications
	return nil
}
func fetchRdsInstanceDbParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBParameterGroups
	return nil
}
func fetchRdsInstanceDbSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBSecurityGroups
	return nil
}
func fetchRdsInstanceDbSubnetGroupSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBSubnetGroup.Subnets
	return nil
}
func fetchRdsInstanceDomainMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DomainMemberships
	return nil
}
func fetchRdsInstanceOptionGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.OptionGroupMemberships
	return nil
}
func resolveRdsInstanceStatusInfos(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance, ok := resource.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	data, err := json.Marshal(instance.StatusInfos)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchRdsInstanceVpcSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.VpcSecurityGroups
	return nil
}
