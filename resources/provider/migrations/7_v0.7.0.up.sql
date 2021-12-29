ALTER TABLE IF EXISTS "aws_iam_password_policies" ADD COLUMN policy_exists boolean;

ALTER TABLE IF EXISTS "aws_directconnect_gateways" RENAME COLUMN "direct_connect_gateway_state" TO "state";

ALTER TABLE IF EXISTS "aws_directconnect_gateways" RENAME COLUMN "direct_connect_gateway_name" TO "name";

ALTER TABLE IF EXISTS "aws_directconnect_gateways" DROP COLUMN "directconnect_gateway_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "directconnect_gateway_cq_id" TO "gateway_cq_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "directconnect_gateway_id" TO "gateway_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_attachments" ADD COLUMN "gateway_id" text;

ALTER TABLE IF EXISTS "aws_elbv2_listeners" ADD COLUMN "load_balancer_cq_id" uuid;

ALTER TABLE IF EXISTS "aws_sns_topics" ADD COLUMN kms_master_key_id text;
