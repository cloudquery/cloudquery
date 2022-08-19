
# Table: aws_elasticache_user_groups
Describes Elasticache user groups
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the user group.|
|engine|text|The current supported value is Redis.|
|minimum_engine_version|text|The minimum engine version required, which is Redis 6.0|
|pending_user_ids_to_add|text[]|The list of user IDs to add.|
|pending_user_ids_to_remove|text[]|The list of user IDs to remove.|
|replication_groups|text[]|A list of replication groups that the user group can access.|
|status|text|Indicates user group status|
|user_group_id|text|The ID of the user group.|
|user_ids|text[]|The list of user IDs that belong to the user group.|
