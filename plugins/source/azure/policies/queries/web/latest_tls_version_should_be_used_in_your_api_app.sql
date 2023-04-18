INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Latest TLS version should be used in your API App',
  subscription_id,
  id,
  'UNKNOWN'
FROM azure_appservice_web_apps
WHERE
  kind LIKE '%api'
  AND (properties -> 'siteConfig' -> 'minTlsVersion' IS NULL
       OR properties -> 'siteConfig' ->> 'minTlsVersion' is distinct from '1.2');
