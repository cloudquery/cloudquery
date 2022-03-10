-- Resource: ec2.eips

ALTER TABLE IF EXISTS aws_ec2_eips
    DROP CONSTRAINT aws_ec2_eips_pk;
ALTER TABLE IF EXISTS aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (cq_fetch_date,account_id,public_ip);
