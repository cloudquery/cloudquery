
# Table: gcp_compute_project_quotas
A quotas entry
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|uuid|Unique ID of gcp_compute_projects table (FK)|
|limit|float|Quota limit for this metric|
|metric|text|Name of the quota metric|
|owner|text|Owning resource This is the resource on which this quota is applied|
|usage|float|Current usage of this metric|
