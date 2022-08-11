insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  '[Preview]: Container Registry should use a virtual network service endpoint',
  r.subscription_id,
  r.id,
  case
    when r.network_rule_set_default_action IS DISTINCT FROM 'Deny'
      OR acrnrsvnr.virtual_network_id IS NULL
      OR anvns.cq_id IS NULL
    then 'fail' else 'pass'
  end
FROM azure_container_registries r
         LEFT JOIN azure_container_registry_network_rule_set_virtual_network_rules acrnrsvnr
                   ON r.cq_id = acrnrsvnr.registry_cq_id
         LEFT JOIN azure_network_virtual_network_subnets anvns ON acrnrsvnr.virtual_network_id = anvns.id
