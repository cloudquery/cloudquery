# Table: aws_ecs_task_definitions

This table shows data for Amazon Elastic Container Service (ECS) Task Definitions.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|compatibilities|`list<item: utf8, nullable>`|
|container_definitions|`json`|
|cpu|`utf8`|
|deregistered_at|`timestamp[us, tz=UTC]`|
|ephemeral_storage|`json`|
|execution_role_arn|`utf8`|
|family|`utf8`|
|inference_accelerators|`json`|
|ipc_mode|`utf8`|
|memory|`utf8`|
|network_mode|`utf8`|
|pid_mode|`utf8`|
|placement_constraints|`json`|
|proxy_configuration|`json`|
|registered_at|`timestamp[us, tz=UTC]`|
|registered_by|`utf8`|
|requires_attributes|`json`|
|requires_compatibilities|`list<item: utf8, nullable>`|
|revision|`int64`|
|runtime_platform|`json`|
|status|`utf8`|
|task_definition_arn|`utf8`|
|task_role_arn|`utf8`|
|volumes|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon ECS task definitions should have secure networking modes and user definitions

```sql
SELECT
  'Amazon ECS task definitions should have secure networking modes and user definitions'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN network_mode = 'host'
  AND (c->>'Privileged')::BOOL IS NOT true
  AND (c->>'User' = 'root' OR (c->>'User') IS NULL)
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ecs_task_definitions,
  jsonb_array_elements(aws_ecs_task_definitions.container_definitions) AS c;
```


