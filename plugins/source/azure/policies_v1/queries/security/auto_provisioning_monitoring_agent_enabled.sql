insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)',
  subscription_id,
  id,
  case
    when auto_provision = 'On'
    then 'pass' else 'fail'
  end
FROM azure_security_auto_provisioning_settings asaps
WHERE "name" = 'default'
