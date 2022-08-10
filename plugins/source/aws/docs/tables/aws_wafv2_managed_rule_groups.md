
# Table: aws_wafv2_managed_rule_groups
High-level information about a managed rule group, returned by ListAvailableManagedRuleGroups
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|scope|text|The scope (Regional or Global) of the resource.|
|available_labels|text[]||
|consumed_labels|text[]||
|capacity|bigint||
|label_namespace|text||
|rules|jsonb||
|description|text|The description of the managed rule group, provided by AWS Managed Rules or the AWS Marketplace seller who manages it.|
|name|text|The name of the managed rule group|
|vendor_name|text|The name of the managed rule group vendor|
