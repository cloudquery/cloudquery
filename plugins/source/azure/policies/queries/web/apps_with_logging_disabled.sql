insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in App Services should be enabled',
  subscription_id,
  id,
  case
    when NOT (properties -> 'siteConfig' -> 'httpLoggingEnabled')::text::bool
      OR NOT (properties -> 'siteConfig' -> 'detailedErrorLoggingEnabled')::text::bool
      OR NOT (properties -> 'siteConfig' -> 'requestTracingEnabled')::text::bool
    then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps