# Table: github_installations



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
|node_id|String|
|app_id|Int|
|app_slug|String|
|target_id|Int|
|account|JSON|
|access_tokens_url|String|
|repositories_url|String|
|html_url|String|
|target_type|String|
|single_file_name|String|
|repository_selection|String|
|events|StringArray|
|single_file_paths|StringArray|
|permissions|JSON|
|created_at|Timestamp|
|updated_at|Timestamp|
|has_multiple_single_files|Bool|
|suspended_by|JSON|
|suspended_at|Timestamp|