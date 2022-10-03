insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that "Auditing" Retention is "greater than 90 days" (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when assdbap.retention_days < 90
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_blob_auditing_policies assdbap ON
        s.id = assdbap.sql_server_id
