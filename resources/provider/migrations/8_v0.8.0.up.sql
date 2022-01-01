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


--aws_cloudfront_distributions
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN "tags" json;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN "alias_icp_recordals" json;
DROP TABLE IF EXISTS aws_cloudfront_distribution_alias_icp_recordals;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN caller_reference text;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN logging_bucket text;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN logging_enabled bool;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN logging_include_cookies bool;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN logging_prefix text;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN in_progress_invalidation_batches int;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN active_trusted_key_groups_enabled bool;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN active_trusted_key_groups json;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN active_trusted_signers_enabled bool;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN active_trusted_signers json;
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    ADD COLUMN default_root_object text;

ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "aliases_items" TO "aliases";

ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_target_origin_id" TO "cache_behavior_target_origin_id";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_viewer_protocol_policy" TO "cache_behavior_viewer_protocol_policy";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_allowed_methods" TO "cache_behavior_allowed_methods";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_allowed_methods_cached_methods" TO "cache_behavior_allowed_methods_cached_methods";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_cache_policy_id" TO "cache_behavior_cache_policy_id";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_compress" TO "cache_behavior_compress";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_default_ttl" TO "cache_behavior_default_ttl";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_field_level_encryption_id" TO "cache_behavior_field_level_encryption_id";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_forwarded_values_cookies_forward" TO "cache_behavior_forwarded_values_cookies_forward";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_forwarded_values_cookies_white_listed_names" TO "cache_behavior_forwarded_values_cookies_whitelisted_names";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_forwarded_values_query_string" TO "cache_behavior_forwarded_values_query_string";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_forwarded_values_headers" TO "cache_behavior_forwarded_values_headers";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_forwarded_values_query_string_cache_keys" TO "cache_behavior_forwarded_values_query_string_cache_keys";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_max_ttl" TO "cache_behavior_max_ttl";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_min_ttl" TO "cache_behavior_min_ttl";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_origin_request_policy_id" TO "cache_behavior_origin_request_policy_id";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_realtime_log_config_arn" TO "cache_behavior_realtime_log_config_arn";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_smooth_streaming" TO "cache_behavior_smooth_streaming";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_trusted_key_groups_enabled" TO "cache_behavior_trusted_key_groups_enabled";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_trusted_key_groups" TO "cache_behavior_trusted_key_groups";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_trusted_signers_enabled" TO "cache_behavior_trusted_signers_enabled";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "cache_behaviour_trusted_signers" TO "cache_behavior_trusted_signers";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "restrictions_geo_restriction_restriction_items" TO "geo_restrictions";
ALTER TABLE IF EXISTS "aws_cloudfront_distributions"
    RENAME COLUMN "restrictions_geo_restriction_restriction_type" TO "geo_restriction_type";


--aws_cloudfront_distribution_cache_behaviours
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "cache_policy_id" TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "compress" BOOL;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "default_ttl" INT8;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "field_level_encryption_id" TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "forwarded_values_cookies_forward" TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "forwarded_values_cookies_whitelisted_names" _TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "forwarded_values_query_string" BOOL;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "forwarded_values_headers" _TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "forwarded_values_query_string_cache_keys" _TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "max_ttl" INT8;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "min_ttl" INT8;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "origin_request_policy_id" TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "realtime_log_config_arn" TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "smooth_streaming" BOOL;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "trusted_key_groups_enabled" BOOL;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "trusted_key_groups" _TEXT;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "trusted_signers_enabled" BOOL;
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_cache_behaviours"
    ADD COLUMN "trusted_signers" _TEXT;


--aws_cloudfront_distribution_default_cache_behavior_lambda_functions
ALTER TABLE IF EXISTS "aws_cache_behaviour_lambda_function_associations"
    RENAME TO "aws_cloudfront_distribution_default_behaviour_lambda_functions";

--aws_cloudfront_distribution_origin_groups
ALTER TABLE IF EXISTS "aws_cloudfront_distribution_origin_groups"
    RENAME COLUMN "failover_criteria_status_codes_items" TO "failover_criteria_status_codes";

DROP TABLE IF EXISTS public.aws_emr_clusters;

CREATE TABLE public.aws_emr_clusters (
                                         cq_id uuid,
                                         meta jsonb,
                                         account_id text,
                                         region text,
                                         applications jsonb,
                                         auto_scaling_role text,
                                         auto_terminate boolean,
                                         arn text NOT NULL,
                                         configurations jsonb,
                                         custom_ami_id text,
                                         ebs_root_volume_size integer,
                                         ec2_instance_attribute_additional_master_security_groups text[],
                                         ec2_instance_attribute_additional_slave_security_groups text[],
                                         ec2_instance_attribute_availability_zone text,
                                         ec2_instance_attribute_key_name text,
                                         ec2_instance_attribute_subnet_id text,
                                         ec2_instance_attribute_emr_managed_master_security_group text,
                                         ec2_instance_attribute_emr_managed_slave_security_group text,
                                         ec2_instance_attribute_iam_instance_profile text,
                                         ec2_instance_attribute_requested_availability_zones text[],
                                         ec2_instance_attribute_requested_subnet_ids text[],
                                         ec2_instance_attribute_service_access_security_group text,
                                         id text,
                                         instance_collection_type text,
                                         kerberos_kdc_admin_password text,
                                         kerberos_realm text,
                                         kerberos_ad_domain_join_password text,
                                         kerberos_ad_domain_join_user text,
                                         kerberos_cross_realm_trust_principal_password text,
                                         log_encryption_kms_key_id text,
                                         log_uri text,
                                         master_public_dns_name text,
                                         name text,
                                         normalized_instance_hours integer,
                                         outpost_arn text,
                                         placement_groups jsonb,
                                         release_label text,
                                         repo_upgrade_on_boot text,
                                         requested_ami_version text,
                                         running_ami_version text,
                                         scale_down_behavior text,
                                         security_configuration text,
                                         service_role text,
                                         state text,
                                         state_change_reason_code text,
                                         state_change_reason_message text,
                                         creation_date_time timestamp without time zone,
                                         end_date_time timestamp without time zone,
                                         ready_date_time timestamp without time zone,
                                         step_concurrency_level integer,
                                         tags jsonb,
                                         termination_protected boolean,
                                         visible_to_all_users boolean
);

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_cq_id_key UNIQUE (cq_id);

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_pk PRIMARY KEY (arn);

