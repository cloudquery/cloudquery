insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure "Enforce SSL connection" is set to "ENABLED" for MySQL Database Server (Automated)',
  subscription_id,
  id AS server_id,
  case
    when ssl_enforcement != 'Enabled'
      OR ssl_enforcement IS NULL
    then 'fail' else 'pass'
  end
FROM azure_mysql_servers
