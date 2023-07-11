# Table: aws_elbv2_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_load_balancers:
  - [aws_elbv2_listeners](aws_elbv2_listeners)
  - [aws_elbv2_load_balancer_attributes](aws_elbv2_load_balancer_attributes)
  - [aws_elbv2_load_balancer_web_acls](aws_elbv2_load_balancer_web_acls)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|availability_zones|`json`|
|canonical_hosted_zone_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|customer_owned_ipv4_pool|`utf8`|
|dns_name|`utf8`|
|ip_address_type|`utf8`|
|load_balancer_arn|`utf8`|
|load_balancer_name|`utf8`|
|scheme|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|state|`json`|
|type|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Application Load Balancer deletion protection should be enabled

```sql
SELECT
  'Application Load Balancer deletion protection should be enabled' AS title,
  lb.account_id,
  lb.arn AS resource_id,
  CASE
  WHEN lb.type = 'application' AND (a.value)::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv2_load_balancers AS lb
  INNER JOIN aws_elbv2_load_balancer_attributes AS a ON
      a.load_balancer_arn = lb.arn AND a.key = 'deletion_protection.enabled';
```

### Application load balancers should be configured to drop HTTP headers

```sql
SELECT
  'Application load balancers should be configured to drop HTTP headers'
    AS title,
  lb.account_id,
  lb.arn AS resource_id,
  CASE
  WHEN lb.type = 'application' AND (a.value)::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv2_load_balancers AS lb
  INNER JOIN aws_elbv2_load_balancer_attributes AS a ON
      a.load_balancer_arn = lb.arn
      AND a.key = 'routing.http.drop_invalid_header_fields.enabled';
```

### Application and Classic Load Balancers logging should be enabled

```sql
(
  SELECT
    'Application and Classic Load Balancers logging should be enabled' AS title,
    lb.account_id,
    lb.arn AS resource_id,
    CASE
    WHEN lb.type = 'application' AND (a.value)::BOOL IS NOT true THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    aws_elbv2_load_balancers AS lb
    INNER JOIN aws_elbv2_load_balancer_attributes AS a ON
        a.load_balancer_arn = lb.arn AND a.key = 'access_logs.s3.enabled'
)
UNION
  (
    SELECT
      'Application and Classic Load Balancers logging should be enabled'
        AS title,
      account_id,
      arn AS resource_id,
      CASE
      WHEN (attributes->'AccessLog'->>'Enabled')::BOOL IS NOT true THEN 'fail'
      ELSE 'pass'
      END
        AS status
    FROM
      aws_elbv1_load_balancers
  );
```

### Find all ELB V2s that are Internet Facing

```sql
SELECT
  'Find all ELB V2s that are Internet Facing' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN scheme = 'internet-facing' THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_elbv2_load_balancers;
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


