ALTER TABLE "aws_ec2_instance_network_interface_ipv6_addresses" RENAME COLUMN "instance_network_interface_id" TO instance_network_interface_cq_id;

ALTER TABLE "aws_ec2_instance_network_interface_private_ip_addresses" RENAME COLUMN "instance_network_interface_id" TO instance_network_interface_cq_id;


ALTER TABLE "aws_access_analyzer_analyzer_finding_sources" DROP CONSTRAINT "aws_access_analyzer_analyzer_finding_sources_pk";

ALTER TABLE "aws_access_analyzer_analyzer_finding_sources" DROP CONSTRAINT "aws_access_analyzer_analyzer_finding_sources_cq_id_key";

ALTER TABLE aws_access_analyzer_analyzer_finding_sources
    ADD CONSTRAINT "aws_access_analyzer_analyzer_finding_sources_pk"
        PRIMARY KEY(cq_id);

