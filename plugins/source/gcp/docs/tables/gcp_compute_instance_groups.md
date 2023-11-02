# Table: gcp_compute_instance_groups

This table shows data for GCP Compute Instance Groups.

https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups#InstanceGroup

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|fingerprint|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|named_ports|`json`|
|network|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|size|`int64`|
|subnetwork|`utf8`|
|zone|`utf8`|