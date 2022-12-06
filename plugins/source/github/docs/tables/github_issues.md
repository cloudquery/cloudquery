# Table: github_issues



The composite primary key for this table is (**org**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|id (PK)|Int|
|number|Int|
|state|String|
|locked|Bool|
|title|String|
|body|String|
|author_association|String|
|user|JSON|
|labels|JSON|
|assignee|JSON|
|comments|Int|
|closed_at|Timestamp|
|created_at|Timestamp|
|updated_at|Timestamp|
|closed_by|JSON|
|url|String|
|html_url|String|
|comments_url|String|
|events_url|String|
|labels_url|String|
|repository_url|String|
|milestone|JSON|
|pull_request|JSON|
|repository|JSON|
|reactions|JSON|
|assignees|JSON|
|node_id|String|
|text_matches|JSON|
|active_lock_reason|String|