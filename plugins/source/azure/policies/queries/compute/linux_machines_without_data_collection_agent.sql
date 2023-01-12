WITH secured_vms AS ( SELECT vm.id as compute_virtual_machine_id
                      FROM azure_compute_virtual_machines vm left join azure_compute_virtual_machine_extensions ex on vm._cq_id = ex._cq_parent_id
                     WHERE properties ->> 'publisher' = 'DependencyAgentLinux'
                       AND properties ->> 'type' = 'Microsoft.Azure.Monitoring.DependencyAgent'
                       AND properties ->> 'provisioningState' = 'Succeeded')
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  '[Preview]: Network traffic data collection agent should be installed on Linux virtual machines',
  vms.subscription_id,
  vms.id,
  case
    when vms.storage_profile -> 'osDisk' ->> 'osType' = 'Linux'
      and s.compute_virtual_machine_id IS NULL then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machines vms
         LEFT JOIN secured_vms s ON vms.id = s.compute_virtual_machine_id