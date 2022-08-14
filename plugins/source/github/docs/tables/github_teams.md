
# Table: github_teams
Team represents a team within a GitHub organization
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|org|text|The Github Organization of the resource.|
|id|bigint||
|node_id|text||
|name|text||
|description|text||
|url|text||
|slug|text||
|permission|text|Permission specifies the default permission for repositories owned by the team.|
|permissions|jsonb|Permissions identifies the permissions that a team has on a given repository|
|privacy|text|Privacy identifies the level of privacy this team should have. Possible values are:     secret - only visible to organization owners and members of this team     closed - visible to all members of this organization Default is "secret".|
|members_count|bigint||
|repos_count|bigint||
|html_url|text||
|members_url|text||
|repositories_url|text||
|parent|bigint||
|ldapdn|text|LDAPDN is only available in GitHub Enterprise and when the team membership is synchronized with LDAP.|
