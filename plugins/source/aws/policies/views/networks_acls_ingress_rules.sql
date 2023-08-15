CREATE OR REPLACE VIEW view_aws_nacl_ingress_rules AS
SELECT
  account_id,
  region,
  arn,
  is_default,
  vpc_id,
  (entry -> 'PortRange' ->> 'From')::int AS port_range_from,
  (entry -> 'PortRange' ->> 'To')::int   AS port_range_to,
  entry ->> 'Protocol'                   AS protocol,
  entry ->> 'CidrBlock'                  AS cidr_block,
  entry ->> 'Ipv6CidrBlock'              AS ipv6_cidr_block,
  entry ->> 'RuleAction'                 AS rule_action
FROM aws_ec2_network_acls, jsonb_array_elements(entries) AS entry
WHERE NOT (entry ->> 'Egress')::bool
