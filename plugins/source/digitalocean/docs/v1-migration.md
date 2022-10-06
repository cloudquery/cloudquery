# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

## digitalocean_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|reserved_ip_limit|bigint|added|
|team|jsonb|added|

## digitalocean_balance
Moved to JSON column on [digitalocean_balances](#digitalocean_balances)


## digitalocean_balances
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|month_to_date_balance|text|added|
|account_balance|text|added|
|month_to_date_usage|text|added|
|generated_at|timestamp without time zone|added|

## digitalocean_billing_history

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## digitalocean_cdn
Moved to JSON column on [digitalocean_cdns](#digitalocean_cdns)


## digitalocean_cdns
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|added|
|origin|text|added|
|endpoint|text|added|
|created_at|timestamp without time zone|added|
|ttl|bigint|added|
|certificate_id|text|added|
|custom_domain|text|added|

## digitalocean_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created|text|removed|
|created_at|text|added|
|s_h_a1_fingerprint|text|removed|
|sha_1_fingerprint|text|added|

## digitalocean_database_backups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|database_cq_id|uuid|removed|
|size_gigabytes|real|updated|Type changed from float to real

## digitalocean_database_firewall_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|database_cq_id|uuid|removed|

## digitalocean_database_replicas

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|connection|jsonb|added|
|connection_database|text|removed|
|connection_host|text|removed|
|connection_password|text|removed|
|connection_port|bigint|removed|
|connection_ssl|boolean|removed|
|connection_uri|text|removed|
|connection_user|text|removed|
|database_cq_id|uuid|removed|
|private_connection|jsonb|added|
|private_connection_database|text|removed|
|private_connection_host|text|removed|
|private_connection_password|text|removed|
|private_connection_port|bigint|removed|
|private_connection_ssl|boolean|removed|
|private_connection_uri|text|removed|
|private_connection_user|text|removed|

## digitalocean_database_users
Moved to JSON column on [digitalocean_databases](#digitalocean_databases)


## digitalocean_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|connection|jsonb|added|
|connection_database|text|removed|
|connection_host|text|removed|
|connection_password|text|removed|
|connection_port|bigint|removed|
|connection_ssl|boolean|removed|
|connection_uri|text|removed|
|connection_user|text|removed|
|maintenance_window|jsonb|added|
|maintenance_window_day|text|removed|
|maintenance_window_description|text[]|removed|
|maintenance_window_hour|text|removed|
|maintenance_window_pending|boolean|removed|
|private_connection|jsonb|added|
|private_connection_database|text|removed|
|private_connection_host|text|removed|
|private_connection_password|text|removed|
|private_connection_port|bigint|removed|
|private_connection_ssl|boolean|removed|
|private_connection_uri|text|removed|
|private_connection_user|text|removed|
|project_id|text|added|
|region|text|added|
|region_slug|text|removed|
|size|text|added|
|size_slug|text|removed|
|users|jsonb|added|

## digitalocean_domain_records

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|domain_cq_id|uuid|removed|
|id|text|updated|Type changed from bigint to text

## digitalocean_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## digitalocean_droplet_neighbors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|droplet_cq_id|uuid|removed|

## digitalocean_droplet_networks_v4
Moved to JSON column on [digitalocean_droplets](#digitalocean_droplets)


## digitalocean_droplet_networks_v6
Moved to JSON column on [digitalocean_droplets](#digitalocean_droplets)


## digitalocean_droplets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|backup_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|created|text|removed|
|created_at|text|added|
|image|jsonb|added|
|image_created|text|removed|
|image_description|text|removed|
|image_distribution|text|removed|
|image_error_message|text|removed|
|image_id|bigint|removed|
|image_min_disk_size|bigint|removed|
|image_name|text|removed|
|image_public|boolean|removed|
|image_regions|text[]|removed|
|image_size_giga_bytes|float|removed|
|image_slug|text|removed|
|image_status|text|removed|
|image_tags|text[]|removed|
|image_type|text|removed|
|kernel|jsonb|added|
|kernel_id|bigint|removed|
|kernel_name|text|removed|
|kernel_version|text|removed|
|networks|jsonb|added|
|next_backup_window|jsonb|added|
|next_backup_window_end_time|timestamp without time zone|removed|
|next_backup_window_start_time|timestamp without time zone|removed|
|region|jsonb|added|
|region_available|boolean|removed|
|region_features|text[]|removed|
|region_name|text|removed|
|region_sizes|text[]|removed|
|region_slug|text|removed|
|size|jsonb|added|
|size_available|boolean|removed|
|size_description|text|removed|
|size_disk|bigint|removed|
|size_memory|bigint|removed|
|size_price_hourly|float|removed|
|size_price_monthly|float|removed|
|size_regions|text[]|removed|
|size_transfer|float|removed|
|size_vcpus|bigint|removed|
|snapshot_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|volume_ids|bigint[]|updated|Type changed from text[] to bigint[]

## digitalocean_firewall_droplets
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_inbound_rules
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_outbound_rules
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_pending_changes
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewalls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created|text|removed|
|created_at|text|added|
|droplet_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|id|text|updated|Type changed from uuid to text
|inbound_rules|jsonb|added|
|outbound_rules|jsonb|added|
|pending_changes|jsonb|added|

## digitalocean_floating_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|droplet|jsonb|added|
|droplet_id|bigint|removed|
|ip|text|updated|Type changed from cidr to text
|region|jsonb|added|
|region_available|boolean|removed|
|region_features|text[]|removed|
|region_name|text|removed|
|region_sizes|text[]|removed|
|region_slug|text|removed|

## digitalocean_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created|text|removed|
|created_at|text|added|
|size_giga_bytes|float|removed|
|size_gigabytes|real|added|

## digitalocean_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## digitalocean_load_balancer_droplets
This table was removed.


## digitalocean_load_balancer_forwarding_rules
This table was removed.


## digitalocean_load_balancers
This table was removed.


## digitalocean_monitoring_alert_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|uuid|text|added|
|type|text|added|
|description|text|added|
|compare|text|added|
|value|real|added|
|window|text|added|
|entities|text[]|added|
|tags|text[]|added|
|alerts|jsonb|added|
|enabled|boolean|added|

## digitalocean_project_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|assigned_at|text|updated|Type changed from timestamp without time zone to text
|links|jsonb|added|
|links_self|text|removed|
|project_cq_id|uuid|removed|

## digitalocean_projects

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## digitalocean_regions
This table was removed.


## digitalocean_registries
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|name|text|added|
|storage_usage_bytes|bigint|added|
|storage_usage_bytes_updated_at|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|region|text|added|

## digitalocean_registry
Moved to JSON column on [digitalocean_registries](#digitalocean_registries)


## digitalocean_registry_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|latest_tag|jsonb|updated|Type changed from text to jsonb
|latest_tag_compressed_size_bytes|bigint|removed|
|latest_tag_manifest_digest|text|removed|
|latest_tag_registry_name|text|removed|
|latest_tag_repository|text|removed|
|latest_tag_size_bytes|bigint|removed|
|latest_tag_updated_at|timestamp without time zone|removed|
|registry_cq_id|uuid|removed|

## digitalocean_sizes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|price_hourly|real|updated|Type changed from float to real
|price_monthly|real|updated|Type changed from float to real
|transfer|real|updated|Type changed from float to real

## digitalocean_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|created|text|removed|
|created_at|text|added|
|size_giga_bytes|float|removed|
|size_gigabytes|real|added|

## digitalocean_space_acls
Moved to JSON column on [digitalocean_spaces](#digitalocean_spaces)


## digitalocean_space_cors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|expose_headers|text[]|added|
|id|text|added|
|max_age_seconds|bigint|updated|Type changed from integer to bigint
|space_cq_id|uuid|removed|
|space_name|text|removed|

## digitalocean_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|acls|jsonb|added|
|bucket|jsonb|added|
|creation_date|timestamp without time zone|removed|
|name|text|removed|

## digitalocean_storage_volumes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|added|
|droplet_ids|bigint[]|added|
|region|jsonb|added|
|name|text|added|
|size_gigabytes|bigint|added|
|description|text|added|
|created_at|timestamp without time zone|added|
|filesystem_type|text|added|
|filesystem_label|text|added|
|tags|text[]|added|

## digitalocean_volume_droplets
This table was removed.


## digitalocean_volumes
This table was removed.


## digitalocean_vpc_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|type|text|removed|
|vpc_cq_id|uuid|removed|

## digitalocean_vpcs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|ip_range|text|updated|Type changed from cidr to text
|region|text|added|
|region_slug|text|removed|
