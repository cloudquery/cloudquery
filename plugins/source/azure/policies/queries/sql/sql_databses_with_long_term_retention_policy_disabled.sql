insert into azure_policy_result
SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  s.subscription_id,
  d.id,
  case
    when long_term_retention_policy IS NULL
      OR (long_term_retention_policy ->> 'monthlyRetention' = 'PT0S'
      AND long_term_retention_policy ->> 'weeklyRetention' = 'PT0S'
      AND long_term_retention_policy ->> 'yearlyRetention' = 'PT0S')
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
         LEFT JOIN azure_sql_databases d ON s.id = d.sql_server_id