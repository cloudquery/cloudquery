# Table: gcp_containeranalysis_occurrences

This table shows data for GCP Container Analysis Occurrences.

https://cloud.google.com/container-analysis/docs/reference/rest/v1beta1/projects.occurrences#Occurrence

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|resource|`json`|
|note_name|`utf8`|
|kind|`utf8`|
|remediation|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|