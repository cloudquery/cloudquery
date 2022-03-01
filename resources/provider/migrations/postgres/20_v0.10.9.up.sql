-- Resource: directconnect.gateways
TRUNCATE aws_directconnect_gateways CASCADE;

ALTER TABLE IF EXISTS aws_directconnect_gateways
    DROP CONSTRAINT aws_directconnect_gateways_pk;
ALTER TABLE IF EXISTS aws_directconnect_gateways
    ADD CONSTRAINT aws_directconnect_gateways_pk PRIMARY KEY (account_id,id);

ALTER TABLE IF EXISTS "aws_directconnect_gateways" DROP COLUMN IF EXISTS "region";
