insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)' as title,
  subscription_id,
  id,
  case
    when auto_provision = 'On'
    then 'pass' else 'fail'
  end
FROM azure_security_auto_provisioning_settings asaps
WHERE "name" = 'default'
