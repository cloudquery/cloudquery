insert into azure_policy_results
SELECT 
  :'execution_time',
  :'framework',
  :'check_id',
  'Geo-redundant backup should be enabled for Azure Database for MariaDB',
  subscription_id,
  id,
  case
    when geo_redundant_backup IS DISTINCT FROM 'Enabled'
      then 'fail' else 'pass'
  end
FROM azure_mariadb_servers