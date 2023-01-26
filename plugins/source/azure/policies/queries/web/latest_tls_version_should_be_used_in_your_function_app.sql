insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Latest TLS version should be used in your Function App',
  subscription_id,
  id,
  case
    when kind LIKE 'functionapp%' AND (properties -> 'siteConfig' -> 'minTlsVersion' IS NULL
       OR properties -> 'siteConfig' ->> 'minTlsVersion' is distinct from '1.2')
    then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps
