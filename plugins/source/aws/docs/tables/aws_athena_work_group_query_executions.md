
# Table: aws_athena_work_group_query_executions
Information about a single instance of a query execution
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|work_group_cq_id|uuid|Unique CloudQuery ID of aws_athena_work_groups table (FK)|
|effective_engine_version|text|The engine version on which the query runs If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested|
|selected_engine_version|text|The engine version requested by the user|
|execution_parameters|text[]|A list of values for the parameters in a query|
|query|text|The SQL query statements which the query execution ran|
|catalog|text|The name of the data catalog used in the query execution|
|database|text|The name of the database used in the query execution|
|id|text|The unique identifier for each query execution|
|acl_configuration_s3_acl_option|text|The Amazon S3 canned ACL that Athena should specify when storing query results Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL|
|encryption_configuration_encryption_option|text|Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used|
|encryption_configuration_kms_key|text|For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID|
|expected_bucket_owner|text|The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation|
|output_location|text|The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/|
|statement_type|text|The type of query statement that was run|
|data_manifest_location|text|The location and file name of a data manifest file|
|data_scanned_in_bytes|bigint|The number of bytes in the data that was queried|
|engine_execution_time_in_millis|bigint|The number of milliseconds that the query took to execute|
|query_planning_time_in_millis|bigint|The number of milliseconds that Athena took to plan the query processing flow This includes the time spent retrieving table partitions from the data source Note that because the query engine performs the query planning, query planning time is a subset of engine processing time|
|query_queue_time_in_millis|bigint|The number of milliseconds that the query was in your query queue waiting for resources|
|service_processing_time_in_millis|bigint|The number of milliseconds that Athena took to finalize and publish the query results after the query engine finished running the query|
|total_execution_time_in_millis|bigint|The number of milliseconds that Athena took to run the query|
|athena_error_error_category|bigint|An integer value that specifies the category of a query failure error|
|athena_error_error_message|text|Contains a short description of the error that occurred|
|athena_error_error_type|bigint|An integer value that provides specific information about an Athena query error For the meaning of specific values, see the Error Type Reference (https://docsawsamazoncom/athena/latest/ug/error-referencehtml#error-reference-error-type-reference) in the Amazon Athena User Guide|
|athena_error_retryable|boolean|True if the query might succeed if resubmitted|
|completion_date_time|timestamp without time zone|The date and time that the query completed|
|state|text|The state of query execution|
|state_change_reason|text|Further detail about the status of the query|
|submission_date_time|timestamp without time zone|The date and time that the query was submitted|
|work_group|text|The name of the workgroup in which the query ran|
