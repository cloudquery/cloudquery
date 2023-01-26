WITH logging_enabled AS (
  SELECT DISTINCT n.cq_id
  FROM azure_servicebus_namespaces n
           LEFT JOIN azure_monitor_diagnostic_settings s ON n.id = s.resource_id
           LEFT JOIN azure_monitor_diagnostic_setting_logs l
                     ON s.cq_id = l.diagnostic_setting_cq_id
  WHERE l.enabled = TRUE
    AND l.category = 'AuditEvent'
    AND (s.properties -> 'storageAccountId' as properties ->> 'storageAccountId' as storage_account_id IS NOT NULL OR s.storage_account_id IS DISTINCT FROM '')
    AND retention_policy_enabled = TRUE
)
SELECT 
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  subscription_id,
  id,
  case
    when e.cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_servicebus_namespaces n
         LEFT JOIN logging_enabled e ON n.cq_id = e.cq_id
