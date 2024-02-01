# Table: aws_codebuild_projects

This table shows data for AWS CodeBuild Projects.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_codebuild_projects:
  - [aws_codebuild_builds](aws_codebuild_builds.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
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