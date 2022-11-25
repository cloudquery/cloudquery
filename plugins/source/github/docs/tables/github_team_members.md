# Table: github_team_members



The composite primary key for this table is (**org**, **id**, **team_id**).

## Relations
This table depends on [github_teams](github_teams.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|id (PK)|Int|
|team_id (PK)|Int|
|membership|JSON|
|login|String|
|node_id|String|
|avatar_url|String|
|html_url|String|
|gravatar_id|String|
|name|String|
|company|String|
|blog|String|
|location|String|
|email|String|
|hireable|Bool|
|bio|String|
|twitter_username|String|
|public_repos|Int|
|public_gists|Int|
|followers|Int|
|following|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|suspended_at|Timestamp|
|type|String|
|site_admin|Bool|
|total_private_repos|Int|
|owned_private_repos|Int|
|private_gists|Int|
|disk_usage|Int|
|collaborators|Int|
|two_factor_authentication|Bool|
|plan|JSON|
|ldap_dn|String|
|url|String|
|events_url|String|
|following_url|String|
|followers_url|String|
|gists_url|String|
|organizations_url|String|
|received_events_url|String|
|repos_url|String|
|starred_url|String|
|subscriptions_url|String|
|text_matches|JSON|
|permissions|JSON|
|role_name|String|