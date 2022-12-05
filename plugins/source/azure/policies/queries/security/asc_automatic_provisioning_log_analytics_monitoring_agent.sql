insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Auto provisioning of the Log Analytics agent should be enabled on your subscription' as title,
  azure_subscriptions.id AS subscription_id,
	azure_security_auto_provisioning_settings._cq_id,
  case
    when auto_provision IS DISTINCT FROM 'AutoProvisionOn'
    then 'fail' else 'pass'
  end
FROM
	azure_security_auto_provisioning_settings
	RIGHT JOIN azure_subscriptions ON azure_security_auto_provisioning_settings.subscription_id = azure_subscriptions.id
