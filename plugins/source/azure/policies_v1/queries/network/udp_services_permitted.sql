-- noinspection SqlNoDataSourceInspection
insert into azure_policy_results
SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  subscription_id,
  id,
  case
    when source_address_prefix in ('0.0.0.0', '0.0.0.0/0', 'any', 'internet', '<nw>/0', '/0', '*')
      AND protocol = 'Udp'
      AND "access" = 'Allow'
      AND direction = 'Inbound'
      AND ((single_port = '53'
        OR 53 BETWEEN range_start AND range_end)
        OR (single_port = '123'
            OR 123 BETWEEN range_start AND range_end)
        OR (single_port = '161'
            OR 161 BETWEEN range_start AND range_end)
        OR (single_port = '389'
            OR 389 BETWEEN range_start AND range_end))
      then 'fail'
  end
FROM azure_nsg_rules
