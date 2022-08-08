WITH vms_with_interfaces AS (SELECT subscription_id,
                                    id,
                                    JSONB_ARRAY_ELEMENTS(network_profile_network_interfaces) ->>
                                    'network_interface_id' AS network_interface_id
                             FROM azure_compute_virtual_machines vm)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Virtual machines should be connected to an approved virtual network',
  v.subscription_id,
  v.id,
  case
    when a.subnet_id IS NULL then 'fail' else 'pass'
  end
FROM vms_with_interfaces v
         LEFT JOIN azure_network_interfaces i ON v.network_interface_id = i.id
         LEFT JOIN azure_network_interface_ip_configurations a
                   ON i.cq_id = a.interface_cq_id