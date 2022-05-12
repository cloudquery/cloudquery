
# Table: aws_athena_work_groups
A workgroup, which contains a name, description, creation time, state, and other configuration, listed under WorkGroup$Configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|arn|text|ARN of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|name|text|The workgroup name|
|bytes_scanned_cutoff_per_query|bigint|The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan|
|enforce_work_group_configuration|boolean|If set to "true", the settings for the workgroup override client-side settings If set to "false", client-side settings are used|
|effective_engine_version|text|The engine version on which the query runs If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested|
|selected_engine_version|text|The engine version requested by the user|
|publish_cloud_watch_metrics_enabled|boolean|Indicates that the Amazon CloudWatch metrics are enabled for the workgroup|
|requester_pays_enabled|boolean|If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries|
|acl_configuration_s3_acl_option|text|The Amazon S3 canned ACL that Athena should specify when storing query results Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL|
|encryption_configuration_encryption_option|text|Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used|
|encryption_configuration_kms_key|text|For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID|
|expected_bucket_owner|text|The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation|
|output_location|text|The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/|
|creation_time|timestamp without time zone|The date and time the workgroup was created|
|description|text|The workgroup description|
|state|text|The state of the workgroup: ENABLED or DISABLED|
