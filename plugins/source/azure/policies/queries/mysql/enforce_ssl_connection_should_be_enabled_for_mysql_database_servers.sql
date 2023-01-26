insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Enforce SSL connection should be enabled for MySQL database servers',
  subscription_id,
  id,
  case
    when properties->>'sslEnforcement' IS DISTINCT FROM 'Enabled'
    then 'fail' else 'pass'
  end
FROM azure_mysql_servers
