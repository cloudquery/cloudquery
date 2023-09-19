# Table: aws_dynamodb_table_replica_auto_scalings

This table shows data for Amazon DynamoDB Table Replica Auto Scalings.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_dynamodb_tables](aws_dynamodb_tables).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|table_arn|`utf8`|
|global_secondary_indexes|`json`|
|region_name|`utf8`|
|replica_provisioned_read_capacity_auto_scaling_settings|`json`|
|replica_provisioned_write_capacity_auto_scaling_settings|`json`|
|replica_status|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### DynamoDB tables should automatically scale capacity with demand

```sql
SELECT
    'DynamoDB tables should automatically scale capacity with demand' as title,
    t.account_id,
		t.arn AS resource_id,
		  case 
		    WHEN t.billing_mode_summary->>'BillingMode' = 'PAY_PER_REQUEST' then 'pass'
		    WHEN (t.billing_mode_summary->>'BillingMode' = 'PROVISIONED' or t.billing_mode_summary->>'BillingMode' is NULL )and (pr._cq_id IS not NULL and pw._cq_id IS not NULL) then 'pass'
		    ELSE 'fail'
		    END
		        AS status
FROM
  aws_dynamodb_tables AS t
  LEFT JOIN aws_applicationautoscaling_policies AS pr ON
      pr.service_namespace = 'dynamodb'
      AND pr.resource_id = concat('table/', t.table_name)
	  AND pr.policy_type = 'TargetTrackingScaling'
      AND pr.scalable_dimension = 'dynamodb:table:ReadCapacityUnits'
  LEFT JOIN aws_applicationautoscaling_policies AS pw ON
      pw.service_namespace = 'dynamodb'
      AND pw.resource_id = concat('table/', t.table_name)
	  AND pw.policy_type = 'TargetTrackingScaling'
      AND pw.scalable_dimension = 'dynamodb:table:WriteCapacityUnits'
group by t.account_id, t.arn, status
```


