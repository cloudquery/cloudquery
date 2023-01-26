WITH asr_protect AS (
    SELECT properties ->> 'sourceId' as source_id
    FROM azure_resources_links
    WHERE name LIKE 'ASR-Protect-%'
)
insert into azure_policy_results
SELECT 
  :'execution_time',
  :'framework',
  :'check_id',
  'Audit virtual machines without disaster recovery configured.',
  subscription_id,
  id,
  case
    when p.source_id is null then 'fail' else 'pass'
  end
FROM
    azure_compute_virtual_machines vm
    LEFT OUTER JOIN asr_protect p
    ON LOWER(vm.id) = LOWER(p.source_id)
