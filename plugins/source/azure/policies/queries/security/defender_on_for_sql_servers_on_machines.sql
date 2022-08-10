insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that Azure Defender is set to On for SQL servers on machines (Automatic)',
  subscription_id,
  id,
  case
    when pricing_properties_tier = 'Standard'
    then 'pass' else 'fail'
  end
FROM azure_security_pricings asp
WHERE "name" = 'SqlserverVirtualMachines'
