insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Internet-facing virtual machines should be protected with network security groups',
  vm.subscription_id,
  vm.id,
  case
    when a.name IS NULL
      OR (
        a.status->>'code' IS DISTINCT FROM 'NotApplicable'
        AND a.status->>'code' IS DISTINCT FROM 'Healthy'
      )
    then 'fail'
    else 'pass'
  end
FROM
  azure_compute_virtual_machines vm
  LEFT OUTER JOIN azure_security_assessments a
  ON a.name = '483f12ed-ae23-447e-a2de-a67a10db4353' AND a.id like (vm.id || '/' || '%')
  
