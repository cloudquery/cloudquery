
# Table: aws_rds_clusters
Contains the details of an Amazon Aurora DB cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|activity_stream_kinesis_stream_name|text|The name of the Amazon Kinesis data stream used for the database activity stream.|
|activity_stream_kms_key_id|text|The AWS KMS key identifier used for encrypting messages in the database activity stream|
|activity_stream_mode|text|The mode of the database activity stream|
|activity_stream_status|text|The status of the database activity stream.|
|allocated_storage|integer|For all database engines except Amazon Aurora, AllocatedStorage specifies the allocated storage size in gibibytes (GiB)|
|availability_zones|text[]|Provides the list of Availability Zones (AZs) where instances in the DB cluster can be created.|
|backtrack_consumed_change_records|bigint|The number of change records stored for Backtrack.|
|backtrack_window|bigint|The target backtrack window, in seconds|
|backup_retention_period|integer|Specifies the number of days for which automatic DB snapshots are retained.|
|capacity|integer|The current capacity of an Aurora Serverless DB cluster|
|character_set_name|text|If present, specifies the name of the character set that this cluster is associated with.|
|clone_group_id|text|Identifies the clone group to which the DB cluster is associated.|
|cluster_create_time|timestamp without time zone|Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).|
|copy_tags_to_snapshot|boolean|Specifies whether tags are copied from the DB cluster to snapshots of the DB cluster.|
|cross_account_clone|boolean|Specifies whether the DB cluster is a clone of a DB cluster owned by a different AWS account.|
|custom_endpoints|text[]|Identifies all custom endpoints associated with the cluster.|
|arn|text|The Amazon Resource Name (ARN) for the DB cluster.|
|db_cluster_identifier|text|Contains a user-supplied DB cluster identifier|
|db_cluster_parameter_group|text|Specifies the name of the DB cluster parameter group for the DB cluster.|
|db_cluster_option_group_memberships|jsonb|Provides the map of option group memberships for this DB cluster.|
|db_subnet_group|text|Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.|
|database_name|text|Contains the name of the initial database of this DB cluster that was provided at create time, if one was specified when the DB cluster was created|
|id|text|The AWS Region-unique, immutable identifier for the DB cluster|
|deletion_protection|boolean|Indicates if the DB cluster has deletion protection enabled|
|earliest_backtrack_time|timestamp without time zone|The earliest time to which a DB cluster can be backtracked.|
|earliest_restorable_time|timestamp without time zone|The earliest time to which a database can be restored with point-in-time restore.|
|enabled_cloudwatch_logs_exports|text[]|A list of log types that this DB cluster is configured to export to CloudWatch Logs|
|endpoint|text|Specifies the connection endpoint for the primary instance of the DB cluster.|
|engine|text|The name of the database engine to be used for this DB cluster.|
|engine_mode|text|The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery, global, or multimaster|
|engine_version|text|Indicates the database engine version.|
|global_write_forwarding_requested|boolean|Specifies whether you have requested to enable write forwarding for a secondary cluster in an Aurora global database|
|global_write_forwarding_status|text|Specifies whether a secondary cluster in an Aurora global database has write forwarding enabled, not enabled, or is in the process of enabling it.|
|hosted_zone_id|text|Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.|
|http_endpoint_enabled|boolean|A value that indicates whether the HTTP endpoint for an Aurora Serverless DB cluster is enabled|
|iam_database_authentication_enabled|boolean|A value that indicates whether the mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.|
|kms_key_id|text|If StorageEncrypted is enabled, the AWS KMS key identifier for the encrypted DB cluster|
|latest_restorable_time|timestamp without time zone|Specifies the latest time to which a database can be restored with point-in-time restore.|
|master_username|text|Contains the master username for the DB cluster.|
|multi_az|boolean|Specifies whether the DB cluster has instances in multiple Availability Zones.|
|pending_modified_values_db_cluster_identifier|text|The DBClusterIdentifier value for the DB cluster.|
|pending_modified_values_engine_version|text|The database engine version.|
|pending_modified_values_iam_database_authentication_enabled|boolean|A value that indicates whether mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.|
|pending_modified_values_master_user_password|text|The master credentials for the DB cluster.|
|pending_cloudwatch_logs_types_to_disable|text[]|Log types that are in the process of being enabled|
|pending_cloudwatch_logs_types_to_enable|text[]|Log types that are in the process of being deactivated|
|percent_progress|text|Specifies the progress of the operation as a percentage.|
|port|integer|Specifies the port that the database engine is listening on.|
|preferred_backup_window|text|Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.|
|preferred_maintenance_window|text|Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).|
|read_replica_identifiers|text[]|Contains one or more identifiers of the read replicas associated with this DB cluster.|
|reader_endpoint|text|The reader endpoint for the DB cluster|
|replication_source_identifier|text|Contains the identifier of the source DB cluster if this DB cluster is a read replica.|
|scaling_configuration_info_auto_pause|boolean|A value that indicates whether automatic pause is allowed for the Aurora DB cluster in serverless DB engine mode|
|scaling_configuration_info_max_capacity|integer|The maximum capacity for an Aurora DB cluster in serverless DB engine mode.|
|scaling_configuration_info_min_capacity|integer|The maximum capacity for the Aurora DB cluster in serverless DB engine mode.|
|scaling_configuration_info_seconds_until_auto_pause|integer|The remaining amount of time, in seconds, before the Aurora DB cluster in serverless mode is paused|
|scaling_configuration_info_timeout_action|text|The timeout action of a call to ModifyCurrentDBClusterCapacity, either ForceApplyCapacityChange or RollbackCapacityChange.|
|status|text|Specifies the current state of this DB cluster.|
|storage_encrypted|boolean|Specifies whether the DB cluster is encrypted.|
|tags|jsonb|A list of tags|
