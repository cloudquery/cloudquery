insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that "Auditing" is set to "On" (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when assdbap.state != 'Enabled'
      then 'fail'
      else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_db_blob_auditing_policies assdbap ON
        s.cq_id = assdbap.server_cq_id
