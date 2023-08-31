# Table: gitlab_projects

This table shows data for Gitlab Projects.

The composite primary key for this table is (**base_url**, **id**).

## Relations

The following tables depend on gitlab_projects:
  - [gitlab_project_branches](gitlab_project_branches)
  - [gitlab_project_members](gitlab_project_members)
  - [gitlab_projects_releases](gitlab_projects_releases)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|id (PK)|`int64`|
|description|`utf8`|
|default_branch|`utf8`|
|public|`bool`|
|visibility|`utf8`|
|ssh_url_to_repo|`utf8`|
|http_url_to_repo|`utf8`|
|web_url|`utf8`|
|readme_url|`utf8`|
|tag_list|`list<item: utf8, nullable>`|
|topics|`list<item: utf8, nullable>`|
|owner|`json`|
|name|`utf8`|
|name_with_namespace|`utf8`|
|path|`utf8`|
|path_with_namespace|`utf8`|
|issues_enabled|`bool`|
|open_issues_count|`int64`|
|merge_requests_enabled|`bool`|
|approvals_before_merge|`int64`|
|jobs_enabled|`bool`|
|wiki_enabled|`bool`|
|snippets_enabled|`bool`|
|resolve_outdated_diff_discussions|`bool`|
|container_expiration_policy|`json`|
|container_registry_enabled|`bool`|
|container_registry_access_level|`utf8`|
|container_registry_image_prefix|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|last_activity_at|`timestamp[us, tz=UTC]`|
|creator_id|`int64`|
|namespace|`json`|
|permissions|`json`|
|marked_for_deletion_at|`timestamp[us, tz=UTC]`|
|empty_repo|`bool`|
|archived|`bool`|
|avatar_url|`utf8`|
|license_url|`utf8`|
|license|`json`|
|shared_runners_enabled|`bool`|
|group_runners_enabled|`bool`|
|runner_token_expiration_interval|`int64`|
|forks_count|`int64`|
|star_count|`int64`|
|runners_token|`utf8`|
|allow_merge_on_skipped_pipeline|`bool`|
|only_allow_merge_if_pipeline_succeeds|`bool`|
|only_allow_merge_if_all_discussions_are_resolved|`bool`|
|remove_source_branch_after_merge|`bool`|
|printing_merge_request_link_enabled|`bool`|
|lfs_enabled|`bool`|
|repository_storage|`utf8`|
|request_access_enabled|`bool`|
|merge_method|`utf8`|
|can_create_merge_request_in|`bool`|
|forked_from_project|`json`|
|mirror|`bool`|
|mirror_user_id|`int64`|
|mirror_trigger_builds|`bool`|
|only_mirror_protected_branches|`bool`|
|mirror_overwrites_diverged_branches|`bool`|
|packages_enabled|`bool`|
|service_desk_enabled|`bool`|
|service_desk_address|`utf8`|
|issues_access_level|`utf8`|
|releases_access_level|`utf8`|
|repository_access_level|`utf8`|
|merge_requests_access_level|`utf8`|
|forking_access_level|`utf8`|
|wiki_access_level|`utf8`|
|builds_access_level|`utf8`|
|snippets_access_level|`utf8`|
|pages_access_level|`utf8`|
|operations_access_level|`utf8`|
|analytics_access_level|`utf8`|
|environments_access_level|`utf8`|
|feature_flags_access_level|`utf8`|
|infrastructure_access_level|`utf8`|
|monitor_access_level|`utf8`|
|autoclose_referenced_issues|`bool`|
|suggestion_commit_message|`utf8`|
|squash_option|`utf8`|
|enforce_auth_checks_on_uploads|`bool`|
|shared_with_groups|`json`|
|statistics|`json`|
|_links|`json`|
|import_url|`utf8`|
|import_type|`utf8`|
|import_status|`utf8`|
|import_error|`utf8`|
|ci_default_git_depth|`int64`|
|ci_forward_deployment_enabled|`bool`|
|ci_separated_caches|`bool`|
|ci_job_token_scope_enabled|`bool`|
|ci_opt_in_jwt|`bool`|
|ci_allow_fork_pipelines_to_run_in_parent_project|`bool`|
|public_jobs|`bool`|
|build_timeout|`int64`|
|auto_cancel_pending_pipelines|`utf8`|
|ci_config_path|`utf8`|
|custom_attributes|`json`|
|compliance_frameworks|`list<item: utf8, nullable>`|
|build_coverage_regex|`utf8`|
|issues_template|`utf8`|
|merge_requests_template|`utf8`|
|issue_branch_template|`utf8`|
|keep_latest_artifact|`bool`|
|merge_pipelines_enabled|`bool`|
|merge_trains_enabled|`bool`|
|restrict_user_defined_variables|`bool`|
|merge_commit_template|`utf8`|
|squash_commit_template|`utf8`|
|auto_devops_deploy_strategy|`utf8`|
|auto_devops_enabled|`bool`|
|build_git_strategy|`utf8`|
|emails_disabled|`bool`|
|external_authorization_classification_label|`utf8`|
|requirements_enabled|`bool`|
|requirements_access_level|`utf8`|
|security_and_compliance_enabled|`bool`|
|security_and_compliance_access_level|`utf8`|
|mr_default_target_self|`bool`|
|public_builds|`bool`|