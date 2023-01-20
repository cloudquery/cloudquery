insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that "Auditing" is set to "On" (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when assdbap.properties->>'state' != 'Enabled'
      then 'fail'
      else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_database_blob_auditing_policies assdbap ON
        s._cq_id = assdbap._cq_parent_id
