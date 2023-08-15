# Table: aws_codebuild_builds

This table shows data for AWS CodeBuild Builds.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Build.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_codebuild_projects](aws_codebuild_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|artifacts|`json`|
|build_batch_arn|`utf8`|
|build_complete|`bool`|
|build_number|`int64`|
|build_status|`utf8`|
|cache|`json`|
|current_phase|`utf8`|
|debug_session|`json`|
|encryption_key|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|environment|`json`|
|exported_environment_variables|`json`|
|file_system_locations|`json`|
|id|`utf8`|
|initiator|`utf8`|
|logs|`json`|
|network_interface|`json`|
|phases|`json`|
|project_name|`utf8`|
|queued_timeout_in_minutes|`int64`|
|report_arns|`list<item: utf8, nullable>`|
|resolved_source_version|`utf8`|
|secondary_artifacts|`json`|
|secondary_source_versions|`json`|
|secondary_sources|`json`|
|service_role|`utf8`|
|source|`json`|
|source_version|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|timeout_in_minutes|`int64`|
|vpc_config|`json`|