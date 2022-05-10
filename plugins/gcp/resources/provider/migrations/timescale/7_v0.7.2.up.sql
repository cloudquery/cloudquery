-- Resource: cloudbilling.accounts
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_accounts" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"display_name" text,
	"master_billing_account" text,
	"name" text,
	"open" boolean,
	"project_billing_enabled" boolean,
	"project_name" text,
	"project_id" text,
	CONSTRAINT gcp_cloudbilling_accounts_pk PRIMARY KEY(cq_fetch_date,project_id,"name"),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('gcp_cloudbilling_accounts');

-- Resource: cloudbilling.services
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_services" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"business_entity_name" text,
	"display_name" text,
	"name" text,
	"service_id" text,
	CONSTRAINT gcp_cloudbilling_services_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('gcp_cloudbilling_services');
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_skus" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"service_cq_id" uuid,
	"resource_family" text,
	"resource_group" text,
	"service_display_name" text,
	"usage_type" text,
	"description" text,
	"geo_taxonomy_regions" text[],
	"geo_taxonomy_type" text,
	"name" text,
	"service_provider_name" text,
	"service_regions" text[],
	"sku_id" text,
	CONSTRAINT gcp_cloudbilling_service_skus_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON gcp_cloudbilling_service_skus (cq_fetch_date, service_cq_id);
SELECT setup_tsdb_child('gcp_cloudbilling_service_skus', 'service_cq_id', 'gcp_cloudbilling_services', 'cq_id');
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_sku_pricing_info" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"service_sku_cq_id" uuid,
	"aggregation_count" bigint,
	"aggregation_interval" text,
	"aggregation_level" text,
	"currency_conversion_rate" float,
	"effective_time" text,
	"base_unit" text,
	"base_unit_conversion_factor" float,
	"base_unit_description" text,
	"display_quantity" float,
	"usage_unit" text,
	"usage_unit_description" text,
	"summary" text,
	CONSTRAINT gcp_cloudbilling_service_sku_pricing_info_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON gcp_cloudbilling_service_sku_pricing_info (cq_fetch_date, service_sku_cq_id);
SELECT setup_tsdb_child('gcp_cloudbilling_service_sku_pricing_info', 'service_sku_cq_id', 'gcp_cloudbilling_service_skus', 'cq_id');
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_sku_pricing_info_tiered_rates" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"service_sku_pricing_info_cq_id" uuid,
	"start_usage_amount" float,
	"unit_price_currency_code" text,
	"unit_price_nanos" bigint,
	"unit_price_units" bigint,
	CONSTRAINT gcp_cloudbilling_service_sku_pricing_info_tiered_rates_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON gcp_cloudbilling_service_sku_pricing_info_tiered_rates (cq_fetch_date, service_sku_pricing_info_cq_id);
SELECT setup_tsdb_child('gcp_cloudbilling_service_sku_pricing_info_tiered_rates', 'service_sku_pricing_info_cq_id', 'gcp_cloudbilling_service_sku_pricing_info', 'cq_id');

-- Resource: compute.instance_groups
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_groups"
(
    "cq_id"              uuid                        NOT NULL,
    "cq_meta"            jsonb,
    "cq_fetch_date"      timestamp without time zone NOT NULL,
    "project_id"         text,
    "creation_timestamp" timestamp,
    "description"        text,
    "fingerprint"        text,
    "id"                 bigint,
    "kind"               text,
    "name"               text,
    "named_ports"        jsonb,
    "network"            text,
    "region"             text,
    "self_link"          text,
    "size"               bigint,
    "subnetwork"         text,
    "zone"               text,
    CONSTRAINT gcp_compute_instance_groups_pk PRIMARY KEY (cq_fetch_date, project_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('gcp_compute_instance_groups');
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_group_instances"
(
    "cq_id"                uuid                        NOT NULL,
    "cq_meta"              jsonb,
    "cq_fetch_date"        timestamp without time zone NOT NULL,
    "instance_group_cq_id" uuid,
    "instance"             text,
    "named_ports"          jsonb,
    "status"               text,
    CONSTRAINT gcp_compute_instance_group_instances_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON gcp_compute_instance_group_instances (cq_fetch_date, instance_group_cq_id);
SELECT setup_tsdb_child('gcp_compute_instance_group_instances', 'instance_group_cq_id', 'gcp_compute_instance_groups',
                        'cq_id');

