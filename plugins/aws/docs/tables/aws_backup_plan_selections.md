
# Table: aws_backup_plan_selections
Contains metadata about a BackupSelection object.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|plan_cq_id|uuid|Unique CloudQuery ID of aws_backup_plan table (FK)|
|creation_date|timestamp without time zone|The date and time a backup plan is created, in Unix format and Coordinated Universal Time (UTC).|
|creator_request_id|text|A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.|
|iam_role_arn|text|Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery point; for example, arn:aws:iam::123456789012:role/S3Access.|
|selection_id|text|Uniquely identifies a request to assign a set of resources to a backup plan.|
|selection_name|text|The display name of a resource selection document.|
|conditions|jsonb|A list of conditions that you define to assign resources to your backup plans using tags.|
|list_of_tags|jsonb|A list of conditions that you define to assign resources to your backup plans using tags.|
|not_resources|text[]|A list of Amazon Resource Names (ARNs) to exclude from a backup plan.|
|resources|text[]|A list of Amazon Resource Names (ARNs) to assign to a backup plan.|
