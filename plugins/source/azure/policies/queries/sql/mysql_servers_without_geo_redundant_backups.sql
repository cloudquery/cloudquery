
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Geo-redundant backup should be enabled for Azure Database for MySQL' as title,
  subscription_id,
  id,
  case
    when storage_profile->>'geoRedundantBackup' IS DISTINCT FROM 'Enabled'
      then 'fail' else 'pass'
  end
FROM azure_mysql_servers