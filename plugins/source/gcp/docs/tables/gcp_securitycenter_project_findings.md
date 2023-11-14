# Table: gcp_securitycenter_project_findings

This table shows data for GCP Securitycenter Project Findings.

https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult

The composite primary key for this table is (**project_id**, **name**).
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|finding|`json`|
|state_change|`utf8`|
|resource|`json`|