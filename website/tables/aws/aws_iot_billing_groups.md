# Table: aws_iot_billing_groups

https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeBillingGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|things_in_group|StringArray|
|tags|JSON|
|arn (PK)|String|
|billing_group_arn|String|
|billing_group_id|String|
|billing_group_metadata|JSON|
|billing_group_name|String|
|billing_group_properties|JSON|
|version|Int|
|result_metadata|JSON|