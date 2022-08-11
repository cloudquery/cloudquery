
# Table: aws_backup_plans
Contains metadata about a backup plan.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|An Amazon Resource Name (ARN) that uniquely identifies a backup plan.|
|id|text|Uniquely identifies a backup plan.|
|name|text|The display name of a saved backup plan.|
|creation_date|timestamp without time zone|The date and time a resource backup plan is created, in Unix format and Coordinated Universal Time (UTC).|
|creator_request_id|text|A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.|
|last_execution_date|timestamp without time zone|The last time a job to back up resources was run with this rule.|
|version_id|text|Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long.|
|advanced_backup_settings|jsonb|Contains a list of backup options for a resource type.|
|tags|jsonb|Resource tags|
