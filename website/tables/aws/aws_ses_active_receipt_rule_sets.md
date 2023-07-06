# Table: aws_ses_active_receipt_rule_sets

This table shows data for Amazon Simple Email Service (SES) Active Receipt Rule Sets.

https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|rules|`json`|