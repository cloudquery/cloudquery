--vms created using old manager have type 'Microsoft.ClassicCompute/virtualMachines'
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Virtual machines should be migrated to new Azure Resource Manager resources',
  subscription_id,
  id,
  case
    when type IS DISTINCT FROM 'Microsoft.Compute/virtualMachines'
    then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machines