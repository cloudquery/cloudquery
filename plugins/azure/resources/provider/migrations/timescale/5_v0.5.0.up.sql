-- Resource: redis.services
CREATE TABLE IF NOT EXISTS "azure_redis_services"
(
    "cq_id"                        uuid                        NOT NULL,
    "cq_meta"                      jsonb,
    "cq_fetch_date"                timestamp without time zone NOT NULL,
    "subscription_id"              text,
    "tags"                         jsonb,
    "location"                     text,
    "id"                           text,
    "name"                         text,
    "type"                         text,
    "provisioning_state"           text,
    "hostname"                     text,
    "port"                         integer,
    "ssl_port"                     integer,
    "linked_server_ids"            text[],
    "instances"                    jsonb,
    "private_endpoint_connections" jsonb,
    "sku_name"                     text,
    "sku_family"                   text,
    "sku_capacity"                 integer,
    "subnet_id"                    text,
    "static_ip"                    inet,
    "configuration"                jsonb,
    "version"                      text,
    "enable_non_ssl_port"          boolean,
    "replicas_per_master"          integer,
    "replicas_per_primary"         integer,
    "tenant_settings"              jsonb,
    "shard_count"                  integer,
    "minimum_tls_version"          text,
    "public_network_access"        boolean,
    CONSTRAINT azure_redis_services_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_redis_services');


-- Resource: sql.servers
CREATE TABLE IF NOT EXISTS "azure_sql_server_virtual_network_rules"
(
    "cq_id"                                uuid                        NOT NULL,
    "cq_meta"                              jsonb,
    "cq_fetch_date"                        timestamp without time zone NOT NULL,
    "server_cq_id"                         uuid,
    "id"                                   text,
    "name"                                 text,
    "type"                                 text,
    "subnet_id"                            text,
    "ignore_missing_vnet_service_endpoint" boolean,
    "state"                                text,
    CONSTRAINT azure_sql_server_virtual_network_rules_pk PRIMARY KEY (cq_fetch_date, server_cq_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_server_virtual_network_rules (cq_fetch_date, server_cq_id);
SELECT setup_tsdb_child('azure_sql_server_virtual_network_rules', 'server_cq_id', 'azure_sql_servers', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_sql_server_security_alert_policy"
(
    "cq_id"                      uuid                        NOT NULL,
    "cq_meta"                    jsonb,
    "cq_fetch_date"              timestamp without time zone NOT NULL,
    "server_cq_id"               uuid,
    "id"                         text,
    "name"                       text,
    "type"                       text,
    "state"                      text,
    "disabled_alerts"            text[],
    "email_addresses"            text[],
    "email_account_admins"       boolean,
    "storage_endpoint"           text,
    "storage_account_access_key" text,
    "retention_days"             integer,
    "creation_time"              timestamp without time zone,
    CONSTRAINT azure_sql_server_security_alert_policy_pk PRIMARY KEY (cq_fetch_date, server_cq_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_server_security_alert_policy (cq_fetch_date, server_cq_id);
SELECT setup_tsdb_child('azure_sql_server_security_alert_policy', 'server_cq_id', 'azure_sql_servers', 'cq_id');

