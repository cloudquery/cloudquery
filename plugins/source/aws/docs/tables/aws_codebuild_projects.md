# Table: aws_codebuild_projects

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|artifacts|JSON|
|badge|JSON|
|build_batch_config|JSON|
|cache|JSON|
|concurrent_build_limit|Int|
|created|Timestamp|
|description|String|
|encryption_key|String|
|environment|JSON|
|file_system_locations|JSON|
|last_modified|Timestamp|
|logs_config|JSON|
|name|String|
|project_visibility|String|
|public_project_alias|String|
|queued_timeout_in_minutes|Int|
|resource_access_role|String|
|secondary_artifacts|JSON|
|secondary_source_versions|JSON|
|secondary_sources|JSON|
|service_role|String|
|source|JSON|
|source_version|String|
|tags|JSON|
|timeout_in_minutes|Int|
|vpc_config|JSON|
|webhook|JSON|