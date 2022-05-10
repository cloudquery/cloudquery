-- Resource: cloudbilling.accounts
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_accounts" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"display_name" text,
	"master_billing_account" text,
	"name" text,
	"open" boolean,
	"project_billing_enabled" boolean,
	"project_name" text,
	"project_id" text,
	CONSTRAINT gcp_cloudbilling_accounts_pk PRIMARY KEY(project_id, "name"),
	UNIQUE(cq_id)
);

-- Resource: cloudbilling.services
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_services" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"business_entity_name" text,
	"display_name" text,
	"name" text,
	"service_id" text,
	CONSTRAINT gcp_cloudbilling_services_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_skus" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
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
	CONSTRAINT gcp_cloudbilling_service_skus_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_cloudbilling_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_sku_pricing_info" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
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
	CONSTRAINT gcp_cloudbilling_service_sku_pricing_info_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_sku_cq_id) REFERENCES gcp_cloudbilling_service_skus(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_cloudbilling_service_sku_pricing_info_tiered_rates" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_sku_pricing_info_cq_id" uuid,
	"start_usage_amount" float,
	"unit_price_currency_code" text,
	"unit_price_nanos" bigint,
	"unit_price_units" bigint,
	CONSTRAINT gcp_cloudbilling_service_sku_pricing_info_tiered_rates_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_sku_pricing_info_cq_id) REFERENCES gcp_cloudbilling_service_sku_pricing_info(cq_id) ON DELETE CASCADE
);
-- Resource: compute.instance_groups
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_groups"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
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
    CONSTRAINT gcp_compute_instance_groups_pk PRIMARY KEY (project_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_group_instances"
(
    "cq_id"                uuid NOT NULL,
    "cq_meta"              jsonb,
    "instance_group_cq_id" uuid,
    "instance"             text,
    "named_ports"          jsonb,
    "status"               text,
    CONSTRAINT gcp_compute_instance_group_instances_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (instance_group_cq_id) REFERENCES gcp_compute_instance_groups (cq_id) ON DELETE CASCADE
);
