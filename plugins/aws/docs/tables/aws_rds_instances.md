
# Table: aws_rds_instances
Contains the details of an Amazon RDS DB instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allocated_storage|integer|Specifies the allocated storage size specified in gibibytes.|
|auto_minor_version_upgrade|boolean|A value that indicates that minor version patches are applied automatically.|
|availability_zone|text|Specifies the name of the Availability Zone the DB instance is located in.|
|aws_backup_recovery_point_arn|text|The Amazon Resource Name (ARN) of the recovery point in AWS Backup.|
|backup_retention_period|integer|Specifies the number of days for which automatic DB snapshots are retained.|
|ca_certificate_identifier|text|The identifier of the CA certificate for this DB instance.|
|character_set_name|text|If present, specifies the name of the character set that this instance is associated with.|
|copy_tags_to_snapshot|boolean|Specifies whether tags are copied from the DB instance to snapshots of the DB instance|
|customer_owned_ip_enabled|boolean|Specifies whether a customer-owned IP address (CoIP) is enabled for an RDS on Outposts DB instance|
|cluster_identifier|text|If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.|
|arn|text|The Amazon Resource Name (ARN) for the DB instance.|
|db_instance_class|text|Contains the name of the compute and memory capacity class of the DB instance.|
|user_instance_id|text|Contains a user-supplied database identifier|
|db_instance_status|text|Specifies the current state of this database|
|db_name|text|The meaning of this parameter differs according to the database engine you use. MySQL, MariaDB, SQL Server, PostgreSQL Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created|
|subnet_group_arn|text|The Amazon Resource Name (ARN) for the DB subnet group.|
|subnet_group_description|text|Provides the description of the DB subnet group.|
|subnet_group_name|text|The name of the DB subnet group.|
|subnet_group_subnet_group_status|text|Provides the status of the DB subnet group.|
|subnet_group_vpc_id|text|Provides the VpcId of the DB subnet group.|
|instance_port|integer|Specifies the port that the DB instance listens on|
|id|text|The AWS Region-unique, immutable identifier for the DB instance|
|deletion_protection|boolean|Indicates if the DB instance has deletion protection enabled|
|enabled_cloudwatch_logs_exports|text[]|A list of log types that this DB instance is configured to export to CloudWatch Logs|
|endpoint_address|text|Specifies the DNS address of the DB instance.|
|endpoint_hosted_zone_id|text|Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.|
|endpoint_port|integer|Specifies the port that the database engine is listening on.|
|engine|text|The name of the database engine to be used for this DB instance.|
|engine_version|text|Indicates the database engine version.|
|enhanced_monitoring_resource_arn|text|The Amazon Resource Name (ARN) of the Amazon CloudWatch Logs log stream that receives the Enhanced Monitoring metrics data for the DB instance.|
|iam_database_authentication_enabled|boolean|True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false|
|instance_create_time|timestamp without time zone|Provides the date and time the DB instance was created.|
|iops|integer|Specifies the Provisioned IOPS (I/O operations per second) value.|
|kms_key_id|text|If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB instance|
|latest_restorable_time|timestamp without time zone|Specifies the latest time to which a database can be restored with point-in-time restore.|
|license_model|text|License model information for this DB instance.|
|listener_endpoint_address|text|Specifies the DNS address of the DB instance.|
|listener_endpoint_hosted_zone_id|text|Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.|
|listener_endpoint_port|integer|Specifies the port that the database engine is listening on.|
|master_username|text|Contains the master username for the DB instance.|
|max_allocated_storage|integer|The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.|
|monitoring_interval|integer|The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.|
|monitoring_role_arn|text|The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.|
|multi_az|boolean|Specifies if the DB instance is a Multi-AZ deployment.|
|nchar_character_set_name|text|The name of the NCHAR character set for the Oracle DB instance|
|pending_modified_values_allocated_storage|integer|The allocated storage size for the DB instance specified in gigabytes .|
|pending_modified_values_backup_retention_period|integer|The number of days for which automated backups are retained.|
|pending_modified_values_ca_certificate_identifier|text|The identifier of the CA certificate for the DB instance.|
|pending_modified_values_db_instance_class|text|The name of the compute and memory capacity class for the DB instance.|
|pending_modified_values_db_instance_identifier|text|The database identifier for the DB instance.|
|pending_modified_values_db_subnet_group_name|text|The DB subnet group for the DB instance.|
|pending_modified_values_engine_version|text|The database engine version.|
|pending_modified_values_iam_database_authentication_enabled|boolean|Whether mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.|
|pending_modified_values_iops|integer|The Provisioned IOPS value for the DB instance.|
|pending_modified_values_license_model|text|The license model for the DB instance|
|pending_modified_values_master_user_password|text|The master credentials for the DB instance.|
|pending_modified_values_multi_az|boolean|A value that indicates that the Single-AZ DB instance will change to a Multi-AZ deployment.|
|pending_cloudwatch_logs_types_to_disable|text[]|Log types that are in the process of being enabled|
|pending_cloudwatch_logs_types_to_enable|text[]|Log types that are in the process of being deactivated|
|pending_modified_values_port|integer|The port for the DB instance.|
|pending_modified_values_processor_features|jsonb|The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.|
|pending_modified_values_storage_type|text|The storage type of the DB instance.|
|performance_insights_enabled|boolean|True if Performance Insights is enabled for the DB instance, and otherwise false.|
|performance_insights_kms_key_id|text|The AWS KMS key identifier for encryption of Performance Insights data|
|performance_insights_retention_period|integer|The amount of time, in days, to retain Performance Insights data|
|preferred_backup_window|text|Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.|
|preferred_maintenance_window|text|Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).|
|processor_features|jsonb|The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.|
|promotion_tier|integer|A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance|
|publicly_accessible|boolean|Specifies the accessibility options for the DB instance|
|read_replica_db_cluster_identifiers|text[]|Contains one or more identifiers of Aurora DB clusters to which the RDS DB instance is replicated as a read replica|
|read_replica_db_instance_identifiers|text[]|Contains one or more identifiers of the read replicas associated with this DB instance.|
|read_replica_source_db_instance_identifier|text|Contains the identifier of the source DB instance if this DB instance is a read replica.|
|replica_mode|text|The open mode of an Oracle read replica|
|secondary_availability_zone|text|If present, specifies the name of the secondary Availability Zone for a DB instance with multi-AZ support.|
|storage_encrypted|boolean|Specifies whether the DB instance is encrypted.|
|storage_type|text|Specifies the storage type associated with DB instance.|
|tags|jsonb|A list of tags|
|tde_credential_arn|text|The ARN from the key store with which the instance is associated for TDE encryption.|
|timezone|text|The time zone of the DB instance|
|status_infos|jsonb|The status of a read replica. If the instance isn't a read replica, this is  blank.|
