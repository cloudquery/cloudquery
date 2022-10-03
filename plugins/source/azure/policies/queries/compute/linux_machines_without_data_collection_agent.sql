WITH secured_vms AS (SELECT virtual_machine_cq_id
                     FROM azure_compute_virtual_machine_resources
                     WHERE extension_type = 'DependencyAgentLinux'
                       AND publisher = 'Microsoft.Azure.Monitoring.DependencyAgent'
                       AND provisioning_state = 'Succeeded')
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
      and s.virtual_machine_cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machines vms
         LEFT JOIN secured_vms s ON vms.cq_id = s.virtual_machine_cq_id