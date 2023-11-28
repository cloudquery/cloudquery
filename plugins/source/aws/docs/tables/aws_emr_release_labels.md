# Table: aws_emr_release_labels

This table shows data for Amazon EMR Release Labels.

https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeReleaseLabel.html

The composite primary key for this table is (**account_id**, **region**, **release_label**).

## Relations

The following tables depend on aws_emr_release_labels:
  - [aws_emr_supported_instance_types](aws_emr_supported_instance_types.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|applications|`json`|
|available_os_releases|`json`|
|release_label (PK)|`utf8`|