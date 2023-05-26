# Table: oracle_networkfirewall_network_firewall_policies

This table shows data for Oracle Network Firewall Network Firewall Policies.

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
|display_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|time_updated|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|system_tags|`json`|