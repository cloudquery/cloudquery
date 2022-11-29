# Table: github_teams



The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_teams:
  - [github_team_members](github_team_members.md)
  - [github_team_repositories](github_team_repositories.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|id (PK)|Int|
|node_id|String|
|name|String|
|description|String|
|url|String|
|slug|String|
|permission|String|
|permissions|JSON|
|privacy|String|
|members_count|Int|
|repos_count|Int|
|organization|JSON|
|html_url|String|
|members_url|String|
|repositories_url|String|
|parent|JSON|
|ldap_dn|String|