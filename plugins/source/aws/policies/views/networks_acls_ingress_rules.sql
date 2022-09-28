CREATE OR REPLACE VIEW view_aws_nacl_allow_ingress_rules AS
WITH rules AS (SELECT aena.arn,
                      aena.account_id,
                      (jsonb_array_elements(entries) -> 'PortRange' ->> 'From')::int AS port_range_from,
                      (jsonb_array_elements(entries) -> 'PortRange' ->> 'To')::int   AS port_range_to,
                      jsonb_array_elements(entries) ->> 'Protocol'                   AS protocol,
                      jsonb_array_elements(entries) ->> 'CidrBlock'                  AS cidr_block,
                      jsonb_array_elements(entries) ->> 'Ipv6CidrBlock'              AS ipv6_cidr_block,
                      jsonb_array_elements(entries) ->> 'Egress'                     AS egress,
                      jsonb_array_elements(entries) ->> 'RuleAction'                 AS rule_action
               FROM aws_ec2_network_acls aena)
SELECT arn, account_id, port_range_from, port_range_to, protocol, cidr_block, ipv6_cidr_block
FROM rules
WHERE egress IS DISTINCT FROM 'true'
  AND rule_action = 'allow';
