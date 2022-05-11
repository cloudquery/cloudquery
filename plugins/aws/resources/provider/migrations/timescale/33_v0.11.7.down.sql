-- Resource: rds.instances
TRUNCATE TABLE aws_rds_instances CASCADE;
ALTER TABLE IF EXISTS aws_rds_instances DROP CONSTRAINT aws_rds_instances_pk;
ALTER TABLE IF EXISTS aws_rds_instances ADD CONSTRAINT aws_rds_instances_pk PRIMARY KEY (cq_fetch_date,account_id,id);

-- Resource: ssm.instances
TRUNCATE TABLE aws_ssm_instance_compliance_items CASCADE;
ALTER TABLE IF EXISTS aws_ssm_instance_compliance_items DROP CONSTRAINT aws_ssm_instance_compliance_items_pk;
ALTER TABLE IF EXISTS aws_ssm_instance_compliance_items ADD CONSTRAINT aws_ssm_instance_compliance_items_pk PRIMARY KEY (cq_fetch_date,instance_cq_id,resource_id,id);
