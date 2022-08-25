CREATE
OR REPLACE VIEW view_aws_nacl_allow_ingress_rules AS
SELECT aena.arn,
       aena.account_id,
       aenae.port_range_from,
       aenae.port_range_to,
       aenae.protocol,
       aenae.cidr_block,
       aenae.ipv6_cidr_block
FROM aws_ec2_network_acls aena
         LEFT JOIN aws_ec2_network_acl_entries aenae ON
    aena.cq_id = aenae.network_acl_cq_id
WHERE aenae.egress = FALSE
  AND aenae.rule_action = 'allow';