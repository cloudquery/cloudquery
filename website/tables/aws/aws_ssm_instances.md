# Table: aws_ssm_instances

This table shows data for AWS Systems Manager (SSM) Instances.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ssm_instances:
  - [aws_ssm_instance_compliance_items](aws_ssm_instance_compliance_items)
  - [aws_ssm_instance_patches](aws_ssm_instance_patches)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|activation_id|`utf8`|
|agent_version|`utf8`|
|association_overview|`json`|
|association_status|`utf8`|
|computer_name|`utf8`|
|ip_address|`utf8`|
|iam_role|`utf8`|
|instance_id|`utf8`|
|is_latest_version|`bool`|
|last_association_execution_date|`timestamp[us, tz=UTC]`|
|last_ping_date_time|`timestamp[us, tz=UTC]`|
|last_successful_association_execution_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|ping_status|`utf8`|
|platform_name|`utf8`|
|platform_type|`utf8`|
|platform_version|`utf8`|
|registration_date|`timestamp[us, tz=UTC]`|
|resource_type|`utf8`|
|source_id|`utf8`|
|source_type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon EC2 instances should be managed by AWS Systems Manager

```sql
SELECT
  'Amazon EC2 instances should be managed by AWS Systems Manager' AS title,
  aws_ec2_instances.account_id,
  aws_ec2_instances.arn AS resource_id,
  CASE
  WHEN aws_ssm_instances.instance_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ec2_instances
  LEFT JOIN aws_ssm_instances ON
      aws_ec2_instances.instance_id = aws_ssm_instances.instance_id;
```

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


