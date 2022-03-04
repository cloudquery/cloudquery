-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" ADD COLUMN IF NOT EXISTS "network_rule_set" jsonb;
