# Table: github_teams

This table shows data for Github Teams.

The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_teams:
  - [github_team_members](github_team_members)
  - [github_team_repositories](github_team_repositories)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|url|`utf8`|
|slug|`utf8`|
|permission|`utf8`|
|permissions|`json`|
|privacy|`utf8`|
|members_count|`int64`|
|repos_count|`int64`|
|organization|`json`|
|html_url|`utf8`|
|members_url|`utf8`|
|repositories_url|`utf8`|
|parent|`json`|
|ldap_dn|`utf8`|