# Table: gcp_videotranscoder_job_templates

This table shows data for GCP Video Transcoder Job Templates.

https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobTemplates

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|config|`json`|
|labels|`json`|