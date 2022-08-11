WITH valid_namespaces AS (
  SELECT id
  FROM azure_eventhub_namespaces, jsonb_array_elements(network_rule_set -> 'properties' -> 'virtualNetworkRules') AS rule
  WHERE rule -> 'subnet' ->> 'id' IS NOT NULL
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Event Hub should use a virtual network service endpoint',
  n.subscription_id,
  n.id,
  case
    when v.id IS NULL then 'fail' else 'pass'
  end
FROM
  azure_eventhub_namespaces n
  LEFT OUTER JOIN valid_namespaces v
  ON n.id = v.id
