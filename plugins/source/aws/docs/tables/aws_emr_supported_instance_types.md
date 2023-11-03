# Table: aws_emr_supported_instance_types

This table shows data for Amazon EMR Supported Instance Types.

https://docs.aws.amazon.com/emr/latest/APIReference/API_SupportedInstanceType.html

The composite primary key for this table is (**account_id**, **region**, **release_label**, **type**).

## Relations

This table depends on [aws_emr_release_labels](aws_emr_release_labels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|release_label (PK)|`utf8`|
|architecture|`utf8`|
|ebs_optimized_available|`bool`|
|ebs_optimized_by_default|`bool`|
|ebs_storage_only|`bool`|
|instance_family_id|`utf8`|
|is64_bits_only|`bool`|
|memory_gb|`float64`|
|number_of_disks|`int64`|
|storage_gb|`int64`|
|type (PK)|`utf8`|
|vcpu|`int64`|