WITH secured_vms AS (SELECT virtual_machine_cq_id
                     FROM azure_compute_virtual_machine_resources
                     WHERE extension_type IN ('MicrosoftMonitoringAgent', 'OmsAgentForLinux')
                       AND publisher = 'Microsoft.EnterpriseCloud.Monitoring'
                       AND provisioning_state = 'Succeeded'
                       AND settings ->> 'workspaceId' IS NOT NULL)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Virtual machines should have the Log Analytics extension installed',
  vms.subscription_id,
  vms.id,
  case
    when s.virtual_machine_cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machines vms
         LEFT JOIN secured_vms s ON vms.cq_id = s.virtual_machine_cq_id