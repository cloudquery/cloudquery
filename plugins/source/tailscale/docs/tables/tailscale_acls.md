# Table: tailscale_acls

https://github.com/tailscale/tailscale/blob/main/api.md#acl

The primary key for this table is **tailnet**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|tailnet (PK)|String|
|acls|JSON|
|derp_map|JSON|
|ssh|JSON|
|disable_ipv4|Bool|
|one_cgnat_route|String|
|autoapprovers|JSON|
|groups|JSON|
|hosts|JSON|
|tagowners|JSON|
|tests|JSON|
|node_attrs|JSON|
|randomize_client_port|Bool|