WITH sets_with_logs AS (
    SELECT compute_virtual_machine_id
    FROM azure_compute_virtual_machine_extensions
    WHERE publisher = 'Microsoft.EnterpriseCloud.Monitoring'
      AND type IN ('MicrosoftMonitoringAgent', 'OmsAgentForLinux')
      AND provisioning_state = 'Succeeded'
      -- AND settings ->> 'workspaceId' IS NOT NULL -- TODO FIXME missing?
      )
insert into azure_policy_results
SELECT 
  :'execution_time',
  :'framework',
  :'check_id',
  'The Log Analytics extension should be installed on Virtual Machine Scale Sets',
  s.subscription_id,
  id,
  case
    when ss.compute_virtual_machine_id IS NULL then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machine_scale_sets s
  LEFT JOIN sets_with_logs ss ON s.id = ss.compute_virtual_machine_id -- TODO check id match