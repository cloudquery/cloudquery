insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure any of the ASC Default policy setting is not set to "Disabled" (Automated)',
  subscription_id,
  id,
  case
    when value = 'Disabled'
    then 'fail' else 'pass'
  end
FROM view_azure_security_policy_parameters
