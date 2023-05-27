# Table: oracle_compute_vnic_attachments

This table shows data for Oracle Compute Virtual Network Interface Card (VNIC) Attachments.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|availability_domain|`utf8`|
|instance_id|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|nic_index|`int64`|
|subnet_id|`utf8`|
|vlan_id|`utf8`|
|vlan_tag|`int64`|
|vnic_id|`utf8`|