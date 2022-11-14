WITH valid_namespaces AS (
  SELECT n.id
  FROM azure_eventhub_namespaces n
    LEFT JOIN azure_eventhub_network_rule_sets r ON r.eventhub_namespace_id = n.id
  WHERE r.id IS NOT NULL
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
