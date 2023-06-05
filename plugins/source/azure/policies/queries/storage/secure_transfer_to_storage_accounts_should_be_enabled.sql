insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Secure transfer to storage accounts should be enabled',
  subscription_id,
  id,
  case
    when properties ->> 'supportsHttpsTrafficOnly' IS DISTINCT FROM 'true'
      then 'fail' else 'pass'
  end
FROM azure_storage_accounts
