WITH
    settings_with_logs AS (
        SELECT resource_id, properties ->> 'storageAccountId' as storage_account_id, JSONB_ARRAY_ELEMENTS(properties -> 'logs') AS logs FROM azure_monitor_diagnostic_settings
    ),
    logging_enabled AS (
        SELECT DISTINCT a._cq_id
  FROM azure_keyvault_keyvault_managed_hsms a
    LEFT JOIN settings_with_logs s ON a.id = s.resource_id
    WHERE (s.logs->>'enabled')::boolean IS TRUE
    AND s.logs->>'category' = 'AuditEvent'
    AND (s.storage_account_id IS NOT NULL OR s.storage_account_id IS DISTINCT FROM '')
    AND (s.logs->'retentionPolicy'->>'enabled')::boolean IS TRUE
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Azure Key Vault Managed HSM should be enabled',  
  subscription_id,
  id,
  case
    when e._cq_id is null then 'fail' else 'pass'
  end
FROM azure_keyvault_keyvault_managed_hsms a
     LEFT JOIN logging_enabled e ON a._cq_id = e._cq_id
