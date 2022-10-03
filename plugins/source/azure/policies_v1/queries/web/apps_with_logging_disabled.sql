insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in App Services should be enabled',
  subscription_id,
  id,
  case
    when NOT (site_config -> 'httpLoggingEnabled')::text::bool
      OR NOT (site_config -> 'detailedErrorLoggingEnabled')::text::bool
      OR NOT (site_config -> 'requestTracingEnabled')::text::bool
    then 'fail' else 'pass'
  end
FROM azure_web_apps