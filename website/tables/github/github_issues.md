# Table: github_issues

This table shows data for Github Issues.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|id (PK)|`int64`|
|number|`int64`|
|state|`utf8`|
|state_reason|`utf8`|
|locked|`bool`|
|title|`utf8`|
|body|`utf8`|
|author_association|`utf8`|
|user|`json`|
|labels|`json`|
|assignee|`json`|
|comments|`int64`|
|closed_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|closed_by|`json`|
|url|`utf8`|
|html_url|`utf8`|
|comments_url|`utf8`|
|events_url|`utf8`|
|labels_url|`utf8`|
|repository_url|`utf8`|
|milestone|`json`|
|pull_request|`json`|
|repository|`json`|
|reactions|`json`|
|assignees|`json`|
|node_id|`utf8`|
|text_matches|`json`|
|active_lock_reason|`utf8`|