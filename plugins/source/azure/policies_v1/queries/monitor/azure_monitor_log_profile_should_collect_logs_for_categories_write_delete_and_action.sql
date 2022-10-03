insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Azure Monitor log profile should collect logs for categories ''write,'' ''delete,'' and ''action''',
  s.id,
  s.id
FROM
  azure_subscriptions s
  LEFT OUTER JOIN azure_monitor_log_profiles p
  ON s.id = '/subscriptions/' || p.subscription_id
WHERE
  categories IS NULL
  OR NOT categories @> '{"Write", "Action", "Delete"}';
