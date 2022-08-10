insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Require encryption on Data Lake Store accounts',
	b.account_id,
	sub.id,
  case
    when encryption_state IS DISTINCT FROM 'Enabled'
    then 'fail' else 'pass'
  end
FROM
	azure_datalake_storage_accounts AS b,
	azure_subscription_subscriptions AS sub 
WHERE
	sub.subscription_id = b.subscription_id 
