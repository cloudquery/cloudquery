insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Latest TLS version should be used in your API App',
  subscription_id,
  id
FROM azure_web_apps
WHERE
  kind LIKE '%api'
  AND (site_config ->> 'minTlsVersion' IS NULL
       OR site_config ->> 'minTlsVersion' != '1.2');
