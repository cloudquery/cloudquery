# Table: tailscale_acls

https://pkg.go.dev/github.com/tailscale/tailscale-client-go/tailscale#ACL

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
|auto_approvers|JSON|
|groups|JSON|
|hosts|JSON|
|tag_owners|JSON|
|derp_map|JSON|
|tests|JSON|
|ssh|JSON|
|node_attrs|JSON|
|disable_ipv4|Bool|
|one_cgnat_route|String|
|randomize_client_port|Bool|