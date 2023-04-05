# Table: snyk_projects

This table shows data for Snyk Projects.

https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Project

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|id (PK)|String|
|name|String|
|origin|String|