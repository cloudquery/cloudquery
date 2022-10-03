insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'CORS should not allow every resource to access your Web Applications',
  subscription_id,
  id,
  case
    when site_config -> 'cors' -> 'allowedOrigins' @> '["*"]'
      AND kind LIKE 'app%'
    then 'fail' else 'pass'
  end
FROM azure_web_apps
