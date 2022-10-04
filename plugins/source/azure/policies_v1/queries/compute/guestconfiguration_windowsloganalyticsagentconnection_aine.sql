WITH installed AS (
	SELECT
		DISTINCT _cq_id
    FROM azure_compute_virtual_machines, jsonb_array_elements(resources) AS res
	WHERE
		res->>'publisher' = 'Microsoft.EnterpriseCloud.Monitoring'
		AND res->>'type' IN ( 'MicrosoftMonitoringAgent', 'OmsAgentForLinux' )
		AND res->>'provisioningState' = 'Succeeded'
		AND res->'settings'->>'workspaceId' IS NOT NULL -- TODO check
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Audit Windows machines on which the Log Analytics agent is not connected as expected',
  azure_compute_virtual_machines.subscription_id,
  azure_compute_virtual_machines.vm_id AS id,
  case
    when azure_compute_virtual_machines.storage_profile -> 'osDisk' ->> 'osType' = 'Windows'
      AND installed._cq_id IS NULL
    then 'fail'
    else 'pass'
  end
FROM
	azure_compute_virtual_machines
	LEFT JOIN installed ON azure_compute_virtual_machines._cq_id = installed._cq_id
