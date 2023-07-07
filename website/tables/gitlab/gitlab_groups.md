# Table: gitlab_groups

This table shows data for Gitlab Groups.

The composite primary key for this table is (**base_url**, **id**, **name**).

## Relations

The following tables depend on gitlab_groups:
  - [gitlab_group_billable_members](gitlab_group_billable_members)
  - [gitlab_group_members](gitlab_group_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|id (PK)|`int64`|
|name (PK)|`utf8`|
|path|`utf8`|
|description|`utf8`|
|membership_lock|`bool`|
|visibility|`utf8`|
|lfs_enabled|`bool`|
|default_branch_protection|`int64`|
|avatar_url|`utf8`|
|web_url|`utf8`|
|request_access_enabled|`bool`|
|full_name|`utf8`|
|full_path|`utf8`|
|file_template_project_id|`int64`|
|parent_id|`int64`|
|projects|`json`|
|statistics|`json`|
|custom_attributes|`json`|
|share_with_group_lock|`bool`|
|require_two_factor_authentication|`bool`|
|two_factor_grace_period|`int64`|
|project_creation_level|`utf8`|
|auto_devops_enabled|`bool`|
|subgroup_creation_level|`utf8`|
|emails_disabled|`bool`|
|mentions_disabled|`bool`|
|runners_token|`utf8`|
|shared_projects|`json`|
|shared_runners_enabled|`bool`|
|shared_with_groups|`json`|
|ldap_cn|`utf8`|
|ldap_access|`int64`|
|ldap_group_links|`json`|
|saml_group_links|`json`|
|shared_runners_minutes_limit|`int64`|
|extra_shared_runners_minutes_limit|`int64`|
|prevent_forking_outside_group|`bool`|
|marked_for_deletion_on|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|ip_restriction_ranges|`utf8`|