INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Function App should only be accessible over HTTPS',
  subscription_id,
  id,
  case
    when kind LIKE 'functionapp%' AND (properties ->> 'httpsOnly')::boolean IS NOT TRUE
    then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps
