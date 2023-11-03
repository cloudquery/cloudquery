# Table: aws_elbv1_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v1 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv1_load_balancers:
  - [aws_elbv1_load_balancer_policies](aws_elbv1_load_balancer_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|availability_zones|`list<item: utf8, nullable>`|
|backend_server_descriptions|`json`|
|canonical_hosted_zone_name|`utf8`|
|canonical_hosted_zone_name_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|health_check|`json`|
|instances|`json`|
|listener_descriptions|`json`|
|load_balancer_name|`utf8`|
|policies|`json`|
|scheme|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|source_security_group|`json`|
|subnets|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|
|tags|`json`|
|attributes|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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

### Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager

```sql
SELECT
  'Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager'
    AS title,
  lb.account_id,
  lb.arn AS resource_id,
  CASE
  WHEN li->'Listener'->>'Protocol' = 'HTTPS'
  AND aws_acm_certificates.arn IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv1_load_balancers AS lb,
  jsonb_array_elements(lb.listener_descriptions) AS li
  LEFT JOIN aws_acm_certificates ON
      aws_acm_certificates.arn = li->'Listener'->>'SSLCertificateId';
```

### Classic Load Balancers should have connection draining enabled

```sql
SELECT
  'Classic Load Balancers should have connection draining enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (attributes->'ConnectionDraining'->>'Enabled')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv1_load_balancers;
```

### Classic Load Balancer listeners should be configured with HTTPS or TLS termination

```sql
SELECT
  'Classic Load Balancer listeners should be configured with HTTPS or TLS termination'
    AS title,
  lb.account_id,
  lb.arn AS resource_id,
  CASE
  WHEN li->'Listener'->>'Protocol' NOT IN ('HTTPS', 'SSL') THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elbv1_load_balancers AS lb,
  jsonb_array_elements(lb.listener_descriptions) AS li;
```

### Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration

```sql
WITH
  flat_listeners
    AS (
      SELECT
        arn,
        account_id,
        li->'Listener'->>'Protocol' AS protocol,
        li->'PolicyNames' AS policies_arr
      FROM
        aws_elbv1_load_balancers AS lb,
        jsonb_array_elements(lb.listener_descriptions) AS li
    ),
  violations
    AS (
      SELECT
        fl.arn,
        fl.account_id,
        CASE
        WHEN fl.protocol IN ('HTTPS', 'SSL')
        AND NOT
            EXISTS(
              SELECT
                1
              FROM
                aws_elbv1_load_balancer_policies AS pol
              WHERE
                fl.policies_arr @> ('["' || pol.policy_name || '"]')::JSONB
                AND pol.policy_attribute_descriptions->>'Reference-Security-Policy'
                  = 'ELBSecurityPolicy-TLS-1-2-2017-01'
            )
        THEN 1
        ELSE 0
        END
          AS flag
      FROM
        flat_listeners AS fl
    )
SELECT
  DISTINCT
  'Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration'
    AS title,
  v.account_id,
  v.arn AS resource_id,
  CASE
  WHEN max(flag) OVER (PARTITION BY arn) = 1 THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  violations AS v;
```

### Find all Classic ELBs that are Internet Facing

```sql
SELECT
  'Find all Classic ELBs that are Internet Facing' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN scheme = 'internet-facing' THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_elbv1_load_balancers;
```


