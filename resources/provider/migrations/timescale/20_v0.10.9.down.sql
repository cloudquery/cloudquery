-- Resource: directconnect.gateways
ALTER TABLE IF EXISTS "aws_directconnect_gateways" ADD COLUMN IF NOT EXISTS "region" text;

ALTER TABLE IF EXISTS aws_directconnect_gateways
DROP CONSTRAINT aws_directconnect_gateways_pk;
ALTER TABLE IF EXISTS aws_directconnect_gateways
    ADD CONSTRAINT aws_directconnect_gateways_pk PRIMARY KEY (cq_fetch_date,account_id,region,id);
