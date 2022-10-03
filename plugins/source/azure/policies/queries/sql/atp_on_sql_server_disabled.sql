insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when p.state != 'Enabled'
    then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_databases d ON
        s.id = d.sql_server_id
    LEFT JOIN azure_sql_database_threat_detection_policies p ON
        d.id = p.sql_database_id
