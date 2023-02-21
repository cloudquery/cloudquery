WITH secured_accounts AS (SELECT a._cq_id
                          FROM azure_storage_accounts a, jsonb_array_elements(a.properties->'networkAcls'->'virtualNetworkRules') AS vnet
                          WHERE a.properties->'networkAcls'->>'defaultAction' = 'Deny'
                            AND vnet->>'id' IS NOT NULL
                            AND vnet->>'state' = 'succeeded')
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Storage Accounts should use a virtual network service endpoint',
  subscription_id,
  id,
  case
    when s._cq_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_storage_accounts a
  LEFT JOIN secured_accounts s ON a._cq_id = s._cq_id