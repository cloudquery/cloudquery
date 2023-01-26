WITH sets_with_logs AS (
    SELECT vm.id as compute_virtual_machine_id
    FROM azure_compute_virtual_machines vm left join azure_compute_virtual_machine_extensions ex on vm._cq_id = ex._cq_parent_id
    WHERE (ex.properties ->> 'publisher' = 'Microsoft.Azure.Diagnostics' AND
           ex.properties ->> 'type' = 'IaaSDiagnostics')
       OR (ex.properties ->> 'publisher' IN ('Microsoft.OSTCExtensions', 'Microsoft.Azure.Diagnostics') AND
           ex.properties ->> 'type' = 'LinuxDiagnostic'))
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Virtual Machine Scale Sets should be enabled',
  subscription_id,
  id,
  case
    when ss.compute_virtual_machine_id IS NULL
    then 'fail' else 'pass'
  end
FROM azure_compute_virtual_machine_scale_sets s
         LEFT JOIN sets_with_logs ss ON s.id = ss.compute_virtual_machine_id -- TODO check id match

