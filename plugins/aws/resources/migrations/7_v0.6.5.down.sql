ALTER TABLE IF EXISTS "aws_iam_password_policies" DROP COLUMN policy_exists;

ALTER TABLE IF EXISTS "aws_directconnect_gateways" RENAME COLUMN "state" TO "direct_connect_gateway_state";

ALTER TABLE IF EXISTS "aws_directconnect_gateways" RENAME COLUMN "name" TO "direct_connect_gateway_name";

ALTER TABLE IF EXISTS "aws_directconnect_gateways" ADD COLUMN "direct_connect_gateway_id";

UPDATE "aws_directconnect_gateways" SET "direct_connect_gateway_id" = "id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "gateway_cq_id" TO "directconnect_gateway_cq_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "gateway_id" TO "directconnect_gateway_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_attachments" DROP COLUMN "gateway_id";

ALTER TABLE IF EXISTS "aws_elbv2_listeners" DROP COLUMN "load_balancer_cq_id";
