insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)',
  s.subscription_id,
  asd.id AS database_id,
  case
    when (asd.transparent_data_encryption -> 'properties' ->> 'status') is distinct from 'Enabled'
    then 'fail'
    else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_databases asd ON
        s.cq_id = asd.server_cq_id
