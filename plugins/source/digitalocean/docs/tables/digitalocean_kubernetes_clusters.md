# Table: digitalocean_kubernetes_clusters

This table shows data for DigitalOcean Kubernetes Clusters.

https://docs.digitalocean.com/reference/api/api-reference/#operation/kubernetes_list_clusters

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|region|`utf8`|
|version|`utf8`|
|cluster_subnet|`utf8`|
|service_subnet|`utf8`|
|ipv4|`utf8`|
|endpoint|`utf8`|
|tags|`list<item: utf8, nullable>`|
|vpc_uuid|`utf8`|
|ha|`bool`|
|node_pools|`json`|
|maintenance_policy|`json`|
|auto_upgrade|`bool`|
|surge_upgrade|`bool`|
|registry_enabled|`bool`|
|status|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|