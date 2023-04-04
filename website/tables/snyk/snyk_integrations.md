# Table: snyk_integrations

This table shows data for Snyk Integrations.

https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Integration

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|settings|JSON|
|credentials|JSON|
|id (PK)|String|
|type|String|