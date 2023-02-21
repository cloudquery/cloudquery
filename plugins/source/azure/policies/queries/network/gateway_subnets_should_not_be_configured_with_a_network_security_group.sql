WITH subs AS (
    SELECT subscription_id, jsonb_array_elements(properties->'subnets') AS subnet
    FROM azure_network_virtual_networks
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Gateway subnets should not be configured with a network security group',
  subscription_id,
  subnet->>'id',
  case
    when subnet->>'name' = 'GatewaySubnet' AND subnet->'networkSecurityGroup'->>'id' IS NOT NULL
    then 'fail' else 'pass'
  end
FROM subs

