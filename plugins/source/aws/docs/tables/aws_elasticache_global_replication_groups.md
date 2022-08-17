
# Table: aws_elasticache_global_replication_groups
Consists of a primary cluster that accepts writes and an associated secondary cluster that resides in a different Amazon region
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|arn|text|The ARN (Amazon Resource Name) of the global replication group.|
|at_rest_encryption_enabled|boolean|A flag that enables encryption at rest when set to true|
|auth_token_enabled|boolean|A flag that enables using an AuthToken (password) when issuing Redis commands. Default: false|
|cache_node_type|text|The cache node type of the Global datastore|
|cluster_enabled|boolean|A flag that indicates whether the Global datastore is cluster enabled.|
|engine|text|The Elasticache engine|
|engine_version|text|The Elasticache Redis engine version.|
|global_replication_group_description|text|The optional description of the Global datastore|
|global_replication_group_id|text|The name of the Global datastore|
|status|text|The status of the Global datastore|
|transit_encryption_enabled|boolean|A flag that enables in-transit encryption when set to true|
