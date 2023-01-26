insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Require encryption on Data Lake Store accounts',
    subscription_id,
	id,
  case
    when properties ->> 'encryptionState' IS DISTINCT FROM 'Enabled'
    then 'fail' else 'pass'
  end
FROM
	azure_datalakestore_accounts
