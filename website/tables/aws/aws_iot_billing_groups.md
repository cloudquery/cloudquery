# Table: aws_iot_billing_groups

This table shows data for AWS IoT Billing Groups.

https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeBillingGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|things_in_group|`list<item: utf8, nullable>`|
|tags|`json`|
|arn (PK)|`utf8`|
|billing_group_arn|`utf8`|
|billing_group_id|`utf8`|
|billing_group_metadata|`json`|
|billing_group_name|`utf8`|
|billing_group_properties|`json`|
|version|`int64`|
|result_metadata|`json`|