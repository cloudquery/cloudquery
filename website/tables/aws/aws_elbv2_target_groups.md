# Table: aws_elbv2_target_groups

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Target Groups.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_target_groups:
  - [aws_elbv2_target_group_target_health_descriptions](aws_elbv2_target_group_target_health_descriptions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|health_check_enabled|`bool`|
|health_check_interval_seconds|`int64`|
|health_check_path|`utf8`|
|health_check_port|`utf8`|
|health_check_protocol|`utf8`|
|health_check_timeout_seconds|`int64`|
|healthy_threshold_count|`int64`|
|ip_address_type|`utf8`|
|load_balancer_arns|`list<item: utf8, nullable>`|
|matcher|`json`|
|port|`int64`|
|protocol|`utf8`|
|protocol_version|`utf8`|
|target_group_arn|`utf8`|
|target_group_name|`utf8`|
|target_type|`utf8`|
|unhealthy_threshold_count|`int64`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused ELB load balancer

```sql
WITH
  listener AS (SELECT DISTINCT load_balancer_arn FROM aws_elbv2_listeners),
  target_group
    AS (
      SELECT
        DISTINCT unnest(load_balancer_arns) AS load_balancer_arn
      FROM
        aws_elbv2_target_groups
    )
SELECT
  'Unused ELB load balancer' AS title,
  lb.account_id,
  lb.arn AS resource_id,
  'fail' AS status
FROM
  aws_elbv2_load_balancers AS lb
  LEFT JOIN listener ON listener.load_balancer_arn = lb.arn
  LEFT JOIN target_group ON target_group.load_balancer_arn = lb.arn
WHERE
  listener.load_balancer_arn IS NULL OR target_group.load_balancer_arn IS NULL;
```

### Unused ELB target group

```sql
SELECT
  'Unused ELB target group' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_elbv2_target_groups
WHERE
  array_length(load_balancer_arns, 1) = 0;
```


