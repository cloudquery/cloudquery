insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Storage accounts should restrict network access',
  az_sub.subscription_id,
  az_stor.id,
  case
    when az_stor.properties -> 'networkAcls' ->>'defaultAction' IS DISTINCT FROM 'Deny'
      then 'fail' else 'pass'
  end
FROM azure_storage_accounts as az_stor
LEFT JOIN azure_subscription_subscriptions as az_sub
ON az_sub.subscription_id = SUBSTRING(az_stor.id,16,36)