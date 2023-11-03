# Table: aws_waf_web_acls

This table shows data for WAF Web ACLs.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|default_action|`json`|
|rules|`json`|
|web_acl_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|
|web_acl_arn|`utf8`|
|logging_configuration|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### AWS WAF Classic global web ACL logging should be enabled

```sql
-- WAF Classic
SELECT
  'AWS WAF Classic global web ACL logging should be enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN logging_configuration IS NULL OR logging_configuration = '{}' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_waf_web_acls;
```

### AWS WAF Classic global web ACL logging should be enabled

```sql
(
  SELECT
    'AWS WAF Classic global web ACL logging should be enabled' AS title,
    account_id,
    arn AS resource_id,
    CASE
    WHEN logging_configuration IS NULL OR logging_configuration = '{}'
    THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    aws_waf_web_acls
)
UNION
  (
    SELECT
      'AWS WAF Classic global web ACL logging should be enabled' AS title,
      account_id,
      arn AS resource_id,
      CASE
      WHEN logging_configuration IS NULL OR logging_configuration = '{}'
      THEN 'fail'
      ELSE 'pass'
      END
        AS status
    FROM
      aws_wafv2_web_acls
  );
```


