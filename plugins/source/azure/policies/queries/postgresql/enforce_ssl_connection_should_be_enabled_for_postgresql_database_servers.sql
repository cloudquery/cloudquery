INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Enforce SSL connection should be enabled for PostgreSQL database servers',
  subscription_id,
  id,
  case
    when properties->>'sslEnforcement' IS DISTINCT FROM 'Enabled'
    then 'fail' else 'pass'
  end
FROM azure_postgresql_servers
