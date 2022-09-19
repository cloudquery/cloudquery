WITH
    settings_with_logs AS (
        SELECT "ResourceURI" AS resource_uri, storage_account_id, JSON_ARRAY_ELEMENTS(logs) AS logs FROM azure_monitor_diagnostic_settings
    ),
    logging_enabled AS (
        SELECT DISTINCT j._cq_id
  FROM azure_streamanalytics_streaming_jobs j
    LEFT JOIN settings_with_logs s ON j.id = s.resource_uri
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
  'Resource logs in Azure Stream Analytics should be enabled',
  subscription_id,
  id,
  case
    when e._cq_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_streamanalytics_streaming_jobs j
         LEFT JOIN logging_enabled e ON j._cq_id = e._cq_id