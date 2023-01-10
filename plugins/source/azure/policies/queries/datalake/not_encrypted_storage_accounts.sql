insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Require encryption on Data Lake Store accounts',
    subscription_id,
	id,
  case
    when encryption_state IS DISTINCT FROM 'Enabled'
    then 'fail' else 'pass'
  end
FROM
	azure_datalake_store_accounts
