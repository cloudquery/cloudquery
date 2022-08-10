insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Long-term geo-redundant backup should be enabled for Azure SQL Databases',
  -- TODO: replace this with the subscription_id of the resource
  id,
  id,
  case
    when backup_long_term_retention_policy -> 'properties' ->> 'weeklyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND backup_long_term_retention_policy -> 'properties' ->> 'monthlyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND backup_long_term_retention_policy -> 'properties' ->> 'yearlyRetention' IS NOT DISTINCT FROM 'PT0S'
    then 'fail' else 'pass'
  end
FROM azure_sql_databases
