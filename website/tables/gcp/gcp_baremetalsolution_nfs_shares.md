# Table: gcp_baremetalsolution_nfs_shares

This table shows data for GCP Bare Metal Solution NFS Shares.

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.nfsShares#NfsShare

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|nfs_share_id|`utf8`|
|state|`utf8`|
|volume|`utf8`|
|allowed_clients|`json`|
|labels|`json`|