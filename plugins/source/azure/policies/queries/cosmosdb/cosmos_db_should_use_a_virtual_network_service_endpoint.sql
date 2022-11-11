WITH valid_accounts AS (
  SELECT id
  FROM azure_cosmosdb_accounts, jsonb_array_elements(virtual_network_rules) AS rule
  WHERE rule ->> 'id' IS NOT NULL
) -- TODO check
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Cosmos DB should use a virtual network service endpoint',
  a.subscription_id,
  a.id,
  case
    when v.id IS NULL then 'fail' else 'pass'
  end
FROM
  azure_cosmosdb_accounts a
  LEFT OUTER JOIN valid_accounts v
  ON a.id = v.id

