-- Resource: ec2.nat_gateways
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses DROP CONSTRAINT aws_ec2_nat_gateway_addresses_pk;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (nat_gateway_cq_id,network_interface_id);
ALTER TABLE aws_ec2_nat_gateway_addresses ALTER COLUMN allocation_id DROP NOT NULL;
