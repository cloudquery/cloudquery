# Table: aws_sagemaker_notebook_instances

This table shows data for Amazon SageMaker Notebook Instances.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeNotebookInstance.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|accelerator_types|`list<item: utf8, nullable>`|
|additional_code_repositories|`list<item: utf8, nullable>`|
|creation_time|`timestamp[us, tz=UTC]`|
|default_code_repository|`utf8`|
|direct_internet_access|`utf8`|
|failure_reason|`utf8`|
|instance_metadata_service_configuration|`json`|
|instance_type|`utf8`|
|kms_key_id|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|network_interface_id|`utf8`|
|notebook_instance_arn|`utf8`|
|notebook_instance_lifecycle_config_name|`utf8`|
|notebook_instance_name|`utf8`|
|notebook_instance_status|`utf8`|
|platform_identifier|`utf8`|
|role_arn|`utf8`|
|root_access|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|subnet_id|`utf8`|
|url|`utf8`|
|volume_size_in_gb|`int64`|