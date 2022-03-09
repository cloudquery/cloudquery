-- Resource: web.apps
ALTER TABLE IF EXISTS "azure_web_apps" ADD COLUMN IF NOT EXISTS "vnet_connection" jsonb;
