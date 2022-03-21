-- Resource: ec2.instance_statuses
DROP TABLE IF EXISTS aws_ec2_instance_status_events;
DROP TABLE IF EXISTS aws_ec2_instance_statuses;
-- Resource: ec2.nat_gateways
ALTER TABLE aws_ec2_nat_gateway_addresses
    ALTER COLUMN allocation_id SET NOT NULL;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses
    DROP CONSTRAINT aws_ec2_nat_gateway_addresses_pk;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses
    ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (cq_fetch_date, nat_gateway_cq_id, allocation_id, network_interface_id);