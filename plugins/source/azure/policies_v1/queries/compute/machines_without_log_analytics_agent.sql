WITH secured_vms AS (SELECT _cq_id
                     FROM azure_compute_virtual_machines, jsonb_array_elements(resources) AS res
                     WHERE res->>'type' IN ('MicrosoftMonitoringAgent', 'OmsAgentForLinux')
                       AND res->>'publisher' = 'Microsoft.EnterpriseCloud.Monitoring'
                       AND res->>'provisioningState' = 'Succeeded'
                       AND res->'settings'->>'workspaceId' IS NOT NULL) -- TODO check
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Virtual machines should have the Log Analytics extension installed',
  vms.subscription_id,
  vms.id,
  case
    when s._cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machines vms
         LEFT JOIN secured_vms s ON vms._cq_id = s._cq_id