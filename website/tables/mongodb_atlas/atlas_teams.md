# Table: atlas_teams

This table shows data for Atlas Teams.

The composite primary key for this table is (**group_id**, **team_id**).

## Relations

The following tables depend on atlas_teams:
  - [atlas_team_users](atlas_team_users.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|group_id (PK)|`utf8`|
|links|`json`|
|role_names|`list<item: utf8, nullable>`|
|team_id (PK)|`utf8`|