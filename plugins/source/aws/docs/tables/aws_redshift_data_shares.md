# Table: aws_redshift_data_shares

This table shows data for Redshift Data Shares.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_DataShare.html

The composite primary key for this table is (**arn**, **producer_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|allow_publicly_accessible_consumers|`bool`|
|data_share_arn|`utf8`|
|data_share_associations|`json`|
|managed_by|`utf8`|
|producer_arn (PK)|`utf8`|