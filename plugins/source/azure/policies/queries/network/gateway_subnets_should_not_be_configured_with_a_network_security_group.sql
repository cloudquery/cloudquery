insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Gateway subnets should not be configured with a network security group',
  -- todo change to subscription id
  id,
  id,
  case
    when name = 'GatewaySubnet' AND network_security_group_id IS NOT NULL
    then 'fail' else 'pass'
  end
FROM azure_network_virtual_network_subnets
