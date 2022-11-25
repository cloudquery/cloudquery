# Table: aws_fsx_storage_virtual_machines

https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachine.html

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
|active_directory_configuration|JSON|
|creation_time|Timestamp|
|endpoints|JSON|
|file_system_id|String|
|lifecycle|String|
|lifecycle_transition_reason|JSON|
|name|String|
|root_volume_security_style|String|
|storage_virtual_machine_id|String|
|subtype|String|
|tags|JSON|
|uuid|String|