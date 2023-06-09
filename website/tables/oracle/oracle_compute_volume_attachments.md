# Table: oracle_compute_volume_attachments

This table shows data for Oracle Compute Volume Attachments.

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
|availability_domain|`utf8`|
|id (PK)|`utf8`|
|instance_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|volume_id|`utf8`|
|ipv4|`utf8`|
|iqn|`utf8`|
|port|`int64`|
|device|`utf8`|
|display_name|`utf8`|
|is_read_only|`bool`|
|is_shareable|`bool`|
|is_pv_encryption_in_transit_enabled|`bool`|
|is_multipath|`bool`|
|chap_secret|`utf8`|
|chap_username|`utf8`|
|multipath_devices|`json`|
|is_agent_auto_iscsi_login_enabled|`bool`|
|lifecycle_state|`utf8`|
|iscsi_login_state|`utf8`|
|encryption_in_transit_type|`utf8`|