
# Table: aws_iot_billing_groups
Billing groups are groups of things created for billing purposes that collect billable information for the things
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|things_in_group|text[]|Lists the things in the specified group|
|tags|jsonb|Tags of the resource|
|arn|text|The ARN of the billing group.|
|id|text|The ID of the billing group.|
|creation_date|timestamp without time zone|The date the billing group was created.|
|name|text|The name of the billing group.|
|description|text|The description of the billing group.|
|version|bigint|The version of the billing group.|
