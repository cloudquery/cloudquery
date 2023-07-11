# Table: aws_codebuild_projects

This table shows data for AWS CodeBuild Projects.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_codebuild_projects:
  - [aws_codebuild_builds](aws_codebuild_builds)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|artifacts|`json`|
|badge|`json`|
|build_batch_config|`json`|
|cache|`json`|
|concurrent_build_limit|`int64`|
|created|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|encryption_key|`utf8`|
|environment|`json`|
|file_system_locations|`json`|
|last_modified|`timestamp[us, tz=UTC]`|
|logs_config|`json`|
|name|`utf8`|
|project_visibility|`utf8`|
|public_project_alias|`utf8`|
|queued_timeout_in_minutes|`int64`|
|resource_access_role|`utf8`|
|secondary_artifacts|`json`|
|secondary_source_versions|`json`|
|secondary_sources|`json`|
|service_role|`utf8`|
|source|`json`|
|source_version|`utf8`|
|timeout_in_minutes|`int64`|
|vpc_config|`json`|
|webhook|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### CodeBuild project environment variables should not contain clear text credentials

```sql
SELECT
  DISTINCT
  'CodeBuild project environment variables should not contain clear text credentials'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN e->>'Type' = 'PLAINTEXT'
  AND (
      upper(e->>'Name') LIKE '%ACCESS_KEY%'
      OR upper(e->>'Name') LIKE '%SECRET%'
      OR upper(e->>'Name') LIKE '%PASSWORD%'
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_codebuild_projects,
  jsonb_array_elements(environment->'EnvironmentVariables') AS e;
```

### CodeBuild GitHub or Bitbucket source repository URLs should use OAuth

```sql
SELECT
  'CodeBuild GitHub or Bitbucket source repository URLs should use OAuth'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN source->>'Type' IN ('GITHUB', 'BITBUCKET')
  AND source->'Auth'->>'Type' != 'OAUTH'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_codebuild_projects;
```


