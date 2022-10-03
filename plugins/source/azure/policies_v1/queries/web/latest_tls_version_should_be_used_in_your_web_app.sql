insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Latest TLS version should be used in your Web App',
  subscription_id,
  id,
  case
    when kind LIKE 'app%'
      AND (site_config ->> 'minTlsVersion' IS NULL
      OR site_config ->> 'minTlsVersion' != '1.2')
    then 'fail' else 'pass'
  end
FROM azure_web_apps
