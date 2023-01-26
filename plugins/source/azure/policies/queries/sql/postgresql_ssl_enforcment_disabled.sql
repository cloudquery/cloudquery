insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure "Enforce SSL connection" is set to "ENABLED" for PostgreSQL Database Server (Automated)' as title,
  subscription_id,
  id AS server_id,
  case
    when properties->>'sslEnforcement'  is distinct from 'Enabled'
      then 'fail' else 'pass'
  end
FROM azure_postgresql_servers
