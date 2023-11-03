# Table: aws_dax_clusters

This table shows data for Dax Clusters.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html

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
|active_nodes|`int64`|
|cluster_arn|`utf8`|
|cluster_discovery_endpoint|`json`|
|cluster_endpoint_encryption_type|`utf8`|
|cluster_name|`utf8`|
|description|`utf8`|
|iam_role_arn|`utf8`|
|node_ids_to_remove|`list<item: utf8, nullable>`|
|node_type|`utf8`|
|nodes|`json`|
|notification_configuration|`json`|
|parameter_group|`json`|
|preferred_maintenance_window|`utf8`|
|sse_description|`json`|
|security_groups|`json`|
|status|`utf8`|
|subnet_group|`utf8`|
|total_nodes|`int64`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### DynamoDB Accelerator (DAX) clusters should be encrypted at rest

```sql
SELECT
  'DynamoDB Accelerator (DAX) clusters should be encrypted at rest' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN sse_description->>'Status' IS DISTINCT FROM 'ENABLED' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_dax_clusters;
```


