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
        s.cq_id = d.server_cq_id
    LEFT JOIN azure_sql_database_db_threat_detection_policies p ON
        d.cq_id = p.database_cq_id
