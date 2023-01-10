insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Long-term geo-redundant backup should be enabled for Azure SQL Databases' as title,
  asd.subscription_id,
  rp.id,
  case
    when rp.id IS NULL OR (rp.weekly_retention IS NOT DISTINCT FROM 'PT0S'
      AND rp.monthly_retention IS NOT DISTINCT FROM 'PT0S'
      AND rp.yearly_retention IS NOT DISTINCT FROM 'PT0S')
    then 'fail' else 'pass'
  end
FROM azure_sql_databases asd
    LEFT JOIN azure_sql_backup_long_term_retention_policies rp ON asd.long_term_retention_backup_resource_id = rp.id AND asd.id = rp.sql_database_id
