
# Table: github_organization_members
User represents a GitHub user.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|organization_cq_id|uuid|Unique CloudQuery ID of github_organizations table (FK)|
|org|text|The Github Organization of the resource.|
|login|text||
|id|bigint||
|node_id|text||
|avatar_url|text||
|html_url|text||
|gravatar_id|text||
|name|text||
|company|text||
|blog|text||
|location|text||
|email|text||
|hireable|boolean||
|bio|text||
|twitter_username|text||
|public_repos|bigint||
|public_gists|bigint||
|followers|bigint||
|following|bigint||
|created_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|suspended_at_time|timestamp without time zone||
|type|text||
|site_admin|boolean||
|total_private_repos|bigint||
|owned_private_repos|bigint||
|private_gists|bigint||
|disk_usage|bigint||
|collaborators|bigint||
|two_factor_authentication|boolean||
|plan_name|text||
|plan_space|bigint||
|plan_collaborators|bigint||
|plan_private_repos|bigint||
|plan_filled_seats|bigint||
|plan_seats|bigint||
|ldap_dn|text||
|url|text|API URLs|
|events_url|text||
|following_url|text||
|followers_url|text||
|gists_url|text||
|organizations_url|text||
|received_events_url|text||
|repos_url|text||
|starred_url|text||
|subscriptions_url|text||
|text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|role_name|text||
