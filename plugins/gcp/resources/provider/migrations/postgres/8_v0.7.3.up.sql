-- Resource: serviceusage.services
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_services" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"project_id" text,
	"name" text,
	"authentication" jsonb,
	"documentation" jsonb,
	"title" text,
	"usage_producer_notification_channel" text,
	"usage_requirements" text[],
	"parent" text,
	"state" text,
	CONSTRAINT gcp_serviceusage_services_pk PRIMARY KEY(name),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_apis" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"methods" jsonb,
	"mixins" jsonb,
	"name" text,
	"options" jsonb,
	"source_context_file_name" text,
	"syntax" text,
	"version" text,
	CONSTRAINT gcp_serviceusage_service_apis_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_endpoints" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"allow_cors" boolean,
	"name" text,
	"target" text,
	CONSTRAINT gcp_serviceusage_service_endpoints_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_monitored_resources" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"description" text,
	"display_name" text,
	"labels" jsonb,
	"launch_stage" text,
	"name" text,
	"type" text,
	CONSTRAINT gcp_serviceusage_service_monitored_resources_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_monitoring_consumer_destinations" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"metrics" text[],
	"monitored_resource" text,
	CONSTRAINT gcp_serviceusage_service_monitoring_consumer_destinations_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_monitoring_producer_destinations" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"metrics" text[],
	"monitored_resource" text,
	CONSTRAINT gcp_serviceusage_service_monitoring_producer_destinations_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_quota_limits" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"default_limit" int,
	"description" text,
	"display_name" text,
	"duration" text,
	"free_tier" bigint,
	"max_limit" bigint,
	"metric" text,
	"name" text,
	"unit" text,
	"values" jsonb,
	CONSTRAINT gcp_serviceusage_service_quota_limits_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_quota_metric_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"metric_costs" jsonb,
	"selector" text,
	CONSTRAINT gcp_serviceusage_service_quota_metric_rules_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "gcp_serviceusage_service_usage_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"service_cq_id" uuid,
	"allow_unregistered_calls" boolean,
	"selector" text,
	"skip_service_control" boolean,
	CONSTRAINT gcp_serviceusage_service_usage_rules_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (service_cq_id) REFERENCES gcp_serviceusage_services(cq_id) ON DELETE CASCADE
);