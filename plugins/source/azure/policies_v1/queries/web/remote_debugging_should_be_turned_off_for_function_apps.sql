insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Remote debugging should be turned off for Function Apps',
  subscription_id,
  id,
  case
    when kind LIKE 'functionapp%'
      AND site_config ->> 'remoteDebuggingEnabled' = 'true'
    then 'fail' else 'pass'
  end
FROM azure_web_apps
