WITH vms_with_interfaces AS (SELECT subscription_id,
                                    id,
                                    jsonb_array_elements(network_profile->'networkInterfaces') AS nics
                             FROM azure_compute_virtual_machines vm),
nics_with_subnets AS (
    SELECT id, jsonb_array_elements(ip_configurations->'ipConfigurations') AS ip_config FROM azure_network_interfaces
    )
-- TODO check
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Virtual machines should be connected to an approved virtual network',
  v.subscription_id,
  v.id,
  case
    when i.ip_config->>'subnet_id' IS NULL then 'fail' else 'pass'
  end
FROM vms_with_interfaces v
         LEFT JOIN nics_with_subnets i ON v.nics->>'id' = i.id
