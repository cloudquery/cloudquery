# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

## github_action_billing
This table was removed.


## github_billing_action
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|org|text|added|
|total_minutes_used|bigint|added|
|total_paid_minutes_used|real|added|
|included_minutes|bigint|added|
|minutes_used_breakdown|jsonb|added|

## github_billing_package
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|org|text|added|
|total_gigabytes_bandwidth_used|bigint|added|
|total_paid_gigabytes_bandwidth_used|bigint|added|
|included_gigabytes_bandwidth|bigint|added|

## github_billing_storage
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|org|text|added|
|days_left_in_billing_cycle|bigint|added|
|estimated_paid_storage_for_month|real|added|
|estimated_storage_for_month|bigint|added|

## github_external_group_members
Moved to JSON column on [github_external_groups](#github_external_groups)


## github_external_group_teams
Moved to JSON column on [github_external_groups](#github_external_groups)


## github_external_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|members|jsonb|added|
|teams|jsonb|added|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_hook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|delivered_at|timestamp without time zone|added|
|delivered_at_time|timestamp without time zone|removed|
|duration|real|updated|Type changed from float to real
|hook_cq_id|uuid|removed|
|hook_id|bigint|added|
|org|text|added|
|request|text|added|
|request_headers|jsonb|removed|
|request_raw_payload|bytea|removed|
|response|text|added|
|response_headers|jsonb|removed|
|response_raw_payload|bytea|removed|

## github_hooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## github_installations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account|jsonb|added|
|account_avatar_url|text|removed|
|account_bio|text|removed|
|account_blog|text|removed|
|account_collaborators|bigint|removed|
|account_company|text|removed|
|account_created_at_time|timestamp without time zone|removed|
|account_disk_usage|bigint|removed|
|account_email|text|removed|
|account_events_url|text|removed|
|account_followers|bigint|removed|
|account_followers_url|text|removed|
|account_following|bigint|removed|
|account_following_url|text|removed|
|account_gists_url|text|removed|
|account_gravatar_id|text|removed|
|account_hireable|boolean|removed|
|account_html_url|text|removed|
|account_id|bigint|removed|
|account_ldap_dn|text|removed|
|account_location|text|removed|
|account_login|text|removed|
|account_name|text|removed|
|account_node_id|text|removed|
|account_organizations_url|text|removed|
|account_owned_private_repos|bigint|removed|
|account_permissions|jsonb|removed|
|account_plan_collaborators|bigint|removed|
|account_plan_filled_seats|bigint|removed|
|account_plan_name|text|removed|
|account_plan_private_repos|bigint|removed|
|account_plan_seats|bigint|removed|
|account_plan_space|bigint|removed|
|account_private_gists|bigint|removed|
|account_public_gists|bigint|removed|
|account_public_repos|bigint|removed|
|account_received_events_url|text|removed|
|account_repos_url|text|removed|
|account_role_name|text|removed|
|account_site_admin|boolean|removed|
|account_starred_url|text|removed|
|account_subscriptions_url|text|removed|
|account_suspended_at_time|timestamp without time zone|removed|
|account_text_matches|jsonb|removed|
|account_total_private_repos|bigint|removed|
|account_twitter_username|text|removed|
|account_two_factor_authentication|boolean|removed|
|account_type|text|removed|
|account_updated_at_time|timestamp without time zone|removed|
|account_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|permissions|jsonb|added|
|permissions_actions|text|removed|
|permissions_administration|text|removed|
|permissions_blocking|text|removed|
|permissions_checks|text|removed|
|permissions_content_references|text|removed|
|permissions_contents|text|removed|
|permissions_deployments|text|removed|
|permissions_emails|text|removed|
|permissions_environments|text|removed|
|permissions_followers|text|removed|
|permissions_issues|text|removed|
|permissions_members|text|removed|
|permissions_metadata|text|removed|
|permissions_organization_administration|text|removed|
|permissions_organization_hooks|text|removed|
|permissions_organization_plan|text|removed|
|permissions_organization_pre_receive_hooks|text|removed|
|permissions_organization_projects|text|removed|
|permissions_organization_secrets|text|removed|
|permissions_organization_self_hosted_runners|text|removed|
|permissions_organization_user_blocking|text|removed|
|permissions_packages|text|removed|
|permissions_pages|text|removed|
|permissions_pull_requests|text|removed|
|permissions_repository_hooks|text|removed|
|permissions_repository_pre_receive_hooks|text|removed|
|permissions_repository_projects|text|removed|
|permissions_secret_scanning_alerts|text|removed|
|permissions_secrets|text|removed|
|permissions_security_events|text|removed|
|permissions_single_file|text|removed|
|permissions_statuses|text|removed|
|permissions_team_discussions|text|removed|
|permissions_vulnerability_alerts|text|removed|
|permissions_workflows|text|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|suspended_by|jsonb|added|
|suspended_by_avatar_url|text|removed|
|suspended_by_bio|text|removed|
|suspended_by_blog|text|removed|
|suspended_by_collaborators|bigint|removed|
|suspended_by_company|text|removed|
|suspended_by_created_at_time|timestamp without time zone|removed|
|suspended_by_disk_usage|bigint|removed|
|suspended_by_email|text|removed|
|suspended_by_events_url|text|removed|
|suspended_by_followers|bigint|removed|
|suspended_by_followers_url|text|removed|
|suspended_by_following|bigint|removed|
|suspended_by_following_url|text|removed|
|suspended_by_gists_url|text|removed|
|suspended_by_gravatar_id|text|removed|
|suspended_by_hireable|boolean|removed|
|suspended_by_html_url|text|removed|
|suspended_by_id|bigint|removed|
|suspended_by_ldap_dn|text|removed|
|suspended_by_location|text|removed|
|suspended_by_login|text|removed|
|suspended_by_name|text|removed|
|suspended_by_node_id|text|removed|
|suspended_by_organizations_url|text|removed|
|suspended_by_owned_private_repos|bigint|removed|
|suspended_by_permissions|jsonb|removed|
|suspended_by_plan_collaborators|bigint|removed|
|suspended_by_plan_filled_seats|bigint|removed|
|suspended_by_plan_name|text|removed|
|suspended_by_plan_private_repos|bigint|removed|
|suspended_by_plan_seats|bigint|removed|
|suspended_by_plan_space|bigint|removed|
|suspended_by_private_gists|bigint|removed|
|suspended_by_public_gists|bigint|removed|
|suspended_by_public_repos|bigint|removed|
|suspended_by_received_events_url|text|removed|
|suspended_by_repos_url|text|removed|
|suspended_by_role_name|text|removed|
|suspended_by_site_admin|boolean|removed|
|suspended_by_starred_url|text|removed|
|suspended_by_subscriptions_url|text|removed|
|suspended_by_suspended_at_time|timestamp without time zone|removed|
|suspended_by_text_matches|jsonb|removed|
|suspended_by_total_private_repos|bigint|removed|
|suspended_by_twitter_username|text|removed|
|suspended_by_two_factor_authentication|boolean|removed|
|suspended_by_type|text|removed|
|suspended_by_updated_at_time|timestamp without time zone|removed|
|suspended_by_url|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_issue_assignees
Moved to JSON column on [github_issues](#github_issues)


## github_issue_labels
Moved to JSON column on [github_issues](#github_issues)


## github_issues

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|assignee|jsonb|added|
|assignee_avatar_url|text|removed|
|assignee_bio|text|removed|
|assignee_blog|text|removed|
|assignee_collaborators|bigint|removed|
|assignee_company|text|removed|
|assignee_created_at_time|timestamp without time zone|removed|
|assignee_disk_usage|bigint|removed|
|assignee_email|text|removed|
|assignee_events_url|text|removed|
|assignee_followers|bigint|removed|
|assignee_followers_url|text|removed|
|assignee_following|bigint|removed|
|assignee_following_url|text|removed|
|assignee_gists_url|text|removed|
|assignee_gravatar_id|text|removed|
|assignee_hireable|boolean|removed|
|assignee_html_url|text|removed|
|assignee_id|bigint|removed|
|assignee_ldap_dn|text|removed|
|assignee_location|text|removed|
|assignee_login|text|removed|
|assignee_name|text|removed|
|assignee_node_id|text|removed|
|assignee_organizations_url|text|removed|
|assignee_owned_private_repos|bigint|removed|
|assignee_permissions|jsonb|removed|
|assignee_plan_collaborators|bigint|removed|
|assignee_plan_filled_seats|bigint|removed|
|assignee_plan_name|text|removed|
|assignee_plan_private_repos|bigint|removed|
|assignee_plan_seats|bigint|removed|
|assignee_plan_space|bigint|removed|
|assignee_private_gists|bigint|removed|
|assignee_public_gists|bigint|removed|
|assignee_public_repos|bigint|removed|
|assignee_received_events_url|text|removed|
|assignee_repos_url|text|removed|
|assignee_role_name|text|removed|
|assignee_site_admin|boolean|removed|
|assignee_starred_url|text|removed|
|assignee_subscriptions_url|text|removed|
|assignee_suspended_at_time|timestamp without time zone|removed|
|assignee_text_matches|jsonb|removed|
|assignee_total_private_repos|bigint|removed|
|assignee_twitter_username|text|removed|
|assignee_two_factor_authentication|boolean|removed|
|assignee_type|text|removed|
|assignee_updated_at_time|timestamp without time zone|removed|
|assignee_url|text|removed|
|assignees|jsonb|added|
|closed_by|jsonb|added|
|closed_by_avatar_url|text|removed|
|closed_by_bio|text|removed|
|closed_by_blog|text|removed|
|closed_by_collaborators|bigint|removed|
|closed_by_company|text|removed|
|closed_by_created_at_time|timestamp without time zone|removed|
|closed_by_disk_usage|bigint|removed|
|closed_by_email|text|removed|
|closed_by_events_url|text|removed|
|closed_by_followers|bigint|removed|
|closed_by_followers_url|text|removed|
|closed_by_following|bigint|removed|
|closed_by_following_url|text|removed|
|closed_by_gists_url|text|removed|
|closed_by_gravatar_id|text|removed|
|closed_by_hireable|boolean|removed|
|closed_by_html_url|text|removed|
|closed_by_id|bigint|removed|
|closed_by_ldap_dn|text|removed|
|closed_by_location|text|removed|
|closed_by_login|text|removed|
|closed_by_name|text|removed|
|closed_by_node_id|text|removed|
|closed_by_organizations_url|text|removed|
|closed_by_owned_private_repos|bigint|removed|
|closed_by_permissions|jsonb|removed|
|closed_by_plan_collaborators|bigint|removed|
|closed_by_plan_filled_seats|bigint|removed|
|closed_by_plan_name|text|removed|
|closed_by_plan_private_repos|bigint|removed|
|closed_by_plan_seats|bigint|removed|
|closed_by_plan_space|bigint|removed|
|closed_by_private_gists|bigint|removed|
|closed_by_public_gists|bigint|removed|
|closed_by_public_repos|bigint|removed|
|closed_by_received_events_url|text|removed|
|closed_by_repos_url|text|removed|
|closed_by_role_name|text|removed|
|closed_by_site_admin|boolean|removed|
|closed_by_starred_url|text|removed|
|closed_by_subscriptions_url|text|removed|
|closed_by_suspended_at_time|timestamp without time zone|removed|
|closed_by_text_matches|jsonb|removed|
|closed_by_total_private_repos|bigint|removed|
|closed_by_twitter_username|text|removed|
|closed_by_two_factor_authentication|boolean|removed|
|closed_by_type|text|removed|
|closed_by_updated_at_time|timestamp without time zone|removed|
|closed_by_url|text|removed|
|labels|jsonb|added|
|milestone|jsonb|added|
|milestone_closed_at|timestamp without time zone|removed|
|milestone_closed_issues|bigint|removed|
|milestone_created_at|timestamp without time zone|removed|
|milestone_creator_avatar_url|text|removed|
|milestone_creator_bio|text|removed|
|milestone_creator_blog|text|removed|
|milestone_creator_collaborators|bigint|removed|
|milestone_creator_company|text|removed|
|milestone_creator_created_at_time|timestamp without time zone|removed|
|milestone_creator_disk_usage|bigint|removed|
|milestone_creator_email|text|removed|
|milestone_creator_events_url|text|removed|
|milestone_creator_followers|bigint|removed|
|milestone_creator_followers_url|text|removed|
|milestone_creator_following|bigint|removed|
|milestone_creator_following_url|text|removed|
|milestone_creator_gists_url|text|removed|
|milestone_creator_gravatar_id|text|removed|
|milestone_creator_hireable|boolean|removed|
|milestone_creator_html_url|text|removed|
|milestone_creator_id|bigint|removed|
|milestone_creator_ldap_dn|text|removed|
|milestone_creator_location|text|removed|
|milestone_creator_login|text|removed|
|milestone_creator_name|text|removed|
|milestone_creator_node_id|text|removed|
|milestone_creator_organizations_url|text|removed|
|milestone_creator_owned_private_repos|bigint|removed|
|milestone_creator_permissions|jsonb|removed|
|milestone_creator_plan_collaborators|bigint|removed|
|milestone_creator_plan_filled_seats|bigint|removed|
|milestone_creator_plan_name|text|removed|
|milestone_creator_plan_private_repos|bigint|removed|
|milestone_creator_plan_seats|bigint|removed|
|milestone_creator_plan_space|bigint|removed|
|milestone_creator_private_gists|bigint|removed|
|milestone_creator_public_gists|bigint|removed|
|milestone_creator_public_repos|bigint|removed|
|milestone_creator_received_events_url|text|removed|
|milestone_creator_repos_url|text|removed|
|milestone_creator_role_name|text|removed|
|milestone_creator_site_admin|boolean|removed|
|milestone_creator_starred_url|text|removed|
|milestone_creator_subscriptions_url|text|removed|
|milestone_creator_suspended_at_time|timestamp without time zone|removed|
|milestone_creator_text_matches|jsonb|removed|
|milestone_creator_total_private_repos|bigint|removed|
|milestone_creator_twitter_username|text|removed|
|milestone_creator_two_factor_authentication|boolean|removed|
|milestone_creator_type|text|removed|
|milestone_creator_updated_at_time|timestamp without time zone|removed|
|milestone_creator_url|text|removed|
|milestone_description|text|removed|
|milestone_due_on|timestamp without time zone|removed|
|milestone_html_url|text|removed|
|milestone_id|bigint|removed|
|milestone_labels_url|text|removed|
|milestone_node_id|text|removed|
|milestone_number|bigint|removed|
|milestone_open_issues|bigint|removed|
|milestone_state|text|removed|
|milestone_title|text|removed|
|milestone_updated_at|timestamp without time zone|removed|
|milestone_url|text|removed|
|pull_request|jsonb|added|
|pull_request_links_diff_url|text|removed|
|pull_request_links_html_url|text|removed|
|pull_request_links_patch_url|text|removed|
|pull_request_links_url|text|removed|
|reactions|jsonb|added|
|reactions_confused|bigint|removed|
|reactions_eyes|bigint|removed|
|reactions_heart|bigint|removed|
|reactions_hooray|bigint|removed|
|reactions_laugh|bigint|removed|
|reactions_plus_one|bigint|removed|
|reactions_rocket|bigint|removed|
|reactions_total_count|bigint|removed|
|reactions_url|text|removed|
|repository|jsonb|added|
|repository_id|bigint|removed|
|user|jsonb|added|
|user_avatar_url|text|removed|
|user_bio|text|removed|
|user_blog|text|removed|
|user_collaborators|bigint|removed|
|user_company|text|removed|
|user_created_at_time|timestamp without time zone|removed|
|user_disk_usage|bigint|removed|
|user_email|text|removed|
|user_events_url|text|removed|
|user_followers|bigint|removed|
|user_followers_url|text|removed|
|user_following|bigint|removed|
|user_following_url|text|removed|
|user_gists_url|text|removed|
|user_gravatar_id|text|removed|
|user_hireable|boolean|removed|
|user_html_url|text|removed|
|user_id|bigint|removed|
|user_ldap_dn|text|removed|
|user_location|text|removed|
|user_login|text|removed|
|user_name|text|removed|
|user_node_id|text|removed|
|user_organizations_url|text|removed|
|user_owned_private_repos|bigint|removed|
|user_permissions|jsonb|removed|
|user_plan_collaborators|bigint|removed|
|user_plan_filled_seats|bigint|removed|
|user_plan_name|text|removed|
|user_plan_private_repos|bigint|removed|
|user_plan_seats|bigint|removed|
|user_plan_space|bigint|removed|
|user_private_gists|bigint|removed|
|user_public_gists|bigint|removed|
|user_public_repos|bigint|removed|
|user_received_events_url|text|removed|
|user_repos_url|text|removed|
|user_role_name|text|removed|
|user_site_admin|boolean|removed|
|user_starred_url|text|removed|
|user_subscriptions_url|text|removed|
|user_suspended_at_time|timestamp without time zone|removed|
|user_text_matches|jsonb|removed|
|user_total_private_repos|bigint|removed|
|user_twitter_username|text|removed|
|user_two_factor_authentication|boolean|removed|
|user_type|text|removed|
|user_updated_at_time|timestamp without time zone|removed|
|user_url|text|removed|

## github_organization_member_membership
Moved to JSON column on [github_organizations](#github_organizations)


## github_organization_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|membership|jsonb|added|
|organization_cq_id|uuid|removed|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_organizations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|default_repo_permission|text|removed|
|default_repo_settings|text|removed|
|default_repository_permission|text|added|
|default_repository_settings|text|added|
|members_can_create_internal_repos|boolean|removed|
|members_can_create_internal_repositories|boolean|added|
|members_can_create_private_repos|boolean|removed|
|members_can_create_private_repositories|boolean|added|
|members_can_create_public_repos|boolean|removed|
|members_can_create_public_repositories|boolean|added|
|members_can_create_repos|boolean|removed|
|members_can_create_repositories|boolean|added|
|members_can_fork_private_repos|boolean|removed|
|members_can_fork_private_repositories|boolean|added|
|org|text|added|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|

## github_package_billing
This table was removed.


## github_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|code_of_conduct|jsonb|added|
|code_of_conduct_body|text|removed|
|code_of_conduct_key|text|removed|
|code_of_conduct_name|text|removed|
|code_of_conduct_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|license|jsonb|added|
|license_body|text|removed|
|license_conditions|text[]|removed|
|license_description|text|removed|
|license_featured|boolean|removed|
|license_html_url|text|removed|
|license_implementation|text|removed|
|license_key|text|removed|
|license_limitations|text[]|removed|
|license_name|text|removed|
|license_permissions|text[]|removed|
|license_spdx_id|text|removed|
|license_url|text|removed|
|organization|jsonb|added|
|organization_avatar_url|text|removed|
|organization_billing_email|text|removed|
|organization_blog|text|removed|
|organization_collaborators|bigint|removed|
|organization_company|text|removed|
|organization_created_at|timestamp without time zone|removed|
|organization_default_repo_permission|text|removed|
|organization_default_repo_settings|text|removed|
|organization_description|text|removed|
|organization_disk_usage|bigint|removed|
|organization_email|text|removed|
|organization_events_url|text|removed|
|organization_followers|bigint|removed|
|organization_following|bigint|removed|
|organization_has_organization_projects|boolean|removed|
|organization_has_repository_projects|boolean|removed|
|organization_hooks_url|text|removed|
|organization_html_url|text|removed|
|organization_id|bigint|removed|
|organization_is_verified|boolean|removed|
|organization_issues_url|text|removed|
|organization_location|text|removed|
|organization_login|text|removed|
|organization_members_allowed_repository_creation_type|text|removed|
|organization_members_can_create_internal_repos|boolean|removed|
|organization_members_can_create_pages|boolean|removed|
|organization_members_can_create_private_pages|boolean|removed|
|organization_members_can_create_private_repos|boolean|removed|
|organization_members_can_create_public_pages|boolean|removed|
|organization_members_can_create_public_repos|boolean|removed|
|organization_members_can_create_repos|boolean|removed|
|organization_members_can_fork_private_repos|boolean|removed|
|organization_members_url|text|removed|
|organization_name|text|removed|
|organization_node_id|text|removed|
|organization_owned_private_repos|bigint|removed|
|organization_plan_collaborators|bigint|removed|
|organization_plan_filled_seats|bigint|removed|
|organization_plan_name|text|removed|
|organization_plan_private_repos|bigint|removed|
|organization_plan_seats|bigint|removed|
|organization_plan_space|bigint|removed|
|organization_private_gists|bigint|removed|
|organization_public_gists|bigint|removed|
|organization_public_members_url|text|removed|
|organization_public_repos|bigint|removed|
|organization_repos_url|text|removed|
|organization_total_private_repos|bigint|removed|
|organization_twitter_username|text|removed|
|organization_two_factor_requirement_enabled|boolean|removed|
|organization_type|text|removed|
|organization_updated_at|timestamp without time zone|removed|
|organization_url|text|removed|
|owner|jsonb|added|
|owner_avatar_url|text|removed|
|owner_bio|text|removed|
|owner_blog|text|removed|
|owner_collaborators|bigint|removed|
|owner_company|text|removed|
|owner_created_at_time|timestamp without time zone|removed|
|owner_disk_usage|bigint|removed|
|owner_email|text|removed|
|owner_events_url|text|removed|
|owner_followers|bigint|removed|
|owner_followers_url|text|removed|
|owner_following|bigint|removed|
|owner_following_url|text|removed|
|owner_gists_url|text|removed|
|owner_gravatar_id|text|removed|
|owner_hireable|boolean|removed|
|owner_html_url|text|removed|
|owner_id|bigint|removed|
|owner_ldap_dn|text|removed|
|owner_location|text|removed|
|owner_login|text|removed|
|owner_name|text|removed|
|owner_node_id|text|removed|
|owner_organizations_url|text|removed|
|owner_owned_private_repos|bigint|removed|
|owner_permissions|jsonb|removed|
|owner_plan_collaborators|bigint|removed|
|owner_plan_filled_seats|bigint|removed|
|owner_plan_name|text|removed|
|owner_plan_private_repos|bigint|removed|
|owner_plan_seats|bigint|removed|
|owner_plan_space|bigint|removed|
|owner_private_gists|bigint|removed|
|owner_public_gists|bigint|removed|
|owner_public_repos|bigint|removed|
|owner_received_events_url|text|removed|
|owner_repos_url|text|removed|
|owner_role_name|text|removed|
|owner_site_admin|boolean|removed|
|owner_starred_url|text|removed|
|owner_subscriptions_url|text|removed|
|owner_suspended_at_time|timestamp without time zone|removed|
|owner_text_matches|jsonb|removed|
|owner_total_private_repos|bigint|removed|
|owner_twitter_username|text|removed|
|owner_two_factor_authentication|boolean|removed|
|owner_type|text|removed|
|owner_updated_at_time|timestamp without time zone|removed|
|owner_url|text|removed|
|parent|jsonb|updated|Type changed from bigint to jsonb
|pushed_at|timestamp without time zone|added|
|pushed_at_time|timestamp without time zone|removed|
|security_and_analysis|jsonb|added|
|security_and_analysis_advanced_security_status|text|removed|
|security_and_analysis_secret_scanning_status|text|removed|
|source|jsonb|updated|Type changed from bigint to jsonb
|template_repository|jsonb|updated|Type changed from bigint to jsonb
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_storage_billing
This table was removed.


## github_team_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|membership|jsonb|added|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|team_cq_id|uuid|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_team_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|code_of_conduct|jsonb|added|
|code_of_conduct_body|text|removed|
|code_of_conduct_key|text|removed|
|code_of_conduct_name|text|removed|
|code_of_conduct_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|license|jsonb|added|
|license_body|text|removed|
|license_conditions|text[]|removed|
|license_description|text|removed|
|license_featured|boolean|removed|
|license_html_url|text|removed|
|license_implementation|text|removed|
|license_key|text|removed|
|license_limitations|text[]|removed|
|license_name|text|removed|
|license_permissions|text[]|removed|
|license_s_p_d_x_id|text|removed|
|license_url|text|removed|
|org|text|added|
|organization|jsonb|added|
|organization_avatar_url|text|removed|
|organization_billing_email|text|removed|
|organization_blog|text|removed|
|organization_collaborators|bigint|removed|
|organization_company|text|removed|
|organization_created_at|timestamp without time zone|removed|
|organization_default_repo_permission|text|removed|
|organization_default_repo_settings|text|removed|
|organization_description|text|removed|
|organization_disk_usage|bigint|removed|
|organization_email|text|removed|
|organization_events_url|text|removed|
|organization_followers|bigint|removed|
|organization_following|bigint|removed|
|organization_has_organization_projects|boolean|removed|
|organization_has_repository_projects|boolean|removed|
|organization_hooks_url|text|removed|
|organization_html_url|text|removed|
|organization_id|bigint|removed|
|organization_is_verified|boolean|removed|
|organization_issues_url|text|removed|
|organization_location|text|removed|
|organization_login|text|removed|
|organization_members_allowed_repository_creation_type|text|removed|
|organization_members_can_create_internal_repos|boolean|removed|
|organization_members_can_create_pages|boolean|removed|
|organization_members_can_create_private_pages|boolean|removed|
|organization_members_can_create_private_repos|boolean|removed|
|organization_members_can_create_public_pages|boolean|removed|
|organization_members_can_create_public_repos|boolean|removed|
|organization_members_can_create_repos|boolean|removed|
|organization_members_can_fork_private_repos|boolean|removed|
|organization_members_url|text|removed|
|organization_name|text|removed|
|organization_node_id|text|removed|
|organization_owned_private_repos|bigint|removed|
|organization_plan_collaborators|bigint|removed|
|organization_plan_filled_seats|bigint|removed|
|organization_plan_name|text|removed|
|organization_plan_private_repos|bigint|removed|
|organization_plan_seats|bigint|removed|
|organization_plan_space|bigint|removed|
|organization_private_gists|bigint|removed|
|organization_public_gists|bigint|removed|
|organization_public_members_url|text|removed|
|organization_public_repos|bigint|removed|
|organization_repos_url|text|removed|
|organization_total_private_repos|bigint|removed|
|organization_twitter_username|text|removed|
|organization_two_factor_requirement_enabled|boolean|removed|
|organization_type|text|removed|
|organization_updated_at|timestamp without time zone|removed|
|organization_url|text|removed|
|owner|jsonb|added|
|owner_avatar_url|text|removed|
|owner_bio|text|removed|
|owner_blog|text|removed|
|owner_collaborators|bigint|removed|
|owner_company|text|removed|
|owner_created_at_time|timestamp without time zone|removed|
|owner_disk_usage|bigint|removed|
|owner_email|text|removed|
|owner_events_url|text|removed|
|owner_followers|bigint|removed|
|owner_followers_url|text|removed|
|owner_following|bigint|removed|
|owner_following_url|text|removed|
|owner_gists_url|text|removed|
|owner_gravatar_id|text|removed|
|owner_hireable|boolean|removed|
|owner_html_url|text|removed|
|owner_id|bigint|removed|
|owner_ldap_dn|text|removed|
|owner_location|text|removed|
|owner_login|text|removed|
|owner_name|text|removed|
|owner_node_id|text|removed|
|owner_organizations_url|text|removed|
|owner_owned_private_repos|bigint|removed|
|owner_permissions|jsonb|removed|
|owner_plan_collaborators|bigint|removed|
|owner_plan_filled_seats|bigint|removed|
|owner_plan_name|text|removed|
|owner_plan_private_repos|bigint|removed|
|owner_plan_seats|bigint|removed|
|owner_plan_space|bigint|removed|
|owner_private_gists|bigint|removed|
|owner_public_gists|bigint|removed|
|owner_public_repos|bigint|removed|
|owner_received_events_url|text|removed|
|owner_repos_url|text|removed|
|owner_role_name|text|removed|
|owner_site_admin|boolean|removed|
|owner_starred_url|text|removed|
|owner_subscriptions_url|text|removed|
|owner_suspended_at_time|timestamp without time zone|removed|
|owner_text_matches|jsonb|removed|
|owner_total_private_repos|bigint|removed|
|owner_twitter_username|text|removed|
|owner_two_factor_authentication|boolean|removed|
|owner_type|text|removed|
|owner_updated_at_time|timestamp without time zone|removed|
|owner_url|text|removed|
|parent|jsonb|updated|Type changed from bigint to jsonb
|pushed_at|timestamp without time zone|added|
|pushed_at_time|timestamp without time zone|removed|
|s_v_n_url|text|removed|
|security_and_analysis|jsonb|added|
|security_and_analysis_advanced_security_status|text|removed|
|security_and_analysis_secret_scanning_status|text|removed|
|source|jsonb|updated|Type changed from bigint to jsonb
|svn_url|text|added|
|team_cq_id|uuid|removed|
|template_repository|jsonb|updated|Type changed from bigint to jsonb
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|
|use_squash_p_r_title_as_default|boolean|removed|
|use_squash_pr_title_as_default|boolean|added|

## github_teams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|ldap_dn|text|added|
|ldapdn|text|removed|
|organization|jsonb|added|
|parent|jsonb|updated|Type changed from bigint to jsonb
