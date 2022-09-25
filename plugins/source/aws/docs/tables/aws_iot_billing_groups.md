# Table: aws_iot_billing_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|things_in_group|StringArray|
|tags|JSON|
|arn (PK)|String|
|billing_group_id|String|
|billing_group_metadata|JSON|
|billing_group_name|String|
|billing_group_properties|JSON|
|version|Int|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|