ALTER TABLE IF EXISTS aws_rds_certificates
    DROP CONSTRAINT aws_rds_certificates_pk;
ALTER TABLE IF EXISTS aws_rds_certificates
    ADD CONSTRAINT aws_rds_certificates_pk PRIMARY KEY (cq_fetch_date, account_id, arn);