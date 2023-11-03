# Table: aws_elbv1_load_balancer_policies

This table shows data for Amazon Elastic Load Balancer (ELB) v1 Load Balancer Policies.

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elbv1_load_balancers](aws_elbv1_load_balancers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn|`utf8`|
|load_balancer_name|`utf8`|
|policy_attribute_descriptions|`json`|
|policy_name|`utf8`|
|policy_type_name|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


