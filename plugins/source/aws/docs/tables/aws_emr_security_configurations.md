# Table: aws_emr_security_configurations

https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|security_configuration|JSON|
|creation_date_time|Timestamp|
|name (PK)|String|