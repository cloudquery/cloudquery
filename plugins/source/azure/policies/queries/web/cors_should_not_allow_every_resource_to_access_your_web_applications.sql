insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'CORS should not allow every resource to access your Web Applications',
  subscription_id,
  id,
  case
    when array(select jsonb_array_elements_text(properties -> 'siteConfig' -> 'cors' -> 'allowedOrigins')) && ARRAY['*']
      AND kind LIKE 'app%'
    then 'fail' else 'pass'
  end
FROM azure_appservice_web_apps
