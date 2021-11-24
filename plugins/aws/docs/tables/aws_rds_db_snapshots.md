
# Table: aws_rds_db_snapshots
Contains the details of an Amazon RDS DB snapshot
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allocated_storage|integer|Specifies the allocated storage size in gibibytes (GiB).|
|availability_zone|text|Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.|
|db_instance_identifier|text|Specifies the DB instance identifier of the DB instance this DB snapshot was created from.|
|arn|text|The Amazon Resource Name (ARN) for the DB snapshot.|
|db_snapshot_identifier|text|Specifies the identifier for the DB snapshot.|
|dbi_resource_id|text|The identifier for the source DB instance, which can't be changed and which is unique to an AWS Region.|
|encrypted|boolean|Specifies whether the DB snapshot is encrypted.|
|engine|text|Specifies the name of the database engine.|
|engine_version|text|Specifies the version of the database engine.|
|iam_database_authentication_enabled|boolean|True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.|
|instance_create_time|timestamp without time zone|Specifies the time in Coordinated Universal Time (UTC) when the DB instance, from which the snapshot was taken, was created.|
|iops|integer|Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.|
|kms_key_id|text|If Encrypted is true, the AWS KMS key identifier for the encrypted DB snapshot. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the AWS KMS customer master key (CMK).|
|license_model|text|License model information for the restored DB instance.|
|master_username|text|Provides the master username for the DB snapshot.|
|option_group_name|text|Provides the option group name for the DB snapshot.|
|percent_progress|integer|The percentage of the estimated data that has been transferred.|
|port|integer|Specifies the port that the database engine was listening on at the time of the snapshot.|
|processor_features|jsonb|The number of CPU cores and the number of threads per core for the DB instance class of the DB instance when the DB snapshot was created.|
|snapshot_create_time|timestamp without time zone|Specifies when the snapshot was taken in Coordinated Universal Time (UTC).|
|snapshot_type|text|Provides the type of the DB snapshot.|
|source_db_snapshot_identifier|text|The DB snapshot Amazon Resource Name (ARN) that the DB snapshot was copied from. It only has value in case of cross-customer or cross-region copy.|
|source_region|text|The AWS Region that the DB snapshot was created in or copied from.|
|status|text|Specifies the status of this DB snapshot.|
|storage_type|text|Specifies the storage type associated with DB snapshot.|
|tde_credential_arn|text|The ARN from the key store with which to associate the instance for TDE encryption.|
|timezone|text|The time zone of the DB snapshot|
|vpc_id|text|Provides the VPC ID associated with the DB snapshot.|
|tags|jsonb|Resource tags.|
|attributes|jsonb|Snapshot attribute names and values|
