# Table: azure_storage_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage#Account

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_containers](azure_storage_containers.md)
  - [azure_storage_blob_services](azure_storage_blob_services.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|kind|String|
|identity|JSON|
|extended_location|JSON|
|provisioning_state|String|
|primary_endpoints|JSON|
|primary_location|String|
|status_of_primary|String|
|last_geo_failover_time|Timestamp|
|secondary_location|String|
|status_of_secondary|String|
|creation_time|Timestamp|
|custom_domain|JSON|
|secondary_endpoints|JSON|
|encryption|JSON|
|access_tier|String|
|azure_files_identity_based_authentication|JSON|
|supports_https_traffic_only|Bool|
|network_acls|JSON|
|is_hns_enabled|Bool|
|geo_replication_stats|JSON|
|failover_in_progress|Bool|
|large_file_shares_state|String|
|private_endpoint_connections|JSON|
|routing_preference|JSON|
|blob_restore_status|JSON|
|allow_blob_public_access|Bool|
|minimum_tls_version|String|
|allow_shared_key_access|Bool|
|tags|JSON|
|location|String|
|id (PK)|String|
|name|String|
|type|String|
|blob_logging_settings|JSON|
|queue_logging_settings|JSON|
|is_nfs_v3_enabled|Bool|