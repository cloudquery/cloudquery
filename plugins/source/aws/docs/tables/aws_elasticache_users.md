
# Table: aws_elasticache_users
Describes Elasticache users
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the user.|
|access_string|text|Access permissions string used for this user.|
|authentication_password_count|bigint|The number of passwords belonging to the user|
|authentication_type|text|Indicates whether the user requires a password to authenticate.|
|engine|text|The current supported value is Redis.|
|minimum_engine_version|text|The minimum engine version required, which is Redis 6.0|
|status|text|Indicates the user status|
|user_group_ids|text[]|Returns a list of the user group IDs the user belongs to.|
|user_id|text|The ID of the user.|
|user_name|text|The username of the user.|
