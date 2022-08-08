insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Windows machines should meet requirements for ''User Rights Assignment''',
  subscription_id,
  nsg_id,
  case
    when source_address_prefix in ('0.0.0.0', '0.0.0.0/0', 'any', 'internet', '<nw>/0', '/0', '*')
      AND (single_port = '3389' OR 3389 BETWEEN range_start AND range_end)
      AND protocol = 'Tcp'
      AND "access" = 'Allow'
      AND direction = 'Inbound'
    then 'fail' else 'pass'
  end
FROM view_azure_nsg_rules