insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'System updates should be installed on your machines',
  azure_subscription_subscriptions.subscription_id,
	azure_compute_virtual_machines.vm_id AS vm_id,
  case
    when windows_configuration_enable_automatic_updates is distinct from true then 'fail' else 'pass'
  end
FROM
	azure_compute_virtual_machines
	JOIN azure_subscription_subscriptions ON azure_subscription_subscriptions.subscription_id = azure_compute_virtual_machines.subscription_id
