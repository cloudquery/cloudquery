# Table: aws_emr_security_configurations

This table shows data for Amazon EMR Security Configurations.

https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|security_configuration|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|name (PK)|`utf8`|