# Schema Changes from v1 to v2
This guide summarizes schema changes from CloudQuery v1 to v2. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v2.

Last updated Fri Sep 23 15:18:12 BST 2022.

## heroku_account_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_region_capabilities

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_webhook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|removed|
|event|jsonb|removed|
|id|text|removed|
|last_attempt|jsonb|removed|
|next_attempt_at|timestamp without time zone|removed|
|num_attempts|integer|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|webhook|jsonb|removed|

## heroku_add_on_webhook_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|include|text|removed|
|payload|jsonb|removed|
|updated_at|timestamp without time zone|removed|

## heroku_add_on_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon|jsonb|removed|

## heroku_add_ons

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_transfers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_webhook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|num_attempts|bigint|updated|Type changed from integer to bigint

## heroku_app_webhook_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|include|text[]|removed|
|level|text|removed|
|updated_at|timestamp without time zone|removed|
|url|text|removed|

## heroku_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|repo_size|bigint|updated|Type changed from integer to bigint
|slug_size|bigint|updated|Type changed from integer to bigint

## heroku_buildpack_installations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ordinal|bigint|updated|Type changed from integer to bigint

## heroku_builds

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_collaborators

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_credits

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|amount|real|updated|Type changed from float to real
|balance|real|updated|Type changed from float to real

## heroku_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_dyno_sizes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|compute|bigint|updated|Type changed from integer to bigint
|dyno_units|bigint|updated|Type changed from integer to bigint
|memory|real|updated|Type changed from float to real

## heroku_dynos

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_enterprise_account_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_enterprise_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_formations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|quantity|bigint|updated|Type changed from integer to bigint

## heroku_inbound_rulesets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_invoices

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|charges_total|real|updated|Type changed from float to real
|credits_total|real|updated|Type changed from float to real
|number|bigint|updated|Type changed from integer to bigint
|state|bigint|updated|Type changed from integer to bigint
|total|real|updated|Type changed from float to real

## heroku_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_log_drains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_oauth_authorizations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_oauth_clients

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_outbound_rulesets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_peerings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_permission_entities

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_pipeline_builds

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|app|jsonb|removed|
|buildpacks|jsonb|removed|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|output_stream_url|text|removed|
|release|jsonb|removed|
|slug|jsonb|removed|
|source_blob|jsonb|removed|
|stack|text|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|

## heroku_pipeline_couplings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_pipeline_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon_plan_names|text[]|removed|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|current|boolean|removed|
|description|text|removed|
|id|text|removed|
|output_stream_url|text|removed|
|slug|jsonb|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|
|version|integer|removed|

## heroku_pipeline_releases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon_plan_names|text[]|removed|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|current|boolean|removed|
|description|text|removed|
|id|text|removed|
|output_stream_url|text|removed|
|slug|jsonb|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|
|version|integer|removed|

## heroku_pipelines

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_regions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_releases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|version|bigint|updated|Type changed from integer to bigint

## heroku_review_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|pr_number|bigint|updated|Type changed from integer to bigint

## heroku_space_app_accesses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_stacks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_app_permissions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_invitations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_invoices

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addons_total|bigint|updated|Type changed from integer to bigint
|charges_total|bigint|updated|Type changed from integer to bigint
|credits_total|bigint|updated|Type changed from integer to bigint
|database_total|bigint|updated|Type changed from integer to bigint
|dyno_units|real|updated|Type changed from float to real
|number|bigint|updated|Type changed from integer to bigint
|platform_total|bigint|updated|Type changed from integer to bigint
|state|bigint|updated|Type changed from integer to bigint
|total|bigint|updated|Type changed from integer to bigint
|weighted_dyno_hours|real|updated|Type changed from float to real

## heroku_team_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cidr|text|removed|
|created_at|timestamp without time zone|removed|
|data_cidr|text|removed|
|id|text|removed|
|name|text|removed|
|organization|jsonb|removed|
|region|jsonb|removed|
|shield|boolean|removed|
|state|text|removed|
|team|jsonb|removed|
|updated_at|timestamp without time zone|removed|

## heroku_teams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|membership_limit|real|updated|Type changed from float to real

## heroku_vpn_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ike_version|bigint|updated|Type changed from integer to bigint
