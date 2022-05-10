
# Table: aws_shield_protection_groups
A grouping of protected resources that you and Shield Advanced can monitor as a collective
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|tags|jsonb||
|aggregation|text|Defines how Shield combines resource data for the group in order to detect, mitigate, and report events  * Sum - Use the total traffic across the group This is a good choice for most cases|
|members|text[]|The Amazon Resource Names (ARNs) of the resources to include in the protection group|
|pattern|text|The criteria to use to choose the protected resources for inclusion in the group|
|id|text|The name of the protection group|
|arn|text|The ARN (Amazon Resource Name) of the protection group|
|resource_type|text|The resource type to include in the protection group|
