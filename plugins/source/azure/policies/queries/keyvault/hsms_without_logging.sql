WITH hsms_with_logging_enabled AS (SELECT DISTINCT h.cq_id
   FROM azure_keyvault_managed_hsm h
            LEFT JOIN azure_monitor_diagnostic_settings s ON h.id = s.resource_uri
            LEFT JOIN azure_monitor_diagnostic_setting_logs l
                      ON s.cq_id = l.diagnostic_setting_cq_id
   WHERE l.enabled = TRUE
     AND l.category = 'AuditEvent'
     AND (s.storage_account_id IS NOT NULL OR s.storage_account_id IS DISTINCT FROM '')
     AND retention_policy_enabled = TRUE)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Azure Key Vault Managed HSM should be enabled',  
  subscription_id,
  id,
  case
    when e.cq_id is null then 'fail' else 'pass'
  end
FROM azure_keyvault_managed_hsm a
     LEFT JOIN hsms_with_logging_enabled e ON a.cq_id = e.cq_id
