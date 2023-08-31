# Table: tailscale_acls

This table shows data for Tailscale Access Control Lists (ACLs).

https://github.com/tailscale/tailscale/blob/main/api.md#acl

The primary key for this table is **tailnet**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|tailnet (PK)|`utf8`|
|acls|`json`|
|derp_map|`json`|
|ssh|`json`|
|disable_ipv4|`bool`|
|one_cgnat_route|`utf8`|
|autoapprovers|`json`|
|groups|`json`|
|hosts|`json`|
|tagowners|`json`|
|tests|`json`|
|node_attrs|`json`|
|randomize_client_port|`bool`|