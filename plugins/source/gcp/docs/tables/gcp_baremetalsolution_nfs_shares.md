# Table: gcp_baremetalsolution_nfs_shares

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.nfsShares#NfsShare

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
|nfs_share_id|String|
|state|String|
|volume|String|
|allowed_clients|JSON|
|labels|JSON|