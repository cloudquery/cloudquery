# Table: aws_emr_supported_instance_types

This table shows data for Amazon EMR Supported Instance Types.

https://docs.aws.amazon.com/emr/latest/APIReference/API_SupportedInstanceType.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **release_label**, **type**).
## Relations

This table depends on [aws_emr_release_labels](aws_emr_release_labels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|release_label|`utf8`|
|architecture|`utf8`|
|ebs_optimized_available|`bool`|
|ebs_optimized_by_default|`bool`|
|ebs_storage_only|`bool`|
|instance_family_id|`utf8`|
|is64_bits_only|`bool`|
|memory_gb|`float64`|
|number_of_disks|`int64`|
|storage_gb|`int64`|
|type|`utf8`|
|vcpu|`int64`|