# Table: aws_glue_dev_endpoints



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|arguments|JSON|
|availability_zone|String|
|created_timestamp|Timestamp|
|endpoint_name|String|
|extra_jars_s3_path|String|
|extra_python_libs_s3_path|String|
|failure_reason|String|
|glue_version|String|
|last_modified_timestamp|Timestamp|
|last_update_status|String|
|number_of_nodes|Int|
|number_of_workers|Int|
|private_address|String|
|public_address|String|
|public_key|String|
|public_keys|StringArray|
|role_arn|String|
|security_configuration|String|
|security_group_ids|StringArray|
|status|String|
|subnet_id|String|
|vpc_id|String|
|worker_type|String|
|yarn_endpoint_address|String|
|zeppelin_remote_spark_interpreter_port|Int|