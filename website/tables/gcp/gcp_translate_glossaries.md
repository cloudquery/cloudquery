# Table: gcp_translate_glossaries

https://cloud.google.com/translate/docs/reference/rest/v3/projects.locations.glossaries#resource:-glossary

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|input_config|JSON|
|entry_count|Int|
|submit_time|Timestamp|
|end_time|Timestamp|