insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Storage accounts should restrict network access',
  subscription_id,
  id,
  case
    when network_acls->>'defaultAction' IS DISTINCT FROM 'Deny'
      then 'fail' else 'pass'
  end
FROM azure_storage_accounts