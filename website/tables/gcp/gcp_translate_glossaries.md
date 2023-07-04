# Table: gcp_translate_glossaries

This table shows data for GCP Translate Glossaries.

https://cloud.google.com/translate/docs/reference/rest/v3/projects.locations.glossaries#resource:-glossary

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|input_config|`json`|
|entry_count|`int64`|
|submit_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|