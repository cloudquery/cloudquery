# Table: atlas_database_clusters

This table shows data for Atlas Database Clusters.

The composite primary key for this table is (**org_id**, **group_id**, **cluster_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org_id (PK)|`utf8`|
|group_id (PK)|`utf8`|
|alert_count|`int64`|
|auth_enabled|`bool`|
|availability|`utf8`|
|backup_enabled|`bool`|
|cluster_id (PK)|`utf8`|
|data_size_bytes|`int64`|
|name|`utf8`|
|node_count|`int64`|
|ssl_enabled|`bool`|
|type|`utf8`|
|versions|`list<item: utf8, nullable>`|