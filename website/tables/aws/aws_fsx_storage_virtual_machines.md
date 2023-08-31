# Table: aws_fsx_storage_virtual_machines

This table shows data for Amazon FSx Storage Virtual Machines.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachine.html

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
|active_directory_configuration|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|endpoints|`json`|
|file_system_id|`utf8`|
|lifecycle|`utf8`|
|lifecycle_transition_reason|`json`|
|name|`utf8`|
|resource_arn|`utf8`|
|root_volume_security_style|`utf8`|
|storage_virtual_machine_id|`utf8`|
|subtype|`utf8`|
|uuid|`utf8`|