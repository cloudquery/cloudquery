insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Secure transfer to storage accounts should be enabled',
  subscription_id,
  id,
  case
    when supports_https_traffic_only IS NOT TRUE
      then 'fail' else 'pass'
  end
FROM azure_storage_accounts
