# Table: aws_sagemaker_notebook_instances



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
|accelerator_types|StringArray|
|additional_code_repositories|StringArray|
|creation_time|Timestamp|
|default_code_repository|String|
|direct_internet_access|String|
|failure_reason|String|
|instance_metadata_service_configuration|JSON|
|instance_type|String|
|kms_key_id|String|
|last_modified_time|Timestamp|
|network_interface_id|String|
|notebook_instance_lifecycle_config_name|String|
|notebook_instance_name|String|
|notebook_instance_status|String|
|platform_identifier|String|
|role_arn|String|
|root_access|String|
|security_groups|StringArray|
|subnet_id|String|
|url|String|
|volume_size_in_gb|Int|
|result_metadata|JSON|