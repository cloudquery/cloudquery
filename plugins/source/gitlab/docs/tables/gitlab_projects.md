# Table: gitlab_projects

The composite primary key for this table is (**base_url**, **id**).

## Relations

The following tables depend on gitlab_projects:
  - [gitlab_project_branches](gitlab_project_branches.md)
  - [gitlab_projects_releases](gitlab_projects_releases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|base_url (PK)|String|
|id (PK)|Int|
|marked_for_deletion_at|Timestamp|
|description|String|
|default_branch|String|
|public|Bool|
|visibility|String|
|ssh_url_to_repo|String|
|http_url_to_repo|String|
|web_url|String|
|readme_url|String|
|tag_list|StringArray|
|topics|StringArray|
|owner|JSON|
|name|String|
|name_with_namespace|String|
|path|String|
|path_with_namespace|String|
|issues_enabled|Bool|
|open_issues_count|Int|
|merge_requests_enabled|Bool|
|approvals_before_merge|Int|
|jobs_enabled|Bool|
|wiki_enabled|Bool|
|snippets_enabled|Bool|
|resolve_outdated_diff_discussions|Bool|
|container_expiration_policy|JSON|
|container_registry_enabled|Bool|
|container_registry_access_level|String|
|container_registry_image_prefix|String|
|created_at|Timestamp|
|last_activity_at|Timestamp|
|creator_id|Int|
|namespace|JSON|
|import_status|String|
|import_error|String|
|permissions|JSON|
|empty_repo|Bool|
|archived|Bool|
|avatar_url|String|
|license_url|String|
|license|JSON|
|shared_runners_enabled|Bool|
|forks_count|Int|
|star_count|Int|
|runners_token|String|
|public_jobs|Bool|
|allow_merge_on_skipped_pipeline|Bool|
|only_allow_merge_if_pipeline_succeeds|Bool|
|only_allow_merge_if_all_discussions_are_resolved|Bool|
|remove_source_branch_after_merge|Bool|
|printing_merge_request_link_enabled|Bool|
|lfs_enabled|Bool|
|repository_storage|String|
|request_access_enabled|Bool|
|merge_method|String|
|forked_from_project|JSON|
|mirror|Bool|
|mirror_user_id|Int|
|mirror_trigger_builds|Bool|
|only_mirror_protected_branches|Bool|
|mirror_overwrites_diverged_branches|Bool|
|packages_enabled|Bool|
|service_desk_enabled|Bool|
|service_desk_address|String|
|issues_access_level|String|
|repository_access_level|String|
|merge_requests_access_level|String|
|forking_access_level|String|
|wiki_access_level|String|
|builds_access_level|String|
|snippets_access_level|String|
|pages_access_level|String|
|operations_access_level|String|
|analytics_access_level|String|
|autoclose_referenced_issues|Bool|
|suggestion_commit_message|String|
|auto_cancel_pending_pipelines|String|
|ci_forward_deployment_enabled|Bool|
|squash_option|String|
|enforce_auth_checks_on_uploads|Bool|
|shared_with_groups|JSON|
|statistics|JSON|
|_links|JSON|
|ci_config_path|String|
|ci_default_git_depth|Int|
|ci_separated_caches|Bool|
|custom_attributes|JSON|
|compliance_frameworks|StringArray|
|build_coverage_regex|String|
|build_timeout|Int|
|issues_template|String|
|merge_requests_template|String|
|keep_latest_artifact|Bool|
|merge_pipelines_enabled|Bool|
|merge_trains_enabled|Bool|
|restrict_user_defined_variables|Bool|
|merge_commit_template|String|
|squash_commit_template|String|
|auto_devops_deploy_strategy|String|
|auto_devops_enabled|Bool|
|build_git_strategy|String|
|emails_disabled|Bool|
|external_authorization_classification_label|String|
|requirements_access_level|String|
|security_and_compliance_access_level|String|
|public_builds|Bool|