# Table: aws_codebuild_projects

This table shows data for Codebuild Projects.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|artifacts|json|
|badge|json|
|build_batch_config|json|
|cache|json|
|concurrent_build_limit|int64|
|created|timestamp[us, tz=UTC]|
|description|utf8|
|encryption_key|utf8|
|environment|json|
|file_system_locations|json|
|last_modified|timestamp[us, tz=UTC]|
|logs_config|json|
|name|utf8|
|project_visibility|utf8|
|public_project_alias|utf8|
|queued_timeout_in_minutes|int64|
|resource_access_role|utf8|
|secondary_artifacts|json|
|secondary_source_versions|json|
|secondary_sources|json|
|service_role|utf8|
|source|json|
|source_version|utf8|
|timeout_in_minutes|int64|
|vpc_config|json|
|webhook|json|