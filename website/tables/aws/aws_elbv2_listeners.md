# Table: aws_elbv2_listeners

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listeners.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers).

The following tables depend on aws_elbv2_listeners:
  - [aws_elbv2_listener_certificates](aws_elbv2_listener_certificates)
  - [aws_elbv2_listener_rules](aws_elbv2_listener_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|alpn_policy|`list<item: utf8, nullable>`|
|certificates|`json`|
|default_actions|`json`|
|listener_arn|`utf8`|
|load_balancer_arn|`utf8`|
|port|`int64`|
|protocol|`utf8`|
|ssl_policy|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Application Load Balancer should be configured to redirect all HTTP requests to HTTPS

```sql
SELECT
  'Application Load Balancer should be configured to redirect all HTTP requests to HTTPS'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN protocol = 'HTTP'
  AND (
      da->>'Type' != 'REDIRECT'
      OR da->'RedirectConfig'->>'Protocol' != 'HTTPS'
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv2_listeners, jsonb_array_elements(default_actions) AS da;
```

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


