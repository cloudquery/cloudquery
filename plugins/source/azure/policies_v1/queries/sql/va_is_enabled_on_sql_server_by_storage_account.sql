insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that Vulnerability Assessment (VA) is enabled on a SQL server by setting a Storage Account (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when a.storage_container_path IS NULL OR a.storage_container_path = ''
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments a ON
        s.id = a.sql_server_id
