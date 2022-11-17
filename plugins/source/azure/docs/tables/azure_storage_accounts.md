# Table: azure_storage_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage#Account

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_blob_services](azure_storage_blob_services.md)
  - [azure_storage_containers](azure_storage_containers.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|allow_blob_public_access|Bool|
|allow_cross_tenant_replication|Bool|
|allow_shared_key_access|Bool|
|allowed_copy_scope|String|
|azure_files_identity_based_authentication|JSON|
|dns_endpoint_type|String|
|default_to_o_auth_authentication|Bool|
|supports_https_traffic_only|Bool|
|is_nfs_v3_enabled|Bool|
|immutable_storage_with_versioning|JSON|
|is_hns_enabled|Bool|
|is_local_user_enabled|Bool|
|is_sftp_enabled|Bool|
|large_file_shares_state|String|
|minimum_tls_version|String|
|public_network_access|String|
|routing_preference|JSON|
|storage_account_sku_conversion_status|JSON|
|access_tier|String|
|blob_restore_status|JSON|
|creation_time|Timestamp|
|custom_domain|JSON|
|encryption|JSON|
|failover_in_progress|Bool|
|geo_replication_stats|JSON|
|key_creation_time|JSON|
|key_policy|JSON|
|last_geo_failover_time|Timestamp|
|network_acls|JSON|
|primary_endpoints|JSON|
|primary_location|String|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|sas_policy|JSON|
|secondary_endpoints|JSON|
|secondary_location|String|
|status_of_primary|String|
|status_of_secondary|String|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|sku|JSON|
|type|String|