insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure any of the ASC Default policy setting is not set to "Disabled" (Automated)' as title,
  subscription_id,
  id,
  case
    when value = 'Disabled'
    then 'fail' else 'pass'
  end
FROM view_azure_security_policy_parameters
