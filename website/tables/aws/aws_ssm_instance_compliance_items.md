# Table: aws_ssm_instance_compliance_items

This table shows data for AWS Systems Manager (SSM) Instance Compliance Items.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html

The composite primary key for this table is (**id**, **instance_arn**).

## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id (PK)|`utf8`|
|instance_arn (PK)|`utf8`|
|compliance_type|`utf8`|
|details|`json`|
|execution_summary|`json`|
|resource_id|`utf8`|
|resource_type|`utf8`|
|severity|`utf8`|
|status|`utf8`|
|title|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon EC2 instances managed by Systems Manager should have an association compliance status of COMPLIANT

```sql
SELECT
  'Amazon EC2 instances managed by Systems Manager should have an association compliance status of COMPLIANT'
    AS title,
  aws_ssm_instances.account_id,
  aws_ssm_instances.arn,
  CASE
  WHEN aws_ssm_instance_compliance_items.compliance_type = 'Association'
  AND aws_ssm_instance_compliance_items.status IS DISTINCT FROM 'COMPLIANT'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ssm_instances
  INNER JOIN aws_ssm_instance_compliance_items ON
      aws_ssm_instances.arn = aws_ssm_instance_compliance_items.instance_arn;
```

### Amazon EC2 instances managed by Systems Manager should have a patch compliance status of COMPLIANT after a patch installation

```sql
SELECT
  'Amazon EC2 instances managed by Systems Manager should have a patch compliance status of COMPLIANT after a patch installation'
    AS title,
  aws_ssm_instances.account_id,
  aws_ssm_instances.arn,
  CASE
  WHEN aws_ssm_instance_compliance_items.compliance_type = 'Patch'
  AND aws_ssm_instance_compliance_items.status IS DISTINCT FROM 'COMPLIANT'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ssm_instances
  INNER JOIN aws_ssm_instance_compliance_items ON
      aws_ssm_instances.arn = aws_ssm_instance_compliance_items.instance_arn;
```


