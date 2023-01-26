insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Web Application should only be accessible over HTTPS',
  subscription_id,
  id,
  case
    when kind LIKE 'app%' AND (properties ->> 'httpsOnly')::boolean IS NOT TRUE
    then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps
