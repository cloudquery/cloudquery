# Table: aws_emr_release_labels

This table shows data for Amazon EMR Release Labels.

https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeReleaseLabel.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **release_label**).
## Relations

The following tables depend on aws_emr_release_labels:
  - [aws_emr_supported_instance_types](aws_emr_supported_instance_types.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|applications|`json`|
|available_os_releases|`json`|
|release_label|`utf8`|