# Table: github_organization_members

This table shows data for Github Organization Members.

The composite primary key for this table is (**org**, **id**).

## Relations

This table depends on [github_organizations](github_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|membership|`json`|
|login|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|avatar_url|`utf8`|
|html_url|`utf8`|
|gravatar_id|`utf8`|
|name|`utf8`|
|company|`utf8`|
|blog|`utf8`|
|location|`utf8`|
|email|`utf8`|
|hireable|`bool`|
|bio|`utf8`|
|twitter_username|`utf8`|
|public_repos|`int64`|
|public_gists|`int64`|
|followers|`int64`|
|following|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|suspended_at|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|site_admin|`bool`|
|total_private_repos|`int64`|
|owned_private_repos|`int64`|
|private_gists|`int64`|
|disk_usage|`int64`|
|collaborators|`int64`|
|two_factor_authentication|`bool`|
|plan|`json`|
|ldap_dn|`utf8`|
|url|`utf8`|
|events_url|`utf8`|
|following_url|`utf8`|
|followers_url|`utf8`|
|gists_url|`utf8`|
|organizations_url|`utf8`|
|received_events_url|`utf8`|
|repos_url|`utf8`|
|starred_url|`utf8`|
|subscriptions_url|`utf8`|
|text_matches|`json`|
|permissions|`json`|
|role_name|`utf8`|