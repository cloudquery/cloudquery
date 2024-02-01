# Table: aws_emr_security_configurations

This table shows data for Amazon EMR Security Configurations.

https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|security_configuration|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|