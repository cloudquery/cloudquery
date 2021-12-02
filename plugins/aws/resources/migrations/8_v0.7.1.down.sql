ALTER TABLE IF EXISTS "aws_config_configuration_recorders" DROP COLUMN status_last_error_code,
                                                           DROP COLUMN status_last_error_message,
                                                           DROP COLUMN status_last_start_time,
                                                           DROP COLUMN status_last_status,
                                                           DROP COLUMN status_last_status_change_time,
                                                           DROP COLUMN status_last_stop_time,
                                                           DROP COLUMN status_recording;

ALTER TABLE IF EXISTS "aws_wafv2_web_acls" DROP COLUMN logging_configuration;
ALTER TABLE IF EXISTS "aws_waf_web_acls" DROP COLUMN logging_configuration;

ALTER TABLE IF EXISTS "aws_redshift_clusters" DROP COLUMN logging_status;

--ec2-instances
ALTER TABLE IF EXISTS "aws_ec2_instances" DROP COLUMN state_transition_reason_time;
ALTER TABLE IF EXISTS "aws_ec2_instances" DROP COLUMN boot_mode;
ALTER TABLE IF EXISTS "aws_ec2_instances" DROP COLUMN metadata_options_http_protocol_ipv6;
ALTER TABLE IF EXISTS "aws_ec2_instance_network_interfaces" DROP COLUMN ipv4_prefixes;
ALTER TABLE IF EXISTS "aws_ec2_instance_network_interfaces" DROP COLUMN ipv6_prefixes;