insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Remote debugging should be turned off for API Apps',
  subscription_id,
  id,
  case
    when kind LIKE '%api'
      AND site_config ->> 'remoteDebuggingEnabled' = 'true'
    then 'fail' else 'pass'
  end
FROM azure_web_apps
