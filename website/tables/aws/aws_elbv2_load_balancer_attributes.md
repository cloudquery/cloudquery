# Table: aws_elbv2_load_balancer_attributes

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancer Attributes.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerAttribute.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn|`utf8`|
|key|`utf8`|
|value|`utf8`|

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


