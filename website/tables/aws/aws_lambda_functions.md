# Table: aws_lambda_functions

This table shows data for AWS Lambda Functions.

https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lambda_functions:
  - [aws_lambda_function_aliases](aws_lambda_function_aliases)
  - [aws_lambda_function_concurrency_configs](aws_lambda_function_concurrency_configs)
  - [aws_lambda_function_event_invoke_configs](aws_lambda_function_event_invoke_configs)
  - [aws_lambda_function_event_source_mappings](aws_lambda_function_event_source_mappings)
  - [aws_lambda_function_versions](aws_lambda_function_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|policy_revision_id|`utf8`|
|policy_document|`json`|
|code_signing_config|`json`|
|code_repository_type|`utf8`|
|update_runtime_on|`utf8`|
|runtime_version_arn|`utf8`|
|code|`json`|
|concurrency|`json`|
|configuration|`json`|
|tags|`json`|
|result_metadata|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Find all lambda functions that have unrestricted access to the internet

```sql
SELECT
  DISTINCT
  'Find all lambda functions that have unrestricted access to the internet'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_lambda_functions,
  jsonb_array_elements_text(configuration->'VpcConfig'->'SecurityGroupIds')
    AS sgs,
  jsonb_array_elements_text(configuration->'VpcConfig'->' SubnetIds') AS sns
WHERE
  sns
  IN (
      SELECT
        a->>'SubnetId'
      FROM
        public.aws_ec2_route_tables,
        jsonb_array_elements(associations) AS a,
        jsonb_array_elements(routes) AS r
      WHERE
        r->>'DestinationCidrBlock' = '0.0.0.0/0'
        OR r->>'DestinationIpv6CidrBlock' = '::/0'
    )
  AND sgs
    IN (
        SELECT
          id
        FROM
          view_aws_security_group_egress_rules
        WHERE
          ip = '0.0.0.0/0' OR ip6 = '::/0'
      )
UNION
  SELECT
    DISTINCT
    'Find all lambda functions that have unrestricted access to the internet'
      AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status
  FROM
    aws_lambda_functions
  WHERE
    (configuration->'VpcConfig'->>'VpcId') IS NULL;
```

### Lambda functions should be in a VPC

```sql
SELECT
  'Lambda functions should be in a VPC' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (configuration->'VpcConfig'->>'VpcId') IS NULL
  OR configuration->'VpcConfig'->>'VpcId' = ''
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_lambda_functions;
```

### Lambda functions should prohibit public access

```sql
SELECT
  'Lambda functions should prohibit public access' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_lambda_functions,
  jsonb_array_elements(
    CASE jsonb_typeof(policy_document->'Statement')
    WHEN 'string' THEN jsonb_build_array(policy_document->>'Statement')
    WHEN 'array' THEN policy_document->'Statement'
    END
  )
    AS statement
WHERE
  statement->>'Effect' = 'Allow'
  AND (
      statement->>'Principal' = '*'
      OR statement->'Principal'->>'AWS' = '*'
      OR (
          CASE jsonb_typeof(statement->'Principal'->'AWS')
          WHEN 'string' THEN jsonb_build_array(statement->'Principal'->>'AWS')
          WHEN 'array' THEN (statement->'Principal'->>'AWS')::JSONB
          END
        )::JSONB
        ? '*'
    );
```

### Lambda functions should use supported runtimes

```sql
SELECT
  'Lambda functions should use supported runtimes' AS title,
  f.account_id,
  f.arn AS resource_id,
  CASE WHEN r.name IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_lambda_functions AS f
  LEFT JOIN aws_lambda_runtimes AS r ON r.name = f.configuration->>'Runtime'
WHERE
  f.configuration->>'PackageType' != 'Image';
```


