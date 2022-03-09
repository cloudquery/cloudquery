-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" DROP COLUMN IF EXISTS "network_rule_set";

-- Resource: security.assessments
DROP TABLE IF EXISTS azure_security_assessments;

-- Resource: account.locations
DROP TABLE IF EXISTS azure_account_locations;

-- Resource: streamanalytics.jobs
DROP TABLE IF EXISTS azure_streamanalytics_jobs;

-- Resource: logic.app_workflows
DROP TABLE IF EXISTS azure_logic_app_workflows;

-- Resource: web.apps
ALTER TABLE IF EXISTS "azure_web_apps" DROP COLUMN IF EXISTS "vnet_connection";

-- Resource: search.services
DROP TABLE IF EXISTS azure_search_service_private_endpoint_connections;
DROP TABLE IF EXISTS azure_search_service_shared_private_link_resources;
DROP TABLE IF EXISTS azure_search_services;

-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases" DROP COLUMN IF EXISTS "backup_long_term_retention_policy";

-- Resource: batch.accounts
DROP TABLE IF EXISTS azure_batch_account_private_endpoint_connections;
DROP TABLE IF EXISTS azure_batch_accounts;
