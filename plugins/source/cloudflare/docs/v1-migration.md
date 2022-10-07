# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

## cloudflare_access_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## cloudflare_account_member_roles
Moved to JSON column on [cloudflare_accounts](#cloudflare_accounts)


## cloudflare_account_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_cq_id|uuid|removed|
|roles|jsonb|added|
|user|jsonb|added|
|user_email|text|removed|
|user_first_name|text|removed|
|user_id|text|removed|
|user_last_name|text|removed|
|user_two_factor_authentication_enabled|boolean|removed|

## cloudflare_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|enforce_two_factor|boolean|removed|
|settings|jsonb|added|

## cloudflare_certificate_pack_certificates
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_pack_validation_errors
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_pack_validation_records
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_packs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|certificates|jsonb|added|
|validation_errors|jsonb|added|
|validation_records|jsonb|added|

## cloudflare_dns_records

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|priority|bigint|updated|Type changed from integer to bigint

## cloudflare_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|require_signed_ur_ls|boolean|added|
|require_signed_url_s|boolean|removed|
|variants|text[]|updated|Type changed from jsonb to text[]

## cloudflare_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## cloudflare_waf
This table was removed.


## cloudflare_waf_groups
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|waf_package_id|text|added|
|id|text|added|
|name|text|added|
|description|text|added|
|rules_count|bigint|added|
|modified_rules_count|bigint|added|
|package_id|text|added|
|mode|text|added|
|allowed_modes|text[]|added|

## cloudflare_waf_overrides

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|url_s|text[]|removed|
|urls|text[]|added|

## cloudflare_waf_packages
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|id|text|added|
|name|text|added|
|description|text|added|
|zone_id|text|added|
|detection_mode|text|added|
|sensitivity|text|added|
|action_mode|text|added|

## cloudflare_waf_rule_groups
Moved to JSON column on [cloudflare_waf_rules](#cloudflare_waf_rules)


## cloudflare_waf_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|removed|
|waf_cq_id|uuid|removed|
|waf_package_id|text|added|
|zone_id|text|removed|

## cloudflare_worker_cron_triggers
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|worker_meta_data_id|text|added|
|cron|text|added|
|created_on|timestamp without time zone|added|
|modified_on|timestamp without time zone|added|

## cloudflare_worker_meta_data
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|id|text|added|
|etag|text|added|
|size|bigint|added|
|created_on|timestamp without time zone|added|
|modified_on|timestamp without time zone|added|

## cloudflare_worker_routes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|zone_id|text|added|
|id|text|added|
|pattern|text|added|
|enabled|boolean|added|
|script|text|added|

## cloudflare_workers_routes
This table was removed.


## cloudflare_workers_script_cron_triggers
This table was removed.


## cloudflare_workers_script_secrets
This table was removed.


## cloudflare_workers_scripts
This table was removed.


## cloudflare_workers_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|worker_meta_data_id|text|added|
|name|text|added|
|type|text|added|

## cloudflare_zones

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account|jsonb|added|
|deact_reason|text|added|
|deactivation_reason|text|removed|
|host|jsonb|added|
|host_name|text|removed|
|host_website|text|removed|
|meta|jsonb|added|
|owner|jsonb|added|
|owner_email|text|removed|
|owner_id|text|removed|
|owner_name|text|removed|
|owner_type|text|removed|
|page_rule_quota|bigint|removed|
|phishing_detected|boolean|removed|
|plan|jsonb|added|
|plan_can_subscribe|boolean|removed|
|plan_currency|text|removed|
|plan_externally_managed|boolean|removed|
|plan_frequency|text|removed|
|plan_id|text|removed|
|plan_is_subscribed|boolean|removed|
|plan_legacy_discount|boolean|removed|
|plan_legacy_id|text|removed|
|plan_name|text|removed|
|plan_pending|jsonb|added|
|plan_pending_can_subscribe|boolean|removed|
|plan_pending_currency|text|removed|
|plan_pending_externally_managed|boolean|removed|
|plan_pending_frequency|text|removed|
|plan_pending_id|text|removed|
|plan_pending_is_subscribed|boolean|removed|
|plan_pending_legacy_discount|boolean|removed|
|plan_pending_legacy_id|text|removed|
|plan_pending_name|text|removed|
|plan_pending_price|bigint|removed|
|plan_price|bigint|removed|
|wildcard_proxiable|boolean|removed|
