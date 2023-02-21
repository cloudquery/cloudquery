insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'API App should only be accessible over HTTPS',
  subscription_id,
  id,
  case
    when kind LIKE '%api' AND (properties ->> 'httpsOnly')::boolean IS NOT TRUE
      then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps
