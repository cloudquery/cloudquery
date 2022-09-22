insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  '[Preview]: Container Registry should use a virtual network service endpoint',
  r.subscription_id,
  r.id,
  case
    when r.network_rule_set->>'defaultAction' IS DISTINCT FROM 'Deny'
      OR vrr->>'id' IS NULL
      OR anvns._cq_id IS NULL
    then 'fail' else 'pass'
  end
FROM azure_container_registries r LEFT JOIN JSONB_ARRAY_ELEMENTS(r.network_rule_set->'virtualNetworkRules') AS vrr ON TRUE
         LEFT JOIN azure_network_virtual_networks anvns ON anvns.id = vrr->>'id'
