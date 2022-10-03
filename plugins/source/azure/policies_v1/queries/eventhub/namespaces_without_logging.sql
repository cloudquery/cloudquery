WITH
    settings_with_logs AS (
        SELECT resource_uri, JSONB_ARRAY_ELEMENTS(logs) AS logs FROM azure_monitor_diagnostic_settings
    ),
    namespaces_with_logging_enabled AS (
        SELECT DISTINCT n._cq_id
  FROM azure_eventhub_namespaces n
    LEFT JOIN settings_with_logs s ON n.id = s.resource_uri
    WHERE (s.logs->>'enabled')::boolean IS TRUE
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Event Hub should be enabled',
  subscription_id,
  id,
  case when
    e._cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_eventhub_namespaces a
     LEFT JOIN namespaces_with_logging_enabled e ON a._cq_id = e._cq_id
