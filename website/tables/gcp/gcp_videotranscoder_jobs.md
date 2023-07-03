# Table: gcp_videotranscoder_jobs

This table shows data for GCP Video Transcoder Jobs.

https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobs

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|input_uri|`utf8`|
|output_uri|`utf8`|
|state|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|ttl_after_completion_days|`int64`|
|labels|`json`|
|error|`json`|