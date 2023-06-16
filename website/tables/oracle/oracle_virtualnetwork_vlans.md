# Table: oracle_virtualnetwork_vlans

This table shows data for Oracle Virtual Network Virtual LANs (VLANs).

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
|cidr_block|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|vcn_id|`utf8`|
|availability_domain|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|nsg_ids|`list<item: utf8, nullable>`|
|vlan_tag|`int64`|
|route_table_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|