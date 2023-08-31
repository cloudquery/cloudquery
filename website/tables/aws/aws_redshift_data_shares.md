# Table: aws_redshift_data_shares

This table shows data for Redshift Data Shares.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_DataShare.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|allow_publicly_accessible_consumers|`bool`|
|data_share_arn|`utf8`|
|data_share_associations|`json`|
|managed_by|`utf8`|
|producer_arn|`utf8`|