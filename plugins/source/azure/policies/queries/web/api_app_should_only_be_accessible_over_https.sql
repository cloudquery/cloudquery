insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'API App should only be accessible over HTTPS',
  subscription_id,
  id,
  case
    when kind LIKE '%api' AND https_only IS NOT TRUE
      then 'fail' else 'pass'
  end
FROM azure_web_apps
