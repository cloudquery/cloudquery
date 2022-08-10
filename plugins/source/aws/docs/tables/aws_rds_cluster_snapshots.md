
# Table: aws_rds_cluster_snapshots
Contains the details for an Amazon RDS DB cluster snapshot This data type is used as a response element in the DescribeDBClusterSnapshots action.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allocated_storage|integer|Specifies the allocated storage size in gibibytes (GiB).|
|availability_zones|text[]|Provides the list of Availability Zones (AZs) where instances in the DB cluster snapshot can be restored.|
|cluster_create_time|timestamp without time zone|Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).|
|db_cluster_identifier|text|Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.|
|arn|text|The Amazon Resource Name (ARN) for the DB cluster snapshot.|
|db_cluster_snapshot_identifier|text|Specifies the identifier for the DB cluster snapshot.|
|engine|text|Specifies the name of the database engine for this DB cluster snapshot.|
|engine_mode|text|Provides the engine mode of the database engine for this DB cluster snapshot.|
|engine_version|text|Provides the version of the database engine for this DB cluster snapshot.|
|iam_database_authentication_enabled|boolean|True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.|
|kms_key_id|text|If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot|
|license_model|text|Provides the license model information for this DB cluster snapshot.|
|master_username|text|Provides the master username for this DB cluster snapshot.|
|percent_progress|integer|Specifies the percentage of the estimated data that has been transferred.|
|port|integer|Specifies the port that the DB cluster was listening on at the time of the snapshot.|
|snapshot_create_time|timestamp without time zone|Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).|
|snapshot_type|text|Provides the type of the DB cluster snapshot.|
|source_db_cluster_snapshot_arn|text|If the DB cluster snapshot was copied from a source DB cluster snapshot, the Amazon Resource Name (ARN) for the source DB cluster snapshot, otherwise, a null value.|
|status|text|Specifies the status of this DB cluster snapshot.|
|storage_encrypted|boolean|Specifies whether the DB cluster snapshot is encrypted.|
|vpc_id|text|Provides the VPC ID associated with the DB cluster snapshot.|
|tags|jsonb|Resource tags.|
|attributes|jsonb|Snapshot attribute names and values|
