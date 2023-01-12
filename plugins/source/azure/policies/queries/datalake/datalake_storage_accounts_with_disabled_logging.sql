WITH
settings_with_logs AS (
        SELECT resource_id, JSONB_ARRAY_ELEMENTS(diagnostic_settings_resource -> 'properties' -> 'logs') AS logs FROM azure_monitor_diagnostic_settings
),
accounts_with_logging_enabled AS (SELECT DISTINCT d._cq_id
    FROM azure_datalakestore_accounts d
        LEFT JOIN settings_with_logs s ON d.id = s.resource_id
    WHERE (s.logs->>'enabled')::boolean IS TRUE
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Azure Data Lake Store should be enabled',
  subscription_id,
  id,
  case
    when e._cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_datalakestore_accounts a
    LEFT JOIN accounts_with_logging_enabled e ON a._cq_id = e._cq_id
