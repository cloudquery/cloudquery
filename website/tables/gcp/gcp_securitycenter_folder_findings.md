# Table: gcp_securitycenter_folder_findings

This table shows data for GCP Securitycenter Folder Findings.

https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult

The composite primary key for this table is (**folder_id**, **name**).
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|folder_id (PK)|String|
|name (PK)|String|
|finding|JSON|
|state_change|String|
|resource|JSON|