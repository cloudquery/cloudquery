insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)' as title,
  s.subscription_id,
  asd.id AS database_id,
  case
    when tde.properties->>'state' is distinct from 'Enabled'
    then 'fail'
    else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_databases asd ON
        s._cq_id = asd._cq_parent_id
    LEFT JOIN azure_sql_transparent_data_encryptions tde ON
        asd._cq_id = tde._cq_parent_id
