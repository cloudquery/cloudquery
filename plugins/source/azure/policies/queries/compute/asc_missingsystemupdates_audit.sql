insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'System updates should be installed on your machines',
    azure_compute_virtual_machines.subscription_id,
	azure_compute_virtual_machines.vm_id AS vm_id,
  case
    when (os_profile->'windowsConfiguration'->>'enableAutomaticUpdates')::boolean is distinct from true then 'fail' else 'pass' -- TODO check
  end
FROM
	azure_compute_virtual_machines
