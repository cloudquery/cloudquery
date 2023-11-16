# Table: aws_glue_dev_endpoints

This table shows data for Glue Dev Endpoints.

https://docs.aws.amazon.com/glue/latest/webapi/API_DevEndpoint.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|arguments|`json`|
|availability_zone|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|endpoint_name|`utf8`|
|extra_jars_s3_path|`utf8`|
|extra_python_libs_s3_path|`utf8`|
|failure_reason|`utf8`|
|glue_version|`utf8`|
|last_modified_timestamp|`timestamp[us, tz=UTC]`|
|last_update_status|`utf8`|
|number_of_nodes|`int64`|
|number_of_workers|`int64`|
|private_address|`utf8`|
|public_address|`utf8`|
|public_key|`utf8`|
|public_keys|`list<item: utf8, nullable>`|
|role_arn|`utf8`|
|security_configuration|`utf8`|
|security_group_ids|`list<item: utf8, nullable>`|
|status|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|
|worker_type|`utf8`|
|yarn_endpoint_address|`utf8`|
|zeppelin_remote_spark_interpreter_port|`int64`|