ALTER TABLE IF EXISTS "aws_config_configuration_recorders" ADD COLUMN status_last_error_code text,
                                                           ADD COLUMN status_last_error_message text,
                                                           ADD COLUMN status_last_start_time timestamp without time zone,
                                                           ADD COLUMN status_last_status text,
                                                           ADD COLUMN status_last_status_change_time timestamp without time zone,
                                                           ADD COLUMN status_last_stop_time timestamp without time zone,
                                                           ADD COLUMN status_recording boolean;


ALTER TABLE IF EXISTS "aws_wafv2_web_acls" ADD COLUMN logging_configuration text[];
ALTER TABLE IF EXISTS "aws_waf_web_acls" ADD COLUMN logging_configuration text[];

ALTER TABLE IF EXISTS "aws_redshift_clusters" ADD COLUMN logging_status jsonb;

--ec2-instances
ALTER TABLE IF EXISTS "aws_ec2_instances" ADD COLUMN state_transition_reason_time timestamp;
ALTER TABLE IF EXISTS "aws_ec2_instances" ADD COLUMN boot_mode text;
ALTER TABLE IF EXISTS "aws_ec2_instances" ADD COLUMN metadata_options_http_protocol_ipv6 text;
ALTER TABLE IF EXISTS "aws_ec2_instance_network_interfaces" ADD COLUMN ipv4_prefixes _text;
ALTER TABLE IF EXISTS "aws_ec2_instance_network_interfaces" ADD COLUMN ipv6_prefixes _text;