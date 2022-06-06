--
-- PostgreSQL database dump
--

-- Dumped from database version 10.21 (Debian 10.21-1.pgdg90+1)
-- Dumped by pg_dump version 14.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: cloudquery; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA cloudquery;


ALTER SCHEMA cloudquery OWNER TO postgres;

--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA cloudquery;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: calculate_policy_executions_stats(); Type: FUNCTION; Schema: cloudquery; Owner: postgres
--

CREATE FUNCTION cloudquery.calculate_policy_executions_stats() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
		UPDATE cloudquery.policy_executions SET
			checks_total = checks_total + 1,
			checks_failed = checks_failed + CASE WHEN NEW.status = 'failed' THEN 1 ELSE 0 END,
			checks_passed = checks_passed + CASE WHEN NEW.status = 'passed' THEN 1 ELSE 0 END
			WHERE id = NEW.execution_id;
    RETURN NEW;
END;
$$;


ALTER FUNCTION cloudquery.calculate_policy_executions_stats() OWNER TO postgres;

SET default_tablespace = '';

--
-- Name: check_results; Type: TABLE; Schema: cloudquery; Owner: postgres
--

CREATE TABLE cloudquery.check_results (
    execution_id uuid NOT NULL,
    execution_timestamp timestamp without time zone,
    name text,
    selector text NOT NULL,
    description text,
    status text,
    raw_results jsonb,
    error text
);


ALTER TABLE cloudquery.check_results OWNER TO postgres;

--
-- Name: cloudquery_core_schema_migrations; Type: TABLE; Schema: cloudquery; Owner: postgres
--

CREATE TABLE cloudquery.cloudquery_core_schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE cloudquery.cloudquery_core_schema_migrations OWNER TO postgres;

--
-- Name: fetches; Type: TABLE; Schema: cloudquery; Owner: postgres
--

CREATE TABLE cloudquery.fetches (
    id uuid NOT NULL,
    fetch_id uuid NOT NULL,
    start timestamp without time zone,
    finish timestamp without time zone,
    total_resource_count bigint,
    total_errors_count bigint,
    provider_name text,
    provider_version text,
    is_success boolean,
    results jsonb,
    provider_alias text,
    core_version text,
    created_at timestamp without time zone,
    CONSTRAINT non_nil_fetch_id CHECK ((fetch_id <> '00000000-0000-0000-0000-000000000000'::uuid))
);


ALTER TABLE cloudquery.fetches OWNER TO postgres;

--
-- Name: policy_executions; Type: TABLE; Schema: cloudquery; Owner: postgres
--

CREATE TABLE cloudquery.policy_executions (
    id uuid NOT NULL,
    "timestamp" timestamp without time zone,
    scheme text,
    location text,
    policy_name text,
    selector text,
    sha256_hash text,
    version text,
    checks_total integer,
    checks_failed integer,
    checks_passed integer
);


ALTER TABLE cloudquery.policy_executions OWNER TO postgres;

--
-- Name: providers; Type: TABLE; Schema: cloudquery; Owner: postgres
--

CREATE TABLE cloudquery.providers (
    source text NOT NULL,
    name text NOT NULL,
    version text NOT NULL,
    v_major integer NOT NULL,
    v_minor integer NOT NULL,
    v_patch integer NOT NULL,
    v_pre text NOT NULL,
    v_meta text NOT NULL,
    tables jsonb NOT NULL,
    signatures jsonb NOT NULL
);


ALTER TABLE cloudquery.providers OWNER TO postgres;

--
-- Name: aws_access_analyzer_analyzer_archive_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_access_analyzer_analyzer_archive_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    analyzer_cq_id uuid,
    created_at timestamp without time zone,
    filter jsonb,
    rule_name text,
    updated_at timestamp without time zone
);


ALTER TABLE public.aws_access_analyzer_analyzer_archive_rules OWNER TO postgres;

--
-- Name: aws_access_analyzer_analyzer_finding_sources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_access_analyzer_analyzer_finding_sources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    analyzer_finding_cq_id uuid,
    type text,
    detail_access_point_arn text
);


ALTER TABLE public.aws_access_analyzer_analyzer_finding_sources OWNER TO postgres;

--
-- Name: aws_access_analyzer_analyzer_findings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_access_analyzer_analyzer_findings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    analyzer_cq_id uuid,
    analyzed_at timestamp without time zone,
    condition jsonb,
    created_at timestamp without time zone,
    id text,
    resource_owner_account text,
    resource_type text,
    status text,
    updated_at timestamp without time zone,
    action text[],
    error text,
    is_public boolean,
    principal jsonb,
    resource text
);


ALTER TABLE public.aws_access_analyzer_analyzer_findings OWNER TO postgres;

--
-- Name: aws_access_analyzer_analyzers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_access_analyzer_analyzers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    created_at timestamp without time zone,
    name text,
    status text,
    type text,
    last_resource_analyzed text,
    last_resource_analyzed_at timestamp without time zone,
    status_reason_code text,
    tags jsonb
);


ALTER TABLE public.aws_access_analyzer_analyzers OWNER TO postgres;

--
-- Name: aws_accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_accounts (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    users integer,
    users_quota integer,
    groups integer,
    groups_quota integer,
    server_certificates integer,
    server_certificates_quota integer,
    user_policy_size_quota integer,
    group_policy_size_quota integer,
    groups_per_user_quota integer,
    signing_certificates_per_user_quota integer,
    access_keys_per_user_quota integer,
    mfa_devices integer,
    mfa_devices_in_use integer,
    account_mfa_enabled boolean,
    account_access_keys_present boolean,
    account_signing_certificates_present boolean,
    attached_policies_per_group_quota integer,
    policies integer,
    policies_quota integer,
    policy_size_quota integer,
    policy_versions_in_use integer,
    policy_versions_in_use_quota integer,
    versions_per_policy_quota integer,
    global_endpoint_token_version integer,
    aliases text[]
);


ALTER TABLE public.aws_accounts OWNER TO postgres;

--
-- Name: aws_acm_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_acm_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    certificate_authority_arn text,
    created_at timestamp without time zone,
    domain_name text,
    domain_validation_options jsonb,
    extended_key_usages jsonb,
    failure_reason text,
    imported_at timestamp without time zone,
    in_use_by text[],
    issued_at timestamp without time zone,
    issuer text,
    key_algorithm text,
    key_usages text[],
    not_after timestamp without time zone,
    not_before timestamp without time zone,
    certificate_transparency_logging_preference text,
    renewal_eligibility text,
    renewal_summary_domain_validation_options jsonb,
    renewal_summary_status text,
    renewal_summary_updated_at timestamp without time zone,
    renewal_summary_failure_reason text,
    revocation_reason text,
    revoked_at timestamp without time zone,
    serial text,
    signature_algorithm text,
    status text,
    subject text,
    subject_alternative_names text[],
    type text,
    tags jsonb
);


ALTER TABLE public.aws_acm_certificates OWNER TO postgres;

--
-- Name: aws_apigateway_api_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_api_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    created_date timestamp without time zone,
    customer_id text,
    description text,
    enabled boolean,
    id text NOT NULL,
    last_updated_date timestamp without time zone,
    name text,
    stage_keys text[],
    tags jsonb,
    value text
);


ALTER TABLE public.aws_apigateway_api_keys OWNER TO postgres;

--
-- Name: aws_apigateway_client_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_client_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    created_date timestamp without time zone,
    description text,
    expiration_date timestamp without time zone,
    pem_encoded_certificate text,
    tags jsonb
);


ALTER TABLE public.aws_apigateway_client_certificates OWNER TO postgres;

--
-- Name: aws_apigateway_domain_name_base_path_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_domain_name_base_path_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    domain_name_cq_id uuid,
    arn text,
    domain_name text,
    base_path text,
    rest_api_id text,
    stage text
);


ALTER TABLE public.aws_apigateway_domain_name_base_path_mappings OWNER TO postgres;

--
-- Name: aws_apigateway_domain_names; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_domain_names (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    certificate_arn text,
    certificate_name text,
    certificate_upload_date timestamp without time zone,
    distribution_domain_name text,
    distribution_hosted_zone_id text,
    domain_name text NOT NULL,
    domain_name_status text,
    domain_name_status_message text,
    endpoint_configuration_types text[],
    endpoint_configuration_vpc_endpoint_ids text[],
    mutual_tls_authentication_truststore_uri text,
    mutual_tls_authentication_truststore_version text,
    mutual_tls_authentication_truststore_warnings text[],
    regional_certificate_arn text,
    regional_certificate_name text,
    regional_domain_name text,
    regional_hosted_zone_id text,
    security_policy text,
    tags jsonb
);


ALTER TABLE public.aws_apigateway_domain_names OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_authorizers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_authorizers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    auth_type text,
    authorizer_credentials text,
    authorizer_result_ttl_in_seconds integer,
    authorizer_uri text,
    id text,
    identity_source text,
    identity_validation_expression text,
    name text,
    provider_arns text[],
    type text
);


ALTER TABLE public.aws_apigateway_rest_api_authorizers OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_deployments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_deployments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    api_summary jsonb,
    created_date timestamp without time zone,
    description text,
    id text
);


ALTER TABLE public.aws_apigateway_rest_api_deployments OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_documentation_parts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_documentation_parts (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    id text,
    location_type text,
    location_method text,
    location_name text,
    location_path text,
    location_status_code text,
    properties text
);


ALTER TABLE public.aws_apigateway_rest_api_documentation_parts OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_documentation_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_documentation_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    created_date timestamp without time zone,
    description text,
    version text
);


ALTER TABLE public.aws_apigateway_rest_api_documentation_versions OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_gateway_responses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_gateway_responses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    default_response boolean,
    response_parameters jsonb,
    response_templates jsonb,
    response_type text,
    status_code text
);


ALTER TABLE public.aws_apigateway_rest_api_gateway_responses OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_models; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_models (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    model_template text,
    content_type text,
    description text,
    id text,
    name text,
    schema text
);


ALTER TABLE public.aws_apigateway_rest_api_models OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_request_validators; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_request_validators (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    id text,
    name text,
    validate_request_body boolean,
    validate_request_parameters boolean
);


ALTER TABLE public.aws_apigateway_rest_api_request_validators OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_resources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    id text,
    parent_id text,
    path text,
    path_part text,
    resource_methods jsonb
);


ALTER TABLE public.aws_apigateway_rest_api_resources OWNER TO postgres;

--
-- Name: aws_apigateway_rest_api_stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_api_stages (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rest_api_cq_id uuid,
    rest_api_id text,
    arn text,
    access_log_settings_destination_arn text,
    access_log_settings_format text,
    cache_cluster_enabled boolean,
    cache_cluster_size text,
    cache_cluster_status text,
    canary_settings_deployment_id text,
    canary_settings_percent_traffic double precision,
    canary_settings_stage_variable_overrides jsonb,
    canary_settings_use_stage_cache boolean,
    client_certificate_id text,
    created_date timestamp without time zone,
    deployment_id text,
    description text,
    documentation_version text,
    last_updated_date timestamp without time zone,
    method_settings jsonb,
    stage_name text,
    tags jsonb,
    tracing_enabled boolean,
    variables jsonb,
    web_acl_arn text
);


ALTER TABLE public.aws_apigateway_rest_api_stages OWNER TO postgres;

--
-- Name: aws_apigateway_rest_apis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_rest_apis (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    api_key_source text,
    binary_media_types text[],
    created_date timestamp without time zone,
    description text,
    disable_execute_api_endpoint boolean,
    endpoint_configuration_types text[],
    endpoint_configuration_vpc_endpoint_ids text[],
    id text NOT NULL,
    minimum_compression_size integer,
    name text,
    policy text,
    tags jsonb,
    version text,
    warnings text[]
);


ALTER TABLE public.aws_apigateway_rest_apis OWNER TO postgres;

--
-- Name: aws_apigateway_usage_plan_api_stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_usage_plan_api_stages (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    usage_plan_cq_id uuid,
    usage_plan_id text,
    api_id text,
    stage text,
    throttle jsonb
);


ALTER TABLE public.aws_apigateway_usage_plan_api_stages OWNER TO postgres;

--
-- Name: aws_apigateway_usage_plan_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_usage_plan_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    usage_plan_cq_id uuid,
    usage_plan_id text,
    arn text,
    id text,
    name text,
    type text,
    value text
);


ALTER TABLE public.aws_apigateway_usage_plan_keys OWNER TO postgres;

--
-- Name: aws_apigateway_usage_plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_usage_plans (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    description text,
    id text NOT NULL,
    name text,
    product_code text,
    quota_limit integer,
    quota_offset integer,
    quota_period text,
    tags jsonb,
    throttle_burst_limit integer,
    throttle_rate_limit double precision
);


ALTER TABLE public.aws_apigateway_usage_plans OWNER TO postgres;

--
-- Name: aws_apigateway_vpc_links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigateway_vpc_links (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    description text,
    id text NOT NULL,
    name text,
    status text,
    status_message text,
    tags jsonb,
    target_arns text[]
);


ALTER TABLE public.aws_apigateway_vpc_links OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_authorizers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_authorizers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    name text,
    authorizer_credentials_arn text,
    authorizer_id text,
    authorizer_payload_format_version text,
    authorizer_result_ttl_in_seconds integer,
    authorizer_type text,
    authorizer_uri text,
    enable_simple_responses boolean,
    identity_source text[],
    identity_validation_expression text,
    jwt_configuration_audience text[],
    jwt_configuration_issuer text
);


ALTER TABLE public.aws_apigatewayv2_api_authorizers OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_deployments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_deployments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    auto_deployed boolean,
    created_date timestamp without time zone,
    deployment_id text,
    deployment_status text,
    deployment_status_message text,
    description text
);


ALTER TABLE public.aws_apigatewayv2_api_deployments OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_integration_responses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_integration_responses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_integration_cq_id uuid,
    integration_id text,
    arn text,
    integration_response_key text,
    content_handling_strategy text,
    integration_response_id text,
    response_parameters jsonb,
    response_templates jsonb,
    template_selection_expression text
);


ALTER TABLE public.aws_apigatewayv2_api_integration_responses OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_integrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_integrations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    api_gateway_managed boolean,
    connection_id text,
    connection_type text,
    content_handling_strategy text,
    credentials_arn text,
    description text,
    integration_id text,
    integration_method text,
    integration_response_selection_expression text,
    integration_subtype text,
    integration_type text,
    integration_uri text,
    passthrough_behavior text,
    payload_format_version text,
    request_parameters jsonb,
    request_templates jsonb,
    response_parameters jsonb,
    template_selection_expression text,
    timeout_in_millis integer,
    tls_config_server_name_to_verify text
);


ALTER TABLE public.aws_apigatewayv2_api_integrations OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_models; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_models (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    model_template text,
    name text,
    content_type text,
    description text,
    model_id text,
    schema text
);


ALTER TABLE public.aws_apigatewayv2_api_models OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_route_responses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_route_responses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_route_cq_id uuid,
    route_id text,
    arn text,
    route_response_key text,
    model_selection_expression text,
    response_models jsonb,
    response_parameters jsonb,
    route_response_id text
);


ALTER TABLE public.aws_apigatewayv2_api_route_responses OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_routes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_routes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    route_key text,
    api_gateway_managed boolean,
    api_key_required boolean,
    authorization_scopes text[],
    authorization_type text,
    authorizer_id text,
    model_selection_expression text,
    operation_name text,
    request_models jsonb,
    request_parameters jsonb,
    route_id text,
    route_response_selection_expression text,
    target text
);


ALTER TABLE public.aws_apigatewayv2_api_routes OWNER TO postgres;

--
-- Name: aws_apigatewayv2_api_stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_api_stages (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    api_cq_id uuid,
    api_id text,
    arn text,
    stage_name text,
    access_log_settings_destination_arn text,
    access_log_settings_format text,
    api_gateway_managed boolean,
    auto_deploy boolean,
    client_certificate_id text,
    created_date timestamp without time zone,
    route_settings_data_trace_enabled boolean,
    route_settings_detailed_metrics_enabled boolean,
    route_settings_logging_level text,
    route_settings_throttling_burst_limit integer,
    route_settings_throttling_rate_limit double precision,
    deployment_id text,
    description text,
    last_deployment_status_message text,
    last_updated_date timestamp without time zone,
    route_settings jsonb,
    stage_variables jsonb,
    tags jsonb
);


ALTER TABLE public.aws_apigatewayv2_api_stages OWNER TO postgres;

--
-- Name: aws_apigatewayv2_apis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_apis (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    name text,
    protocol_type text,
    route_selection_expression text,
    api_endpoint text,
    api_gateway_managed boolean,
    id text NOT NULL,
    api_key_selection_expression text,
    cors_configuration_allow_credentials boolean,
    cors_configuration_allow_headers text[],
    cors_configuration_allow_methods text[],
    cors_configuration_allow_origins text[],
    cors_configuration_expose_headers text[],
    cors_configuration_max_age integer,
    created_date timestamp without time zone,
    description text,
    disable_execute_api_endpoint boolean,
    disable_schema_validation boolean,
    import_info text[],
    tags jsonb,
    version text,
    warnings text[]
);


ALTER TABLE public.aws_apigatewayv2_apis OWNER TO postgres;

--
-- Name: aws_apigatewayv2_domain_name_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_domain_name_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    domain_name_cq_id uuid,
    api_gateway_domain_name text,
    certificate_arn text,
    certificate_name text,
    certificate_upload_date timestamp without time zone,
    domain_name_status text,
    domain_name_status_message text,
    endpoint_type text,
    hosted_zone_id text,
    security_policy text
);


ALTER TABLE public.aws_apigatewayv2_domain_name_configurations OWNER TO postgres;

--
-- Name: aws_apigatewayv2_domain_name_rest_api_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_domain_name_rest_api_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    domain_name_cq_id uuid,
    api_id text,
    arn text,
    stage text,
    api_mapping_id text,
    api_mapping_key text
);


ALTER TABLE public.aws_apigatewayv2_domain_name_rest_api_mappings OWNER TO postgres;

--
-- Name: aws_apigatewayv2_domain_names; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_domain_names (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    domain_name text NOT NULL,
    api_mapping_selection_expression text,
    mutual_tls_authentication_truststore_uri text,
    mutual_tls_authentication_truststore_version text,
    mutual_tls_authentication_truststore_warnings text[],
    tags jsonb
);


ALTER TABLE public.aws_apigatewayv2_domain_names OWNER TO postgres;

--
-- Name: aws_apigatewayv2_vpc_links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_apigatewayv2_vpc_links (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    name text,
    security_group_ids text[],
    subnet_ids text[],
    id text NOT NULL,
    created_date timestamp without time zone,
    tags jsonb,
    vpc_link_status text,
    vpc_link_status_message text,
    vpc_link_version text
);


ALTER TABLE public.aws_apigatewayv2_vpc_links OWNER TO postgres;

--
-- Name: aws_applicationautoscaling_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_applicationautoscaling_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    namespace text,
    creation_time timestamp without time zone,
    arn text NOT NULL,
    name text,
    type text,
    resource_id text,
    scalable_dimension text,
    service_namespace text,
    alarms jsonb,
    step_scaling_policy_configuration jsonb,
    target_tracking_scaling_policy_configuration jsonb
);


ALTER TABLE public.aws_applicationautoscaling_policies OWNER TO postgres;

--
-- Name: aws_athena_data_catalog_database_table_columns; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_data_catalog_database_table_columns (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    data_catalog_database_table_cq_id uuid,
    name text,
    comment text,
    type text
);


ALTER TABLE public.aws_athena_data_catalog_database_table_columns OWNER TO postgres;

--
-- Name: aws_athena_data_catalog_database_table_partition_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_data_catalog_database_table_partition_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    data_catalog_database_table_cq_id uuid,
    name text,
    comment text,
    type text
);


ALTER TABLE public.aws_athena_data_catalog_database_table_partition_keys OWNER TO postgres;

--
-- Name: aws_athena_data_catalog_database_tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_data_catalog_database_tables (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    data_catalog_database_cq_id uuid,
    name text,
    create_time timestamp without time zone,
    last_access_time timestamp without time zone,
    parameters jsonb,
    table_type text
);


ALTER TABLE public.aws_athena_data_catalog_database_tables OWNER TO postgres;

--
-- Name: aws_athena_data_catalog_databases; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_data_catalog_databases (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    data_catalog_cq_id uuid,
    name text,
    description text,
    parameters jsonb
);


ALTER TABLE public.aws_athena_data_catalog_databases OWNER TO postgres;

--
-- Name: aws_athena_data_catalogs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_data_catalogs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    tags jsonb,
    name text,
    type text,
    description text,
    parameters jsonb
);


ALTER TABLE public.aws_athena_data_catalogs OWNER TO postgres;

--
-- Name: aws_athena_work_group_named_queries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_work_group_named_queries (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    work_group_cq_id uuid,
    database text,
    name text,
    query_string text,
    description text,
    named_query_id text,
    work_group text
);


ALTER TABLE public.aws_athena_work_group_named_queries OWNER TO postgres;

--
-- Name: aws_athena_work_group_prepared_statements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_work_group_prepared_statements (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    work_group_cq_id uuid,
    description text,
    last_modified_time timestamp without time zone,
    query_statement text,
    statement_name text,
    work_group_name text
);


ALTER TABLE public.aws_athena_work_group_prepared_statements OWNER TO postgres;

--
-- Name: aws_athena_work_group_query_executions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_work_group_query_executions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    work_group_cq_id uuid,
    effective_engine_version text,
    selected_engine_version text,
    query text,
    catalog text,
    database text,
    id text,
    acl_configuration_s3_acl_option text,
    encryption_configuration_encryption_option text,
    encryption_configuration_kms_key text,
    expected_bucket_owner text,
    output_location text,
    statement_type text,
    data_manifest_location text,
    data_scanned_in_bytes bigint,
    engine_execution_time_in_millis bigint,
    query_planning_time_in_millis bigint,
    query_queue_time_in_millis bigint,
    service_processing_time_in_millis bigint,
    total_execution_time_in_millis bigint,
    athena_error_error_category integer,
    athena_error_error_message text,
    athena_error_error_type integer,
    athena_error_retryable boolean,
    completion_date_time timestamp without time zone,
    state text,
    state_change_reason text,
    submission_date_time timestamp without time zone,
    work_group text
);


ALTER TABLE public.aws_athena_work_group_query_executions OWNER TO postgres;

--
-- Name: aws_athena_work_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_athena_work_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    arn text NOT NULL,
    region text,
    tags jsonb,
    name text,
    bytes_scanned_cutoff_per_query bigint,
    enforce_work_group_configuration boolean,
    effective_engine_version text,
    selected_engine_version text,
    publish_cloud_watch_metrics_enabled boolean,
    requester_pays_enabled boolean,
    acl_configuration_s3_acl_option text,
    encryption_configuration_encryption_option text,
    encryption_configuration_kms_key text,
    expected_bucket_owner text,
    output_location text,
    creation_time timestamp without time zone,
    description text,
    state text
);


ALTER TABLE public.aws_athena_work_groups OWNER TO postgres;

--
-- Name: aws_autoscaling_group_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_group_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    group_cq_id uuid,
    availability_zone text,
    health_status text,
    id text,
    lifecycle_state text,
    protected_from_scale_in boolean,
    type text,
    launch_configuration_name text,
    launch_template_id text,
    launch_template_name text,
    launch_template_version text,
    weighted_capacity text
);


ALTER TABLE public.aws_autoscaling_group_instances OWNER TO postgres;

--
-- Name: aws_autoscaling_group_lifecycle_hooks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_group_lifecycle_hooks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    group_cq_id uuid,
    auto_scaling_group_name text,
    default_result text,
    global_timeout integer,
    heartbeat_timeout integer,
    lifecycle_hook_name text,
    lifecycle_transition text,
    notification_metadata text,
    notification_target_arn text,
    role_arn text
);


ALTER TABLE public.aws_autoscaling_group_lifecycle_hooks OWNER TO postgres;

--
-- Name: aws_autoscaling_group_scaling_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_group_scaling_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    group_cq_id uuid,
    adjustment_type text,
    alarms jsonb,
    auto_scaling_group_name text,
    cooldown integer,
    enabled boolean,
    estimated_instance_warmup integer,
    metric_aggregation_type text,
    min_adjustment_magnitude integer,
    min_adjustment_step integer,
    arn text,
    name text,
    type text,
    scaling_adjustment integer,
    step_adjustments jsonb,
    target_tracking_configuration_target_value double precision,
    target_tracking_configuration_customized_metric_name text,
    target_tracking_configuration_customized_metric_namespace text,
    target_tracking_configuration_customized_metric_statistic text,
    target_tracking_configuration_customized_metric_dimensions jsonb,
    target_tracking_configuration_customized_metric_unit text,
    target_tracking_configuration_disable_scale_in boolean,
    target_tracking_configuration_predefined_metric_type text,
    target_tracking_configuration_predefined_metric_resource_label text
);


ALTER TABLE public.aws_autoscaling_group_scaling_policies OWNER TO postgres;

--
-- Name: aws_autoscaling_group_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_group_tags (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    group_cq_id uuid,
    key text,
    propagate_at_launch boolean,
    resource_id text,
    resource_type text,
    value text
);


ALTER TABLE public.aws_autoscaling_group_tags OWNER TO postgres;

--
-- Name: aws_autoscaling_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    load_balancers jsonb,
    load_balancer_target_groups jsonb,
    notifications_configurations jsonb,
    name text,
    availability_zones text[],
    created_time timestamp without time zone,
    default_cooldown integer,
    desired_capacity integer,
    health_check_type text,
    max_size integer,
    min_size integer,
    arn text NOT NULL,
    capacity_rebalance boolean,
    enabled_metrics jsonb,
    health_check_grace_period integer,
    launch_configuration_name text,
    launch_template_id text,
    launch_template_name text,
    launch_template_version text,
    load_balancer_names text[],
    max_instance_lifetime integer,
    mixed_instances_policy jsonb,
    new_instances_protected_from_scale_in boolean,
    placement_group text,
    service_linked_role_arn text,
    status text,
    suspended_processes jsonb,
    target_group_arns text[],
    termination_policies text[],
    vpc_zone_identifier text
);


ALTER TABLE public.aws_autoscaling_groups OWNER TO postgres;

--
-- Name: aws_autoscaling_launch_configuration_block_device_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_launch_configuration_block_device_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    launch_configuration_cq_id uuid,
    device_name text,
    ebs_delete_on_termination boolean,
    ebs_encrypted boolean,
    ebs_iops integer,
    ebs_snapshot_id text,
    ebs_volume_size integer,
    ebs_volume_type text,
    no_device boolean,
    virtual_name text
);


ALTER TABLE public.aws_autoscaling_launch_configuration_block_device_mappings OWNER TO postgres;

--
-- Name: aws_autoscaling_launch_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_launch_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    created_time timestamp without time zone,
    image_id text,
    instance_type text,
    launch_configuration_name text,
    associate_public_ip_address boolean,
    classic_link_vpc_id text,
    classic_link_vpc_security_groups text[],
    ebs_optimized boolean,
    iam_instance_profile text,
    instance_monitoring_enabled boolean,
    kernel_id text,
    key_name text,
    arn text NOT NULL,
    metadata_options_http_endpoint text,
    metadata_options_http_put_response_hop_limit integer,
    metadata_options_http_tokens text,
    placement_tenancy text,
    ramdisk_id text,
    security_groups text[],
    spot_price text,
    user_data text
);


ALTER TABLE public.aws_autoscaling_launch_configurations OWNER TO postgres;

--
-- Name: aws_autoscaling_scheduled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_autoscaling_scheduled_actions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    auto_scaling_group_name text,
    desired_capacity integer,
    end_time timestamp without time zone,
    max_size integer,
    min_size integer,
    recurrence text,
    arn text NOT NULL,
    name text,
    start_time timestamp without time zone,
    "time" timestamp without time zone,
    time_zone text
);


ALTER TABLE public.aws_autoscaling_scheduled_actions OWNER TO postgres;

--
-- Name: aws_backup_global_settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_global_settings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    global_settings jsonb,
    last_update_time timestamp without time zone
);


ALTER TABLE public.aws_backup_global_settings OWNER TO postgres;

--
-- Name: aws_backup_plan_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_plan_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    plan_cq_id uuid,
    name text,
    target_backup_vault_name text,
    completion_window_minutes bigint,
    copy_actions jsonb,
    enable_continuous_backup boolean,
    delete_after_days bigint,
    move_to_cold_storage_after_days bigint,
    recovery_point_tags jsonb,
    id text,
    schedule_expression text,
    start_window_minutes bigint
);


ALTER TABLE public.aws_backup_plan_rules OWNER TO postgres;

--
-- Name: aws_backup_plan_selections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_plan_selections (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    plan_cq_id uuid,
    creation_date timestamp without time zone,
    creator_request_id text,
    iam_role_arn text,
    selection_id text,
    selection_name text,
    conditions jsonb,
    list_of_tags jsonb,
    not_resources text[],
    resources text[]
);


ALTER TABLE public.aws_backup_plan_selections OWNER TO postgres;

--
-- Name: aws_backup_plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_plans (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    id text,
    name text,
    creation_date timestamp without time zone,
    creator_request_id text,
    last_execution_date timestamp without time zone,
    version_id text,
    advanced_backup_settings jsonb,
    tags jsonb
);


ALTER TABLE public.aws_backup_plans OWNER TO postgres;

--
-- Name: aws_backup_region_settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_region_settings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    resource_type_management_preference jsonb,
    resource_type_opt_in_preference jsonb
);


ALTER TABLE public.aws_backup_region_settings OWNER TO postgres;

--
-- Name: aws_backup_vault_recovery_points; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_vault_recovery_points (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vault_cq_id uuid,
    backup_size bigint,
    calculated_delete_at timestamp without time zone,
    calculated_move_to_cold_storage_at timestamp without time zone,
    completion_date timestamp without time zone,
    created_by jsonb,
    creation_date timestamp without time zone,
    encryption_key_arn text,
    iam_role_arn text,
    is_encrypted boolean,
    last_restore_time timestamp without time zone,
    delete_after bigint,
    move_to_cold_storage_after bigint,
    arn text,
    resource_arn text,
    resource_type text,
    source_backup_vault_arn text,
    status text,
    status_message text,
    tags jsonb
);


ALTER TABLE public.aws_backup_vault_recovery_points OWNER TO postgres;

--
-- Name: aws_backup_vaults; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_backup_vaults (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    name text,
    creation_date timestamp without time zone,
    creator_request_id text,
    encryption_key_arn text,
    lock_date timestamp without time zone,
    locked boolean,
    max_retention_days bigint,
    min_retention_days bigint,
    number_of_recovery_points bigint,
    access_policy jsonb,
    notification_events text[],
    notification_sns_topic_arn text,
    tags jsonb
);


ALTER TABLE public.aws_backup_vaults OWNER TO postgres;

--
-- Name: aws_cloudformation_stack_outputs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudformation_stack_outputs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    stack_cq_id uuid,
    description text,
    export_name text,
    output_key text,
    output_value text
);


ALTER TABLE public.aws_cloudformation_stack_outputs OWNER TO postgres;

--
-- Name: aws_cloudformation_stack_resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudformation_stack_resources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    stack_cq_id uuid,
    last_updated_timestamp timestamp without time zone,
    logical_resource_id text,
    resource_status text,
    resource_type text,
    stack_resource_drift_status text,
    drift_last_check_timestamp timestamp without time zone,
    module_info_logical_id_hierarchy text,
    module_info_type_hierarchy text,
    physical_resource_id text,
    resource_status_reason text
);


ALTER TABLE public.aws_cloudformation_stack_resources OWNER TO postgres;

--
-- Name: aws_cloudformation_stacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudformation_stacks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text,
    creation_time timestamp without time zone,
    stack text,
    status text,
    capabilities text[],
    change_set_id text,
    deletion_time timestamp without time zone,
    description text,
    disable_rollback boolean,
    stack_drift_status text,
    drift_last_check_timestamp timestamp without time zone,
    enable_termination_protection boolean,
    last_updated_time timestamp without time zone,
    notification_arns text[],
    parameters jsonb,
    parent_id text,
    role_arn text,
    rollback_configuration_monitoring_time_in_minutes integer,
    rollback_configuration_rollback_triggers jsonb,
    root_id text,
    id text NOT NULL,
    stack_status_reason text,
    tags jsonb,
    timeout_in_minutes integer
);


ALTER TABLE public.aws_cloudformation_stacks OWNER TO postgres;

--
-- Name: aws_cloudfront_cache_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_cache_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    min_ttl bigint,
    name text,
    comment text,
    default_ttl bigint,
    max_ttl bigint,
    cookies_behavior text,
    cookies_quantity integer,
    cookies text[],
    enable_accept_encoding_gzip boolean,
    headers_behavior text,
    headers_quantity integer,
    headers text[],
    query_strings_behavior text,
    query_strings_quantity integer,
    query_strings text[],
    enable_accept_encoding_brotli boolean,
    id text NOT NULL,
    last_modified_time timestamp without time zone,
    type text
);


ALTER TABLE public.aws_cloudfront_cache_policies OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_cache_behavior_lambda_functions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_cache_behavior_lambda_functions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cache_behavior_cq_id uuid,
    event_type text,
    lambda_function_arn text,
    include_body boolean
);


ALTER TABLE public.aws_cloudfront_distribution_cache_behavior_lambda_functions OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_cache_behaviors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_cache_behaviors (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cq_id uuid,
    path_pattern text,
    target_origin_id text,
    viewer_protocol_policy text,
    allowed_methods text[],
    cached_methods text[],
    cache_policy_id text,
    compress boolean,
    default_ttl bigint,
    field_level_encryption_id text,
    forwarded_values_cookies_forward text,
    forwarded_values_cookies_whitelisted_names text[],
    forwarded_values_query_string boolean,
    forwarded_values_headers text[],
    forwarded_values_query_string_cache_keys text[],
    max_ttl bigint,
    min_ttl bigint,
    origin_request_policy_id text,
    realtime_log_config_arn text,
    smooth_streaming boolean,
    trusted_key_groups_enabled boolean,
    trusted_key_groups text[],
    trusted_signers_enabled boolean,
    trusted_signers text[]
);


ALTER TABLE public.aws_cloudfront_distribution_cache_behaviors OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_custom_error_responses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_custom_error_responses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cq_id uuid,
    error_code integer,
    error_caching_min_ttl bigint,
    response_code text,
    response_page_path text
);


ALTER TABLE public.aws_cloudfront_distribution_custom_error_responses OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_default_cache_behavior_functions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_default_cache_behavior_functions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cq_id uuid,
    event_type text,
    lambda_function_arn text,
    include_body boolean
);


ALTER TABLE public.aws_cloudfront_distribution_default_cache_behavior_functions OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_origin_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_origin_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cq_id uuid,
    failover_criteria_status_codes integer[],
    id text,
    members_origin_ids text[]
);


ALTER TABLE public.aws_cloudfront_distribution_origin_groups OWNER TO postgres;

--
-- Name: aws_cloudfront_distribution_origins; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distribution_origins (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    distribution_cq_id uuid,
    domain_name text,
    id text,
    connection_attempts integer,
    connection_timeout integer,
    custom_headers jsonb,
    custom_origin_config_http_port integer,
    custom_origin_config_https_port integer,
    custom_origin_config_protocol_policy text,
    custom_origin_config_keepalive_timeout integer,
    custom_origin_config_read_timeout integer,
    custom_origin_config_ssl_protocols text[],
    origin_path text,
    origin_shield_enabled boolean,
    origin_shield_region text,
    s3_origin_config_origin_access_identity text
);


ALTER TABLE public.aws_cloudfront_distribution_origins OWNER TO postgres;

--
-- Name: aws_cloudfront_distributions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudfront_distributions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    tags jsonb,
    arn text NOT NULL,
    caller_reference text,
    comment text,
    cache_behavior_target_origin_id text,
    cache_behavior_viewer_protocol_policy text,
    cache_behavior_allowed_methods text[],
    cache_behavior_allowed_methods_cached_methods text[],
    cache_behavior_cache_policy_id text,
    cache_behavior_compress boolean,
    cache_behavior_default_ttl bigint,
    cache_behavior_field_level_encryption_id text,
    cache_behavior_forwarded_values_cookies_forward text,
    cache_behavior_forwarded_values_cookies_whitelisted_names text[],
    cache_behavior_forwarded_values_query_string boolean,
    cache_behavior_forwarded_values_headers text[],
    cache_behavior_forwarded_values_query_string_cache_keys text[],
    cache_behavior_max_ttl bigint,
    cache_behavior_min_ttl bigint,
    cache_behavior_origin_request_policy_id text,
    cache_behavior_realtime_log_config_arn text,
    cache_behavior_smooth_streaming boolean,
    cache_behavior_trusted_key_groups_enabled boolean,
    cache_behavior_trusted_key_groups text[],
    cache_behavior_trusted_signers_enabled boolean,
    cache_behavior_trusted_signers text[],
    enabled boolean,
    aliases text[],
    default_root_object text,
    http_version text,
    ipv6_enabled boolean,
    logging_bucket text,
    logging_enabled boolean,
    logging_include_cookies boolean,
    logging_prefix text,
    price_class text,
    geo_restriction_type text,
    geo_restrictions text[],
    viewer_certificate_acm_certificate_arn text,
    viewer_certificate text,
    viewer_certificate_source text,
    viewer_certificate_cloudfront_default_certificate boolean,
    viewer_certificate_iam_certificate_id text,
    viewer_certificate_minimum_protocol_version text,
    viewer_certificate_ssl_support_method text,
    web_acl_id text,
    domain_name text,
    id text,
    in_progress_invalidation_batches integer,
    last_modified_time timestamp without time zone,
    status text,
    active_trusted_key_groups_enabled boolean,
    active_trusted_key_groups jsonb,
    active_trusted_signers_enabled boolean,
    active_trusted_signers jsonb,
    alias_icp_recordals jsonb
);


ALTER TABLE public.aws_cloudfront_distributions OWNER TO postgres;

--
-- Name: aws_cloudtrail_trail_event_selectors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudtrail_trail_event_selectors (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    trail_cq_id uuid,
    trail_arn text,
    exclude_management_event_sources text[],
    include_management_events boolean,
    read_write_type text
);


ALTER TABLE public.aws_cloudtrail_trail_event_selectors OWNER TO postgres;

--
-- Name: aws_cloudtrail_trails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudtrail_trails (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    tags jsonb,
    cloudwatch_logs_log_group_name text,
    is_logging boolean,
    latest_cloud_watch_logs_delivery_error text,
    latest_cloud_watch_logs_delivery_time timestamp without time zone,
    latest_delivery_error text,
    latest_delivery_time timestamp without time zone,
    latest_digest_delivery_error text,
    latest_digest_delivery_time timestamp without time zone,
    latest_notification_error text,
    latest_notification_time timestamp without time zone,
    start_logging_time timestamp without time zone,
    stop_logging_time timestamp without time zone,
    cloud_watch_logs_log_group_arn text,
    cloud_watch_logs_role_arn text,
    has_custom_event_selectors boolean,
    has_insight_selectors boolean,
    region text,
    include_global_service_events boolean,
    is_multi_region_trail boolean,
    is_organization_trail boolean,
    kms_key_id text,
    log_file_validation_enabled boolean,
    name text,
    s3_bucket_name text,
    s3_key_prefix text,
    sns_topic_arn text,
    sns_topic_name text,
    arn text NOT NULL
);


ALTER TABLE public.aws_cloudtrail_trails OWNER TO postgres;

--
-- Name: aws_cloudwatch_alarm_metrics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudwatch_alarm_metrics (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    alarm_cq_id uuid,
    alarm_arn text,
    alarm_name text,
    id text,
    expression text,
    label text,
    metric_stat_metric_dimensions jsonb,
    metric_stat_metric_name text,
    metric_stat_metric_namespace text,
    metric_stat_period integer,
    metric_stat text,
    metric_stat_unit text,
    period integer,
    return_data boolean
);


ALTER TABLE public.aws_cloudwatch_alarm_metrics OWNER TO postgres;

--
-- Name: aws_cloudwatch_alarms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudwatch_alarms (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    actions_enabled boolean,
    actions text[],
    arn text NOT NULL,
    configuration_updated_timestamp timestamp without time zone,
    description text,
    name text,
    comparison_operator text,
    datapoints_to_alarm integer,
    dimensions jsonb,
    evaluate_low_sample_count_percentile text,
    evaluation_periods integer,
    extended_statistic text,
    insufficient_data_actions text[],
    metric_name text,
    namespace text,
    ok_actions text[],
    period integer,
    state_reason text,
    state_reason_data text,
    state_updated_timestamp timestamp without time zone,
    state_value text,
    statistic text,
    threshold double precision,
    threshold_metric_id text,
    treat_missing_data text,
    unit text
);


ALTER TABLE public.aws_cloudwatch_alarms OWNER TO postgres;

--
-- Name: aws_cloudwatchlogs_filter_metric_transformations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudwatchlogs_filter_metric_transformations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    filter_cq_id uuid,
    metric_name text,
    metric_namespace text,
    metric_value text,
    default_value double precision
);


ALTER TABLE public.aws_cloudwatchlogs_filter_metric_transformations OWNER TO postgres;

--
-- Name: aws_cloudwatchlogs_filters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cloudwatchlogs_filters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    creation_time bigint,
    name text NOT NULL,
    pattern text,
    log_group_name text NOT NULL
);


ALTER TABLE public.aws_cloudwatchlogs_filters OWNER TO postgres;

--
-- Name: aws_codebuild_project_environment_variables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codebuild_project_environment_variables (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    project_cq_id uuid,
    name text,
    value text,
    type text
);


ALTER TABLE public.aws_codebuild_project_environment_variables OWNER TO postgres;

--
-- Name: aws_codebuild_project_file_system_locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codebuild_project_file_system_locations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    project_cq_id uuid,
    identifier text,
    location text,
    mount_options text,
    mount_point text,
    type text
);


ALTER TABLE public.aws_codebuild_project_file_system_locations OWNER TO postgres;

--
-- Name: aws_codebuild_project_secondary_artifacts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codebuild_project_secondary_artifacts (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    project_cq_id uuid,
    type text,
    artifact_identifier text,
    bucket_owner_access text,
    encryption_disabled boolean,
    location text,
    name text,
    namespace_type text,
    override_artifact_name boolean,
    packaging text,
    path text
);


ALTER TABLE public.aws_codebuild_project_secondary_artifacts OWNER TO postgres;

--
-- Name: aws_codebuild_project_secondary_sources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codebuild_project_secondary_sources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    project_cq_id uuid,
    type text,
    auth_type text,
    auth_resource text,
    build_status_config_context text,
    build_status_config_target_url text,
    buildspec text,
    git_clone_depth integer,
    git_submodules_config_fetch_submodules boolean,
    insecure_ssl boolean,
    location text,
    report_build_status boolean,
    source_identifier text
);


ALTER TABLE public.aws_codebuild_project_secondary_sources OWNER TO postgres;

--
-- Name: aws_codebuild_projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codebuild_projects (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    artifacts_type text,
    artifacts_artifact_identifier text,
    artifacts_bucket_owner_access text,
    artifacts_encryption_disabled boolean,
    artifacts_location text,
    artifacts_name text,
    artifacts_namespace_type text,
    artifacts_override_artifact_name boolean,
    artifacts_packaging text,
    artifacts_path text,
    badge_enabled boolean,
    badge_request_url text,
    build_batch_config_batch_report_mode text,
    build_batch_config_combine_artifacts boolean,
    build_batch_config_restrictions_compute_types_allowed text[],
    build_batch_config_restrictions_maximum_builds_allowed integer,
    build_batch_config_service_role text,
    build_batch_config_timeout_in_mins integer,
    cache_type text,
    cache_location text,
    cache_modes text[],
    concurrent_build_limit integer,
    created timestamp without time zone,
    description text,
    encryption_key text,
    environment_compute_type text,
    environment_image text,
    environment_type text,
    environment_certificate text,
    environment_image_pull_credentials_type text,
    environment_privileged_mode boolean,
    environment_registry_credential text,
    environment_registry_credential_credential_provider text,
    last_modified timestamp without time zone,
    logs_config_cloud_watch_logs_status text,
    logs_config_cloud_watch_logs_group_name text,
    logs_config_cloud_watch_logs_stream_name text,
    logs_config_s3_logs_status text,
    logs_config_s3_logs_bucket_owner_access text,
    logs_config_s3_logs_encryption_disabled boolean,
    logs_config_s3_logs_location text,
    name text,
    project_visibility text,
    public_project_alias text,
    queued_timeout_in_minutes integer,
    resource_access_role text,
    secondary_source_versions jsonb,
    service_role text,
    source_type text,
    source_auth_type text,
    source_auth_resource text,
    source_build_status_config_context text,
    source_build_status_config_target_url text,
    source_buildspec text,
    source_git_clone_depth integer,
    source_git_submodules_config_fetch_submodules boolean,
    source_insecure_ssl boolean,
    source_location text,
    source_report_build_status boolean,
    source_identifier text,
    source_version text,
    tags jsonb,
    timeout_in_minutes integer,
    vpc_config_security_group_ids text[],
    vpc_config_subnets text[],
    vpc_config_vpc_id text,
    webhook_branch_filter text,
    webhook_build_type text,
    webhook_filter_groups jsonb,
    webhook_last_modified_secret timestamp without time zone,
    webhook_payload_url text,
    webhook_secret text,
    webhook_url text
);


ALTER TABLE public.aws_codebuild_projects OWNER TO postgres;

--
-- Name: aws_codepipeline_pipeline_stage_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codepipeline_pipeline_stage_actions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    pipeline_stage_cq_id uuid,
    category text,
    owner text,
    provider text,
    version text,
    name text,
    configuration jsonb,
    input_artifacts text[],
    namespace text,
    output_artifacts text[],
    region text,
    role_arn text,
    run_order integer
);


ALTER TABLE public.aws_codepipeline_pipeline_stage_actions OWNER TO postgres;

--
-- Name: aws_codepipeline_pipeline_stages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codepipeline_pipeline_stages (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    pipeline_cq_id uuid,
    stage_order integer,
    name text,
    blockers jsonb
);


ALTER TABLE public.aws_codepipeline_pipeline_stages OWNER TO postgres;

--
-- Name: aws_codepipeline_pipelines; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codepipeline_pipelines (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    created timestamp without time zone,
    arn text NOT NULL,
    updated timestamp without time zone,
    name text,
    role_arn text,
    artifact_store_location text,
    artifact_store_type text,
    artifact_store_encryption_key_id text,
    artifact_store_encryption_key_type text,
    artifact_stores jsonb,
    version integer
);


ALTER TABLE public.aws_codepipeline_pipelines OWNER TO postgres;

--
-- Name: aws_codepipeline_webhook_filters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codepipeline_webhook_filters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    webhook_cq_id uuid,
    json_path text,
    match_equals text
);


ALTER TABLE public.aws_codepipeline_webhook_filters OWNER TO postgres;

--
-- Name: aws_codepipeline_webhooks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_codepipeline_webhooks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    authentication text,
    authentication_allowed_ip_range text,
    authentication_secret_token text,
    name text,
    target_action text,
    target_pipeline text,
    url text,
    arn text NOT NULL,
    error_code text,
    error_message text,
    last_triggered timestamp without time zone,
    tags jsonb
);


ALTER TABLE public.aws_codepipeline_webhooks OWNER TO postgres;

--
-- Name: aws_cognito_identity_pool_cognito_identity_providers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cognito_identity_pool_cognito_identity_providers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    identity_pool_cq_id uuid,
    identity_pool_id text,
    client_id text,
    provider_name text,
    server_side_token_check boolean
);


ALTER TABLE public.aws_cognito_identity_pool_cognito_identity_providers OWNER TO postgres;

--
-- Name: aws_cognito_identity_pools; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cognito_identity_pools (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    allow_unauthenticated_identities boolean,
    id text NOT NULL,
    identity_pool_name text,
    allow_classic_flow boolean,
    developer_provider_name text,
    identity_pool_tags jsonb,
    open_id_connect_provider_arns text[],
    saml_provider_arns text[],
    supported_login_providers jsonb
);


ALTER TABLE public.aws_cognito_identity_pools OWNER TO postgres;

--
-- Name: aws_cognito_user_pool_identity_providers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cognito_user_pool_identity_providers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_pool_cq_id uuid,
    user_pool_id text,
    account_id text,
    region text,
    attribute_mapping jsonb,
    creation_date timestamp without time zone,
    idp_identifiers text[],
    last_modified_date timestamp without time zone,
    provider_details jsonb,
    provider_name text,
    provider_type text
);


ALTER TABLE public.aws_cognito_user_pool_identity_providers OWNER TO postgres;

--
-- Name: aws_cognito_user_pool_schema_attributes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cognito_user_pool_schema_attributes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_pool_cq_id uuid,
    user_pool_id text,
    attribute_data_type text,
    developer_only_attribute boolean,
    mutable boolean,
    name text,
    number_attribute_constraints_max_value text,
    number_attribute_constraints_min_value text,
    required boolean,
    string_attribute_constraints_max_length text,
    string_attribute_constraints_min_length text
);


ALTER TABLE public.aws_cognito_user_pool_schema_attributes OWNER TO postgres;

--
-- Name: aws_cognito_user_pools; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_cognito_user_pools (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    account_recovery_setting jsonb,
    admin_create_user_admin_only boolean,
    admin_create_user_invite_email_message text,
    admin_create_user_invite_email_subject text,
    admin_create_user_invite_sms text,
    admin_create_user_config_unused_account_validity_days integer,
    alias_attributes text[],
    arn text,
    auto_verified_attributes text[],
    creation_date timestamp without time zone,
    custom_domain text,
    challenge_required_on_new_device boolean,
    device_only_remembered_on_user_prompt boolean,
    domain text,
    email_configuration_set text,
    email_configuration_sending_account text,
    email_configuration_from text,
    email_configuration_reply_to_address text,
    email_configuration_source_arn text,
    email_configuration_failure text,
    email_verification_message text,
    email_verification_subject text,
    estimated_number_of_users integer,
    id text NOT NULL,
    lambda_config_create_auth_challenge text,
    lambda_config_custom_email_sender_lambda_arn text,
    lambda_config_custom_email_sender_lambda_version text,
    lambda_config_custom_message text,
    lambda_config_custom_sms_sender_lambda_arn text,
    lambda_config_custom_sms_sender_lambda_version text,
    lambda_config_define_auth_challenge text,
    lambda_config_kms_key_id text,
    lambda_config_post_authentication text,
    lambda_config_post_confirmation text,
    lambda_config_pre_authentication text,
    lambda_config_pre_sign_up text,
    lambda_config_pre_token_generation text,
    lambda_config_user_migration text,
    lambda_config_verify_auth_challenge_response text,
    last_modified_date timestamp without time zone,
    mfa_configuration text,
    name text,
    policies_password_policy_minimum_length integer,
    policies_password_policy_require_lowercase boolean,
    policies_password_policy_require_numbers boolean,
    policies_password_policy_require_symbols boolean,
    policies_password_policy_require_uppercase boolean,
    policies_password_policy_temporary_password_validity_days integer,
    sms_authentication_message text,
    sms_configuration_sns_caller_arn text,
    sms_configuration_external_id text,
    sms_configuration_failure text,
    sms_verification_message text,
    status text,
    user_pool_add_ons_advanced_security_mode text,
    user_pool_tags jsonb,
    username_attributes text[],
    username_configuration_case_sensitive boolean,
    verification_message_template_default_email_option text,
    verification_message_template_email_message text,
    verification_message_template_email_message_by_link text,
    verification_message_template_email_subject text,
    verification_message_template_email_subject_by_link text,
    verification_message_template_sms_message text
);


ALTER TABLE public.aws_cognito_user_pools OWNER TO postgres;

--
-- Name: aws_config_configuration_recorders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_config_configuration_recorders (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    name text,
    recording_group_all_supported boolean,
    recording_group_include_global_resource_types boolean,
    recording_group_resource_types text[],
    role_arn text,
    status_last_error_code text,
    status_last_error_message text,
    status_last_start_time timestamp without time zone,
    status_last_status text,
    status_last_status_change_time timestamp without time zone,
    status_last_stop_time timestamp without time zone,
    status_recording boolean
);


ALTER TABLE public.aws_config_configuration_recorders OWNER TO postgres;

--
-- Name: aws_config_conformance_pack_rule_compliances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_config_conformance_pack_rule_compliances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    conformance_pack_cq_id uuid,
    compliance_type text,
    config_rule_name text,
    controls text[],
    config_rule_invoked_time timestamp without time zone,
    resource_id text,
    resource_type text,
    ordering_timestamp timestamp without time zone,
    result_recorded_time timestamp without time zone,
    annotation text
);


ALTER TABLE public.aws_config_conformance_pack_rule_compliances OWNER TO postgres;

--
-- Name: aws_config_conformance_packs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_config_conformance_packs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    conformance_pack_id text,
    conformance_pack_name text,
    conformance_pack_input_parameters jsonb,
    created_by text,
    delivery_s3_bucket text,
    delivery_s3_key_prefix text,
    last_update_requested_time timestamp without time zone
);


ALTER TABLE public.aws_config_conformance_packs OWNER TO postgres;

--
-- Name: aws_dax_cluster_nodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dax_cluster_nodes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    availability_zone text,
    endpoint_address text,
    endpoint_port integer,
    endpoint_url text,
    node_create_time timestamp without time zone,
    node_id text,
    node_status text,
    parameter_group_status text
);


ALTER TABLE public.aws_dax_cluster_nodes OWNER TO postgres;

--
-- Name: aws_dax_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dax_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    active_nodes integer,
    arn text NOT NULL,
    cluster_discovery_endpoint_address text,
    cluster_discovery_endpoint_port integer,
    cluster_discovery_endpoint_url text,
    cluster_endpoint_encryption_type text,
    name text,
    description text,
    iam_role_arn text,
    node_ids_to_remove text[],
    node_type text,
    notification_configuration_topic_arn text,
    notification_configuration_topic_status text,
    node_ids_to_reboot text[],
    parameter_apply_status text,
    parameter_group_name text,
    preferred_maintenance_window text,
    sse_description_status text,
    security_groups jsonb,
    status text,
    subnet_group text,
    total_nodes integer
);


ALTER TABLE public.aws_dax_clusters OWNER TO postgres;

--
-- Name: aws_directconnect_connection_mac_sec_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_connection_mac_sec_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    connection_cq_id uuid,
    connection_id text,
    ckn text,
    secret_arn text,
    start_on text,
    state text
);


ALTER TABLE public.aws_directconnect_connection_mac_sec_keys OWNER TO postgres;

--
-- Name: aws_directconnect_connections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_connections (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    aws_device_v2 text,
    bandwidth text,
    id text NOT NULL,
    name text,
    connection_state text,
    encryption_mode text,
    has_logical_redundancy text,
    jumbo_frame_capable boolean,
    lag_id text,
    loa_issue_time timestamp without time zone,
    location text,
    mac_sec_capable boolean,
    owner_account text,
    partner_name text,
    port_encryption_status text,
    provider_name text,
    tags jsonb,
    vlan integer
);


ALTER TABLE public.aws_directconnect_connections OWNER TO postgres;

--
-- Name: aws_directconnect_gateway_associations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_gateway_associations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    gateway_cq_id uuid,
    gateway_id text,
    allowed_prefixes_to_direct_connect_gateway text[],
    associated_gateway_id text,
    associated_gateway_owner_account text,
    associated_gateway_region text,
    associated_gateway_type text,
    association_id text,
    association_state text,
    direct_connect_gateway_owner_account text,
    state_change_error text,
    virtual_gateway_id text,
    virtual_gateway_owner_account text,
    resource_id text
);


ALTER TABLE public.aws_directconnect_gateway_associations OWNER TO postgres;

--
-- Name: aws_directconnect_gateway_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_gateway_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    gateway_cq_id uuid,
    gateway_id text,
    attachment_state text,
    attachment_type text,
    state_change_error text,
    virtual_interface_id text,
    virtual_interface_owner_account text,
    virtual_interface_region text
);


ALTER TABLE public.aws_directconnect_gateway_attachments OWNER TO postgres;

--
-- Name: aws_directconnect_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    amazon_side_asn bigint,
    id text NOT NULL,
    name text,
    state text,
    owner_account text,
    state_change_error text
);


ALTER TABLE public.aws_directconnect_gateways OWNER TO postgres;

--
-- Name: aws_directconnect_lag_mac_sec_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_lag_mac_sec_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    lag_cq_id uuid,
    lag_id text,
    ckn text,
    secret_arn text,
    start_on text,
    state text
);


ALTER TABLE public.aws_directconnect_lag_mac_sec_keys OWNER TO postgres;

--
-- Name: aws_directconnect_lags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_lags (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    allows_hosted_connections boolean,
    aws_device_v2 text,
    connection_ids text[],
    connections_bandwidth text,
    encryption_mode text,
    has_logical_redundancy text,
    jumbo_frame_capable boolean,
    id text NOT NULL,
    name text,
    state text,
    location text,
    mac_sec_capable boolean,
    minimum_links integer,
    number_of_connections integer,
    owner_account text,
    provider_name text,
    tags jsonb
);


ALTER TABLE public.aws_directconnect_lags OWNER TO postgres;

--
-- Name: aws_directconnect_virtual_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_virtual_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    id text NOT NULL,
    state text
);


ALTER TABLE public.aws_directconnect_virtual_gateways OWNER TO postgres;

--
-- Name: aws_directconnect_virtual_interface_bgp_peers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_virtual_interface_bgp_peers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    virtual_interface_cq_id uuid,
    virtual_interface_id text,
    address_family text,
    amazon_address text,
    asn integer,
    auth_key text,
    aws_device_v2 text,
    bgp_peer_id text,
    bgp_peer_state text,
    bgp_status text,
    customer_address text
);


ALTER TABLE public.aws_directconnect_virtual_interface_bgp_peers OWNER TO postgres;

--
-- Name: aws_directconnect_virtual_interfaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_directconnect_virtual_interfaces (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    address_family text,
    amazon_address text,
    amazon_side_asn bigint,
    asn integer,
    auth_key text,
    aws_device_v2 text,
    connection_id text,
    customer_address text,
    customer_router_config text,
    direct_connect_gateway_id text,
    jumbo_frame_capable boolean,
    location text,
    mtu integer,
    owner_account text,
    region text,
    route_filter_prefixes text[],
    tags jsonb,
    virtual_gateway_id text,
    id text NOT NULL,
    virtual_interface_name text,
    virtual_interface_state text,
    virtual_interface_type text,
    vlan integer
);


ALTER TABLE public.aws_directconnect_virtual_interfaces OWNER TO postgres;

--
-- Name: aws_dms_replication_instance_replication_subnet_group_subnets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dms_replication_instance_replication_subnet_group_subnets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    replication_instance_cq_id uuid,
    subnet_availability_zone_name text,
    subnet_identifier text,
    subnet_status text
);


ALTER TABLE public.aws_dms_replication_instance_replication_subnet_group_subnets OWNER TO postgres;

--
-- Name: aws_dms_replication_instance_vpc_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dms_replication_instance_vpc_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    replication_instance_cq_id uuid,
    status text,
    vpc_security_group_id text
);


ALTER TABLE public.aws_dms_replication_instance_vpc_security_groups OWNER TO postgres;

--
-- Name: aws_dms_replication_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dms_replication_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    tags jsonb,
    allocated_storage integer,
    auto_minor_version_upgrade boolean,
    availability_zone text,
    dns_name_servers text,
    engine_version text,
    free_until timestamp without time zone,
    instance_create_time timestamp without time zone,
    kms_key_id text,
    multi_az boolean,
    pending_modified_values_allocated_storage integer,
    pending_modified_values_engine_version text,
    pending_modified_values_multi_az boolean,
    pending_modified_values_class text,
    preferred_maintenance_window text,
    publicly_accessible boolean,
    arn text NOT NULL,
    class text,
    identifier text,
    private_ip_address inet,
    private_ip_addresses inet[],
    public_ip_address inet,
    public_ip_addresses inet[],
    status text,
    replication_subnet_group_description text,
    replication_subnet_group_identifier text,
    replication_subnet_group_subnet_group_status text,
    replication_subnet_group_vpc_id text,
    secondary_availability_zone text
);


ALTER TABLE public.aws_dms_replication_instances OWNER TO postgres;

--
-- Name: aws_dynamodb_table_continuous_backups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_table_continuous_backups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    table_cq_id uuid,
    continuous_backups_status text,
    earliest_restorable_date_time timestamp without time zone,
    latest_restorable_date_time timestamp without time zone,
    point_in_time_recovery_status text
);


ALTER TABLE public.aws_dynamodb_table_continuous_backups OWNER TO postgres;

--
-- Name: aws_dynamodb_table_global_secondary_indexes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_table_global_secondary_indexes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    table_cq_id uuid,
    backfilling boolean,
    arn text,
    name text,
    index_size_bytes bigint,
    status text,
    item_count bigint,
    key_schema jsonb,
    projection_non_key_attributes text[],
    projection_type text,
    provisioned_throughput_last_decrease_date_time timestamp without time zone,
    provisioned_throughput_last_increase_date_time timestamp without time zone,
    provisioned_throughput_number_of_decreases_today bigint,
    provisioned_throughput_read_capacity_units bigint,
    provisioned_throughput_write_capacity_units bigint
);


ALTER TABLE public.aws_dynamodb_table_global_secondary_indexes OWNER TO postgres;

--
-- Name: aws_dynamodb_table_local_secondary_indexes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_table_local_secondary_indexes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    table_cq_id uuid,
    arn text,
    name text,
    index_size_bytes bigint,
    item_count bigint,
    key_schema jsonb,
    projection_non_key_attributes text[],
    projection_type text
);


ALTER TABLE public.aws_dynamodb_table_local_secondary_indexes OWNER TO postgres;

--
-- Name: aws_dynamodb_table_replica_auto_scalings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_table_replica_auto_scalings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    table_cq_id uuid,
    global_secondary_indexes jsonb,
    region_name text,
    read_capacity jsonb,
    write_capacity jsonb,
    replica_status text
);


ALTER TABLE public.aws_dynamodb_table_replica_auto_scalings OWNER TO postgres;

--
-- Name: aws_dynamodb_table_replicas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_table_replicas (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    table_cq_id uuid,
    global_secondary_indexes jsonb,
    kms_master_key_id text,
    provisioned_throughput_override_read_capacity_units bigint,
    region_name text,
    replica_inaccessible_date_time timestamp without time zone,
    replica_status text,
    replica_status_description text,
    replica_status_percent_progress text,
    summary_last_update_date_time timestamp without time zone,
    summary_table_class text
);


ALTER TABLE public.aws_dynamodb_table_replicas OWNER TO postgres;

--
-- Name: aws_dynamodb_tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_dynamodb_tables (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    archival_summary jsonb,
    attribute_definitions jsonb,
    billing_mode_summary jsonb,
    creation_date_time timestamp without time zone,
    global_table_version text,
    item_count bigint,
    key_schema jsonb,
    latest_stream_arn text,
    latest_stream_label text,
    provisioned_throughput_last_decrease_date_time timestamp without time zone,
    provisioned_throughput_last_increase_date_time timestamp without time zone,
    provisioned_throughput_number_of_decreases_today bigint,
    provisioned_throughput_read_capacity_units bigint,
    provisioned_throughput_write_capacity_units bigint,
    restore_summary jsonb,
    inaccessible_encryption_date_time timestamp without time zone,
    kms_master_key_arn text,
    sse_type text,
    sse_status text,
    stream_specification jsonb,
    arn text NOT NULL,
    table_class_last_update timestamp without time zone,
    table_class text,
    id text,
    name text,
    size_bytes bigint,
    status text
);


ALTER TABLE public.aws_dynamodb_tables OWNER TO postgres;

--
-- Name: aws_ec2_byoip_cidrs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_byoip_cidrs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    cidr text NOT NULL,
    description text,
    state text,
    status_message text
);


ALTER TABLE public.aws_ec2_byoip_cidrs OWNER TO postgres;

--
-- Name: aws_ec2_customer_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_customer_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    id text NOT NULL,
    bgp_asn text,
    certificate_arn text,
    arn text,
    device_name text,
    ip_address text,
    state text,
    tags jsonb,
    type text
);


ALTER TABLE public.aws_ec2_customer_gateways OWNER TO postgres;

--
-- Name: aws_ec2_ebs_snapshots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_ebs_snapshots (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    create_volume_permissions jsonb,
    data_encryption_key_id text,
    description text,
    encrypted boolean,
    kms_key_id text,
    outpost_arn text,
    owner_alias text,
    owner_id text,
    progress text,
    snapshot_id text NOT NULL,
    start_time timestamp without time zone,
    state text,
    state_message text,
    tags jsonb,
    volume_id text,
    volume_size integer
);


ALTER TABLE public.aws_ec2_ebs_snapshots OWNER TO postgres;

--
-- Name: aws_ec2_ebs_volume_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_ebs_volume_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    ebs_volume_cq_id uuid,
    attach_time timestamp without time zone,
    delete_on_termination boolean,
    device text,
    instance_id text,
    state text,
    volume_id text
);


ALTER TABLE public.aws_ec2_ebs_volume_attachments OWNER TO postgres;

--
-- Name: aws_ec2_ebs_volumes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_ebs_volumes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    id text NOT NULL,
    arn text,
    availability_zone text,
    create_time timestamp without time zone,
    encrypted boolean,
    fast_restored boolean,
    iops integer,
    kms_key_id text,
    multi_attach_enabled boolean,
    outpost_arn text,
    size integer,
    snapshot_id text,
    state text,
    tags jsonb,
    throughput integer,
    volume_type text
);


ALTER TABLE public.aws_ec2_ebs_volumes OWNER TO postgres;

--
-- Name: aws_ec2_egress_only_internet_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_egress_only_internet_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    attachments jsonb,
    id text,
    tags jsonb
);


ALTER TABLE public.aws_ec2_egress_only_internet_gateways OWNER TO postgres;

--
-- Name: aws_ec2_eips; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_eips (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    allocation_id text NOT NULL,
    association_id text,
    carrier_ip inet,
    customer_owned_ip inet,
    customer_owned_ipv4_pool text,
    domain text,
    instance_id text,
    network_border_group text,
    network_interface_id text,
    network_interface_owner_id text,
    private_ip_address inet,
    public_ip inet,
    public_ipv4_pool text,
    tags jsonb
);


ALTER TABLE public.aws_ec2_eips OWNER TO postgres;

--
-- Name: aws_ec2_flow_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_flow_logs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    creation_time timestamp without time zone,
    deliver_logs_error_message text,
    deliver_logs_permission_arn text,
    deliver_logs_status text,
    flow_log_id text,
    flow_log_status text,
    log_destination text,
    log_destination_type text,
    log_format text,
    log_group_name text,
    max_aggregation_interval integer,
    resource_id text,
    tags jsonb,
    traffic_type text
);


ALTER TABLE public.aws_ec2_flow_logs OWNER TO postgres;

--
-- Name: aws_ec2_host_available_instance_capacity; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_host_available_instance_capacity (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    host_cq_id uuid,
    available_capacity integer,
    instance_type text,
    total_capacity integer
);


ALTER TABLE public.aws_ec2_host_available_instance_capacity OWNER TO postgres;

--
-- Name: aws_ec2_host_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_host_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    host_cq_id uuid,
    instance_id text,
    instance_type text,
    owner_id text
);


ALTER TABLE public.aws_ec2_host_instances OWNER TO postgres;

--
-- Name: aws_ec2_hosts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_hosts (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    allocation_time timestamp without time zone,
    allows_multiple_instance_types text,
    auto_placement text,
    availability_zone text,
    availability_zone_id text,
    available_vcpus integer,
    client_token text,
    id text,
    cores integer,
    instance_family text,
    instance_type text,
    sockets integer,
    total_vcpus integer,
    host_recovery text,
    reservation_id text,
    member_of_service_linked_resource_group boolean,
    owner_id text,
    release_time timestamp without time zone,
    state text,
    tags jsonb
);


ALTER TABLE public.aws_ec2_hosts OWNER TO postgres;

--
-- Name: aws_ec2_image_block_device_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_image_block_device_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    image_cq_id uuid,
    device_name text,
    ebs_delete_on_termination boolean,
    ebs_encrypted boolean,
    ebs_iops integer,
    ebs_kms_key_id text,
    ebs_outpost_arn text,
    ebs_snapshot_id text,
    ebs_throughput integer,
    ebs_volume_size integer,
    ebs_volume_type text,
    no_device text,
    virtual_name text
);


ALTER TABLE public.aws_ec2_image_block_device_mappings OWNER TO postgres;

--
-- Name: aws_ec2_images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_images (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    architecture text,
    creation_date timestamp without time zone,
    description text,
    ena_support boolean,
    hypervisor text,
    image_location text,
    image_owner_alias text,
    image_type text,
    kernel_id text,
    name text,
    owner_id text,
    platform text,
    platform_details text,
    product_codes jsonb,
    public boolean,
    ramdisk_id text,
    root_device_name text,
    root_device_type text,
    sriov_net_support text,
    state text,
    state_reason_code text,
    state_reason_message text,
    tags jsonb,
    usage_operation text,
    virtualization_type text,
    last_launched_time timestamp without time zone
);


ALTER TABLE public.aws_ec2_images OWNER TO postgres;

--
-- Name: aws_ec2_instance_block_device_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_block_device_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    device_name text,
    ebs_attach_time timestamp without time zone,
    ebs_delete_on_termination boolean,
    ebs_status text,
    ebs_volume_id text
);


ALTER TABLE public.aws_ec2_instance_block_device_mappings OWNER TO postgres;

--
-- Name: aws_ec2_instance_elastic_gpu_associations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_elastic_gpu_associations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    elastic_gpu_association_id text,
    elastic_gpu_association_state text,
    elastic_gpu_association_time timestamp without time zone,
    elastic_gpu_id text
);


ALTER TABLE public.aws_ec2_instance_elastic_gpu_associations OWNER TO postgres;

--
-- Name: aws_ec2_instance_elastic_inference_accelerator_associations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_elastic_inference_accelerator_associations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    elastic_inference_accelerator_arn text,
    elastic_inference_accelerator_association_id text,
    elastic_inference_accelerator_association_state text,
    elastic_inference_accelerator_association_time timestamp without time zone
);


ALTER TABLE public.aws_ec2_instance_elastic_inference_accelerator_associations OWNER TO postgres;

--
-- Name: aws_ec2_instance_network_interface_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_network_interface_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_network_interface_cq_id uuid,
    network_interface_id text,
    group_id text,
    group_name text
);


ALTER TABLE public.aws_ec2_instance_network_interface_groups OWNER TO postgres;

--
-- Name: aws_ec2_instance_network_interface_ipv6_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_network_interface_ipv6_addresses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_network_interface_cq_id uuid,
    ipv6_address text
);


ALTER TABLE public.aws_ec2_instance_network_interface_ipv6_addresses OWNER TO postgres;

--
-- Name: aws_ec2_instance_network_interface_private_ip_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_network_interface_private_ip_addresses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_network_interface_cq_id uuid,
    association_carrier_ip text,
    association_ip_owner_id text,
    association_public_dns_name text,
    association_public_ip text,
    is_primary boolean,
    private_dns_name text,
    private_ip_address text
);


ALTER TABLE public.aws_ec2_instance_network_interface_private_ip_addresses OWNER TO postgres;

--
-- Name: aws_ec2_instance_network_interfaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_network_interfaces (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    arn text,
    association_carrier_ip text,
    association_ip_owner_id text,
    association_public_dns_name text,
    association_public_ip text,
    attachment_attach_time timestamp without time zone,
    attachment_id text,
    attachment_delete_on_termination boolean,
    attachment_device_index integer,
    attachment_network_card_index integer,
    attachment_status text,
    description text,
    interface_type text,
    ipv4_prefixes text[],
    ipv6_prefixes text[],
    mac_address text,
    network_interface_id text,
    owner_id text,
    private_dns_name text,
    private_ip_address text,
    source_dest_check boolean,
    status text,
    subnet_id text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_instance_network_interfaces OWNER TO postgres;

--
-- Name: aws_ec2_instance_product_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_product_codes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    product_code_id text,
    product_code_type text
);


ALTER TABLE public.aws_ec2_instance_product_codes OWNER TO postgres;

--
-- Name: aws_ec2_instance_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    group_id text,
    group_name text
);


ALTER TABLE public.aws_ec2_instance_security_groups OWNER TO postgres;

--
-- Name: aws_ec2_instance_status_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_status_events (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_status_cq_id uuid,
    code text,
    description text,
    id text,
    not_after timestamp without time zone,
    not_before timestamp without time zone,
    not_before_deadline timestamp without time zone
);


ALTER TABLE public.aws_ec2_instance_status_events OWNER TO postgres;

--
-- Name: aws_ec2_instance_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instance_statuses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    availability_zone text,
    instance_id text,
    instance_state_code integer,
    instance_state_name text,
    details jsonb,
    status text,
    outpost_arn text,
    system_status text,
    system_status_details jsonb
);


ALTER TABLE public.aws_ec2_instance_statuses OWNER TO postgres;

--
-- Name: aws_ec2_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    state_transition_reason_time timestamp without time zone,
    ami_launch_index integer,
    architecture text,
    boot_mode text,
    capacity_reservation_id text,
    cap_reservation_preference text,
    cap_reservation_target_capacity_reservation_id text,
    cap_reservation_target_capacity_reservation_rg_arn text,
    client_token text,
    cpu_options_core_count integer,
    cpu_options_threads_per_core integer,
    ebs_optimized boolean,
    ena_support boolean,
    enclave_options_enabled boolean,
    hibernation_options_configured boolean,
    hypervisor text,
    iam_instance_profile_arn text,
    iam_instance_profile_id text,
    image_id text,
    id text NOT NULL,
    instance_lifecycle text,
    instance_type text,
    kernel_id text,
    key_name text,
    launch_time timestamp without time zone,
    licenses text[],
    metadata_options_http_endpoint text,
    metadata_options_http_protocol_ipv6 text,
    metadata_options_http_put_response_hop_limit integer,
    metadata_options_http_tokens text,
    metadata_options_state text,
    monitoring_state text,
    outpost_arn text,
    placement_affinity text,
    placement_availability_zone text,
    placement_group_name text,
    placement_host_id text,
    placement_host_resource_group_arn text,
    placement_partition_number integer,
    placement_spread_domain text,
    placement_tenancy text,
    platform text,
    private_dns_name text,
    private_ip_address text,
    public_dns_name text,
    public_ip_address text,
    ramdisk_id text,
    root_device_name text,
    root_device_type text,
    source_dest_check boolean,
    spot_instance_request_id text,
    sriov_net_support text,
    state_code integer,
    state_name text,
    state_reason_code text,
    state_reason_message text,
    state_transition_reason text,
    subnet_id text,
    tags jsonb,
    virtualization_type text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_instances OWNER TO postgres;

--
-- Name: aws_ec2_internet_gateway_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_internet_gateway_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    internet_gateway_cq_id uuid,
    state text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_internet_gateway_attachments OWNER TO postgres;

--
-- Name: aws_ec2_internet_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_internet_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    owner_id text,
    tags jsonb
);


ALTER TABLE public.aws_ec2_internet_gateways OWNER TO postgres;

--
-- Name: aws_ec2_nat_gateway_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_nat_gateway_addresses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    nat_gateway_cq_id uuid,
    allocation_id text,
    network_interface_id text,
    private_ip text,
    public_ip text
);


ALTER TABLE public.aws_ec2_nat_gateway_addresses OWNER TO postgres;

--
-- Name: aws_ec2_nat_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_nat_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    create_time timestamp without time zone,
    delete_time timestamp without time zone,
    failure_code text,
    failure_message text,
    provisioned_bandwidth_provision_time timestamp without time zone,
    provisioned_bandwidth_provisioned text,
    provisioned_bandwidth_request_time timestamp without time zone,
    provisioned_bandwidth_requested text,
    provisioned_bandwidth_status text,
    state text,
    subnet_id text,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_ec2_nat_gateways OWNER TO postgres;

--
-- Name: aws_ec2_network_acl_associations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_network_acl_associations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    network_acl_cq_id uuid,
    network_acl_association_id text,
    subnet_id text
);


ALTER TABLE public.aws_ec2_network_acl_associations OWNER TO postgres;

--
-- Name: aws_ec2_network_acl_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_network_acl_entries (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    network_acl_cq_id uuid,
    cidr_block text,
    egress boolean,
    icmp_type_code integer,
    icmp_type_code_type integer,
    ipv6_cidr_block text,
    port_range_from integer,
    port_range_to integer,
    protocol text,
    rule_action text,
    rule_number integer
);


ALTER TABLE public.aws_ec2_network_acl_entries OWNER TO postgres;

--
-- Name: aws_ec2_network_acls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_network_acls (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    is_default boolean,
    id text NOT NULL,
    owner_id text,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_ec2_network_acls OWNER TO postgres;

--
-- Name: aws_ec2_network_interface_private_ip_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_network_interface_private_ip_addresses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    network_interface_cq_id uuid,
    association_allocation_id text,
    association_id text,
    association_carrier_ip text,
    association_customer_owned_ip text,
    association_ip_owner_id text,
    association_public_dns_name text,
    association_public_ip text,
    "primary" boolean,
    private_dns_name text,
    private_ip_address text
);


ALTER TABLE public.aws_ec2_network_interface_private_ip_addresses OWNER TO postgres;

--
-- Name: aws_ec2_network_interfaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_network_interfaces (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    tags jsonb,
    association_allocation_id text,
    association_id text,
    association_carrier_ip text,
    association_customer_owned_ip text,
    association_ip_owner_id text,
    association_public_dns_name text,
    association_public_ip text,
    attachment_attach_time timestamp without time zone,
    attachment_id text,
    attachment_delete_on_termination boolean,
    attachment_device_index integer,
    attachment_instance_id text,
    attachment_instance_owner_id text,
    attachment_network_card_index integer,
    attachment_status text,
    availability_zone text,
    deny_all_igw_traffic boolean,
    description text,
    groups jsonb,
    interface_type text,
    ipv4_prefixes text[],
    ipv6_address text,
    ipv6_addresses text[],
    ipv6_native boolean,
    ipv6_prefixes text[],
    mac_address text,
    id text,
    outpost_arn text,
    owner_id text,
    private_dns_name text,
    private_ip_address text,
    requester_id text,
    requester_managed boolean,
    source_dest_check boolean,
    status text,
    subnet_id text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_network_interfaces OWNER TO postgres;

--
-- Name: aws_ec2_regional_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_regional_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    ebs_encryption_enabled_by_default boolean,
    ebs_default_kms_key_id text
);


ALTER TABLE public.aws_ec2_regional_config OWNER TO postgres;

--
-- Name: aws_ec2_route_table_associations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_route_table_associations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    route_table_cq_id uuid,
    id text,
    association_state text,
    association_state_status_message text,
    gateway_id text,
    main boolean,
    subnet_id text
);


ALTER TABLE public.aws_ec2_route_table_associations OWNER TO postgres;

--
-- Name: aws_ec2_route_table_propagating_vgws; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_route_table_propagating_vgws (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    route_table_cq_id uuid,
    gateway_id text
);


ALTER TABLE public.aws_ec2_route_table_propagating_vgws OWNER TO postgres;

--
-- Name: aws_ec2_route_table_routes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_route_table_routes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    route_table_cq_id uuid,
    carrier_gateway_id text,
    destination_cidr_block text,
    destination_ipv6_cidr_block text,
    destination_prefix_list_id text,
    egress_only_internet_gateway_id text,
    gateway_id text,
    instance_id text,
    instance_owner_id text,
    local_gateway_id text,
    nat_gateway_id text,
    network_interface_id text,
    origin text,
    state text,
    transit_gateway_id text,
    vpc_peering_connection_id text
);


ALTER TABLE public.aws_ec2_route_table_routes OWNER TO postgres;

--
-- Name: aws_ec2_route_tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_route_tables (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    owner_id text,
    id text NOT NULL,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_ec2_route_tables OWNER TO postgres;

--
-- Name: aws_ec2_security_group_ip_permission_ip_ranges; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_security_group_ip_permission_ip_ranges (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    security_group_ip_permission_cq_id uuid,
    cidr text,
    description text,
    cidr_type text
);


ALTER TABLE public.aws_ec2_security_group_ip_permission_ip_ranges OWNER TO postgres;

--
-- Name: aws_ec2_security_group_ip_permission_prefix_list_ids; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_security_group_ip_permission_prefix_list_ids (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    security_group_ip_permission_cq_id uuid,
    description text,
    prefix_list_id text
);


ALTER TABLE public.aws_ec2_security_group_ip_permission_prefix_list_ids OWNER TO postgres;

--
-- Name: aws_ec2_security_group_ip_permission_user_id_group_pairs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_security_group_ip_permission_user_id_group_pairs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    security_group_ip_permission_cq_id uuid,
    description text,
    group_id text,
    group_name text,
    peering_status text,
    user_id text,
    vpc_id text,
    vpc_peering_connection_id text
);


ALTER TABLE public.aws_ec2_security_group_ip_permission_user_id_group_pairs OWNER TO postgres;

--
-- Name: aws_ec2_security_group_ip_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_security_group_ip_permissions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    security_group_cq_id uuid,
    from_port integer,
    ip_protocol text,
    to_port integer,
    permission_type text
);


ALTER TABLE public.aws_ec2_security_group_ip_permissions OWNER TO postgres;

--
-- Name: aws_ec2_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    description text,
    id text NOT NULL,
    group_name text,
    owner_id text,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_ec2_security_groups OWNER TO postgres;

--
-- Name: aws_ec2_subnet_ipv6_cidr_block_association_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_subnet_ipv6_cidr_block_association_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    subnet_cq_id uuid,
    association_id text,
    ipv6_cidr_block text,
    ipv6_cidr_block_state text,
    ipv6_cidr_block_state_status_message text
);


ALTER TABLE public.aws_ec2_subnet_ipv6_cidr_block_association_sets OWNER TO postgres;

--
-- Name: aws_ec2_subnets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_subnets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    assign_ipv6_address_on_creation boolean,
    availability_zone text,
    availability_zone_id text,
    available_ip_address_count integer,
    cidr_block text,
    customer_owned_ipv4_pool text,
    default_for_az boolean,
    map_customer_owned_ip_on_launch boolean,
    map_public_ip_on_launch boolean,
    outpost_arn text,
    owner_id text,
    state text,
    arn text,
    id text NOT NULL,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_ec2_subnets OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateway_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateway_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    transit_gateway_cq_id uuid,
    association_state text,
    association_route_table_id text,
    creation_time timestamp without time zone,
    resource_id text,
    resource_owner_id text,
    resource_type text,
    state text,
    tags jsonb,
    transit_gateway_owner_id text
);


ALTER TABLE public.aws_ec2_transit_gateway_attachments OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateway_multicast_domains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateway_multicast_domains (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    transit_gateway_cq_id uuid,
    creation_time timestamp without time zone,
    auto_accept_shared_associations text,
    igmpv2_support text,
    static_sources_support text,
    owner_id text,
    state text,
    tags jsonb,
    transit_gateway_multicast_domain_arn text,
    transit_gateway_multicast_domain_id text
);


ALTER TABLE public.aws_ec2_transit_gateway_multicast_domains OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateway_peering_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateway_peering_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    transit_gateway_cq_id uuid,
    accepter_owner_id text,
    accepter_region text,
    accepter_transit_gateway_id text,
    creation_time timestamp without time zone,
    requester_owner_id text,
    requester_region text,
    requester_transit_gateway_id text,
    state text,
    status_code text,
    status_message text,
    tags jsonb,
    transit_gateway_attachment_id text
);


ALTER TABLE public.aws_ec2_transit_gateway_peering_attachments OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateway_route_tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateway_route_tables (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    transit_gateway_cq_id uuid,
    creation_time timestamp without time zone,
    default_association_route_table boolean,
    default_propagation_route_table boolean,
    state text,
    tags jsonb,
    transit_gateway_route_table_id text
);


ALTER TABLE public.aws_ec2_transit_gateway_route_tables OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateway_vpc_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateway_vpc_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    transit_gateway_cq_id uuid,
    creation_time timestamp without time zone,
    appliance_mode_support text,
    dns_support text,
    ipv6_support text,
    state text,
    tags jsonb,
    transit_gateway_attachment_id text,
    vpc_id text,
    vpc_owner_id text
);


ALTER TABLE public.aws_ec2_transit_gateway_vpc_attachments OWNER TO postgres;

--
-- Name: aws_ec2_transit_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_transit_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    amazon_side_asn bigint,
    association_default_route_table_id text,
    auto_accept_shared_attachments text,
    creation_time timestamp without time zone,
    default_route_table_association text,
    default_route_table_propagation text,
    description text,
    dns_support text,
    multicast_support text,
    owner_id text,
    propagation_default_route_table_id text,
    state text,
    tags jsonb,
    arn text,
    transit_gateway_cidr_blocks text[],
    id text NOT NULL,
    vpn_ecmp_support text
);


ALTER TABLE public.aws_ec2_transit_gateways OWNER TO postgres;

--
-- Name: aws_ec2_vpc_attachment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_attachment (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vpn_gateway_cq_id uuid,
    state text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_vpc_attachment OWNER TO postgres;

--
-- Name: aws_ec2_vpc_cidr_block_association_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_cidr_block_association_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vpc_cq_id uuid,
    association_id text,
    cidr_block text,
    cidr_block_state text,
    cidr_block_state_status_message text
);


ALTER TABLE public.aws_ec2_vpc_cidr_block_association_sets OWNER TO postgres;

--
-- Name: aws_ec2_vpc_endpoint_dns_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_endpoint_dns_entries (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vpc_endpoint_cq_id uuid,
    dns_name text,
    hosted_zone_id text
);


ALTER TABLE public.aws_ec2_vpc_endpoint_dns_entries OWNER TO postgres;

--
-- Name: aws_ec2_vpc_endpoint_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_endpoint_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vpc_endpoint_cq_id uuid,
    group_id text,
    group_name text
);


ALTER TABLE public.aws_ec2_vpc_endpoint_groups OWNER TO postgres;

--
-- Name: aws_ec2_vpc_endpoints; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_endpoints (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    creation_timestamp timestamp without time zone,
    last_error_code text,
    last_error_message text,
    network_interface_ids text[],
    owner_id text,
    policy_document text,
    private_dns_enabled boolean,
    requester_managed boolean,
    route_table_ids text[],
    service_name text,
    state text,
    subnet_ids text[],
    tags jsonb,
    id text NOT NULL,
    vpc_endpoint_type text,
    vpc_id text
);


ALTER TABLE public.aws_ec2_vpc_endpoints OWNER TO postgres;

--
-- Name: aws_ec2_vpc_ipv6_cidr_block_association_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_ipv6_cidr_block_association_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    vpc_cq_id uuid,
    association_id text,
    ipv6_cidr_block text,
    ipv6_cidr_block_state text,
    ipv6_cidr_block_state_status_message text,
    ipv6_pool text,
    network_border_group text
);


ALTER TABLE public.aws_ec2_vpc_ipv6_cidr_block_association_sets OWNER TO postgres;

--
-- Name: aws_ec2_vpc_peering_connections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpc_peering_connections (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    accepter_cidr_block text,
    accepter_cidr_block_set text[],
    accepter_ipv6_cidr_block_set text[],
    accepter_owner_id text,
    accepter_allow_dns_resolution_from_remote_vpc boolean,
    accepter_allow_egress_local_classic_link_to_remote_vpc boolean,
    accepter_allow_egress_local_vpc_to_remote_classic_link boolean,
    accepter_vpc_region text,
    accepter_vpc_id text,
    expiration_time timestamp without time zone,
    requester_cidr_block text,
    requester_cidr_block_set text[],
    requester_ipv6_cidr_block_set text[],
    requester_owner_id text,
    requester_allow_dns_resolution_from_remote_vpc boolean,
    requester_allow_egress_local_classic_link_to_remote_vpc boolean,
    requester_allow_egress_local_vpc_to_remote_classic_link boolean,
    requester_vpc_region text,
    requester_vpc_id text,
    status_code text,
    status_message text,
    tags jsonb,
    id text NOT NULL
);


ALTER TABLE public.aws_ec2_vpc_peering_connections OWNER TO postgres;

--
-- Name: aws_ec2_vpcs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpcs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    cidr_block text,
    dhcp_options_id text,
    instance_tenancy text,
    is_default boolean,
    owner_id text,
    state text,
    tags jsonb,
    id text NOT NULL
);


ALTER TABLE public.aws_ec2_vpcs OWNER TO postgres;

--
-- Name: aws_ec2_vpn_gateways; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ec2_vpn_gateways (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    amazon_side_asn bigint,
    availability_zone text,
    state text,
    tags jsonb,
    type text,
    id text NOT NULL
);


ALTER TABLE public.aws_ec2_vpn_gateways OWNER TO postgres;

--
-- Name: aws_ecr_repositories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecr_repositories (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    created_at timestamp without time zone,
    encryption_configuration_encryption_type text,
    encryption_configuration_kms_key text,
    image_scanning_configuration_scan_on_push boolean,
    image_tag_mutability text,
    registry_id text,
    arn text NOT NULL,
    name text,
    uri text
);


ALTER TABLE public.aws_ecr_repositories OWNER TO postgres;

--
-- Name: aws_ecr_repository_images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecr_repository_images (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    repository_cq_id uuid,
    account_id text,
    region text,
    artifact_media_type text,
    image_digest text,
    image_manifest_media_type text,
    image_pushed_at timestamp without time zone,
    image_scan_findings_summary_finding_severity_counts jsonb,
    image_scan_findings_summary_image_scan_completed_at timestamp without time zone,
    image_scan_findings_summary_vulnerability_source_updated_at timestamp without time zone,
    image_scan_status_description text,
    image_scan_status text,
    image_size_in_bytes bigint,
    image_tags text[],
    registry_id text,
    repository_name text
);


ALTER TABLE public.aws_ecr_repository_images OWNER TO postgres;

--
-- Name: aws_ecs_cluster_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    details jsonb,
    id text,
    status text,
    type text
);


ALTER TABLE public.aws_ecs_cluster_attachments OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instance_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instance_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_container_instance_cq_id uuid,
    details jsonb,
    id text,
    status text,
    type text
);


ALTER TABLE public.aws_ecs_cluster_container_instance_attachments OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instance_attributes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instance_attributes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_container_instance_cq_id uuid,
    name text,
    target_id text,
    target_type text,
    value text
);


ALTER TABLE public.aws_ecs_cluster_container_instance_attributes OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instance_health_status_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instance_health_status_details (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_container_instance_cq_id uuid,
    last_status_change timestamp without time zone,
    last_updated timestamp without time zone,
    status text,
    type text
);


ALTER TABLE public.aws_ecs_cluster_container_instance_health_status_details OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instance_registered_resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instance_registered_resources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_container_instance_cq_id uuid,
    double_value double precision,
    integer_value integer,
    long_value bigint,
    name text,
    string_set_value text[],
    type text
);


ALTER TABLE public.aws_ecs_cluster_container_instance_registered_resources OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instance_remaining_resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instance_remaining_resources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_container_instance_cq_id uuid,
    double_value double precision,
    integer_value integer,
    long_value bigint,
    name text,
    string_set_value text[],
    type text
);


ALTER TABLE public.aws_ecs_cluster_container_instance_remaining_resources OWNER TO postgres;

--
-- Name: aws_ecs_cluster_container_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_container_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    agent_connected boolean,
    agent_update_status text,
    capacity_provider_name text,
    container_instance_arn text,
    ec2_instance_id text,
    health_status_overall_status text,
    pending_tasks_count integer,
    registered_at timestamp without time zone,
    running_tasks_count integer,
    status text,
    status_reason text,
    tags jsonb,
    version bigint,
    version_info_agent_hash text,
    version_info_agent_version text,
    version_info_docker_version text
);


ALTER TABLE public.aws_ecs_cluster_container_instances OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_deployments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_deployments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_cq_id uuid,
    capacity_provider_strategy jsonb,
    created_at timestamp without time zone,
    desired_count integer,
    failed_tasks integer,
    id text,
    launch_type text,
    network_configuration_awsvpc_configuration_subnets text[],
    network_configuration_awsvpc_configuration_assign_public_ip text,
    network_configuration_awsvpc_configuration_security_groups text[],
    pending_count integer,
    platform_family text,
    platform_version text,
    rollout_state text,
    rollout_state_reason text,
    running_count integer,
    status text,
    task_definition text,
    updated_at timestamp without time zone
);


ALTER TABLE public.aws_ecs_cluster_service_deployments OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_events (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_cq_id uuid,
    created_at timestamp without time zone,
    id text,
    message text
);


ALTER TABLE public.aws_ecs_cluster_service_events OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_load_balancers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_load_balancers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_cq_id uuid,
    container_name text,
    container_port integer,
    load_balancer_name text,
    target_group_arn text
);


ALTER TABLE public.aws_ecs_cluster_service_load_balancers OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_service_registries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_service_registries (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_cq_id uuid,
    container_name text,
    container_port integer,
    port integer,
    registry_arn text
);


ALTER TABLE public.aws_ecs_cluster_service_service_registries OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_task_set_load_balancers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_task_set_load_balancers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_task_set_cq_id uuid,
    container_name text,
    container_port integer,
    load_balancer_name text,
    target_group_arn text
);


ALTER TABLE public.aws_ecs_cluster_service_task_set_load_balancers OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_task_set_service_registries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_task_set_service_registries (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_task_set_cq_id uuid,
    container_name text,
    container_port integer,
    port integer,
    arn text
);


ALTER TABLE public.aws_ecs_cluster_service_task_set_service_registries OWNER TO postgres;

--
-- Name: aws_ecs_cluster_service_task_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_service_task_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_service_cq_id uuid,
    capacity_provider_strategy jsonb,
    cluster_arn text,
    computed_desired_count integer,
    created_at timestamp without time zone,
    external_id text,
    id text,
    launch_type text,
    network_configuration_awsvpc_configuration_subnets text[],
    network_configuration_awsvpc_configuration_assign_public_ip text,
    network_configuration_awsvpc_configuration_security_groups text[],
    pending_count integer,
    platform_family text,
    platform_version text,
    running_count integer,
    scale_unit text,
    scale_value double precision,
    service_arn text,
    stability_status text,
    stability_status_at timestamp without time zone,
    started_by text,
    status text,
    tags jsonb,
    task_definition text,
    arn text,
    updated_at timestamp without time zone
);


ALTER TABLE public.aws_ecs_cluster_service_task_sets OWNER TO postgres;

--
-- Name: aws_ecs_cluster_services; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_services (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    capacity_provider_strategy jsonb,
    cluster_arn text,
    created_at timestamp without time zone,
    created_by text,
    deployment_configuration_deployment_circuit_breaker_enable boolean,
    deployment_configuration_deployment_circuit_breaker_rollback boolean,
    deployment_configuration_maximum_percent integer,
    deployment_configuration_minimum_healthy_percent integer,
    deployment_controller_type text,
    desired_count integer,
    enable_ecs_managed_tags boolean,
    enable_execute_command boolean,
    health_check_grace_period_seconds integer,
    launch_type text,
    network_configuration_awsvpc_configuration_subnets text[],
    network_configuration_awsvpc_configuration_assign_public_ip text,
    network_configuration_awsvpc_configuration_security_groups text[],
    pending_count integer,
    placement_constraints jsonb,
    placement_strategy jsonb,
    platform_family text,
    platform_version text,
    propagate_tags text,
    role_arn text,
    running_count integer,
    scheduling_strategy text,
    arn text,
    name text,
    status text,
    tags jsonb,
    task_definition text
);


ALTER TABLE public.aws_ecs_cluster_services OWNER TO postgres;

--
-- Name: aws_ecs_cluster_task_attachments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_task_attachments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_task_cq_id uuid,
    details jsonb,
    id text,
    status text,
    type text
);


ALTER TABLE public.aws_ecs_cluster_task_attachments OWNER TO postgres;

--
-- Name: aws_ecs_cluster_task_containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_task_containers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_task_cq_id uuid,
    container_arn text,
    cpu text,
    exit_code integer,
    gpu_ids text[],
    health_status text,
    image text,
    image_digest text,
    last_status text,
    managed_agents jsonb,
    memory text,
    memory_reservation text,
    name text,
    network_bindings jsonb,
    network_interfaces jsonb,
    reason text,
    runtime_id text,
    task_arn text
);


ALTER TABLE public.aws_ecs_cluster_task_containers OWNER TO postgres;

--
-- Name: aws_ecs_cluster_tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_cluster_tasks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    attributes jsonb,
    availability_zone text,
    capacity_provider_name text,
    cluster_arn text,
    connectivity text,
    connectivity_at timestamp without time zone,
    container_instance_arn text,
    cpu text,
    created_at timestamp without time zone,
    desired_status text,
    enable_execute_command boolean,
    ephemeral_storage_size_in_gib integer,
    execution_stopped_at timestamp without time zone,
    "group" text,
    health_status text,
    inference_accelerators jsonb,
    last_status text,
    launch_type text,
    memory text,
    overrides jsonb,
    platform_family text,
    platform_version text,
    pull_started_at timestamp without time zone,
    pull_stopped_at timestamp without time zone,
    started_at timestamp without time zone,
    started_by text,
    stop_code text,
    stopped_at timestamp without time zone,
    stopped_reason text,
    stopping_at timestamp without time zone,
    tags jsonb,
    arn text,
    task_definition_arn text,
    version bigint
);


ALTER TABLE public.aws_ecs_cluster_tasks OWNER TO postgres;

--
-- Name: aws_ecs_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    active_services_count integer,
    attachments_status text,
    capacity_providers text[],
    arn text NOT NULL,
    name text,
    execute_config_kms_key_id text,
    execute_config_logs_cloud_watch_encryption_enabled boolean,
    execute_config_log_cloud_watch_log_group_name text,
    execute_config_log_s3_bucket_name text,
    execute_config_log_s3_encryption_enabled boolean,
    execute_config_log_s3_key_prefix text,
    execute_config_logging text,
    default_capacity_provider_strategy jsonb,
    pending_tasks_count integer,
    registered_container_instances_count integer,
    running_tasks_count integer,
    settings jsonb,
    statistics jsonb,
    status text,
    tags jsonb
);


ALTER TABLE public.aws_ecs_clusters OWNER TO postgres;

--
-- Name: aws_ecs_task_definition_container_definitions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_task_definition_container_definitions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    task_definition_cq_id uuid,
    command text[],
    cpu integer,
    depends_on jsonb,
    disable_networking boolean,
    dns_search_domains text[],
    dns_servers text[],
    docker_labels jsonb,
    docker_security_options text[],
    entry_point text[],
    environment jsonb,
    environment_files jsonb,
    essential boolean,
    extra_hosts jsonb,
    firelens_configuration_type text,
    firelens_configuration_options jsonb,
    health_check_command text[],
    health_check_interval integer,
    health_check_retries integer,
    health_check_start_period integer,
    health_check_timeout integer,
    hostname text,
    image text,
    interactive boolean,
    links text[],
    linux_parameters_capabilities_add text[],
    linux_parameters_capabilities_drop text[],
    linux_parameters_devices jsonb,
    linux_parameters_init_process_enabled boolean,
    linux_parameters_max_swap integer,
    linux_parameters_shared_memory_size integer,
    linux_parameters_swappiness integer,
    linux_parameters_tmpfs jsonb,
    log_configuration_log_driver text,
    log_configuration_options jsonb,
    log_configuration_secret_options jsonb,
    memory integer,
    memory_reservation integer,
    mount_points jsonb,
    name text,
    port_mappings jsonb,
    privileged boolean,
    pseudo_terminal boolean,
    readonly_root_filesystem boolean,
    repository_credentials_parameter text,
    resource_requirements jsonb,
    secrets jsonb,
    start_timeout integer,
    stop_timeout integer,
    system_controls jsonb,
    ulimits jsonb,
    "user" text,
    volumes_from jsonb,
    working_directory text
);


ALTER TABLE public.aws_ecs_task_definition_container_definitions OWNER TO postgres;

--
-- Name: aws_ecs_task_definition_volumes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_task_definition_volumes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    task_definition_cq_id uuid,
    docker_autoprovision boolean,
    docker_driver text,
    docker_driver_opts jsonb,
    docker_labels jsonb,
    docker_scope text,
    efs_file_system_id text,
    efs_authorization_config_access_point_id text,
    efs_authorization_config_iam text,
    efs_root_directory text,
    efs_volume_configuration_transit_encryption text,
    efs_transit_encryption_port integer,
    fsx_wfs_authorization_config_credentials_parameter text,
    fsx_wfs_authorization_config_domain text,
    fsx_wfs_file_system_id text,
    fsx_wfs_root_directory text,
    host_source_path text,
    name text
);


ALTER TABLE public.aws_ecs_task_definition_volumes OWNER TO postgres;

--
-- Name: aws_ecs_task_definitions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ecs_task_definitions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    compatibilities text[],
    cpu text,
    deregistered_at timestamp without time zone,
    ephemeral_storage_size integer,
    execution_role_arn text,
    family text,
    inference_accelerators jsonb,
    ipc_mode text,
    memory text,
    network_mode text,
    pid_mode text,
    placement_constraints jsonb,
    proxy_configuration_container_name text,
    proxy_configuration_properties jsonb,
    proxy_configuration_type text,
    registered_at timestamp without time zone,
    registered_by text,
    requires_attributes jsonb,
    requires_compatibilities text[],
    revision integer,
    runtime_platform_cpu_architecture text,
    runtime_platform_os_family text,
    status text,
    arn text NOT NULL,
    task_role_arn text
);


ALTER TABLE public.aws_ecs_task_definitions OWNER TO postgres;

--
-- Name: aws_efs_filesystems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_efs_filesystems (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    backup_policy_status text,
    creation_time timestamp without time zone,
    creation_token text,
    id text NOT NULL,
    life_cycle_state text,
    number_of_mount_targets integer,
    owner_id text,
    performance_mode text,
    size_in_bytes_value bigint,
    size_in_bytes_timestamp timestamp without time zone,
    size_in_bytes_value_in_ia bigint,
    size_in_bytes_value_in_standard bigint,
    tags jsonb,
    availability_zone_id text,
    availability_zone_name text,
    encrypted boolean,
    arn text,
    kms_key_id text,
    name text,
    provisioned_throughput_in_mibps double precision,
    throughput_mode text
);


ALTER TABLE public.aws_efs_filesystems OWNER TO postgres;

--
-- Name: aws_eks_cluster_encryption_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_eks_cluster_encryption_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    provider_key_arn text,
    resources text[]
);


ALTER TABLE public.aws_eks_cluster_encryption_configs OWNER TO postgres;

--
-- Name: aws_eks_cluster_loggings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_eks_cluster_loggings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    enabled boolean,
    types text[]
);


ALTER TABLE public.aws_eks_cluster_loggings OWNER TO postgres;

--
-- Name: aws_eks_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_eks_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    certificate_authority_data text,
    client_request_token text,
    created_at timestamp without time zone,
    endpoint text,
    identity_oidc_issuer text,
    kubernetes_network_config_service_ipv4_cidr text,
    name text,
    platform_version text,
    resources_vpc_config_cluster_security_group_id text,
    resources_vpc_config_endpoint_private_access boolean,
    resources_vpc_config_endpoint_public_access boolean,
    resources_vpc_config_public_access_cidrs text[],
    resources_vpc_config_security_group_ids text[],
    resources_vpc_config_subnet_ids text[],
    resources_vpc_config_vpc_id text,
    role_arn text,
    status text,
    tags jsonb,
    version text
);


ALTER TABLE public.aws_eks_clusters OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_application_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_application_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    application_name text,
    arn text NOT NULL,
    build_arn text,
    date_created timestamp without time zone,
    date_updated timestamp without time zone,
    description text,
    source_location text,
    source_repository text,
    source_type text,
    source_bundle_s3_bucket text,
    source_bundle_s3_key text,
    status text,
    version_label text
);


ALTER TABLE public.aws_elasticbeanstalk_application_versions OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_applications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_applications (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    name text,
    configuration_templates text[],
    date_created timestamp without time zone NOT NULL,
    date_updated timestamp without time zone,
    description text,
    resource_lifecycle_config_service_role text,
    max_age_rule_enabled boolean,
    max_age_rule_delete_source_from_s3 boolean,
    max_age_rule_max_age_in_days integer,
    max_count_rule_enabled boolean,
    max_count_rule_delete_source_from_s3 boolean,
    max_count_rule_max_count integer,
    versions text[]
);


ALTER TABLE public.aws_elasticbeanstalk_applications OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_configuration_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_configuration_options (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    environment_cq_id uuid,
    application_arn text,
    name text,
    namespace text,
    change_severity text,
    default_value text,
    max_length integer,
    max_value integer,
    min_value integer,
    regex_label text,
    regex_pattern text,
    user_defined boolean,
    value_options text[],
    value_type text
);


ALTER TABLE public.aws_elasticbeanstalk_configuration_options OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_configuration_setting_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_configuration_setting_options (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    configuration_setting_cq_id uuid,
    namespace text,
    option_name text,
    resource_name text,
    value text
);


ALTER TABLE public.aws_elasticbeanstalk_configuration_setting_options OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_configuration_settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_configuration_settings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    environment_cq_id uuid,
    application_name text,
    application_arn text,
    date_created timestamp without time zone,
    date_updated timestamp without time zone,
    deployment_status text,
    description text,
    environment_name text,
    platform_arn text,
    solution_stack_name text,
    template_name text
);


ALTER TABLE public.aws_elasticbeanstalk_configuration_settings OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_environment_links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_environment_links (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    environment_cq_id uuid,
    environment_name text,
    link_name text
);


ALTER TABLE public.aws_elasticbeanstalk_environment_links OWNER TO postgres;

--
-- Name: aws_elasticbeanstalk_environments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticbeanstalk_environments (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    tags jsonb,
    abortable_operation_in_progress boolean,
    application_name text,
    cname text,
    date_created timestamp without time zone,
    date_updated timestamp without time zone,
    description text,
    endpoint_url text,
    arn text,
    id text NOT NULL,
    name text,
    health text,
    health_status text,
    operations_role text,
    platform_arn text,
    load_balancer_domain text,
    listeners jsonb,
    load_balancer_name text,
    solution_stack_name text,
    status text,
    template_name text,
    tier_name text,
    tier_type text,
    tier_version text,
    version_label text
);


ALTER TABLE public.aws_elasticbeanstalk_environments OWNER TO postgres;

--
-- Name: aws_elasticsearch_domains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elasticsearch_domains (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    tags jsonb,
    arn text,
    id text NOT NULL,
    name text,
    cluster_cold_storage_options_enabled boolean,
    cluster_dedicated_master_count integer,
    cluster_dedicated_master_enabled boolean,
    cluster_dedicated_master_type text,
    cluster_instance_count integer,
    cluster_instance_type text,
    cluster_warm_count integer,
    cluster_warm_enabled boolean,
    cluster_warm_type text,
    cluster_zone_awareness_config_availability_zone_count integer,
    cluster_zone_awareness_enabled boolean,
    access_policies text,
    advanced_options jsonb,
    advanced_security_enabled boolean,
    advanced_security_internal_user_database_enabled boolean,
    advanced_security_saml_enabled boolean,
    advanced_security_saml_idp_entity_id text,
    advanced_security_saml_roles_key text,
    advanced_security_options_saml_options_roles_key text,
    advanced_security_saml_session_timeout_minutes integer,
    advanced_security_saml_subject_key text,
    auto_tune_error_message text,
    auto_tune_options_state text,
    cognito_enabled boolean,
    cognito_identity_pool_id text,
    cognito_role_arn text,
    cognito_user_pool_id text,
    created boolean,
    deleted boolean,
    domain_endpoint_custom text,
    domain_endpoint_custom_certificate_arn text,
    domain_endpoint_custom_enabled boolean,
    domain_endpoint_enforce_https boolean,
    domain_endpoint_tls_security_policy text,
    ebs_enabled boolean,
    ebs_iops integer,
    ebs_volume_size integer,
    ebs_volume_type text,
    elasticsearch_version text,
    encryption_at_rest_enabled boolean,
    encryption_at_rest_kms_key_id text,
    endpoint text,
    endpoints jsonb,
    log_publishing_options jsonb,
    node_to_node_encryption_enabled boolean,
    processing boolean,
    service_software_automated_update_date timestamp without time zone,
    service_software_cancellable boolean,
    service_software_current_version text,
    service_software_description text,
    service_software_new_version text,
    service_software_optional_deployment boolean,
    service_software_update_available boolean,
    service_software_update_status text,
    snapshot_options_automated_snapshot_start_hour integer,
    upgrade_processing boolean,
    vpc_availability_zones text[],
    vpc_security_group_ids text[],
    vpc_subnet_ids text[],
    vpc_vpc_id text
);


ALTER TABLE public.aws_elasticsearch_domains OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancer_backend_server_descriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancer_backend_server_descriptions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    name text,
    instance_port integer,
    policy_names text[]
);


ALTER TABLE public.aws_elbv1_load_balancer_backend_server_descriptions OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancer_listeners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancer_listeners (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    load_balance_name text,
    listener_instance_port integer,
    listener_load_balancer_port integer,
    listener_protocol text,
    listener_instance_protocol text,
    listener_ssl_certificate_id text,
    policy_names text[]
);


ALTER TABLE public.aws_elbv1_load_balancer_listeners OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancer_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancer_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    load_balance_name text,
    policy_attribute_descriptions jsonb,
    policy_name text,
    policy_type_name text
);


ALTER TABLE public.aws_elbv1_load_balancer_policies OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancer_policies_app_cookie_stickiness; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancer_policies_app_cookie_stickiness (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    load_balance_name text,
    cookie_name text,
    policy_name text
);


ALTER TABLE public.aws_elbv1_load_balancer_policies_app_cookie_stickiness OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancer_policies_lb_cookie_stickiness; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancer_policies_lb_cookie_stickiness (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    load_balance_name text,
    cookie_expiration_period bigint,
    policy_name text
);


ALTER TABLE public.aws_elbv1_load_balancer_policies_lb_cookie_stickiness OWNER TO postgres;

--
-- Name: aws_elbv1_load_balancers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv1_load_balancers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    attributes_access_log_enabled boolean,
    attributes_access_log_s3_bucket_name text,
    attributes_access_log_s3_bucket_prefix text,
    attributes_access_log_emit_interval integer,
    attributes_connection_settings_idle_timeout integer,
    attributes_cross_zone_load_balancing_enabled boolean,
    attributes_connection_draining_enabled boolean,
    attributes_connection_draining_timeout integer,
    attributes_additional_attributes jsonb,
    tags jsonb,
    availability_zones text[],
    canonical_hosted_zone_name text,
    canonical_hosted_zone_name_id text,
    created_time timestamp without time zone,
    dns_name text,
    health_check_healthy_threshold integer,
    health_check_interval integer,
    health_check_target text,
    health_check_timeout integer,
    health_check_unhealthy_threshold integer,
    instances text[],
    name text NOT NULL,
    other_policies text[],
    scheme text,
    security_groups text[],
    source_security_group_name text,
    source_security_group_owner_alias text,
    subnets text[],
    vpc_id text
);


ALTER TABLE public.aws_elbv1_load_balancers OWNER TO postgres;

--
-- Name: aws_elbv2_listener_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_listener_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    listener_cq_id uuid,
    certificate_arn text,
    is_default boolean
);


ALTER TABLE public.aws_elbv2_listener_certificates OWNER TO postgres;

--
-- Name: aws_elbv2_listener_default_action_forward_config_target_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_listener_default_action_forward_config_target_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    listener_default_action_cq_id uuid,
    target_group_arn text,
    weight integer
);


ALTER TABLE public.aws_elbv2_listener_default_action_forward_config_target_groups OWNER TO postgres;

--
-- Name: aws_elbv2_listener_default_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_listener_default_actions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    listener_cq_id uuid,
    type text,
    auth_cognito_user_pool_arn text,
    auth_cognito_user_pool_client_id text,
    auth_cognito_user_pool_domain text,
    auth_cognito_authentication_request_extra_params jsonb,
    auth_cognito_on_unauthenticated_request text,
    auth_cognito_scope text,
    auth_cognito_session_cookie_name text,
    auth_cognito_session_timeout bigint,
    auth_oidc_authorization_endpoint text,
    auth_oidc_client_id text,
    auth_oidc_issuer text,
    auth_oidc_token_endpoint text,
    auth_oidc_user_info_endpoint text,
    auth_oidc_authentication_request_extra_params jsonb,
    auth_oidc_client_secret text,
    auth_oidc_on_unauthenticated_request text,
    auth_oidc_scope text,
    auth_oidc_session_cookie_name text,
    auth_oidc_session_timeout bigint,
    auth_oidc_use_existing_client_secret boolean,
    fixed_response_config_status_code text,
    fixed_response_config_content_type text,
    fixed_response_config_message_body text,
    forward_config_target_group_stickiness_config_duration_seconds integer,
    forward_config_target_group_stickiness_config_enabled boolean,
    "order" integer,
    redirect_config_status_code text,
    redirect_config_host text,
    redirect_config_path text,
    redirect_config_port text,
    redirect_config_protocol text,
    redirect_config_query text,
    target_group_arn text
);


ALTER TABLE public.aws_elbv2_listener_default_actions OWNER TO postgres;

--
-- Name: aws_elbv2_listeners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_listeners (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    load_balancer_cq_id uuid,
    tags jsonb,
    alpn_policy text[],
    arn text,
    load_balancer_arn text,
    port integer,
    protocol text,
    ssl_policy text
);


ALTER TABLE public.aws_elbv2_listeners OWNER TO postgres;

--
-- Name: aws_elbv2_load_balancer_attributes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_load_balancer_attributes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    access_logs_s3_enabled boolean,
    access_logs_s3_bucket text,
    access_logs_s3_prefix text,
    deletion_protection boolean,
    idle_timeout integer,
    routing_http_desync_mitigation_mode text,
    routing_http_drop_invalid_header_fields boolean,
    routing_http_xamzntls_enabled boolean,
    routing_http_xff_client_port boolean,
    routing_http2 boolean,
    waf_fail_open boolean,
    load_balancing_cross_zone boolean
);


ALTER TABLE public.aws_elbv2_load_balancer_attributes OWNER TO postgres;

--
-- Name: aws_elbv2_load_balancer_availability_zone_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_load_balancer_availability_zone_addresses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_availability_zone_cq_id uuid,
    zone_name text,
    allocation_id text,
    ipv6_address text,
    ip_address text,
    private_ipv4_address text
);


ALTER TABLE public.aws_elbv2_load_balancer_availability_zone_addresses OWNER TO postgres;

--
-- Name: aws_elbv2_load_balancer_availability_zones; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_load_balancer_availability_zones (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    load_balancer_cq_id uuid,
    load_balance_name text,
    outpost_id text,
    subnet_id text,
    zone_name text
);


ALTER TABLE public.aws_elbv2_load_balancer_availability_zones OWNER TO postgres;

--
-- Name: aws_elbv2_load_balancers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_load_balancers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    web_acl_arn text,
    tags jsonb,
    canonical_hosted_zone_id text,
    created_time timestamp without time zone,
    customer_owned_ipv4_pool text,
    dns_name text,
    ip_address_type text,
    arn text NOT NULL,
    name text,
    scheme text,
    security_groups text[],
    state_code text,
    state_reason text,
    type text,
    vpc_id text
);


ALTER TABLE public.aws_elbv2_load_balancers OWNER TO postgres;

--
-- Name: aws_elbv2_target_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_elbv2_target_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    health_check_enabled boolean,
    health_check_interval_seconds integer,
    health_check_path text,
    health_check_port text,
    health_check_protocol text,
    health_check_timeout_seconds integer,
    healthy_threshold_count integer,
    load_balancer_arns text[],
    matcher_grpc_code text,
    matcher_http_code text,
    port integer,
    protocol text,
    protocol_version text,
    arn text NOT NULL,
    name text,
    target_type text,
    unhealthy_threshold_count integer,
    vpc_id text
);


ALTER TABLE public.aws_elbv2_target_groups OWNER TO postgres;

--
-- Name: aws_emr_block_public_access_config_port_ranges; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_emr_block_public_access_config_port_ranges (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    block_public_access_config_cq_id uuid,
    min_range integer,
    max_range integer
);


ALTER TABLE public.aws_emr_block_public_access_config_port_ranges OWNER TO postgres;

--
-- Name: aws_emr_block_public_access_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_emr_block_public_access_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    block_public_security_group_rules boolean,
    classification text,
    configurations jsonb,
    properties jsonb,
    created_by_arn text,
    creation_date_time timestamp without time zone
);


ALTER TABLE public.aws_emr_block_public_access_configs OWNER TO postgres;

--
-- Name: aws_emr_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_emr_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
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


ALTER TABLE public.aws_emr_clusters OWNER TO postgres;

--
-- Name: aws_fsx_backups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_fsx_backups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    id text NOT NULL,
    creation_time timestamp without time zone,
    lifecycle text,
    type text,
    directory_information_active_directory_id text,
    directory_information_domain_name text,
    failure_details_message text,
    kms_key_id text,
    progress_percent integer,
    arn text,
    tags jsonb
);


ALTER TABLE public.aws_fsx_backups OWNER TO postgres;

--
-- Name: aws_guardduty_detector_members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_guardduty_detector_members (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    detector_cq_id uuid,
    account_id text,
    email text,
    master_id text,
    relationship_status text,
    updated_at timestamp without time zone,
    detector_id text,
    invited_at timestamp without time zone
);


ALTER TABLE public.aws_guardduty_detector_members OWNER TO postgres;

--
-- Name: aws_guardduty_detectors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_guardduty_detectors (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    id text NOT NULL,
    service_role text,
    status text,
    created_at timestamp without time zone,
    data_sources_cloud_trail_status text,
    data_sources_dns_logs_status text,
    data_sources_flow_logs_status text,
    data_sources_s3_logs_status text,
    finding_publishing_frequency text,
    tags jsonb,
    updated_at timestamp without time zone
);


ALTER TABLE public.aws_guardduty_detectors OWNER TO postgres;

--
-- Name: aws_iam_group_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_group_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    group_cq_id uuid,
    group_id text,
    group_name text,
    policy_document jsonb,
    policy_name text
);


ALTER TABLE public.aws_iam_group_policies OWNER TO postgres;

--
-- Name: aws_iam_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    policies jsonb,
    arn text,
    create_date timestamp without time zone,
    id text NOT NULL,
    name text,
    path text
);


ALTER TABLE public.aws_iam_groups OWNER TO postgres;

--
-- Name: aws_iam_openid_connect_identity_providers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_openid_connect_identity_providers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    arn text NOT NULL,
    client_id_list text[],
    create_date timestamp without time zone,
    tags jsonb,
    thumbprint_list text[],
    url text
);


ALTER TABLE public.aws_iam_openid_connect_identity_providers OWNER TO postgres;

--
-- Name: aws_iam_password_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_password_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    allow_users_to_change_password boolean,
    expire_passwords boolean,
    hard_expiry boolean,
    max_password_age integer,
    minimum_password_length integer,
    password_reuse_prevention integer,
    require_lowercase_characters boolean,
    require_numbers boolean,
    require_symbols boolean,
    require_uppercase_characters boolean,
    policy_exists boolean
);


ALTER TABLE public.aws_iam_password_policies OWNER TO postgres;

--
-- Name: aws_iam_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    attachment_count integer,
    create_date timestamp without time zone,
    default_version_id text,
    description text,
    is_attachable boolean,
    path text,
    permissions_boundary_usage_count integer,
    id text NOT NULL,
    name text,
    update_date timestamp without time zone
);


ALTER TABLE public.aws_iam_policies OWNER TO postgres;

--
-- Name: aws_iam_policy_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_policy_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    policy_cq_id uuid,
    policy_id text,
    create_date timestamp without time zone,
    document jsonb,
    is_default_version boolean,
    version_id text
);


ALTER TABLE public.aws_iam_policy_versions OWNER TO postgres;

--
-- Name: aws_iam_role_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_role_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    role_cq_id uuid,
    role_id text,
    account_id text,
    policy_document jsonb,
    policy_name text,
    role_name text
);


ALTER TABLE public.aws_iam_role_policies OWNER TO postgres;

--
-- Name: aws_iam_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_roles (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    policies jsonb,
    arn text,
    create_date timestamp without time zone,
    path text,
    id text NOT NULL,
    name text,
    assume_role_policy_document jsonb,
    description text,
    max_session_duration integer,
    permissions_boundary_arn text,
    permissions_boundary_type text,
    role_last_used_last_used_date timestamp without time zone,
    role_last_used_region text,
    tags jsonb
);


ALTER TABLE public.aws_iam_roles OWNER TO postgres;

--
-- Name: aws_iam_saml_identity_providers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_saml_identity_providers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    arn text NOT NULL,
    create_date timestamp without time zone,
    saml_metadata_document text,
    tags jsonb,
    valid_until timestamp without time zone
);


ALTER TABLE public.aws_iam_saml_identity_providers OWNER TO postgres;

--
-- Name: aws_iam_server_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_server_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    id text NOT NULL,
    arn text,
    path text,
    name text,
    expiration timestamp without time zone,
    upload_date timestamp without time zone
);


ALTER TABLE public.aws_iam_server_certificates OWNER TO postgres;

--
-- Name: aws_iam_user_access_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_user_access_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_cq_id uuid,
    user_id text,
    access_key_id text,
    create_date timestamp without time zone,
    status text,
    last_used timestamp without time zone,
    last_rotated timestamp without time zone,
    last_used_service_name text
);


ALTER TABLE public.aws_iam_user_access_keys OWNER TO postgres;

--
-- Name: aws_iam_user_attached_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_user_attached_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_cq_id uuid,
    user_id text,
    policy_arn text,
    policy_name text
);


ALTER TABLE public.aws_iam_user_attached_policies OWNER TO postgres;

--
-- Name: aws_iam_user_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_user_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_cq_id uuid,
    user_id text,
    group_arn text,
    create_date timestamp without time zone,
    group_id text,
    group_name text,
    path text
);


ALTER TABLE public.aws_iam_user_groups OWNER TO postgres;

--
-- Name: aws_iam_user_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_user_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    user_cq_id uuid,
    account_id text,
    user_id text,
    policy_document jsonb,
    policy_name text,
    user_name text
);


ALTER TABLE public.aws_iam_user_policies OWNER TO postgres;

--
-- Name: aws_iam_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_users (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    id text NOT NULL,
    password_last_used timestamp without time zone,
    arn text,
    password_enabled boolean,
    password_status text,
    password_last_changed timestamp without time zone,
    password_next_rotation timestamp without time zone,
    mfa_active boolean,
    create_date timestamp without time zone,
    path text,
    permissions_boundary_arn text,
    permissions_boundary_type text,
    tags jsonb,
    user_id text,
    user_name text,
    access_key_1_active boolean,
    access_key_1_last_rotated timestamp without time zone,
    access_key_2_active boolean,
    access_key_2_last_rotated timestamp without time zone,
    cert_1_active boolean,
    cert_1_last_rotated timestamp without time zone,
    cert_2_active boolean,
    cert_2_last_rotated timestamp without time zone
);


ALTER TABLE public.aws_iam_users OWNER TO postgres;

--
-- Name: aws_iam_virtual_mfa_devices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iam_virtual_mfa_devices (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    serial_number text NOT NULL,
    base32_string_seed bytea,
    enable_date timestamp without time zone,
    qr_code_png bytea,
    tags jsonb,
    user_arn text,
    user_create_date timestamp without time zone,
    user_path text,
    user_id text,
    user_name text,
    user_password_last_used timestamp without time zone,
    user_permissions_boundary_permissions_boundary_arn text,
    user_permissions_boundary_permissions_boundary_type text,
    user_tags jsonb
);


ALTER TABLE public.aws_iam_virtual_mfa_devices OWNER TO postgres;

--
-- Name: aws_iot_billing_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_billing_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    things_in_group text[],
    tags jsonb,
    arn text NOT NULL,
    id text,
    creation_date timestamp without time zone,
    name text,
    description text,
    version bigint
);


ALTER TABLE public.aws_iot_billing_groups OWNER TO postgres;

--
-- Name: aws_iot_ca_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_ca_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    certificates text[],
    auto_registration_status text,
    arn text NOT NULL,
    id text,
    pem text,
    creation_date timestamp without time zone,
    customer_version integer,
    generation_id text,
    last_modified_date timestamp without time zone,
    owned_by text,
    status text,
    validity_not_after timestamp without time zone,
    validity_not_before timestamp without time zone
);


ALTER TABLE public.aws_iot_ca_certificates OWNER TO postgres;

--
-- Name: aws_iot_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    policies text[],
    ca_certificate_id text,
    arn text NOT NULL,
    id text,
    mode text,
    pem text,
    creation_date timestamp without time zone,
    customer_version integer,
    generation_id text,
    last_modified_date timestamp without time zone,
    owned_by text,
    previous_owned_by text,
    status text,
    transfer_data_accept_date timestamp without time zone,
    transfer_data_reject_date timestamp without time zone,
    transfer_data_reject_reason text,
    transfer_data_transfer_date timestamp without time zone,
    transfer_data_transfer_message text,
    validity_not_after timestamp without time zone,
    validity_not_before timestamp without time zone
);


ALTER TABLE public.aws_iot_certificates OWNER TO postgres;

--
-- Name: aws_iot_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    creation_date timestamp without time zone,
    default_version_id text,
    generation_id text,
    last_modified_date timestamp without time zone,
    arn text NOT NULL,
    document text,
    name text
);


ALTER TABLE public.aws_iot_policies OWNER TO postgres;

--
-- Name: aws_iot_stream_files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_stream_files (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    stream_cq_id uuid,
    file_id integer,
    s3_location_bucket text,
    s3_location_key text,
    s3_location_version text
);


ALTER TABLE public.aws_iot_stream_files OWNER TO postgres;

--
-- Name: aws_iot_streams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_streams (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    created_at timestamp without time zone,
    description text,
    last_updated_at timestamp without time zone,
    role_arn text,
    arn text NOT NULL,
    id text,
    version integer
);


ALTER TABLE public.aws_iot_streams OWNER TO postgres;

--
-- Name: aws_iot_thing_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_thing_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    things_in_group text[],
    policies text[],
    tags jsonb,
    index_name text,
    query_string text,
    query_version text,
    status text,
    arn text NOT NULL,
    id text,
    creation_date timestamp without time zone,
    parent_group_name text,
    root_to_parent_thing_groups jsonb,
    name text,
    attribute_payload_attributes jsonb,
    attribute_payload_merge boolean,
    thing_group_description text,
    version bigint
);


ALTER TABLE public.aws_iot_thing_groups OWNER TO postgres;

--
-- Name: aws_iot_thing_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_thing_types (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    arn text NOT NULL,
    creation_date timestamp without time zone,
    deprecated boolean,
    deprecation_date timestamp without time zone,
    name text,
    searchable_attributes text[],
    description text
);


ALTER TABLE public.aws_iot_thing_types OWNER TO postgres;

--
-- Name: aws_iot_things; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_things (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    principals text[],
    attributes jsonb,
    arn text NOT NULL,
    name text,
    type_name text,
    version bigint
);


ALTER TABLE public.aws_iot_things OWNER TO postgres;

--
-- Name: aws_iot_topic_rule_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_topic_rule_actions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    topic_rule_cq_id uuid,
    cloudwatch_alarm_alarm_name text,
    cloudwatch_alarm_role_arn text,
    cloudwatch_alarm_state_reason text,
    cloudwatch_alarm_state_value text,
    cloudwatch_logs_log_group_name text,
    cloudwatch_logs_role_arn text,
    cloudwatch_metric_metric_name text,
    cloudwatch_metric_metric_namespace text,
    cloudwatch_metric_metric_unit text,
    cloudwatch_metric_metric_value text,
    cloudwatch_metric_role_arn text,
    cloudwatch_metric_metric_timestamp text,
    dynamo_db_hash_key_field text,
    dynamo_db_hash_key_value text,
    dynamo_db_role_arn text,
    dynamo_db_table_name text,
    dynamo_db_hash_key_type text,
    dynamo_db_operation text,
    dynamo_db_payload_field text,
    dynamo_db_range_key_field text,
    dynamo_db_range_key_type text,
    dynamo_db_range_key_value text,
    dynamo_db_v2_put_item_table_name text,
    dynamo_db_v2_role_arn text,
    elasticsearch_endpoint text,
    elasticsearch_id text,
    elasticsearch_index text,
    elasticsearch_role_arn text,
    elasticsearch_type text,
    firehose_delivery_stream_name text,
    firehose_role_arn text,
    firehose_batch_mode boolean,
    firehose_separator text,
    http_url text,
    http_auth_sigv4_role_arn text,
    http_auth_sigv4_service_name text,
    http_auth_sigv4_signing_region text,
    http_confirmation_url text,
    http_headers jsonb,
    iot_analytics_batch_mode boolean,
    iot_analytics_channel_arn text,
    iot_analytics_channel_name text,
    iot_analytics_role_arn text,
    iot_events_input_name text,
    iot_events_role_arn text,
    iot_events_batch_mode boolean,
    iot_events_message_id text,
    iot_site_wise jsonb,
    kafka_client_properties jsonb,
    kafka_destination_arn text,
    kafka_topic text,
    kafka_key text,
    kafka_partition text,
    kinesis_role_arn text,
    kinesis_stream_name text,
    kinesis_partition_key text,
    lambda_function_arn text,
    open_search_endpoint text,
    open_search_id text,
    open_search_index text,
    open_search_role_arn text,
    open_search_type text,
    republish_role_arn text,
    republish_topic text,
    republish_qos integer,
    s3_bucket_name text,
    s3_key text,
    s3_role_arn text,
    s3_canned_acl text,
    salesforce_token text,
    salesforce_url text,
    sns_role_arn text,
    sns_target_arn text,
    sns_message_format text,
    sqs_queue_url text,
    sqs_role_arn text,
    sqs_use_base64 boolean,
    step_functions_role_arn text,
    step_functions_state_machine_name text,
    step_functions_execution_name_prefix text,
    timestream_database_name text,
    timestream_dimensions jsonb,
    timestream_role_arn text,
    timestream_table_name text,
    timestream_timestamp_unit text,
    timestream_timestamp_value text
);


ALTER TABLE public.aws_iot_topic_rule_actions OWNER TO postgres;

--
-- Name: aws_iot_topic_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_iot_topic_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    aws_iot_sql_version text,
    created_at timestamp without time zone,
    description text,
    error_action_cloudwatch_alarm_name text,
    error_action_cloudwatch_alarm_role_arn text,
    error_action_cloudwatch_alarm_state_reason text,
    error_action_cloudwatch_alarm_state_value text,
    error_action_cloudwatch_logs_log_group_name text,
    error_action_cloudwatch_logs_role_arn text,
    error_action_cloudwatch_metric_metric_name text,
    error_action_cloudwatch_metric_metric_namespace text,
    error_action_cloudwatch_metric_unit text,
    error_action_cloudwatch_metric_value text,
    error_action_cloudwatch_metric_role_arn text,
    error_action_cloudwatch_metric_timestamp text,
    error_action_dynamo_db_hash_key_field text,
    error_action_dynamo_db_hash_key_value text,
    error_action_dynamo_db_role_arn text,
    error_action_dynamo_db_table_name text,
    error_action_dynamo_db_hash_key_type text,
    error_action_dynamo_db_operation text,
    error_action_dynamo_db_payload_field text,
    error_action_dynamo_db_range_key_field text,
    error_action_dynamo_db_range_key_type text,
    error_action_dynamo_db_range_key_value text,
    error_action_dynamo_db_v2_put_item_table_name text,
    error_action_dynamo_db_v2_role_arn text,
    error_action_elasticsearch_endpoint text,
    error_action_elasticsearch_id text,
    error_action_elasticsearch_index text,
    error_action_elasticsearch_role_arn text,
    error_action_elasticsearch_type text,
    error_action_firehose_delivery_stream_name text,
    error_action_firehose_role_arn text,
    error_action_firehose_batch_mode boolean,
    error_action_firehose_separator text,
    error_action_http_url text,
    error_action_http_auth_sigv4_role_arn text,
    error_action_http_auth_sigv4_service_name text,
    error_action_http_auth_sigv4_signing_region text,
    error_action_http_confirmation_url text,
    error_action_http_headers jsonb,
    error_action_iot_analytics_batch_mode boolean,
    error_action_iot_analytics_channel_arn text,
    error_action_iot_analytics_channel_name text,
    error_action_iot_analytics_role_arn text,
    error_action_iot_events_input_name text,
    error_action_iot_events_role_arn text,
    error_action_iot_events_batch_mode boolean,
    error_action_iot_events_message_id text,
    error_action_iot_site_wise jsonb,
    error_action_kafka_client_properties jsonb,
    error_action_kafka_destination_arn text,
    error_action_kafka_topic text,
    error_action_kafka_key text,
    error_action_kafka_partition text,
    error_action_kinesis_role_arn text,
    error_action_kinesis_stream_name text,
    error_action_kinesis_partition_key text,
    error_action_lambda_function_arn text,
    error_action_open_search_endpoint text,
    error_action_open_search_id text,
    error_action_open_search_index text,
    error_action_open_search_role_arn text,
    error_action_open_search_type text,
    error_action_republish_role_arn text,
    error_action_republish_topic text,
    error_action_republish_qos integer,
    error_action_s3_bucket_name text,
    error_action_s3_key text,
    error_action_s3_role_arn text,
    error_action_s3_canned_acl text,
    error_action_salesforce_token text,
    error_action_salesforce_url text,
    error_action_sns_role_arn text,
    error_action_sns_target_arn text,
    error_action_sns_message_format text,
    error_action_sqs_queue_url text,
    error_action_sqs_role_arn text,
    error_action_sqs_use_base64 boolean,
    error_action_step_functions_role_arn text,
    error_action_step_functions_state_machine_name text,
    error_action_step_functions_execution_name_prefix text,
    error_action_timestream_database_name text,
    error_action_timestream_dimensions jsonb,
    error_action_timestream_role_arn text,
    error_action_timestream_table_name text,
    error_action_timestream_timestamp_unit text,
    error_action_timestream_timestamp_value text,
    rule_disabled boolean,
    rule_name text,
    sql text,
    arn text NOT NULL
);


ALTER TABLE public.aws_iot_topic_rules OWNER TO postgres;

--
-- Name: aws_kms_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_kms_keys (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    rotation_enabled boolean,
    tags jsonb,
    id text,
    aws_account_id text,
    arn text NOT NULL,
    cloud_hsm_cluster_id text,
    creation_date timestamp without time zone,
    custom_key_store_id text,
    deletion_date timestamp without time zone,
    description text,
    enabled boolean,
    encryption_algorithms text[],
    expiration_model text,
    manager text,
    key_spec text,
    key_state text,
    key_usage text,
    mac_algorithms text[],
    multi_region boolean,
    multi_region_key_type text,
    primary_key_arn text,
    primary_key_region text,
    replica_keys jsonb,
    origin text,
    pending_deletion_window_in_days integer,
    signing_algorithms text[],
    valid_to timestamp without time zone
);


ALTER TABLE public.aws_kms_keys OWNER TO postgres;

--
-- Name: aws_lambda_function_aliases; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_aliases (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    function_arn text,
    arn text,
    description text,
    function_version text,
    name text,
    revision_id text,
    routing_config_additional_version_weights jsonb,
    url_config_auth_type text,
    url_config_creation_time timestamp without time zone,
    url_config_function_arn text,
    url_config_function_url text,
    url_config_last_modified_time timestamp without time zone,
    url_config_cors jsonb
);


ALTER TABLE public.aws_lambda_function_aliases OWNER TO postgres;

--
-- Name: aws_lambda_function_concurrency_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_concurrency_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    allocated_provisioned_concurrent_executions integer,
    available_provisioned_concurrent_executions integer,
    function_arn text,
    last_modified timestamp without time zone,
    requested_provisioned_concurrent_executions integer,
    status text,
    status_reason text
);


ALTER TABLE public.aws_lambda_function_concurrency_configs OWNER TO postgres;

--
-- Name: aws_lambda_function_event_invoke_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_event_invoke_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    on_failure_destination text,
    on_success_destination text,
    function_arn text,
    last_modified timestamp without time zone,
    maximum_event_age_in_seconds integer,
    maximum_retry_attempts integer
);


ALTER TABLE public.aws_lambda_function_event_invoke_configs OWNER TO postgres;

--
-- Name: aws_lambda_function_event_source_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_event_source_mappings (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    batch_size integer,
    bisect_batch_on_function_error boolean,
    on_failure_destination text,
    on_success_destination text,
    event_source_arn text,
    criteria_filters text[],
    function_arn text,
    function_response_types text[],
    last_modified timestamp without time zone,
    last_processing_result text,
    maximum_batching_window_in_seconds integer,
    maximum_record_age_in_seconds integer,
    maximum_retry_attempts integer,
    parallelization_factor integer,
    queues text[],
    self_managed_event_source_endpoints jsonb,
    source_access_configurations jsonb,
    starting_position text,
    starting_position_timestamp timestamp without time zone,
    state text,
    state_transition_reason text,
    topics text[],
    tumbling_window_in_seconds integer,
    uuid text
);


ALTER TABLE public.aws_lambda_function_event_source_mappings OWNER TO postgres;

--
-- Name: aws_lambda_function_file_system_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_file_system_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    function_arn text,
    arn text,
    local_mount_path text
);


ALTER TABLE public.aws_lambda_function_file_system_configs OWNER TO postgres;

--
-- Name: aws_lambda_function_layers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_layers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    function_arn text,
    arn text,
    code_size bigint,
    signing_job_arn text,
    signing_profile_version_arn text
);


ALTER TABLE public.aws_lambda_function_layers OWNER TO postgres;

--
-- Name: aws_lambda_function_version_file_system_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_version_file_system_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_version_cq_id uuid,
    arn text,
    local_mount_path text
);


ALTER TABLE public.aws_lambda_function_version_file_system_configs OWNER TO postgres;

--
-- Name: aws_lambda_function_version_layers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_version_layers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_version_cq_id uuid,
    arn text,
    code_size bigint,
    signing_job_arn text,
    signing_profile_version_arn text
);


ALTER TABLE public.aws_lambda_function_version_layers OWNER TO postgres;

--
-- Name: aws_lambda_function_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_function_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    function_cq_id uuid,
    architectures text[],
    code_sha256 text,
    code_size bigint,
    dead_letter_config_target_arn text,
    description text,
    environment_error_error_code text,
    environment_error_message text,
    environment_variables jsonb,
    ephemeral_storage_size integer,
    function_arn text,
    function_name text,
    handler text,
    error_code text,
    error_message text,
    image_config_command text[],
    image_config_entry_point text[],
    image_config_working_directory text,
    kms_key_arn text,
    last_modified timestamp without time zone,
    last_update_status text,
    last_update_status_reason text,
    last_update_status_reason_code text,
    master_arn text,
    memory_size integer,
    package_type text,
    revision_id text,
    role text,
    runtime text,
    signing_job_arn text,
    signing_profile_version_arn text,
    state text,
    state_reason text,
    state_reason_code text,
    timeout integer,
    tracing_config_mode text,
    version text,
    vpc_config_security_group_ids text[],
    vpc_config_subnet_ids text[],
    vpc_config_vpc_id text
);


ALTER TABLE public.aws_lambda_function_versions OWNER TO postgres;

--
-- Name: aws_lambda_functions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_functions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    policy_document jsonb,
    policy_revision_id text,
    code_signing_allowed_publishers_version_arns text[],
    code_signing_config_arn text,
    code_signing_config_id text,
    code_signing_policies_untrusted_artifact_on_deployment text,
    code_signing_description text,
    code_signing_last_modified timestamp without time zone,
    code_image_uri text,
    code_location text,
    code_repository_type text,
    code_resolved_image_uri text,
    concurrency_reserved_concurrent_executions integer,
    architectures text[],
    code_sha256 text,
    code_size bigint,
    dead_letter_config_target_arn text,
    description text,
    environment_error_code text,
    environment_error_message text,
    environment_variables jsonb,
    ephemeral_storage_size integer,
    arn text NOT NULL,
    name text,
    handler text,
    error_code text,
    error_message text,
    image_config_command text[],
    image_config_entry_point text[],
    image_config_working_directory text,
    kms_key_arn text,
    last_modified timestamp without time zone,
    last_update_status text,
    last_update_status_reason text,
    last_update_status_reason_code text,
    master_arn text,
    memory_size integer,
    package_type text,
    revision_id text,
    role text,
    runtime text,
    signing_job_arn text,
    signing_profile_version_arn text,
    state text,
    state_reason text,
    state_reason_code text,
    timeout integer,
    tracing_config_mode text,
    version text,
    vpc_config_security_group_ids text[],
    vpc_config_subnet_ids text[],
    vpc_config_vpc_id text,
    tags jsonb
);


ALTER TABLE public.aws_lambda_functions OWNER TO postgres;

--
-- Name: aws_lambda_layer_version_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_layer_version_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    layer_version_cq_id uuid,
    layer_version bigint,
    policy text,
    revision_id text
);


ALTER TABLE public.aws_lambda_layer_version_policies OWNER TO postgres;

--
-- Name: aws_lambda_layer_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_layer_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    layer_cq_id uuid,
    compatible_runtimes text[],
    created_date timestamp without time zone,
    description text,
    layer_version_arn text,
    license_info text,
    version bigint
);


ALTER TABLE public.aws_lambda_layer_versions OWNER TO postgres;

--
-- Name: aws_lambda_layers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_layers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    latest_matching_version_compatible_runtimes text[],
    latest_matching_version_created_date timestamp without time zone,
    latest_matching_version_description text,
    latest_matching_version_layer_version_arn text,
    latest_matching_version_license_info text,
    latest_matching_version bigint,
    arn text NOT NULL,
    name text
);


ALTER TABLE public.aws_lambda_layers OWNER TO postgres;

--
-- Name: aws_lambda_runtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_lambda_runtimes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    name text NOT NULL
);


ALTER TABLE public.aws_lambda_runtimes OWNER TO postgres;

--
-- Name: aws_mq_broker_configuration_revisions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_mq_broker_configuration_revisions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    broker_configuration_cq_id uuid,
    configuration_id text,
    created timestamp without time zone,
    data jsonb,
    description text
);


ALTER TABLE public.aws_mq_broker_configuration_revisions OWNER TO postgres;

--
-- Name: aws_mq_broker_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_mq_broker_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    broker_cq_id uuid,
    account_id text,
    region text,
    arn text,
    authentication_strategy text,
    created timestamp without time zone,
    description text,
    engine_type text,
    engine_version text,
    id text,
    latest_revision_created timestamp without time zone,
    latest_revision integer,
    latest_revision_description text,
    name text,
    tags jsonb
);


ALTER TABLE public.aws_mq_broker_configurations OWNER TO postgres;

--
-- Name: aws_mq_broker_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_mq_broker_users (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    broker_cq_id uuid,
    account_id text,
    region text,
    console_access boolean,
    groups text[],
    pending jsonb,
    username text
);


ALTER TABLE public.aws_mq_broker_users OWNER TO postgres;

--
-- Name: aws_mq_brokers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_mq_brokers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    authentication_strategy text,
    auto_minor_version_upgrade boolean,
    arn text,
    id text NOT NULL,
    broker_instances jsonb,
    broker_name text,
    broker_state text,
    created timestamp without time zone,
    deployment_mode text,
    encryption_options_use_aws_owned_key boolean,
    encryption_options_kms_key_id text,
    engine_type text,
    engine_version text,
    host_instance_type text,
    ldap_server_metadata jsonb,
    logs jsonb,
    maintenance_window_start_time jsonb,
    pending_authentication_strategy text,
    pending_engine_version text,
    pending_host_instance_type text,
    pending_ldap_server_metadata jsonb,
    pending_security_groups text[],
    publicly_accessible boolean,
    security_groups text[],
    storage_type text,
    subnet_ids text[],
    tags jsonb
);


ALTER TABLE public.aws_mq_brokers OWNER TO postgres;

--
-- Name: aws_organizations_accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_organizations_accounts (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    tags jsonb,
    arn text,
    email text,
    id text NOT NULL,
    joined_method text,
    joined_timestamp timestamp without time zone,
    name text,
    status text
);


ALTER TABLE public.aws_organizations_accounts OWNER TO postgres;

--
-- Name: aws_qldb_ledger_journal_kinesis_streams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_qldb_ledger_journal_kinesis_streams (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    ledger_cq_id uuid,
    stream_arn text,
    aggregation_enabled boolean,
    ledger_name text,
    role_arn text,
    status text,
    stream_id text,
    stream_name text,
    arn text,
    creation_time timestamp without time zone,
    error_cause text,
    exclusive_end_time timestamp without time zone,
    inclusive_start_time timestamp without time zone
);


ALTER TABLE public.aws_qldb_ledger_journal_kinesis_streams OWNER TO postgres;

--
-- Name: aws_qldb_ledger_journal_s3_exports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_qldb_ledger_journal_s3_exports (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    ledger_cq_id uuid,
    exclusive_end_time timestamp without time zone,
    export_creation_time timestamp without time zone,
    export_id text,
    inclusive_start_time timestamp without time zone,
    ledger_name text,
    role_arn text,
    bucket text,
    object_encryption_type text,
    kms_key_arn text,
    prefix text,
    status text,
    output_format text
);


ALTER TABLE public.aws_qldb_ledger_journal_s3_exports OWNER TO postgres;

--
-- Name: aws_qldb_ledgers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_qldb_ledgers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    arn text NOT NULL,
    creation_date_time timestamp without time zone,
    deletion_protection boolean,
    encryption_status text,
    kms_key_arn text,
    inaccessible_kms_key_date_time timestamp without time zone,
    name text,
    permissions_mode text,
    state text
);


ALTER TABLE public.aws_qldb_ledgers OWNER TO postgres;

--
-- Name: aws_rds_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_certificates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text NOT NULL,
    certificate_identifier text,
    certificate_type text,
    customer_override boolean,
    customer_override_valid_till timestamp without time zone,
    thumbprint text,
    valid_from timestamp without time zone,
    valid_till timestamp without time zone
);


ALTER TABLE public.aws_rds_certificates OWNER TO postgres;

--
-- Name: aws_rds_cluster_associated_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_associated_roles (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    feature_name text,
    role_arn text,
    status text
);


ALTER TABLE public.aws_rds_cluster_associated_roles OWNER TO postgres;

--
-- Name: aws_rds_cluster_db_cluster_members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_db_cluster_members (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    db_cluster_parameter_group_status text,
    db_instance_identifier text,
    is_cluster_writer boolean,
    promotion_tier integer
);


ALTER TABLE public.aws_rds_cluster_db_cluster_members OWNER TO postgres;

--
-- Name: aws_rds_cluster_domain_memberships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_domain_memberships (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    domain text,
    fqdn text,
    iam_role_name text,
    status text
);


ALTER TABLE public.aws_rds_cluster_domain_memberships OWNER TO postgres;

--
-- Name: aws_rds_cluster_parameter_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_parameter_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    name text,
    family text,
    description text,
    tags jsonb
);


ALTER TABLE public.aws_rds_cluster_parameter_groups OWNER TO postgres;

--
-- Name: aws_rds_cluster_parameters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_parameters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_parameter_group_cq_id uuid,
    allowed_values text,
    apply_method text,
    apply_type text,
    data_type text,
    description text,
    is_modifiable boolean,
    minimum_engine_version text,
    parameter_name text,
    parameter_value text,
    source text,
    supported_engine_modes text[]
);


ALTER TABLE public.aws_rds_cluster_parameters OWNER TO postgres;

--
-- Name: aws_rds_cluster_snapshots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_snapshots (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    allocated_storage integer,
    availability_zones text[],
    cluster_create_time timestamp without time zone,
    db_cluster_identifier text,
    arn text NOT NULL,
    db_cluster_snapshot_identifier text,
    engine text,
    engine_mode text,
    engine_version text,
    iam_database_authentication_enabled boolean,
    kms_key_id text,
    license_model text,
    master_username text,
    percent_progress integer,
    port integer,
    snapshot_create_time timestamp without time zone,
    snapshot_type text,
    source_db_cluster_snapshot_arn text,
    status text,
    storage_encrypted boolean,
    vpc_id text,
    tags jsonb,
    attributes jsonb
);


ALTER TABLE public.aws_rds_cluster_snapshots OWNER TO postgres;

--
-- Name: aws_rds_cluster_vpc_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_cluster_vpc_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    status text,
    vpc_security_group_id text
);


ALTER TABLE public.aws_rds_cluster_vpc_security_groups OWNER TO postgres;

--
-- Name: aws_rds_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    activity_stream_kinesis_stream_name text,
    activity_stream_kms_key_id text,
    activity_stream_mode text,
    activity_stream_status text,
    allocated_storage integer,
    availability_zones text[],
    backtrack_consumed_change_records bigint,
    backtrack_window bigint,
    backup_retention_period integer,
    capacity integer,
    character_set_name text,
    clone_group_id text,
    cluster_create_time timestamp without time zone,
    copy_tags_to_snapshot boolean,
    cross_account_clone boolean,
    custom_endpoints text[],
    arn text,
    db_cluster_identifier text,
    db_cluster_parameter_group text,
    db_cluster_option_group_memberships jsonb,
    db_subnet_group text,
    database_name text,
    id text NOT NULL,
    deletion_protection boolean,
    earliest_backtrack_time timestamp without time zone,
    earliest_restorable_time timestamp without time zone,
    enabled_cloudwatch_logs_exports text[],
    endpoint text,
    engine text,
    engine_mode text,
    engine_version text,
    global_write_forwarding_requested boolean,
    global_write_forwarding_status text,
    hosted_zone_id text,
    http_endpoint_enabled boolean,
    iam_database_authentication_enabled boolean,
    kms_key_id text,
    latest_restorable_time timestamp without time zone,
    master_username text,
    multi_az boolean,
    pending_modified_values_db_cluster_identifier text,
    pending_modified_values_engine_version text,
    pending_modified_values_iam_database_authentication_enabled boolean,
    pending_modified_values_master_user_password text,
    pending_cloudwatch_logs_types_to_disable text[],
    pending_cloudwatch_logs_types_to_enable text[],
    percent_progress text,
    port integer,
    preferred_backup_window text,
    preferred_maintenance_window text,
    read_replica_identifiers text[],
    reader_endpoint text,
    replication_source_identifier text,
    scaling_configuration_info_auto_pause boolean,
    scaling_configuration_info_max_capacity integer,
    scaling_configuration_info_min_capacity integer,
    scaling_configuration_info_seconds_until_auto_pause integer,
    scaling_configuration_info_timeout_action text,
    status text,
    storage_encrypted boolean,
    tags jsonb
);


ALTER TABLE public.aws_rds_clusters OWNER TO postgres;

--
-- Name: aws_rds_db_parameter_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_db_parameter_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    family text,
    name text,
    description text,
    tags jsonb
);


ALTER TABLE public.aws_rds_db_parameter_groups OWNER TO postgres;

--
-- Name: aws_rds_db_parameters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_db_parameters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    db_parameter_group_cq_id uuid,
    allowed_values text,
    apply_method text,
    apply_type text,
    data_type text,
    description text,
    is_modifiable boolean,
    minimum_engine_version text,
    parameter_name text,
    parameter_value text,
    source text,
    supported_engine_modes text[]
);


ALTER TABLE public.aws_rds_db_parameters OWNER TO postgres;

--
-- Name: aws_rds_db_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_db_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    description text,
    name text,
    ec2_security_groups jsonb,
    ip_ranges jsonb,
    owner_id text,
    vpc_id text,
    tags jsonb
);


ALTER TABLE public.aws_rds_db_security_groups OWNER TO postgres;

--
-- Name: aws_rds_db_snapshots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_db_snapshots (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    allocated_storage integer,
    availability_zone text,
    db_instance_identifier text,
    arn text NOT NULL,
    db_snapshot_identifier text,
    dbi_resource_id text,
    encrypted boolean,
    engine text,
    engine_version text,
    iam_database_authentication_enabled boolean,
    instance_create_time timestamp without time zone,
    iops integer,
    kms_key_id text,
    license_model text,
    master_username text,
    option_group_name text,
    percent_progress integer,
    port integer,
    processor_features jsonb,
    snapshot_create_time timestamp without time zone,
    snapshot_type text,
    source_db_snapshot_identifier text,
    source_region text,
    status text,
    storage_type text,
    tde_credential_arn text,
    timezone text,
    vpc_id text,
    tags jsonb,
    attributes jsonb
);


ALTER TABLE public.aws_rds_db_snapshots OWNER TO postgres;

--
-- Name: aws_rds_event_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_event_subscriptions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    cust_subscription_id text,
    customer_aws_id text,
    enabled boolean,
    event_categories_list text[],
    arn text NOT NULL,
    sns_topic_arn text,
    source_ids_list text[],
    source_type text,
    status text,
    subscription_creation_time text,
    tags jsonb
);


ALTER TABLE public.aws_rds_event_subscriptions OWNER TO postgres;

--
-- Name: aws_rds_instance_associated_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_associated_roles (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    feature_name text,
    role_arn text,
    status text
);


ALTER TABLE public.aws_rds_instance_associated_roles OWNER TO postgres;

--
-- Name: aws_rds_instance_db_instance_automated_backups_replications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_db_instance_automated_backups_replications (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    db_instance_automated_backups_arn text
);


ALTER TABLE public.aws_rds_instance_db_instance_automated_backups_replications OWNER TO postgres;

--
-- Name: aws_rds_instance_db_parameter_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_db_parameter_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    db_parameter_group_name text,
    parameter_apply_status text
);


ALTER TABLE public.aws_rds_instance_db_parameter_groups OWNER TO postgres;

--
-- Name: aws_rds_instance_db_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_db_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    db_security_group_name text,
    status text
);


ALTER TABLE public.aws_rds_instance_db_security_groups OWNER TO postgres;

--
-- Name: aws_rds_instance_db_subnet_group_subnets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_db_subnet_group_subnets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    subnet_availability_zone_name text,
    subnet_identifier text,
    subnet_outpost_arn text,
    subnet_status text
);


ALTER TABLE public.aws_rds_instance_db_subnet_group_subnets OWNER TO postgres;

--
-- Name: aws_rds_instance_domain_memberships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_domain_memberships (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    domain text,
    fqdn text,
    iam_role_name text,
    status text
);


ALTER TABLE public.aws_rds_instance_domain_memberships OWNER TO postgres;

--
-- Name: aws_rds_instance_option_group_memberships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_option_group_memberships (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    option_group_name text,
    status text
);


ALTER TABLE public.aws_rds_instance_option_group_memberships OWNER TO postgres;

--
-- Name: aws_rds_instance_vpc_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instance_vpc_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    instance_id text,
    status text,
    vpc_security_group_id text
);


ALTER TABLE public.aws_rds_instance_vpc_security_groups OWNER TO postgres;

--
-- Name: aws_rds_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    allocated_storage integer,
    auto_minor_version_upgrade boolean,
    availability_zone text,
    aws_backup_recovery_point_arn text,
    backup_retention_period integer,
    ca_certificate_identifier text,
    character_set_name text,
    copy_tags_to_snapshot boolean,
    customer_owned_ip_enabled boolean,
    cluster_identifier text,
    arn text NOT NULL,
    db_instance_class text,
    user_instance_id text,
    db_instance_status text,
    db_name text,
    subnet_group_arn text,
    subnet_group_description text,
    subnet_group_name text,
    subnet_group_subnet_group_status text,
    subnet_group_vpc_id text,
    instance_port integer,
    id text,
    deletion_protection boolean,
    enabled_cloudwatch_logs_exports text[],
    endpoint_address text,
    endpoint_hosted_zone_id text,
    endpoint_port integer,
    engine text,
    engine_version text,
    enhanced_monitoring_resource_arn text,
    iam_database_authentication_enabled boolean,
    instance_create_time timestamp without time zone,
    iops integer,
    kms_key_id text,
    latest_restorable_time timestamp without time zone,
    license_model text,
    listener_endpoint_address text,
    listener_endpoint_hosted_zone_id text,
    listener_endpoint_port integer,
    master_username text,
    max_allocated_storage integer,
    monitoring_interval integer,
    monitoring_role_arn text,
    multi_az boolean,
    nchar_character_set_name text,
    pending_modified_values_allocated_storage integer,
    pending_modified_values_backup_retention_period integer,
    pending_modified_values_ca_certificate_identifier text,
    pending_modified_values_db_instance_class text,
    pending_modified_values_db_instance_identifier text,
    pending_modified_values_db_subnet_group_name text,
    pending_modified_values_engine_version text,
    pending_modified_values_iam_database_authentication_enabled boolean,
    pending_modified_values_iops integer,
    pending_modified_values_license_model text,
    pending_modified_values_master_user_password text,
    pending_modified_values_multi_az boolean,
    pending_cloudwatch_logs_types_to_disable text[],
    pending_cloudwatch_logs_types_to_enable text[],
    pending_modified_values_port integer,
    pending_modified_values_processor_features jsonb,
    pending_modified_values_storage_type text,
    performance_insights_enabled boolean,
    performance_insights_kms_key_id text,
    performance_insights_retention_period integer,
    preferred_backup_window text,
    preferred_maintenance_window text,
    processor_features jsonb,
    promotion_tier integer,
    publicly_accessible boolean,
    read_replica_db_cluster_identifiers text[],
    read_replica_db_instance_identifiers text[],
    read_replica_source_db_instance_identifier text,
    replica_mode text,
    secondary_availability_zone text,
    storage_encrypted boolean,
    storage_type text,
    tags jsonb,
    tde_credential_arn text,
    timezone text,
    status_infos jsonb
);


ALTER TABLE public.aws_rds_instances OWNER TO postgres;

--
-- Name: aws_rds_subnet_group_subnets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_subnet_group_subnets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    subnet_group_cq_id uuid,
    subnet_availability_zone_name text,
    subnet_identifier text,
    subnet_outpost_arn text,
    subnet_status text
);


ALTER TABLE public.aws_rds_subnet_group_subnets OWNER TO postgres;

--
-- Name: aws_rds_subnet_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_rds_subnet_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    description text,
    name text,
    status text,
    vpc_id text
);


ALTER TABLE public.aws_rds_subnet_groups OWNER TO postgres;

--
-- Name: aws_redshift_cluster_deferred_maintenance_windows; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_deferred_maintenance_windows (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    defer_maintenance_end_time timestamp without time zone,
    defer_maintenance_identifier text,
    defer_maintenance_start_time timestamp without time zone
);


ALTER TABLE public.aws_redshift_cluster_deferred_maintenance_windows OWNER TO postgres;

--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_endpoint_vpc_endpoint_cq_id uuid,
    availability_zone text,
    network_interface_id text,
    private_ip_address text,
    subnet_id text
);


ALTER TABLE public.aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces OWNER TO postgres;

--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoints; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_endpoint_vpc_endpoints (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    vpc_endpoint_id text,
    vpc_id text
);


ALTER TABLE public.aws_redshift_cluster_endpoint_vpc_endpoints OWNER TO postgres;

--
-- Name: aws_redshift_cluster_iam_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_iam_roles (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    apply_status text,
    iam_role_arn text
);


ALTER TABLE public.aws_redshift_cluster_iam_roles OWNER TO postgres;

--
-- Name: aws_redshift_cluster_nodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_nodes (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    node_role text,
    private_ip_address text,
    public_ip_address text
);


ALTER TABLE public.aws_redshift_cluster_nodes OWNER TO postgres;

--
-- Name: aws_redshift_cluster_parameter_group_status_lists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_parameter_group_status_lists (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_parameter_group_cq_id uuid,
    parameter_apply_error_description text,
    parameter_apply_status text,
    parameter_name text
);


ALTER TABLE public.aws_redshift_cluster_parameter_group_status_lists OWNER TO postgres;

--
-- Name: aws_redshift_cluster_parameter_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_parameter_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    parameter_apply_status text,
    parameter_group_name text
);


ALTER TABLE public.aws_redshift_cluster_parameter_groups OWNER TO postgres;

--
-- Name: aws_redshift_cluster_parameters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_parameters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_parameter_group_cq_id uuid,
    allowed_values text,
    apply_type text,
    data_type text,
    description text,
    is_modifiable boolean,
    minimum_engine_version text,
    parameter_name text,
    parameter_value text,
    source text
);


ALTER TABLE public.aws_redshift_cluster_parameters OWNER TO postgres;

--
-- Name: aws_redshift_cluster_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    cluster_security_group_name text,
    status text
);


ALTER TABLE public.aws_redshift_cluster_security_groups OWNER TO postgres;

--
-- Name: aws_redshift_cluster_vpc_security_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_cluster_vpc_security_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    status text,
    vpc_security_group_id text
);


ALTER TABLE public.aws_redshift_cluster_vpc_security_groups OWNER TO postgres;

--
-- Name: aws_redshift_clusters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_clusters (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    allow_version_upgrade boolean,
    automated_snapshot_retention_period integer,
    availability_zone text,
    availability_zone_relocation_status text,
    cluster_availability_status text,
    cluster_create_time timestamp without time zone,
    id text,
    cluster_namespace_arn text,
    cluster_public_key text,
    cluster_revision_number text,
    cluster_snapshot_copy_status_destination_region text,
    cluster_snapshot_copy_status_manual_snapshot_retention_period integer,
    cluster_snapshot_copy_status_retention_period bigint,
    cluster_snapshot_copy_status_snapshot_copy_grant_name text,
    cluster_status text,
    cluster_subnet_group_name text,
    cluster_version text,
    db_name text,
    data_transfer_progress_current_rate_in_mega_bytes_per_second double precision,
    data_transfer_progress_data_transferred_in_mega_bytes bigint,
    data_transfer_progress_elapsed_time_in_seconds bigint,
    data_transfer_progress_estimated_time_to_completion_in_seconds bigint,
    data_transfer_progress_status text,
    data_transfer_progress_total_data_in_mega_bytes bigint,
    elastic_ip_status_elastic_ip text,
    elastic_ip_status text,
    elastic_resize_number_of_node_options text,
    encrypted boolean,
    endpoint_address text,
    endpoint_port integer,
    enhanced_vpc_routing boolean,
    expected_next_snapshot_schedule_time timestamp without time zone,
    expected_next_snapshot_schedule_time_status text,
    hsm_status_hsm_client_certificate_identifier text,
    hsm_status_hsm_configuration_identifier text,
    hsm_status text,
    kms_key_id text,
    maintenance_track_name text,
    manual_snapshot_retention_period integer,
    master_username text,
    modify_status text,
    next_maintenance_window_start_time timestamp without time zone,
    node_type text,
    number_of_nodes integer,
    pending_actions text[],
    pending_modified_values_automated_snapshot_retention_period integer,
    pending_modified_values_cluster_identifier text,
    pending_modified_values_cluster_type text,
    pending_modified_values_cluster_version text,
    pending_modified_values_encryption_type text,
    pending_modified_values_enhanced_vpc_routing boolean,
    pending_modified_values_maintenance_track_name text,
    pending_modified_values_master_user_password text,
    pending_modified_values_node_type text,
    pending_modified_values_number_of_nodes integer,
    pending_modified_values_publicly_accessible boolean,
    preferred_maintenance_window text,
    publicly_accessible boolean,
    resize_info_allow_cancel_resize boolean,
    resize_info_resize_type text,
    restore_status_current_restore_rate_in_mega_bytes_per_second double precision,
    restore_status_elapsed_time_in_seconds bigint,
    restore_status_estimated_time_to_completion_in_seconds bigint,
    restore_status_progress_in_mega_bytes bigint,
    restore_status_snapshot_size_in_mega_bytes bigint,
    restore_status text,
    snapshot_schedule_identifier text,
    snapshot_schedule_state text,
    tags jsonb,
    total_storage_capacity_in_mega_bytes bigint,
    vpc_id text,
    logging_status jsonb
);


ALTER TABLE public.aws_redshift_clusters OWNER TO postgres;

--
-- Name: aws_redshift_event_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_event_subscriptions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    arn text NOT NULL,
    account_id text,
    region text,
    id text,
    customer_aws_id text,
    enabled boolean,
    event_categories_list text[],
    severity text,
    sns_topic_arn text,
    source_ids_list text[],
    source_type text,
    status text,
    subscription_creation_time timestamp without time zone,
    tags jsonb
);


ALTER TABLE public.aws_redshift_event_subscriptions OWNER TO postgres;

--
-- Name: aws_redshift_snapshot_accounts_with_restore_access; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_snapshot_accounts_with_restore_access (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    snapshot_cq_id uuid,
    account_alias text,
    account_id text
);


ALTER TABLE public.aws_redshift_snapshot_accounts_with_restore_access OWNER TO postgres;

--
-- Name: aws_redshift_snapshots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_snapshots (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    cluster_cq_id uuid,
    arn text,
    actual_incremental_backup_size double precision,
    availability_zone text,
    backup_progress double precision,
    cluster_create_time timestamp without time zone,
    cluster_identifier text,
    cluster_version text,
    current_backup_rate double precision,
    db_name text,
    elapsed_time bigint,
    encrypted boolean,
    encrypted_with_hsm boolean,
    engine_full_version text,
    enhanced_vpc_routing boolean,
    estimated_seconds_to_completion bigint,
    kms_key_id text,
    maintenance_track_name text,
    manual_snapshot_remaining_days integer,
    manual_snapshot_retention_period integer,
    master_username text,
    node_type text,
    number_of_nodes integer,
    owner_account text,
    port integer,
    restorable_node_types text[],
    snapshot_create_time timestamp without time zone,
    snapshot_identifier text,
    snapshot_retention_start_time timestamp without time zone,
    snapshot_type text,
    source_region text,
    status text,
    total_backup_size_in_mega_bytes double precision,
    vpc_id text,
    tags jsonb
);


ALTER TABLE public.aws_redshift_snapshots OWNER TO postgres;

--
-- Name: aws_redshift_subnet_group_subnets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_subnet_group_subnets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    subnet_group_cq_id uuid,
    subnet_availability_zone_name text,
    subnet_availability_zone_supported_platforms text[],
    subnet_identifier text,
    subnet_status text
);


ALTER TABLE public.aws_redshift_subnet_group_subnets OWNER TO postgres;

--
-- Name: aws_redshift_subnet_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_redshift_subnet_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    cluster_subnet_group_name text,
    description text,
    subnet_group_status text,
    tags jsonb,
    vpc_id text
);


ALTER TABLE public.aws_redshift_subnet_groups OWNER TO postgres;

--
-- Name: aws_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_regions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    enabled boolean,
    endpoint text,
    opt_in_status text,
    region text,
    partition text
);


ALTER TABLE public.aws_regions OWNER TO postgres;

--
-- Name: aws_route53_domain_nameservers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_domain_nameservers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    domain_cq_id uuid,
    name text,
    glue_ips text[]
);


ALTER TABLE public.aws_route53_domain_nameservers OWNER TO postgres;

--
-- Name: aws_route53_domains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_domains (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    admin_contact_address_line1 text,
    admin_contact_address_line2 text,
    admin_contact_city text,
    admin_contact_type text,
    admin_contact_country_code text,
    admin_contact_email text,
    admin_contact_fax text,
    admin_contact_first_name text,
    admin_contact_last_name text,
    admin_contact_organization_name text,
    admin_contact_phone_number text,
    admin_contact_state text,
    admin_contact_zip_code text,
    admin_contact_extra_params jsonb,
    domain_name text NOT NULL,
    registrant_contact_address_line1 text,
    registrant_contact_address_line2 text,
    registrant_contact_city text,
    registrant_contact_type text,
    registrant_contact_country_code text,
    registrant_contact_email text,
    registrant_contact_fax text,
    registrant_contact_first_name text,
    registrant_contact_last_name text,
    registrant_contact_organization_name text,
    registrant_contact_phone_number text,
    registrant_contact_state text,
    registrant_contact_zip_code text,
    registrant_contact_extra_params jsonb,
    tech_contact_address_line1 text,
    tech_contact_address_line2 text,
    tech_contact_city text,
    tech_contact_type text,
    tech_contact_country_code text,
    tech_contact_email text,
    tech_contact_fax text,
    tech_contact_first_name text,
    tech_contact_last_name text,
    tech_contact_organization_name text,
    tech_contact_phone_number text,
    tech_contact_state text,
    tech_contact_zip_code text,
    tech_contact_extra_params jsonb,
    abuse_contact_email text,
    abuse_contact_phone text,
    admin_privacy boolean,
    auto_renew boolean,
    creation_date timestamp without time zone,
    dns_sec text,
    expiration_date timestamp without time zone,
    registrant_privacy boolean,
    registrar_name text,
    registrar_url text,
    registry_domain_id text,
    reseller text,
    status_list text[],
    tech_privacy boolean,
    updated_date timestamp without time zone,
    who_is_server text,
    tags jsonb
);


ALTER TABLE public.aws_route53_domains OWNER TO postgres;

--
-- Name: aws_route53_health_checks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_health_checks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    cloud_watch_alarm_configuration_dimensions jsonb,
    tags jsonb,
    caller_reference text,
    type text,
    alarm_identifier_name text,
    alarm_identifier_region text,
    child_health_checks text[],
    disabled boolean,
    enable_sni boolean,
    failure_threshold integer,
    fully_qualified_domain_name text,
    health_threshold integer,
    ip_address text,
    insufficient_data_health_status text,
    inverted boolean,
    measure_latency boolean,
    port integer,
    regions text[],
    request_interval integer,
    resource_path text,
    search_string text,
    health_check_version bigint,
    id text NOT NULL,
    cloud_watch_alarm_config_comparison_operator text,
    cloud_watch_alarm_config_evaluation_periods integer,
    cloud_watch_alarm_config_metric_name text,
    cloud_watch_alarm_config_namespace text,
    cloud_watch_alarm_config_period integer,
    cloud_watch_alarm_config_statistic text,
    cloud_watch_alarm_config_threshold double precision,
    linked_service_description text,
    linked_service_service_principal text,
    arn text
);


ALTER TABLE public.aws_route53_health_checks OWNER TO postgres;

--
-- Name: aws_route53_hosted_zone_query_logging_configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_hosted_zone_query_logging_configs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    hosted_zone_cq_id uuid,
    cloud_watch_logs_log_group_arn text,
    id text,
    arn text
);


ALTER TABLE public.aws_route53_hosted_zone_query_logging_configs OWNER TO postgres;

--
-- Name: aws_route53_hosted_zone_resource_record_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_hosted_zone_resource_record_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    hosted_zone_cq_id uuid,
    resource_records text[],
    name text,
    type text,
    dns_name text,
    evaluate_target_health boolean,
    failover text,
    geo_location_continent_code text,
    geo_location_country_code text,
    geo_location_subdivision_code text,
    health_check_id text,
    multi_value_answer boolean,
    region text,
    set_identifier text,
    ttl bigint,
    traffic_policy_instance_id text,
    weight bigint
);


ALTER TABLE public.aws_route53_hosted_zone_resource_record_sets OWNER TO postgres;

--
-- Name: aws_route53_hosted_zone_traffic_policy_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_hosted_zone_traffic_policy_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    hosted_zone_cq_id uuid,
    id text,
    message text,
    name text,
    state text,
    ttl bigint,
    traffic_policy_id text,
    traffic_policy_type text,
    traffic_policy_version integer,
    arn text
);


ALTER TABLE public.aws_route53_hosted_zone_traffic_policy_instances OWNER TO postgres;

--
-- Name: aws_route53_hosted_zone_vpc_association_authorizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_hosted_zone_vpc_association_authorizations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    hosted_zone_cq_id uuid,
    vpc_id text,
    vpc_region text,
    vpc_arn text
);


ALTER TABLE public.aws_route53_hosted_zone_vpc_association_authorizations OWNER TO postgres;

--
-- Name: aws_route53_hosted_zones; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_hosted_zones (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    tags jsonb,
    arn text,
    delegation_set_id text,
    caller_reference text,
    id text NOT NULL,
    name text,
    config_comment text,
    config_private_zone boolean,
    linked_service_description text,
    linked_service_principal text,
    resource_record_set_count bigint
);


ALTER TABLE public.aws_route53_hosted_zones OWNER TO postgres;

--
-- Name: aws_route53_reusable_delegation_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_reusable_delegation_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    name_servers text[],
    caller_reference text,
    id text NOT NULL
);


ALTER TABLE public.aws_route53_reusable_delegation_sets OWNER TO postgres;

--
-- Name: aws_route53_traffic_policies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_traffic_policies (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    id text NOT NULL,
    latest_version integer,
    name text,
    traffic_policy_count integer,
    type text,
    arn text
);


ALTER TABLE public.aws_route53_traffic_policies OWNER TO postgres;

--
-- Name: aws_route53_traffic_policy_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_route53_traffic_policy_versions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    traffic_policy_cq_id uuid,
    document jsonb,
    id text,
    name text,
    type text,
    version integer,
    comment text
);


ALTER TABLE public.aws_route53_traffic_policy_versions OWNER TO postgres;

--
-- Name: aws_s3_account_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_account_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    config_exists boolean,
    block_public_acls boolean,
    block_public_policy boolean,
    ignore_public_acls boolean,
    restrict_public_buckets boolean
);


ALTER TABLE public.aws_s3_account_config OWNER TO postgres;

--
-- Name: aws_s3_bucket_cors_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_bucket_cors_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    bucket_cq_id uuid,
    allowed_methods text[],
    allowed_origins text[],
    allowed_headers text[],
    expose_headers text[],
    id text,
    max_age_seconds integer
);


ALTER TABLE public.aws_s3_bucket_cors_rules OWNER TO postgres;

--
-- Name: aws_s3_bucket_encryption_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_bucket_encryption_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    bucket_cq_id uuid,
    sse_algorithm text,
    kms_master_key_id text,
    bucket_key_enabled boolean
);


ALTER TABLE public.aws_s3_bucket_encryption_rules OWNER TO postgres;

--
-- Name: aws_s3_bucket_grants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_bucket_grants (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    bucket_cq_id uuid,
    type text,
    display_name text,
    email_address text,
    grantee_id text,
    uri text,
    permission text
);


ALTER TABLE public.aws_s3_bucket_grants OWNER TO postgres;

--
-- Name: aws_s3_bucket_lifecycles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_bucket_lifecycles (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    bucket_cq_id uuid,
    status text,
    abort_incomplete_multipart_upload_days_after_initiation integer,
    expiration_date timestamp without time zone,
    expiration_days integer,
    expiration_expired_object_delete_marker boolean,
    filter jsonb,
    id text,
    noncurrent_version_expiration_days integer,
    noncurrent_version_transitions jsonb,
    prefix text,
    transitions jsonb
);


ALTER TABLE public.aws_s3_bucket_lifecycles OWNER TO postgres;

--
-- Name: aws_s3_bucket_replication_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_bucket_replication_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    bucket_cq_id uuid,
    destination_bucket text,
    destination_access_control_translation_owner text,
    destination_account text,
    destination_encryption_configuration_replica_kms_key_id text,
    destination_metrics_status text,
    destination_metrics_event_threshold_minutes integer,
    destination_replication_time_status text,
    destination_replication_time_minutes integer,
    destination_storage_class text,
    status text,
    delete_marker_replication_status text,
    existing_object_replication_status text,
    filter jsonb,
    id text,
    prefix text,
    priority integer,
    source_replica_modifications_status text,
    source_sse_kms_encrypted_objects_status text
);


ALTER TABLE public.aws_s3_bucket_replication_rules OWNER TO postgres;

--
-- Name: aws_s3_buckets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_s3_buckets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    logging_target_prefix text,
    logging_target_bucket text,
    versioning_status text,
    versioning_mfa_delete text,
    policy jsonb,
    tags jsonb,
    creation_date timestamp without time zone,
    name text NOT NULL,
    block_public_acls boolean,
    block_public_policy boolean,
    ignore_public_acls boolean,
    restrict_public_buckets boolean,
    replication_role text,
    arn text,
    ownership_controls text[]
);


ALTER TABLE public.aws_s3_buckets OWNER TO postgres;

--
-- Name: aws_sagemaker_endpoint_configuration_production_variants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_endpoint_configuration_production_variants (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    endpoint_configuration_cq_id uuid,
    initial_instance_count integer,
    instance_type text,
    model_name text,
    variant_name text,
    accelerator_type text,
    core_dump_config_destination_s3_uri text,
    core_dump_config_kms_key_id text,
    initial_variant_weight double precision
);


ALTER TABLE public.aws_sagemaker_endpoint_configuration_production_variants OWNER TO postgres;

--
-- Name: aws_sagemaker_endpoint_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_endpoint_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    kms_key_id text,
    data_capture_config jsonb,
    tags jsonb,
    creation_time timestamp without time zone,
    arn text NOT NULL,
    name text
);


ALTER TABLE public.aws_sagemaker_endpoint_configurations OWNER TO postgres;

--
-- Name: aws_sagemaker_model_containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_model_containers (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    model_cq_id uuid,
    container_hostname text,
    environment jsonb,
    image text,
    image_config_repository_access_mode text,
    image_config_repository_auth_config_repo_creds_provider_arn text,
    mode text,
    model_data_url text,
    model_package_name text,
    multi_model_config_model_cache_setting text
);


ALTER TABLE public.aws_sagemaker_model_containers OWNER TO postgres;

--
-- Name: aws_sagemaker_model_vpc_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_model_vpc_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    model_cq_id uuid,
    security_group_ids text[],
    subnets text[]
);


ALTER TABLE public.aws_sagemaker_model_vpc_config OWNER TO postgres;

--
-- Name: aws_sagemaker_models; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_models (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    enable_network_isolation boolean,
    execution_role_arn text,
    inference_execution_config jsonb,
    primary_container jsonb,
    tags jsonb,
    creation_time timestamp without time zone,
    arn text NOT NULL,
    name text
);


ALTER TABLE public.aws_sagemaker_models OWNER TO postgres;

--
-- Name: aws_sagemaker_notebook_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_notebook_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    network_interface_id text,
    kms_key_id text,
    subnet_id text,
    volume_size_in_gb integer,
    accelerator_types text[],
    security_groups jsonb,
    direct_internet_access boolean,
    tags jsonb,
    arn text NOT NULL,
    name text,
    additional_code_repositories text[],
    creation_time timestamp without time zone,
    default_code_repository text,
    instance_type text,
    last_modified_time timestamp without time zone,
    notebook_instance_lifecycle_config_name text,
    notebook_instance_status text,
    url text
);


ALTER TABLE public.aws_sagemaker_notebook_instances OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_algorithm_specification; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_algorithm_specification (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    training_input_mode text,
    algorithm_name text,
    enable_sage_maker_metrics_time_series boolean,
    metric_definitions jsonb,
    training_image text
);


ALTER TABLE public.aws_sagemaker_training_job_algorithm_specification OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_debug_hook_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_debug_hook_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    s3_output_path text,
    collection_configurations jsonb,
    hook_parameters jsonb,
    local_path text
);


ALTER TABLE public.aws_sagemaker_training_job_debug_hook_config OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_debug_rule_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_debug_rule_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    rule_configuration_name text,
    rule_evaluator_image text,
    instance_type text,
    local_path text,
    rule_parameters jsonb,
    s3_output_path text,
    volume_size_in_gb integer
);


ALTER TABLE public.aws_sagemaker_training_job_debug_rule_configurations OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_debug_rule_evaluation_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_debug_rule_evaluation_statuses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    last_modified_time timestamp without time zone,
    rule_configuration_name text,
    rule_evaluation_job_arn text,
    rule_evaluation_status text,
    status_details text
);


ALTER TABLE public.aws_sagemaker_training_job_debug_rule_evaluation_statuses OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_input_data_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_input_data_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    channel_name text,
    data_source_file_directory_path text,
    data_source_file_system_access_mode text,
    data_source_file_system_id text,
    data_source_file_system_type text,
    data_source_s3_data_type text,
    data_source_s3_uri text,
    data_source_attribute_names text[],
    data_source_s3_data_distribution_type text,
    compression_type text,
    content_type text,
    input_mode text,
    record_wrapper_type text,
    shuffle_config_seed bigint
);


ALTER TABLE public.aws_sagemaker_training_job_input_data_config OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_profiler_rule_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_profiler_rule_configurations (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    rule_configuration_name text,
    rule_evaluator_image text,
    instance_type text,
    local_path text,
    rule_parameters jsonb,
    s3_output_path text,
    volume_size_in_gb integer
);


ALTER TABLE public.aws_sagemaker_training_job_profiler_rule_configurations OWNER TO postgres;

--
-- Name: aws_sagemaker_training_job_profiler_rule_evaluation_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_job_profiler_rule_evaluation_statuses (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    training_job_cq_id uuid,
    last_modified_time timestamp without time zone,
    rule_configuration_name text,
    rule_evaluation_job_arn text,
    rule_evaluation_status text,
    status_details text
);


ALTER TABLE public.aws_sagemaker_training_job_profiler_rule_evaluation_statuses OWNER TO postgres;

--
-- Name: aws_sagemaker_training_jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sagemaker_training_jobs (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    auto_ml_job_arn text,
    billable_time_in_seconds integer,
    enable_managed_spot_training boolean,
    enable_network_isolation boolean,
    enable_inter_container_traffic_encryption boolean,
    failure_reason text,
    labeling_job_arn text,
    last_modified_time timestamp without time zone,
    profiling_status text,
    role_arn text,
    secondary_status text,
    training_end_time timestamp without time zone,
    training_start_time timestamp without time zone,
    training_time_in_seconds integer,
    tuning_job_arn text,
    checkpoint_config jsonb,
    environment jsonb,
    experiment_config jsonb,
    hyper_parameters jsonb,
    model_artifacts jsonb,
    output_data_config jsonb,
    profiler_config jsonb,
    resource_config jsonb,
    stopping_condition jsonb,
    tensor_board_output_config jsonb,
    vpc_config jsonb,
    tags jsonb,
    creation_time timestamp without time zone,
    arn text NOT NULL,
    name text,
    training_job_status text,
    secondary_status_transitions jsonb,
    final_metric_data_list jsonb
);


ALTER TABLE public.aws_sagemaker_training_jobs OWNER TO postgres;

--
-- Name: aws_secretsmanager_secrets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_secretsmanager_secrets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    policy jsonb,
    replication_status jsonb,
    arn text NOT NULL,
    created_date timestamp without time zone,
    deleted_date timestamp without time zone,
    description text,
    kms_key_id text,
    last_accessed_date timestamp without time zone,
    last_changed_date timestamp without time zone,
    last_rotated_date timestamp without time zone,
    name text,
    owning_service text,
    primary_region text,
    rotation_enabled boolean,
    rotation_lambda_arn text,
    rotation_rules_automatically_after_days bigint,
    secret_versions_to_stages jsonb,
    tags jsonb
);


ALTER TABLE public.aws_secretsmanager_secrets OWNER TO postgres;

--
-- Name: aws_shield_attack_properties; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_attack_properties (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    attack_cq_id uuid,
    attack_layer text,
    attack_property_identifier text,
    top_contributors jsonb,
    total bigint,
    unit text
);


ALTER TABLE public.aws_shield_attack_properties OWNER TO postgres;

--
-- Name: aws_shield_attack_sub_resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_attack_sub_resources (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    attack_cq_id uuid,
    attack_vectors jsonb,
    counters jsonb,
    id text,
    type text
);


ALTER TABLE public.aws_shield_attack_sub_resources OWNER TO postgres;

--
-- Name: aws_shield_attacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_attacks (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    attack_counters jsonb,
    id text NOT NULL,
    end_time timestamp without time zone,
    mitigations text[],
    resource_arn text,
    start_time timestamp without time zone
);


ALTER TABLE public.aws_shield_attacks OWNER TO postgres;

--
-- Name: aws_shield_protection_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_protection_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    tags jsonb,
    aggregation text,
    members text[],
    pattern text,
    id text,
    arn text NOT NULL,
    resource_type text
);


ALTER TABLE public.aws_shield_protection_groups OWNER TO postgres;

--
-- Name: aws_shield_protections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_protections (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    application_automatic_response_configuration_status text,
    health_check_ids text[],
    id text,
    name text,
    arn text NOT NULL,
    resource_arn text
);


ALTER TABLE public.aws_shield_protections OWNER TO postgres;

--
-- Name: aws_shield_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_shield_subscriptions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    protection_group_limits_max_protection_groups integer,
    protection_group_limits_arbitrary_pattern_limits_max_members integer,
    protected_resource_type_limits jsonb,
    auto_renew text,
    end_time timestamp without time zone,
    limits jsonb,
    proactive_engagement_status text,
    start_time timestamp without time zone,
    arn text NOT NULL,
    time_commitment_in_seconds integer
);


ALTER TABLE public.aws_shield_subscriptions OWNER TO postgres;

--
-- Name: aws_sns_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sns_subscriptions (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    endpoint text NOT NULL,
    owner text NOT NULL,
    protocol text NOT NULL,
    arn text NOT NULL,
    topic_arn text NOT NULL
);


ALTER TABLE public.aws_sns_subscriptions OWNER TO postgres;

--
-- Name: aws_sns_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sns_topics (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    owner text,
    policy jsonb,
    delivery_policy jsonb,
    display_name text,
    subscriptions_confirmed bigint,
    subscriptions_deleted bigint,
    subscriptions_pending bigint,
    effective_delivery_policy jsonb,
    fifo_topic boolean,
    content_based_deduplication boolean,
    kms_master_key_id text,
    arn text NOT NULL,
    tags jsonb
);


ALTER TABLE public.aws_sns_topics OWNER TO postgres;

--
-- Name: aws_sqs_queues; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_sqs_queues (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    url text,
    policy jsonb,
    visibility_timeout integer,
    maximum_message_size integer,
    message_retention_period integer,
    approximate_number_of_messages integer,
    approximate_number_of_messages_not_visible integer,
    created_timestamp timestamp without time zone,
    last_modified_timestamp timestamp without time zone,
    arn text NOT NULL,
    approximate_number_of_messages_delayed integer,
    delay_seconds integer,
    receive_message_wait_time_seconds integer,
    redrive_policy jsonb,
    fifo_queue boolean,
    content_based_deduplication boolean,
    kms_master_key_id text,
    kms_data_key_reuse_period_seconds integer,
    deduplication_scope text,
    fifo_throughput_limit text,
    redrive_allow_policy jsonb,
    tags jsonb,
    unknown_fields jsonb
);


ALTER TABLE public.aws_sqs_queues OWNER TO postgres;

--
-- Name: aws_ssm_documents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ssm_documents (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    approved_version text,
    attachments_information jsonb,
    author text,
    created_date timestamp without time zone,
    default_version text,
    description text,
    display_name text,
    document_format text,
    document_type text,
    document_version text,
    hash text,
    hash_type text,
    latest_version text,
    name text,
    owner text,
    parameters jsonb,
    pending_review_version text,
    platform_types text[],
    requires jsonb,
    review_status text,
    schema_version text,
    sha1 text,
    status text,
    status_information text,
    target_type text,
    version_name text,
    review_information jsonb,
    tags jsonb,
    account_ids text[],
    account_sharing_info_list jsonb
);


ALTER TABLE public.aws_ssm_documents OWNER TO postgres;

--
-- Name: aws_ssm_instance_compliance_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ssm_instance_compliance_items (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    instance_cq_id uuid,
    compliance_type text,
    details jsonb,
    execution_summary_execution_time timestamp without time zone,
    execution_summary_execution_id text,
    execution_summary_execution_type text,
    id text,
    resource_id text,
    resource_type text,
    severity text,
    status text,
    title text
);


ALTER TABLE public.aws_ssm_instance_compliance_items OWNER TO postgres;

--
-- Name: aws_ssm_instances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_ssm_instances (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text NOT NULL,
    activation_id text,
    agent_version text,
    association_overview_detailed_status text,
    association_instance_status_aggregated_count jsonb,
    association_status text,
    computer_name text,
    ip_address inet,
    iam_role text,
    instance_id text,
    is_latest_version boolean,
    last_association_execution_date timestamp without time zone,
    last_ping_date_time timestamp without time zone,
    last_successful_association_execution_date timestamp without time zone,
    name text,
    ping_status text,
    platform_name text,
    platform_type text,
    platform_version text,
    registration_date timestamp without time zone,
    resource_type text
);


ALTER TABLE public.aws_ssm_instances OWNER TO postgres;

--
-- Name: aws_waf_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    rule_ids text[],
    tags jsonb,
    id text NOT NULL,
    metric_name text,
    name text
);


ALTER TABLE public.aws_waf_rule_groups OWNER TO postgres;

--
-- Name: aws_waf_rule_predicates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_rule_predicates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rule_cq_id uuid,
    data_id text,
    negated boolean,
    type text
);


ALTER TABLE public.aws_waf_rule_predicates OWNER TO postgres;

--
-- Name: aws_waf_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    arn text,
    tags jsonb,
    id text NOT NULL,
    metric_name text,
    name text
);


ALTER TABLE public.aws_waf_rules OWNER TO postgres;

--
-- Name: aws_waf_subscribed_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_subscribed_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    metric_name text,
    name text,
    rule_group_id text NOT NULL
);


ALTER TABLE public.aws_waf_subscribed_rule_groups OWNER TO postgres;

--
-- Name: aws_waf_web_acl_logging_configuration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_web_acl_logging_configuration (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    log_destination_configs text[],
    resource_arn text,
    redacted_fields jsonb
);


ALTER TABLE public.aws_waf_web_acl_logging_configuration OWNER TO postgres;

--
-- Name: aws_waf_web_acl_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_web_acl_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    priority integer,
    rule_id text,
    action_type text,
    excluded_rules text[],
    override_action_type text,
    type text
);


ALTER TABLE public.aws_waf_web_acl_rules OWNER TO postgres;

--
-- Name: aws_waf_web_acls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_waf_web_acls (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    tags jsonb,
    default_action_type text,
    id text NOT NULL,
    metric_name text,
    name text,
    arn text,
    logging_configuration text[]
);


ALTER TABLE public.aws_waf_web_acls OWNER TO postgres;

--
-- Name: aws_wafregional_rate_based_rule_match_predicates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_rate_based_rule_match_predicates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rate_based_rule_cq_id uuid,
    data_id text,
    negated boolean,
    type text
);


ALTER TABLE public.aws_wafregional_rate_based_rule_match_predicates OWNER TO postgres;

--
-- Name: aws_wafregional_rate_based_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_rate_based_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    tags jsonb,
    rate_key text,
    rate_limit bigint,
    id text NOT NULL,
    metric_name text,
    name text
);


ALTER TABLE public.aws_wafregional_rate_based_rules OWNER TO postgres;

--
-- Name: aws_wafregional_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    tags jsonb,
    id text NOT NULL,
    metric_name text,
    name text
);


ALTER TABLE public.aws_wafregional_rule_groups OWNER TO postgres;

--
-- Name: aws_wafregional_rule_predicates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_rule_predicates (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    rule_cq_id uuid,
    data_id text,
    negated boolean,
    type text
);


ALTER TABLE public.aws_wafregional_rule_predicates OWNER TO postgres;

--
-- Name: aws_wafregional_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    arn text,
    tags jsonb,
    id text NOT NULL,
    metric_name text,
    name text
);


ALTER TABLE public.aws_wafregional_rules OWNER TO postgres;

--
-- Name: aws_wafregional_web_acl_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_web_acl_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    priority integer,
    rule_id text,
    action text,
    excluded_rules text[],
    override_action text,
    type text
);


ALTER TABLE public.aws_wafregional_web_acl_rules OWNER TO postgres;

--
-- Name: aws_wafregional_web_acls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafregional_web_acls (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    tags jsonb,
    default_action text,
    id text NOT NULL,
    metric_name text,
    name text,
    arn text
);


ALTER TABLE public.aws_wafregional_web_acls OWNER TO postgres;

--
-- Name: aws_wafv2_ipsets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_ipsets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    scope text,
    arn text NOT NULL,
    addresses cidr[],
    ip_address_version text,
    id text,
    name text,
    description text,
    tags jsonb
);


ALTER TABLE public.aws_wafv2_ipsets OWNER TO postgres;

--
-- Name: aws_wafv2_managed_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_managed_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    scope text NOT NULL,
    available_labels text[],
    consumed_labels text[],
    capacity bigint,
    label_namespace text,
    rules jsonb,
    description text,
    name text NOT NULL,
    vendor_name text NOT NULL
);


ALTER TABLE public.aws_wafv2_managed_rule_groups OWNER TO postgres;

--
-- Name: aws_wafv2_regex_pattern_sets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_regex_pattern_sets (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    scope text,
    arn text NOT NULL,
    description text,
    id text,
    name text,
    regular_expression_list text[],
    tags jsonb
);


ALTER TABLE public.aws_wafv2_regex_pattern_sets OWNER TO postgres;

--
-- Name: aws_wafv2_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    scope text,
    tags jsonb,
    policy jsonb,
    arn text,
    capacity bigint,
    id text NOT NULL,
    name text,
    visibility_config_cloud_watch_metrics_enabled boolean,
    visibility_config_metric_name text,
    visibility_config_sampled_requests_enabled boolean,
    custom_response_bodies jsonb,
    description text,
    label_namespace text,
    rules jsonb,
    available_labels text[],
    consumed_labels text[]
);


ALTER TABLE public.aws_wafv2_rule_groups OWNER TO postgres;

--
-- Name: aws_wafv2_web_acl_logging_configuration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_web_acl_logging_configuration (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    log_destination_configs text[],
    resource_arn text,
    logging_filter jsonb,
    managed_by_firewall_manager boolean,
    redacted_fields jsonb
);


ALTER TABLE public.aws_wafv2_web_acl_logging_configuration OWNER TO postgres;

--
-- Name: aws_wafv2_web_acl_post_process_firewall_manager_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_web_acl_post_process_firewall_manager_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    statement jsonb,
    name text,
    override_action jsonb,
    priority integer,
    visibility_config_cloud_watch_metrics_enabled boolean,
    visibility_config_metric_name text,
    visibility_config_sampled_requests_enabled boolean
);


ALTER TABLE public.aws_wafv2_web_acl_post_process_firewall_manager_rule_groups OWNER TO postgres;

--
-- Name: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    statement jsonb,
    name text,
    override_action jsonb,
    priority integer,
    visibility_config_cloud_watch_metrics_enabled boolean,
    visibility_config_metric_name text,
    visibility_config_sampled_requests_enabled boolean
);


ALTER TABLE public.aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups OWNER TO postgres;

--
-- Name: aws_wafv2_web_acl_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_web_acl_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    web_acl_cq_id uuid,
    name text,
    priority integer,
    statement jsonb,
    visibility_config_cloud_watch_metrics_enabled boolean,
    visibility_config_metric_name text,
    visibility_config_sampled_requests_enabled boolean,
    action jsonb,
    override_action jsonb,
    labels text[]
);


ALTER TABLE public.aws_wafv2_web_acl_rules OWNER TO postgres;

--
-- Name: aws_wafv2_web_acls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_wafv2_web_acls (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text,
    scope text,
    resources_for_web_acl text[],
    tags jsonb,
    arn text,
    default_action jsonb,
    id text NOT NULL,
    name text,
    visibility_config_cloud_watch_metrics_enabled boolean,
    visibility_config_metric_name text,
    visibility_config_sampled_requests_enabled boolean,
    capacity bigint,
    custom_response_bodies jsonb,
    description text,
    label_namespace text,
    managed_by_firewall_manager boolean,
    logging_configuration text[]
);


ALTER TABLE public.aws_wafv2_web_acls OWNER TO postgres;

--
-- Name: aws_workspaces_directories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_workspaces_directories (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text,
    alias text,
    customer_user_name text,
    id text NOT NULL,
    name text,
    type text,
    dns_ip_addresses text[],
    iam_role_id text,
    ip_group_ids text[],
    registration_code text,
    change_compute_type text,
    increase_volume_size text,
    rebuild_workspace text,
    restart_workspace text,
    switch_running_mode text,
    state text,
    subnet_ids text[],
    tenancy text,
    device_type_android text,
    device_type_chrome_os text,
    device_type_ios text,
    device_type_linux text,
    device_type_osx text,
    device_type_web text,
    device_type_windows text,
    device_type_zero_client text,
    custom_security_group_id text,
    default_ou text,
    enable_internet_access boolean,
    enable_maintenance_mode boolean,
    enable_work_docs boolean,
    user_enabled_as_local_administrator boolean,
    workspace_security_group_id text
);


ALTER TABLE public.aws_workspaces_directories OWNER TO postgres;

--
-- Name: aws_workspaces_workspaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_workspaces_workspaces (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    arn text,
    bundle_id text,
    computer_name text,
    directory_id text,
    error_code text,
    error_message text,
    ip_address text,
    modification_states jsonb,
    root_volume_encryption_enabled boolean,
    state text,
    subnet_id text,
    user_name text,
    user_volume_encryption_enabled boolean,
    volume_encryption_key text,
    id text NOT NULL,
    compute_type_name text,
    root_volume_size_gib integer,
    running_mode text,
    running_mode_auto_stop_timeout_in_minutes integer,
    user_volume_size_gib integer
);


ALTER TABLE public.aws_workspaces_workspaces OWNER TO postgres;

--
-- Name: aws_xray_encryption_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_xray_encryption_config (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text NOT NULL,
    region text NOT NULL,
    key_id text,
    status text,
    type text
);


ALTER TABLE public.aws_xray_encryption_config OWNER TO postgres;

--
-- Name: aws_xray_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_xray_groups (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    filter_expression text,
    arn text NOT NULL,
    group_name text,
    insights_enabled boolean,
    notifications_enabled boolean
);


ALTER TABLE public.aws_xray_groups OWNER TO postgres;

--
-- Name: aws_xray_sampling_rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aws_xray_sampling_rules (
    cq_id uuid NOT NULL,
    cq_meta jsonb,
    account_id text,
    region text,
    tags jsonb,
    created_at timestamp without time zone,
    modified_at timestamp without time zone,
    fixed_rate double precision,
    http_method text,
    host text,
    priority integer,
    reservoir_size integer,
    resource_arn text,
    service_name text,
    service_type text,
    url_path text,
    version integer,
    attributes jsonb,
    arn text NOT NULL,
    rule_name text
);


ALTER TABLE public.aws_xray_sampling_rules OWNER TO postgres;

--
-- Data for Name: check_results; Type: TABLE DATA; Schema: cloudquery; Owner: postgres
--

COPY cloudquery.check_results (execution_id, execution_timestamp, name, selector, description, status, raw_results, error) FROM stdin;
\.


--
-- Data for Name: cloudquery_core_schema_migrations; Type: TABLE DATA; Schema: cloudquery; Owner: postgres
--

COPY cloudquery.cloudquery_core_schema_migrations (version, dirty) FROM stdin;
4	f
\.


--
-- Data for Name: fetches; Type: TABLE DATA; Schema: cloudquery; Owner: postgres
--

COPY cloudquery.fetches (id, fetch_id, start, finish, total_resource_count, total_errors_count, provider_name, provider_version, is_success, results, provider_alias, core_version, created_at) FROM stdin;
9d83c7b8-e58f-11ec-b238-aed9c18def80	7d5995da-ba7c-4336-b667-26f032b2b135	2022-06-06 11:55:32.510867	2022-06-06 11:55:52.738599	11	0	aws	v0.12.4	t	[{"error": "no errors", "status": "COMPLETE", "resource_name": "s3.accounts", "resource_count": 1, "finished_resources": null}, {"error": "no errors", "status": "COMPLETE", "resource_name": "s3.buckets", "resource_count": 10, "finished_resources": null}]	aws	development	2022-06-06 11:55:52.738599
\.


--
-- Data for Name: policy_executions; Type: TABLE DATA; Schema: cloudquery; Owner: postgres
--

COPY cloudquery.policy_executions (id, "timestamp", scheme, location, policy_name, selector, sha256_hash, version, checks_total, checks_failed, checks_passed) FROM stdin;
\.


--
-- Data for Name: providers; Type: TABLE DATA; Schema: cloudquery; Owner: postgres
--

COPY cloudquery.providers (source, name, version, v_major, v_minor, v_patch, v_pre, v_meta, tables, signatures) FROM stdin;
cloudquery	aws	v0.12.4	0	12	4			{"ec2.eips": ["aws_ec2_eips"], "ec2.vpcs": ["aws_ec2_vpcs", "aws_ec2_vpc_cidr_block_association_sets", "aws_ec2_vpc_ipv6_cidr_block_association_sets"], "kms.keys": ["aws_kms_keys"], "ec2.hosts": ["aws_ec2_hosts", "aws_ec2_host_available_instance_capacity", "aws_ec2_host_instances"], "iam.roles": ["aws_iam_roles", "aws_iam_role_policies"], "iam.users": ["aws_iam_users", "aws_iam_user_access_keys", "aws_iam_user_groups", "aws_iam_user_attached_policies", "aws_iam_user_policies"], "waf.rules": ["aws_waf_rules", "aws_waf_rule_predicates"], "ec2.images": ["aws_ec2_images", "aws_ec2_image_block_device_mappings"], "iam.groups": ["aws_iam_groups", "aws_iam_group_policies"], "iot.things": ["aws_iot_things"], "mq.brokers": ["aws_mq_brokers", "aws_mq_broker_configurations", "aws_mq_broker_configuration_revisions", "aws_mq_broker_users"], "s3.buckets": ["aws_s3_buckets", "aws_s3_bucket_grants", "aws_s3_bucket_cors_rules", "aws_s3_bucket_encryption_rules", "aws_s3_bucket_replication_rules", "aws_s3_bucket_lifecycles"], "sns.topics": ["aws_sns_topics"], "sqs.queues": ["aws_sqs_queues"], "aws.regions": ["aws_regions"], "ec2.subnets": ["aws_ec2_subnets", "aws_ec2_subnet_ipv6_cidr_block_association_sets"], "fsx.backups": ["aws_fsx_backups"], "iot.streams": ["aws_iot_streams", "aws_iot_stream_files"], "s3.accounts": ["aws_s3_account_config"], "xray.groups": ["aws_xray_groups"], "backup.plans": ["aws_backup_plans", "aws_backup_plan_rules", "aws_backup_plan_selections"], "dax.clusters": ["aws_dax_clusters", "aws_dax_cluster_nodes"], "ecs.clusters": ["aws_ecs_clusters", "aws_ecs_cluster_attachments", "aws_ecs_cluster_tasks", "aws_ecs_cluster_task_attachments", "aws_ecs_cluster_task_containers", "aws_ecs_cluster_services", "aws_ecs_cluster_service_deployments", "aws_ecs_cluster_service_events", "aws_ecs_cluster_service_load_balancers", "aws_ecs_cluster_service_service_registries", "aws_ecs_cluster_service_task_sets", "aws_ecs_cluster_service_task_set_load_balancers", "aws_ecs_cluster_service_task_set_service_registries", "aws_ecs_cluster_container_instances", "aws_ecs_cluster_container_instance_attachments", "aws_ecs_cluster_container_instance_attributes", "aws_ecs_cluster_container_instance_health_status_details", "aws_ecs_cluster_container_instance_registered_resources", "aws_ecs_cluster_container_instance_remaining_resources"], "eks.clusters": ["aws_eks_clusters", "aws_eks_cluster_encryption_configs", "aws_eks_cluster_loggings"], "emr.clusters": ["aws_emr_clusters"], "iam.accounts": ["aws_accounts"], "iam.policies": ["aws_iam_policies", "aws_iam_policy_versions"], "iot.policies": ["aws_iot_policies"], "qldb.ledgers": ["aws_qldb_ledgers", "aws_qldb_ledger_journal_kinesis_streams", "aws_qldb_ledger_journal_s3_exports"], "rds.clusters": ["aws_rds_clusters", "aws_rds_cluster_associated_roles", "aws_rds_cluster_db_cluster_members", "aws_rds_cluster_domain_memberships", "aws_rds_cluster_vpc_security_groups"], "waf.web_acls": ["aws_waf_web_acls", "aws_waf_web_acl_rules", "aws_waf_web_acl_logging_configuration"], "wafv2.ipsets": ["aws_wafv2_ipsets"], "backup.vaults": ["aws_backup_vaults", "aws_backup_vault_recovery_points"], "ec2.flow_logs": ["aws_ec2_flow_logs"], "ec2.instances": ["aws_ec2_instances", "aws_ec2_instance_block_device_mappings", "aws_ec2_instance_elastic_gpu_associations", "aws_ec2_instance_elastic_inference_accelerator_associations", "aws_ec2_instance_network_interfaces", "aws_ec2_instance_network_interface_groups", "aws_ec2_instance_network_interface_ipv6_addresses", "aws_ec2_instance_network_interface_private_ip_addresses", "aws_ec2_instance_product_codes", "aws_ec2_instance_security_groups"], "lambda.layers": ["aws_lambda_layers", "aws_lambda_layer_versions", "aws_lambda_layer_version_policies"], "rds.instances": ["aws_rds_instances", "aws_rds_instance_associated_roles", "aws_rds_instance_db_instance_automated_backups_replications", "aws_rds_instance_db_parameter_groups", "aws_rds_instance_db_security_groups", "aws_rds_instance_db_subnet_group_subnets", "aws_rds_instance_domain_memberships", "aws_rds_instance_option_group_memberships", "aws_rds_instance_vpc_security_groups"], "ssm.documents": ["aws_ssm_documents"], "ssm.instances": ["aws_ssm_instances", "aws_ssm_instance_compliance_items"], "shield.attacks": ["aws_shield_attacks", "aws_shield_attack_properties", "aws_shield_attack_sub_resources"], "wafv2.web_acls": ["aws_wafv2_web_acls", "aws_wafv2_web_acl_rules", "aws_wafv2_web_acl_post_process_firewall_manager_rule_groups", "aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups", "aws_wafv2_web_acl_logging_configuration"], "dynamodb.tables": ["aws_dynamodb_tables", "aws_dynamodb_table_global_secondary_indexes", "aws_dynamodb_table_local_secondary_indexes", "aws_dynamodb_table_replicas", "aws_dynamodb_table_replica_auto_scalings", "aws_dynamodb_table_continuous_backups"], "ec2.byoip_cidrs": ["aws_ec2_byoip_cidrs"], "ec2.ebs_volumes": ["aws_ec2_ebs_volumes", "aws_ec2_ebs_volume_attachments"], "efs.filesystems": ["aws_efs_filesystems"], "iot.thing_types": ["aws_iot_thing_types"], "iot.topic_rules": ["aws_iot_topic_rules", "aws_iot_topic_rule_actions"], "lambda.runtimes": ["aws_lambda_runtimes"], "route53.domains": ["aws_route53_domains", "aws_route53_domain_nameservers"], "waf.rule_groups": ["aws_waf_rule_groups"], "acm.certificates": ["aws_acm_certificates"], "ec2.nat_gateways": ["aws_ec2_nat_gateways", "aws_ec2_nat_gateway_addresses"], "ec2.network_acls": ["aws_ec2_network_acls", "aws_ec2_network_acl_associations", "aws_ec2_network_acl_entries"], "ec2.route_tables": ["aws_ec2_route_tables", "aws_ec2_route_table_associations", "aws_ec2_route_table_propagating_vgws", "aws_ec2_route_table_routes"], "ec2.vpn_gateways": ["aws_ec2_vpn_gateways", "aws_ec2_vpc_attachment"], "ecr.repositories": ["aws_ecr_repositories", "aws_ecr_repository_images"], "iot.certificates": ["aws_iot_certificates"], "iot.thing_groups": ["aws_iot_thing_groups"], "lambda.functions": ["aws_lambda_functions", "aws_lambda_function_file_system_configs", "aws_lambda_function_layers", "aws_lambda_function_event_invoke_configs", "aws_lambda_function_aliases", "aws_lambda_function_versions", "aws_lambda_function_version_file_system_configs", "aws_lambda_function_version_layers", "aws_lambda_function_concurrency_configs", "aws_lambda_function_event_source_mappings"], "rds.certificates": ["aws_rds_certificates"], "rds.db_snapshots": ["aws_rds_db_snapshots"], "sagemaker.models": ["aws_sagemaker_models", "aws_sagemaker_model_containers", "aws_sagemaker_model_vpc_config"], "apigatewayv2.apis": ["aws_apigatewayv2_apis", "aws_apigatewayv2_api_authorizers", "aws_apigatewayv2_api_deployments", "aws_apigatewayv2_api_integrations", "aws_apigatewayv2_api_integration_responses", "aws_apigatewayv2_api_models", "aws_apigatewayv2_api_routes", "aws_apigatewayv2_api_route_responses", "aws_apigatewayv2_api_stages"], "cloudtrail.trails": ["aws_cloudtrail_trails", "aws_cloudtrail_trail_event_selectors"], "cloudwatch.alarms": ["aws_cloudwatch_alarms", "aws_cloudwatch_alarm_metrics"], "ec2.ebs_snapshots": ["aws_ec2_ebs_snapshots"], "ec2.vpc_endpoints": ["aws_ec2_vpc_endpoints", "aws_ec2_vpc_endpoint_dns_entries", "aws_ec2_vpc_endpoint_groups"], "redshift.clusters": ["aws_redshift_clusters", "aws_redshift_cluster_nodes", "aws_redshift_cluster_parameter_groups", "aws_redshift_cluster_parameters", "aws_redshift_cluster_parameter_group_status_lists", "aws_redshift_cluster_security_groups", "aws_redshift_cluster_deferred_maintenance_windows", "aws_redshift_cluster_endpoint_vpc_endpoints", "aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces", "aws_redshift_cluster_iam_roles", "aws_redshift_cluster_vpc_security_groups", "aws_redshift_snapshots", "aws_redshift_snapshot_accounts_with_restore_access"], "sns.subscriptions": ["aws_sns_subscriptions"], "wafregional.rules": ["aws_wafregional_rules", "aws_wafregional_rule_predicates"], "wafv2.rule_groups": ["aws_wafv2_rule_groups"], "athena.work_groups": ["aws_athena_work_groups", "aws_athena_work_group_prepared_statements", "aws_athena_work_group_query_executions", "aws_athena_work_group_named_queries"], "autoscaling.groups": ["aws_autoscaling_groups", "aws_autoscaling_group_instances", "aws_autoscaling_group_tags", "aws_autoscaling_group_scaling_policies", "aws_autoscaling_group_lifecycle_hooks"], "codebuild.projects": ["aws_codebuild_projects", "aws_codebuild_project_environment_variables", "aws_codebuild_project_file_system_locations", "aws_codebuild_project_secondary_artifacts", "aws_codebuild_project_secondary_sources"], "cognito.user_pools": ["aws_cognito_user_pools", "aws_cognito_user_pool_schema_attributes", "aws_cognito_user_pool_identity_providers"], "directconnect.lags": ["aws_directconnect_lags", "aws_directconnect_lag_mac_sec_keys"], "iot.billing_groups": ["aws_iot_billing_groups"], "shield.protections": ["aws_shield_protections"], "apigateway.api_keys": ["aws_apigateway_api_keys"], "ec2.regional_config": ["aws_ec2_regional_config"], "ec2.security_groups": ["aws_ec2_security_groups", "aws_ec2_security_group_ip_permissions", "aws_ec2_security_group_ip_permission_ip_ranges", "aws_ec2_security_group_ip_permission_prefix_list_ids", "aws_ec2_security_group_ip_permission_user_id_group_pairs"], "elbv2.target_groups": ["aws_elbv2_target_groups"], "guardduty.detectors": ["aws_guardduty_detectors", "aws_guardduty_detector_members"], "iot.ca_certificates": ["aws_iot_ca_certificates"], "xray.sampling_rules": ["aws_xray_sampling_rules"], "apigateway.rest_apis": ["aws_apigateway_rest_apis", "aws_apigateway_rest_api_authorizers", "aws_apigateway_rest_api_deployments", "aws_apigateway_rest_api_documentation_parts", "aws_apigateway_rest_api_documentation_versions", "aws_apigateway_rest_api_gateway_responses", "aws_apigateway_rest_api_models", "aws_apigateway_rest_api_request_validators", "aws_apigateway_rest_api_resources", "aws_apigateway_rest_api_stages"], "apigateway.vpc_links": ["aws_apigateway_vpc_links"], "athena.data_catalogs": ["aws_athena_data_catalogs", "aws_athena_data_catalog_databases", "aws_athena_data_catalog_database_tables", "aws_athena_data_catalog_database_table_columns", "aws_athena_data_catalog_database_table_partition_keys"], "ec2.transit_gateways": ["aws_ec2_transit_gateways", "aws_ec2_transit_gateway_attachments", "aws_ec2_transit_gateway_route_tables", "aws_ec2_transit_gateway_vpc_attachments", "aws_ec2_transit_gateway_peering_attachments", "aws_ec2_transit_gateway_multicast_domains"], "ecs.task_definitions": ["aws_ecs_task_definitions", "aws_ecs_task_definition_container_definitions", "aws_ecs_task_definition_volumes"], "elbv1.load_balancers": ["aws_elbv1_load_balancers", "aws_elbv1_load_balancer_backend_server_descriptions", "aws_elbv1_load_balancer_listeners", "aws_elbv1_load_balancer_policies_app_cookie_stickiness", "aws_elbv1_load_balancer_policies_lb_cookie_stickiness", "aws_elbv1_load_balancer_policies"], "elbv2.load_balancers": ["aws_elbv2_load_balancers", "aws_elbv2_listeners", "aws_elbv2_listener_certificates", "aws_elbv2_listener_default_actions", "aws_elbv2_listener_default_action_forward_config_target_groups", "aws_elbv2_load_balancer_availability_zones", "aws_elbv2_load_balancer_availability_zone_addresses", "aws_elbv2_load_balancer_attributes"], "rds.db_subnet_groups": ["aws_rds_subnet_groups", "aws_rds_subnet_group_subnets"], "route53.hosted_zones": ["aws_route53_hosted_zones", "aws_route53_hosted_zone_query_logging_configs", "aws_route53_hosted_zone_resource_record_sets", "aws_route53_hosted_zone_traffic_policy_instances", "aws_route53_hosted_zone_vpc_association_authorizations"], "shield.subscriptions": ["aws_shield_subscriptions"], "wafregional.web_acls": ["aws_wafregional_web_acls", "aws_wafregional_web_acl_rules"], "cloudformation.stacks": ["aws_cloudformation_stacks", "aws_cloudformation_stack_outputs", "aws_cloudformation_stack_resources"], "codepipeline.webhooks": ["aws_codepipeline_webhooks", "aws_codepipeline_webhook_filters"], "ec2.customer_gateways": ["aws_ec2_customer_gateways"], "ec2.instance_statuses": ["aws_ec2_instance_statuses", "aws_ec2_instance_status_events"], "ec2.internet_gateways": ["aws_ec2_internet_gateways", "aws_ec2_internet_gateway_attachments"], "elasticsearch.domains": ["aws_elasticsearch_domains"], "iam.password_policies": ["aws_iam_password_policies"], "rds.cluster_snapshots": ["aws_rds_cluster_snapshots"], "route53.health_checks": ["aws_route53_health_checks"], "workspaces.workspaces": ["aws_workspaces_workspaces"], "apigateway.usage_plans": ["aws_apigateway_usage_plans", "aws_apigateway_usage_plan_api_stages", "aws_apigateway_usage_plan_keys"], "apigatewayv2.vpc_links": ["aws_apigatewayv2_vpc_links"], "backup.global_settings": ["aws_backup_global_settings"], "backup.region_settings": ["aws_backup_region_settings"], "cloudwatchlogs.filters": ["aws_cloudwatchlogs_filters", "aws_cloudwatchlogs_filter_metric_transformations"], "codepipeline.pipelines": ["aws_codepipeline_pipelines", "aws_codepipeline_pipeline_stages", "aws_codepipeline_pipeline_stage_actions"], "cognito.identity_pools": ["aws_cognito_identity_pools", "aws_cognito_identity_pool_cognito_identity_providers"], "directconnect.gateways": ["aws_directconnect_gateways", "aws_directconnect_gateway_associations", "aws_directconnect_gateway_attachments"], "ec2.network_interfaces": ["aws_ec2_network_interfaces", "aws_ec2_network_interface_private_ip_addresses"], "organizations.accounts": ["aws_organizations_accounts"], "rds.db_security_groups": ["aws_rds_db_security_groups"], "redshift.subnet_groups": ["aws_redshift_subnet_groups", "aws_redshift_subnet_group_subnets"], "secretsmanager.secrets": ["aws_secretsmanager_secrets"], "workspaces.directories": ["aws_workspaces_directories"], "xray.encryption_config": ["aws_xray_encryption_config"], "apigateway.domain_names": ["aws_apigateway_domain_names", "aws_apigateway_domain_name_base_path_mappings"], "iam.server_certificates": ["aws_iam_server_certificates"], "iam.virtual_mfa_devices": ["aws_iam_virtual_mfa_devices"], "rds.db_parameter_groups": ["aws_rds_db_parameter_groups", "aws_rds_db_parameters"], "rds.event_subscriptions": ["aws_rds_event_subscriptions"], "sagemaker.training_jobs": ["aws_sagemaker_training_jobs", "aws_sagemaker_training_job_algorithm_specification", "aws_sagemaker_training_job_debug_hook_config", "aws_sagemaker_training_job_debug_rule_configurations", "aws_sagemaker_training_job_debug_rule_evaluation_statuses", "aws_sagemaker_training_job_input_data_config", "aws_sagemaker_training_job_profiler_rule_configurations", "aws_sagemaker_training_job_profiler_rule_evaluation_statuses"], "wafregional.rule_groups": ["aws_wafregional_rule_groups"], "accessanalyzer.analyzers": ["aws_access_analyzer_analyzers", "aws_access_analyzer_analyzer_findings", "aws_access_analyzer_analyzer_finding_sources", "aws_access_analyzer_analyzer_archive_rules"], "cloudfront.distributions": ["aws_cloudfront_distributions", "aws_cloudfront_distribution_default_cache_behavior_functions", "aws_cloudfront_distribution_origins", "aws_cloudfront_distribution_cache_behaviors", "aws_cloudfront_distribution_cache_behavior_lambda_functions", "aws_cloudfront_distribution_custom_error_responses", "aws_cloudfront_distribution_origin_groups"], "config.conformance_packs": ["aws_config_conformance_packs", "aws_config_conformance_pack_rule_compliances"], "route53.traffic_policies": ["aws_route53_traffic_policies", "aws_route53_traffic_policy_versions"], "wafv2.regex_pattern_sets": ["aws_wafv2_regex_pattern_sets"], "apigatewayv2.domain_names": ["aws_apigatewayv2_domain_names", "aws_apigatewayv2_domain_name_configurations", "aws_apigatewayv2_domain_name_rest_api_mappings"], "cloudfront.cache_policies": ["aws_cloudfront_cache_policies"], "directconnect.connections": ["aws_directconnect_connections", "aws_directconnect_connection_mac_sec_keys"], "dms.replication_instances": ["aws_dms_replication_instances", "aws_dms_replication_instance_replication_subnet_group_subnets", "aws_dms_replication_instance_vpc_security_groups"], "shield.protections_groups": ["aws_shield_protection_groups"], "wafv2.managed_rule_groups": ["aws_wafv2_managed_rule_groups"], "waf.subscribed_rule_groups": ["aws_waf_subscribed_rule_groups"], "ec2.vpc_peering_connections": ["aws_ec2_vpc_peering_connections"], "iam.saml_identity_providers": ["aws_iam_saml_identity_providers"], "rds.cluster_parameter_groups": ["aws_rds_cluster_parameter_groups", "aws_rds_cluster_parameters"], "redshift.event_subscriptions": ["aws_redshift_event_subscriptions"], "sagemaker.notebook_instances": ["aws_sagemaker_notebook_instances"], "wafregional.rate_based_rules": ["aws_wafregional_rate_based_rules", "aws_wafregional_rate_based_rule_match_predicates"], "autoscaling.scheduled_actions": ["aws_autoscaling_scheduled_actions"], "elasticbeanstalk.applications": ["aws_elasticbeanstalk_applications"], "elasticbeanstalk.environments": ["aws_elasticbeanstalk_environments", "aws_elasticbeanstalk_configuration_settings", "aws_elasticbeanstalk_configuration_setting_options", "aws_elasticbeanstalk_configuration_options", "aws_elasticbeanstalk_environment_links"], "apigateway.client_certificates": ["aws_apigateway_client_certificates"], "config.configuration_recorders": ["aws_config_configuration_recorders"], "directconnect.virtual_gateways": ["aws_directconnect_virtual_gateways"], "applicationautoscaling.policies": ["aws_applicationautoscaling_policies"], "emr.block_public_access_configs": ["aws_emr_block_public_access_configs", "aws_emr_block_public_access_config_port_ranges"], "directconnect.virtual_interfaces": ["aws_directconnect_virtual_interfaces", "aws_directconnect_virtual_interface_bgp_peers"], "route53.reusable_delegation_sets": ["aws_route53_reusable_delegation_sets"], "autoscaling.launch_configurations": ["aws_autoscaling_launch_configurations", "aws_autoscaling_launch_configuration_block_device_mappings"], "ec2.egress_only_internet_gateways": ["aws_ec2_egress_only_internet_gateways"], "sagemaker.endpoint_configurations": ["aws_sagemaker_endpoint_configurations", "aws_sagemaker_endpoint_configuration_production_variants"], "elasticbeanstalk.application_versions": ["aws_elasticbeanstalk_application_versions"], "iam.openid_connect_identity_providers": ["aws_iam_openid_connect_identity_providers"]}	{"ec2.eips": "18a4b3427a20f5a0a20a9d73d44ec9c3802ac5b42a4c45c6274610387c01fe0f", "ec2.vpcs": "bda850f70376ada6c62de1fe25d4b7186dfbca0271ec996d8b15050881219c25", "kms.keys": "1ec8a9b4e4651bddd91a32493b06ba295ac332c846d4198d843060c122253a52", "ec2.hosts": "4f479e77941281b2ad141aa6723cd6d818211548ee07ec614164ff4ea0cbea7d", "iam.roles": "3e9a0eeac3f9d593e888ca50f156047640f2988cf4a740f5d11281a39516cdba", "iam.users": "cc5dbe32e7dcf95c7618d2d5415b0bbba9549be2d725aec22706b438553d2f94", "waf.rules": "cd50e600893e36bbc96d91c1944ab6f762388dbce8da70d1e2bc519eaf98f271", "ec2.images": "19b9bb5647d2772e28f337a8e771764626751ddea8b2ec6bc0261270896001ca", "iam.groups": "2b5ce95aaab27b7fa30a369bd681701cffc53531df0dc6dd72ce0467a5d08d17", "iot.things": "4bb9173bcd2c8e1253a8abbd9f537fde9ffd5ec6d71df20015a1591b985ce2a6", "mq.brokers": "389feb2994e9b6d005c400c8ee4db4df5438e953d865d961e58aa588e02ff4e3", "s3.buckets": "cc49c36d342eba5fc5588cd110fa202c63ec591635e7dfa7d8276167f50c6e71", "sns.topics": "406f29d47a733f8d55cade4b26af6918f641f5bcac440285ea7e665a43b98038", "sqs.queues": "cdfb750ba56a9eb7cf58735ad6bbfd3ef10c78ad18a7acfd93f66513833374bc", "aws.regions": "5d368c00193ca98b070aaae30bf07ccb8ef197b83358f17f76d4bb9beda3d511", "ec2.subnets": "48338e5267a31675ad58cfcd5d9f1ad93f3f70f2ab3b1dee170d35d9ed15b1c8", "fsx.backups": "53daed3617987d44e077e95b8dcdf1a65281dc6c398a46dac567207dcf19ab2f", "iot.streams": "fb79e7b91ffe3b19ca93451045da4916e40fc2c94a01ef634daf869e7172c935", "s3.accounts": "acfb4d82bd4b4ad57b99b536bbd9b2242ece56364f96d63697819abedc479c13", "xray.groups": "66cd501e25416d8419df5fa8f6ec127202a8f0d3f89beb30e4e6b1556f5fdba7", "backup.plans": "6b1c7b385b22e6d86d1e7f848e25e3c48d82c1d9c327654ae9a05dc0ff40058a", "dax.clusters": "719fe4d61f47b726564a4a4b0cbbf604a9bce9f15941ce73ffe246f926bc8cf5", "ecs.clusters": "ab08de1a2eb0f32850363ffe6c591d4acd01dad14244b76c4e89acc2d07fb415", "eks.clusters": "29e7baae040adf64499fa8445b754a55ce154674ccd4f445da575de6ed49cc89", "emr.clusters": "8c12df21a8d17b9fbd9fa73b37618244e17d735f7643b09e000da9c67f631d67", "iam.accounts": "11bd6b5c39c1d3f803653b821609326410b823852e4e4aea2fac5dd3f36faddd", "iam.policies": "1f0c90846de348340512b4d0c802fc949ab598a55c8aa9694b62f4783b6e4222", "iot.policies": "40ce8b0d63274b9388b6faec878dbf8a7d52866af1827edb190579aa13537320", "qldb.ledgers": "3b7a0a9ea1e73b4329afec78bafddec2f59eeb088a7d7648e5d7cca6f54bab2f", "rds.clusters": "c53216025236a7c442a64a49e761ad230180b79ddd616959c3e187b0b6368aec", "waf.web_acls": "7681adad76f131b4bec44fc42c3c05d3a57d8f354b204977069b2471f6b946b1", "wafv2.ipsets": "7e02fa8b6121b3fcd3a41e9f3929b9fc570da55de2e7e3bf631ce47a268e94dc", "backup.vaults": "626301c0c379d372a846d350e3624914ec544e2139815d09f5360d35287ccd35", "ec2.flow_logs": "fd5b595e1e2cc5baaa164063669cd7dd4e88b83e02f7046734d5928eae39c762", "ec2.instances": "f5899f5b35c1ed916f45cf8c9b496a7f192f5527453d585cbb7f779b378f3c90", "lambda.layers": "056b320c99bbd41acacfde431e7af6e77a2db47ee9c249e3d90af672142e20d5", "rds.instances": "abe11c82adf29c3826b47a77deca42ed86104f7322afda573180b66c9c7a0615", "ssm.documents": "edd54f5769be03ac77ce16c516c323bc1a97a96fdcf10cc5260b49fc1b1e6126", "ssm.instances": "1b6176ddba20c0d138c3ffb077f6ab1ae0c81268a0e2078fdb3122a3c25b73a5", "shield.attacks": "fe9eb6d5e521461563e06fbbf4a42b10dad4abfdfd4c21105eac3193d5673056", "wafv2.web_acls": "a738e5980870dcc891c2d27e2d76a3295ae62e9cae23f61971456b70ea18d58f", "dynamodb.tables": "55941dd9cd359092e9de36cb0a9016d7a08e0a8aecfd68b62b4a2a0d32f0f560", "ec2.byoip_cidrs": "e9db7c8feef4ff9d168b3efbd1d84e19256cc20a3d501ab31f346efc65ded0a8", "ec2.ebs_volumes": "97ecb1e772d0c256d22fe1220f38496f291c7dabf79091b0bceefa6e39b11fd3", "efs.filesystems": "c879a1596e39b8954e87b52d13ed659a72d109a2bb022e0dde3bdbaf3aee79bb", "iot.thing_types": "2cd46908f37e342e03953bc8c8a16883e42e80f08bbdf4a085c9df9781b2df65", "iot.topic_rules": "5744a0bbdcdc1e0efcf5533573bc704f03b0264d2da025c73548db7539bba92e", "lambda.runtimes": "ababff07e4d12918f7fe1d7d45448d9694635b70d0adbc5fa84d36bccd443e01", "route53.domains": "aab54b80fabb4c21b15b5b0bdb8955435cd55b9f573562656fa4c6d367b463a5", "waf.rule_groups": "bdca5ac013a00ac79b062e5543cc664bfa895bd2a96ee782dfd55a108383494e", "acm.certificates": "4d1dcc7b37b3e7c2a372d41cdfad198449273399457f5cc5512109b0ea2d9e34", "ec2.nat_gateways": "606db5b3cccc559bf2b319d2d7bd4e47157931ff42cb9c8cb8091c5b57181d8e", "ec2.network_acls": "0dbd56dcefbfe50ff2d483963467983a00c8de47a6dbf704aebd5b8da6517a18", "ec2.route_tables": "0d7b34dccc2881ab8f6c959f82691e733340faa0d97e239500427f399d680685", "ec2.vpn_gateways": "46a79bc0ce46d013903ac10e680a400e4785978f7daee8b08b3caf0b3cb4279f", "ecr.repositories": "7e8c343963dd7964a6b65de48a165f3e3879db119a86ed449755a384a6ef29df", "iot.certificates": "8c962927a25abd67dbd8b95c5d97fb990ee0ad1d354a95329e7f7934ccfd2d8d", "iot.thing_groups": "d62c4597f5f7ea9124d7c3fb3483a632e1b92f1a0b63f05c7363f5f1a8665150", "lambda.functions": "ac598474ae87d320c15d1d669b43403ececd0ae198fc007dc07f26114068bacc", "rds.certificates": "335010a7984a5e626afc21a55d85bfd254f67ed2ef729173e8d682011d4aaf64", "rds.db_snapshots": "5e415eb425f8c75fe4fba54e449a7b4893eb2e45b90953b686bec9ce7c508094", "sagemaker.models": "3b2663fd88d89656a09c5b1476d9aab98dc0df13e8218858ef127b480070ebd1", "apigatewayv2.apis": "c17db20dcc117acfec9df7b647dd37f6ec704a19239543cc4f0e7feeeb709ffc", "cloudtrail.trails": "272c9033006b8102745ab65372953aa38c1fecc27f6b8f4c9c787900c9dee38d", "cloudwatch.alarms": "24a0844096266a78db60e04c5168cb576e04e1cb31cad5846d31125693f50364", "ec2.ebs_snapshots": "d2d8e4c0892b666eab6e774374de4902d90a241f52c7fb2c78a8bc1a7bcda851", "ec2.vpc_endpoints": "dcd0284ed77855562796bdcf25909bcd7929a5e837304faf1d8d2a6ff7fc27f7", "redshift.clusters": "36351959a51bf368ae5cd5f772d2e3d4e7fcdfe1f3b65fc98ce29b017f14577d", "sns.subscriptions": "55648603775f351a689fb8009e7ebd336d5ee7e38216b4a8b08e1e474a667fd6", "wafregional.rules": "963c793fe68bdf83281d2c24bff824debd18608bdb9acd701bc34f14bc09f789", "wafv2.rule_groups": "074210315b06a6bc5e2d19b35b940c2689f562e58d3cf8f582574e8a11fd5227", "athena.work_groups": "eb6e898150f6092078314321259b8ba5cf1b7585d0ab6972194585318668bbba", "autoscaling.groups": "68192fa8bd75ade5ff517688c4563f06df9bb84684aa138e9fe76d7b32837747", "codebuild.projects": "043ad3c3fd74e08814b8616ddc3bed2580aa635146c4d6a4ec00f425d65cc60d", "cognito.user_pools": "5916413d6b8e1ffcc745248c989941f4050704f1756052df552610864f27e203", "directconnect.lags": "a7659d813161b80f73aee647258d4762e54abf3d99a9b5cbb52e38308307563a", "iot.billing_groups": "778f44d5dfac33e277acd2a1b8486b4f3f2729bba882623abbd1f499ae6c91e8", "shield.protections": "9e1094ef0d89bf6e8f4212e88a621baee7a19840ce7af68deeffd4bb4f74aa1b", "apigateway.api_keys": "e7690fc5f2846dbb863d2dbafeb540618330cbe73213d51e87e8666539a4d114", "ec2.regional_config": "bed9ae13243c845b65a9bed06bc7b000e62f567bdfe09d01ca74587ff3a38aed", "ec2.security_groups": "c467ac7b5fcc06a09c60eac1de032fa7cc1d6ccaffbc17083e6c2e5a97af3663", "elbv2.target_groups": "529cb273e1c59a88cb1b63c99f8fef7f7d00ee7edbe2ab2cf0a2b2fcde75237b", "guardduty.detectors": "7b45abb5473abe465d397a68772516e4966330093cc9535cb78162d788ca7b6e", "iot.ca_certificates": "b0654cc9691541b90068fca528e30ee2822ef61558e472103f5f43b9adf8bf9a", "xray.sampling_rules": "2a90975716fc12a6fa4f43b3be40771406586d9e5549822f1dd9a1611dd154fa", "apigateway.rest_apis": "3399b2bd34df2398bf8b6ded19ffee3ecf063df32a3bcaf080ec79c2984b89dc", "apigateway.vpc_links": "f965434ce3e625169b03701ef1db60305fdafdf6ded49e78281e9d5cdaed0456", "athena.data_catalogs": "82470d76a1036791bc593dfc3bbfdac8b10282dafbf7a74c9273d26f4007c61a", "ec2.transit_gateways": "78d67210286bd0fd7075b1463c081f0e023af29acc24cb5b343fec7d59f84ede", "ecs.task_definitions": "9ce77502d987b9c4651fe7743f0b7c57992d1299d1429c05491082dea9a04cf9", "elbv1.load_balancers": "13c8a18f4da2fd8786adad29af0354bb866c1b3dcb2d37d9f1b997cd79939116", "elbv2.load_balancers": "cfca99e6c163172b99c2672b294b2c439accdb033a9ea980d8281832ba754003", "rds.db_subnet_groups": "6cd100d21e3b8512d791c5e798c9f5fbe23836f05f7ede3e5e4922b0e40d06aa", "route53.hosted_zones": "08032a460758dded555deb0fd02ae850a561151bc6bcb33c05258ad384a1e709", "shield.subscriptions": "d1ef4cff342b20bb33040197d8e1f3a2017d3f4d3e8892b9fda9e3a2a4f91d9f", "wafregional.web_acls": "4283326fcede87e80fd8e130b6c18f41afdc43209858cf48a42bcb16b4377c2d", "cloudformation.stacks": "bc752e596db129229a1b8eebf6e83b30ce82dc408f3466a26cb944362495c4a1", "codepipeline.webhooks": "a15185254495d695549858d1b3ea82371fd8432bf8aaafafb63b5fae3ee8c845", "ec2.customer_gateways": "05a3a4e27830157538cdf1853a6c21c875e7028d4a1b952dc72bb809e0c2e773", "ec2.instance_statuses": "ac8e0eda2636d8e37326a8e6ae8ce694d61d6a2e505be25320d23ef4a40cc4fd", "ec2.internet_gateways": "b773e02a2632d96be2812c9c0fd506f91c2a0601c4c2b36cb300f47050fe1a98", "elasticsearch.domains": "136c6f9534e17e5c97df615a296fbddf662f1be5e462167f44f548775b6c6690", "iam.password_policies": "f75c7525ed69dc31e59190d37b2b8fb903b24c265a0ee565a98edccc5a2e7d8f", "rds.cluster_snapshots": "678376c4a15abfbf7b5c012f776e4a71467aa9385085715c93dcf54187f8d038", "route53.health_checks": "cdf721ac9fea7b7ed38d895fe473b81e55f66dbd21b51d5d560c33f9e52a7e98", "workspaces.workspaces": "3b09cbccdb0dbc61cd47843a0b4411ad320d566e4d924d9da9c6bf5d72aa6307", "apigateway.usage_plans": "5d83d745b62420117f67f3a782d7065ffcdd2639f38449c07778518dfaf06737", "apigatewayv2.vpc_links": "e48ea45edf8c6775105905a9325c491588c22058b78aa12834a8fd67a2df88da", "backup.global_settings": "5c63ac538fb9dfe3f67e2b83c31fd3a4cb54f66973f743ef3e2a1687872eaae6", "backup.region_settings": "693d08dbd3f35aaefe770779a6af3fdc8e74ac0744ba536d68a5ae37c4b66869", "cloudwatchlogs.filters": "1ff65dceaeb1704e5b40dae65e630559eb7fc195e99523798ee2888cb7d3bc49", "codepipeline.pipelines": "50254d6b8a1161ae1f82c15568c9a9b27e66d10247e6c2030ae401ecc59989e1", "cognito.identity_pools": "de75b3034a03ab1bd9d242142a43f3b88ec37c4cf689aab8f160d4ec91de6e40", "directconnect.gateways": "1e1bf7d38408534b1d9a8fe6e774f862045138f7c9fcf54bc5f6403fddd201d4", "ec2.network_interfaces": "8e91cba5107499ca18da8daad1b5511d4cc69f736a9d432a8d49d816b220c840", "organizations.accounts": "c82fe646c10ce7046bbe3cd792c9f883d9f2f068be816de615545ba66a5f0c47", "rds.db_security_groups": "a8e726bd9429aba0d115d2904f624c0447546eb699c94641fba7e7de68db0095", "redshift.subnet_groups": "58dda9f034a1191938d90a0435ffa1d97810a4ba89c65ecbb738b686ad430ed5", "secretsmanager.secrets": "2cd910bd24ed84590bb6bc2bb834dcb84ef1ab2cf564a975165259b6c01ad9db", "workspaces.directories": "a0328d95cfad712c6337cdd2c1043b84b14e4394d10709998a2b98ff7e28c976", "xray.encryption_config": "d1e2b5e5e687d6bc332a44a74d77afe166b1adc969b4ca9d92c7105640c11b4d", "apigateway.domain_names": "1d7ec1d2ad923a168c24fa2dcfd3fe8b527208fa9677da3173d38a7bb227baa2", "iam.server_certificates": "53e3b0ae026e5d7e6dd6e0ae941093911ea84f647e55fdf89ad04a049e9b51a8", "iam.virtual_mfa_devices": "ad5d0091213d3070b92d2175ca37350bc2ed2bc4ad963e2e76f97592dd3fafe5", "rds.db_parameter_groups": "feb6e4f7f4c68d8f93825034f9be25171187dc9262fbedf9b4a6231a4c0dcc3a", "rds.event_subscriptions": "f747daa0ee2a6e1b2914b0a672f255fdf32d9eb8c1a914ae86a0b93791385502", "sagemaker.training_jobs": "c758382f957d15afef276864951a06b50956a017504f60fbb54810d77b7ca48d", "wafregional.rule_groups": "000a4058bc01753102a8316b6e8188e7a94ab3662692df90bb76b5115b1be4f4", "accessanalyzer.analyzers": "2a86eace92e561a3bcff3005e8a45cbd71b3ba29eec421394c8dbddd64865f1c", "cloudfront.distributions": "27ececdc1b2f71b01120e8106d41790239424c2fbcd432695c361796059a41d0", "config.conformance_packs": "e6b01aed9e18b3dcbff950c54efdba3d0184017a9888a7ff30f6c2fdc7f50cd3", "route53.traffic_policies": "8da00d8b67821405b516959457a54d17d814b380fb6b3eb04576cb32169feb53", "wafv2.regex_pattern_sets": "3a371606cd2e42ea9260c123b605a90e353c8f0eb88da4496f92db7885f72e96", "apigatewayv2.domain_names": "1361a6bf047f24b6ecbbb5b2f5d71807c29c40649777b90e813b520154d7c50d", "cloudfront.cache_policies": "ce82e4159eb589bf9ebe68dc158994191d6d7d5607972e52ee19774d1986e98d", "directconnect.connections": "6cd73d370f0f29b7a0cf966aab92dc682c8642b8a5c42bdc1b824869f048a742", "dms.replication_instances": "e93b98a914948e2d5a817f529cbc0b0e3603778e15b0ca8d3e085a097d2096dc", "shield.protections_groups": "8a336bac08e7c904de50a0b716acedd2cd81e245c7ff8c920981cf431a675460", "wafv2.managed_rule_groups": "ce32619a2612251748ed27a1c4627c25c818a618333b6faaea928b8d7fbc2381", "waf.subscribed_rule_groups": "157c2de1ebe45abaa487e9760490408031e14e122e2ff71cb4e0c8c180b7e46b", "ec2.vpc_peering_connections": "6f4f67c8f8c276f521948d8742f22d06084ffb77092bd3ec9240228e02cc2c75", "iam.saml_identity_providers": "ca06ab367d25f21902369ecba341ebb90658d0e5896fc3d9d121e3b127bc1df2", "rds.cluster_parameter_groups": "f07aac16d84f200a359de7203922fded8950f090374cb764caa3b10298c0eb9a", "redshift.event_subscriptions": "8b1fc6ad506149b6aa5c14bca299cef433fb3e182163dba8d4d1d408c429ae66", "sagemaker.notebook_instances": "c301e4514d59d9035cf1b9c134ddab520cd6262c09f3be7a8318990e13e75423", "wafregional.rate_based_rules": "ae6aadf7606167c1ba60ad43a3d3049e6ed3401cb6f0616e4aec744562b7a513", "autoscaling.scheduled_actions": "22ac46565086dc674fd612c074ca83c595b7ea4b75b607696ddc88aba2d20da5", "elasticbeanstalk.applications": "e4fc6e4b8ce1bc616184b2371427e0d03ae1cb45dde145f2d012d7b9169538fb", "elasticbeanstalk.environments": "be1a35e4117d8f094a2bd28a50c4d65e4c0c863c82fc2bba1ce7ee1f2a707fd7", "apigateway.client_certificates": "e85d6875ab1e8aa0349cc4edaea9cd6f89f701b9d3c6856554be466a89657536", "config.configuration_recorders": "8de94ecbd5c3849dc54034978708973d72049027b314db699f2b0e4ad0347ebc", "directconnect.virtual_gateways": "109d4ddd81c8e4e11d24a8115ca46dc7ea005ebb7cecb90b9fbd772383060e33", "applicationautoscaling.policies": "03c1c839baffae0447a8967400023bf04ee8a0165b4c9474cf9bee3afd7ee055", "emr.block_public_access_configs": "edd6cdf10c2e5e2f9c70f2533d4eb7580307a13e9d5ce4483b6f545cfb2b55eb", "directconnect.virtual_interfaces": "0b7030d17946063c8df37464891c607ebdc660c713f350a440e029f7947a407a", "route53.reusable_delegation_sets": "44aa8910873fd1b7d1fb8ba819983f1e172974c0d04973996feec3e90cfb039b", "autoscaling.launch_configurations": "24f0ff45efb96d49d6a24b1f97d1a06223ad8325798e1423b3ca081403de8120", "ec2.egress_only_internet_gateways": "19c88f1ee064d386292453401aeefa9ee5efcebd9fa20570d0b11f11866ca142", "sagemaker.endpoint_configurations": "c3b107f8aac9f04c9894b5088a97b5bae727296d3ec672c085fa6db19d826530", "elasticbeanstalk.application_versions": "14a89ea5a178e8dd45288b701d57f31abf4636ec27badac738af63733c70b837", "iam.openid_connect_identity_providers": "5081c5c04a0d71e8d95325c45ddb3d18f381ae41b32e5356a0834a531a411023"}
\.


--
-- Data for Name: aws_access_analyzer_analyzer_archive_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_access_analyzer_analyzer_archive_rules (cq_id, cq_meta, analyzer_cq_id, created_at, filter, rule_name, updated_at) FROM stdin;
\.


--
-- Data for Name: aws_access_analyzer_analyzer_finding_sources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_access_analyzer_analyzer_finding_sources (cq_id, cq_meta, analyzer_finding_cq_id, type, detail_access_point_arn) FROM stdin;
\.


--
-- Data for Name: aws_access_analyzer_analyzer_findings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_access_analyzer_analyzer_findings (cq_id, cq_meta, analyzer_cq_id, analyzed_at, condition, created_at, id, resource_owner_account, resource_type, status, updated_at, action, error, is_public, principal, resource) FROM stdin;
\.


--
-- Data for Name: aws_access_analyzer_analyzers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_access_analyzer_analyzers (cq_id, cq_meta, account_id, region, arn, created_at, name, status, type, last_resource_analyzed, last_resource_analyzed_at, status_reason_code, tags) FROM stdin;
\.


--
-- Data for Name: aws_accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_accounts (cq_id, cq_meta, account_id, users, users_quota, groups, groups_quota, server_certificates, server_certificates_quota, user_policy_size_quota, group_policy_size_quota, groups_per_user_quota, signing_certificates_per_user_quota, access_keys_per_user_quota, mfa_devices, mfa_devices_in_use, account_mfa_enabled, account_access_keys_present, account_signing_certificates_present, attached_policies_per_group_quota, policies, policies_quota, policy_size_quota, policy_versions_in_use, policy_versions_in_use_quota, versions_per_policy_quota, global_endpoint_token_version, aliases) FROM stdin;
\.


--
-- Data for Name: aws_acm_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_acm_certificates (cq_id, cq_meta, account_id, region, arn, certificate_authority_arn, created_at, domain_name, domain_validation_options, extended_key_usages, failure_reason, imported_at, in_use_by, issued_at, issuer, key_algorithm, key_usages, not_after, not_before, certificate_transparency_logging_preference, renewal_eligibility, renewal_summary_domain_validation_options, renewal_summary_status, renewal_summary_updated_at, renewal_summary_failure_reason, revocation_reason, revoked_at, serial, signature_algorithm, status, subject, subject_alternative_names, type, tags) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_api_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_api_keys (cq_id, cq_meta, account_id, region, arn, created_date, customer_id, description, enabled, id, last_updated_date, name, stage_keys, tags, value) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_client_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_client_certificates (cq_id, cq_meta, account_id, region, arn, id, created_date, description, expiration_date, pem_encoded_certificate, tags) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_domain_name_base_path_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_domain_name_base_path_mappings (cq_id, cq_meta, domain_name_cq_id, arn, domain_name, base_path, rest_api_id, stage) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_domain_names; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_domain_names (cq_id, cq_meta, account_id, region, arn, certificate_arn, certificate_name, certificate_upload_date, distribution_domain_name, distribution_hosted_zone_id, domain_name, domain_name_status, domain_name_status_message, endpoint_configuration_types, endpoint_configuration_vpc_endpoint_ids, mutual_tls_authentication_truststore_uri, mutual_tls_authentication_truststore_version, mutual_tls_authentication_truststore_warnings, regional_certificate_arn, regional_certificate_name, regional_domain_name, regional_hosted_zone_id, security_policy, tags) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_authorizers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_authorizers (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, auth_type, authorizer_credentials, authorizer_result_ttl_in_seconds, authorizer_uri, id, identity_source, identity_validation_expression, name, provider_arns, type) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_deployments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_deployments (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, api_summary, created_date, description, id) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_documentation_parts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_documentation_parts (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, id, location_type, location_method, location_name, location_path, location_status_code, properties) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_documentation_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_documentation_versions (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, created_date, description, version) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_gateway_responses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_gateway_responses (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, default_response, response_parameters, response_templates, response_type, status_code) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_models; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_models (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, model_template, content_type, description, id, name, schema) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_request_validators; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_request_validators (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, id, name, validate_request_body, validate_request_parameters) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_resources (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, id, parent_id, path, path_part, resource_methods) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_api_stages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_api_stages (cq_id, cq_meta, rest_api_cq_id, rest_api_id, arn, access_log_settings_destination_arn, access_log_settings_format, cache_cluster_enabled, cache_cluster_size, cache_cluster_status, canary_settings_deployment_id, canary_settings_percent_traffic, canary_settings_stage_variable_overrides, canary_settings_use_stage_cache, client_certificate_id, created_date, deployment_id, description, documentation_version, last_updated_date, method_settings, stage_name, tags, tracing_enabled, variables, web_acl_arn) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_rest_apis; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_rest_apis (cq_id, cq_meta, account_id, region, arn, api_key_source, binary_media_types, created_date, description, disable_execute_api_endpoint, endpoint_configuration_types, endpoint_configuration_vpc_endpoint_ids, id, minimum_compression_size, name, policy, tags, version, warnings) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_usage_plan_api_stages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_usage_plan_api_stages (cq_id, cq_meta, usage_plan_cq_id, usage_plan_id, api_id, stage, throttle) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_usage_plan_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_usage_plan_keys (cq_id, cq_meta, usage_plan_cq_id, usage_plan_id, arn, id, name, type, value) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_usage_plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_usage_plans (cq_id, cq_meta, account_id, region, arn, description, id, name, product_code, quota_limit, quota_offset, quota_period, tags, throttle_burst_limit, throttle_rate_limit) FROM stdin;
\.


--
-- Data for Name: aws_apigateway_vpc_links; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigateway_vpc_links (cq_id, cq_meta, account_id, region, arn, description, id, name, status, status_message, tags, target_arns) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_authorizers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_authorizers (cq_id, cq_meta, api_cq_id, api_id, arn, name, authorizer_credentials_arn, authorizer_id, authorizer_payload_format_version, authorizer_result_ttl_in_seconds, authorizer_type, authorizer_uri, enable_simple_responses, identity_source, identity_validation_expression, jwt_configuration_audience, jwt_configuration_issuer) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_deployments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_deployments (cq_id, cq_meta, api_cq_id, api_id, arn, auto_deployed, created_date, deployment_id, deployment_status, deployment_status_message, description) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_integration_responses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_integration_responses (cq_id, cq_meta, api_integration_cq_id, integration_id, arn, integration_response_key, content_handling_strategy, integration_response_id, response_parameters, response_templates, template_selection_expression) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_integrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_integrations (cq_id, cq_meta, api_cq_id, api_id, arn, api_gateway_managed, connection_id, connection_type, content_handling_strategy, credentials_arn, description, integration_id, integration_method, integration_response_selection_expression, integration_subtype, integration_type, integration_uri, passthrough_behavior, payload_format_version, request_parameters, request_templates, response_parameters, template_selection_expression, timeout_in_millis, tls_config_server_name_to_verify) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_models; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_models (cq_id, cq_meta, api_cq_id, api_id, arn, model_template, name, content_type, description, model_id, schema) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_route_responses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_route_responses (cq_id, cq_meta, api_route_cq_id, route_id, arn, route_response_key, model_selection_expression, response_models, response_parameters, route_response_id) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_routes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_routes (cq_id, cq_meta, api_cq_id, api_id, arn, route_key, api_gateway_managed, api_key_required, authorization_scopes, authorization_type, authorizer_id, model_selection_expression, operation_name, request_models, request_parameters, route_id, route_response_selection_expression, target) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_api_stages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_api_stages (cq_id, cq_meta, api_cq_id, api_id, arn, stage_name, access_log_settings_destination_arn, access_log_settings_format, api_gateway_managed, auto_deploy, client_certificate_id, created_date, route_settings_data_trace_enabled, route_settings_detailed_metrics_enabled, route_settings_logging_level, route_settings_throttling_burst_limit, route_settings_throttling_rate_limit, deployment_id, description, last_deployment_status_message, last_updated_date, route_settings, stage_variables, tags) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_apis; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_apis (cq_id, cq_meta, account_id, region, arn, name, protocol_type, route_selection_expression, api_endpoint, api_gateway_managed, id, api_key_selection_expression, cors_configuration_allow_credentials, cors_configuration_allow_headers, cors_configuration_allow_methods, cors_configuration_allow_origins, cors_configuration_expose_headers, cors_configuration_max_age, created_date, description, disable_execute_api_endpoint, disable_schema_validation, import_info, tags, version, warnings) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_domain_name_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_domain_name_configurations (cq_id, cq_meta, domain_name_cq_id, api_gateway_domain_name, certificate_arn, certificate_name, certificate_upload_date, domain_name_status, domain_name_status_message, endpoint_type, hosted_zone_id, security_policy) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_domain_name_rest_api_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_domain_name_rest_api_mappings (cq_id, cq_meta, domain_name_cq_id, api_id, arn, stage, api_mapping_id, api_mapping_key) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_domain_names; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_domain_names (cq_id, cq_meta, account_id, region, arn, domain_name, api_mapping_selection_expression, mutual_tls_authentication_truststore_uri, mutual_tls_authentication_truststore_version, mutual_tls_authentication_truststore_warnings, tags) FROM stdin;
\.


--
-- Data for Name: aws_apigatewayv2_vpc_links; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_apigatewayv2_vpc_links (cq_id, cq_meta, account_id, region, arn, name, security_group_ids, subnet_ids, id, created_date, tags, vpc_link_status, vpc_link_status_message, vpc_link_version) FROM stdin;
\.


--
-- Data for Name: aws_applicationautoscaling_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_applicationautoscaling_policies (cq_id, cq_meta, account_id, region, namespace, creation_time, arn, name, type, resource_id, scalable_dimension, service_namespace, alarms, step_scaling_policy_configuration, target_tracking_scaling_policy_configuration) FROM stdin;
\.


--
-- Data for Name: aws_athena_data_catalog_database_table_columns; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_data_catalog_database_table_columns (cq_id, cq_meta, data_catalog_database_table_cq_id, name, comment, type) FROM stdin;
\.


--
-- Data for Name: aws_athena_data_catalog_database_table_partition_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_data_catalog_database_table_partition_keys (cq_id, cq_meta, data_catalog_database_table_cq_id, name, comment, type) FROM stdin;
\.


--
-- Data for Name: aws_athena_data_catalog_database_tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_data_catalog_database_tables (cq_id, cq_meta, data_catalog_database_cq_id, name, create_time, last_access_time, parameters, table_type) FROM stdin;
\.


--
-- Data for Name: aws_athena_data_catalog_databases; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_data_catalog_databases (cq_id, cq_meta, data_catalog_cq_id, name, description, parameters) FROM stdin;
\.


--
-- Data for Name: aws_athena_data_catalogs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_data_catalogs (cq_id, cq_meta, account_id, region, arn, tags, name, type, description, parameters) FROM stdin;
\.


--
-- Data for Name: aws_athena_work_group_named_queries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_work_group_named_queries (cq_id, cq_meta, work_group_cq_id, database, name, query_string, description, named_query_id, work_group) FROM stdin;
\.


--
-- Data for Name: aws_athena_work_group_prepared_statements; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_work_group_prepared_statements (cq_id, cq_meta, work_group_cq_id, description, last_modified_time, query_statement, statement_name, work_group_name) FROM stdin;
\.


--
-- Data for Name: aws_athena_work_group_query_executions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_work_group_query_executions (cq_id, cq_meta, work_group_cq_id, effective_engine_version, selected_engine_version, query, catalog, database, id, acl_configuration_s3_acl_option, encryption_configuration_encryption_option, encryption_configuration_kms_key, expected_bucket_owner, output_location, statement_type, data_manifest_location, data_scanned_in_bytes, engine_execution_time_in_millis, query_planning_time_in_millis, query_queue_time_in_millis, service_processing_time_in_millis, total_execution_time_in_millis, athena_error_error_category, athena_error_error_message, athena_error_error_type, athena_error_retryable, completion_date_time, state, state_change_reason, submission_date_time, work_group) FROM stdin;
\.


--
-- Data for Name: aws_athena_work_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_athena_work_groups (cq_id, cq_meta, account_id, arn, region, tags, name, bytes_scanned_cutoff_per_query, enforce_work_group_configuration, effective_engine_version, selected_engine_version, publish_cloud_watch_metrics_enabled, requester_pays_enabled, acl_configuration_s3_acl_option, encryption_configuration_encryption_option, encryption_configuration_kms_key, expected_bucket_owner, output_location, creation_time, description, state) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_group_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_group_instances (cq_id, cq_meta, group_cq_id, availability_zone, health_status, id, lifecycle_state, protected_from_scale_in, type, launch_configuration_name, launch_template_id, launch_template_name, launch_template_version, weighted_capacity) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_group_lifecycle_hooks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_group_lifecycle_hooks (cq_id, cq_meta, group_cq_id, auto_scaling_group_name, default_result, global_timeout, heartbeat_timeout, lifecycle_hook_name, lifecycle_transition, notification_metadata, notification_target_arn, role_arn) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_group_scaling_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_group_scaling_policies (cq_id, cq_meta, group_cq_id, adjustment_type, alarms, auto_scaling_group_name, cooldown, enabled, estimated_instance_warmup, metric_aggregation_type, min_adjustment_magnitude, min_adjustment_step, arn, name, type, scaling_adjustment, step_adjustments, target_tracking_configuration_target_value, target_tracking_configuration_customized_metric_name, target_tracking_configuration_customized_metric_namespace, target_tracking_configuration_customized_metric_statistic, target_tracking_configuration_customized_metric_dimensions, target_tracking_configuration_customized_metric_unit, target_tracking_configuration_disable_scale_in, target_tracking_configuration_predefined_metric_type, target_tracking_configuration_predefined_metric_resource_label) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_group_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_group_tags (cq_id, cq_meta, group_cq_id, key, propagate_at_launch, resource_id, resource_type, value) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_groups (cq_id, cq_meta, account_id, region, load_balancers, load_balancer_target_groups, notifications_configurations, name, availability_zones, created_time, default_cooldown, desired_capacity, health_check_type, max_size, min_size, arn, capacity_rebalance, enabled_metrics, health_check_grace_period, launch_configuration_name, launch_template_id, launch_template_name, launch_template_version, load_balancer_names, max_instance_lifetime, mixed_instances_policy, new_instances_protected_from_scale_in, placement_group, service_linked_role_arn, status, suspended_processes, target_group_arns, termination_policies, vpc_zone_identifier) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_launch_configuration_block_device_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_launch_configuration_block_device_mappings (cq_id, cq_meta, launch_configuration_cq_id, device_name, ebs_delete_on_termination, ebs_encrypted, ebs_iops, ebs_snapshot_id, ebs_volume_size, ebs_volume_type, no_device, virtual_name) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_launch_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_launch_configurations (cq_id, cq_meta, account_id, region, created_time, image_id, instance_type, launch_configuration_name, associate_public_ip_address, classic_link_vpc_id, classic_link_vpc_security_groups, ebs_optimized, iam_instance_profile, instance_monitoring_enabled, kernel_id, key_name, arn, metadata_options_http_endpoint, metadata_options_http_put_response_hop_limit, metadata_options_http_tokens, placement_tenancy, ramdisk_id, security_groups, spot_price, user_data) FROM stdin;
\.


--
-- Data for Name: aws_autoscaling_scheduled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_autoscaling_scheduled_actions (cq_id, cq_meta, account_id, region, auto_scaling_group_name, desired_capacity, end_time, max_size, min_size, recurrence, arn, name, start_time, "time", time_zone) FROM stdin;
\.


--
-- Data for Name: aws_backup_global_settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_global_settings (cq_id, cq_meta, account_id, global_settings, last_update_time) FROM stdin;
\.


--
-- Data for Name: aws_backup_plan_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_plan_rules (cq_id, cq_meta, plan_cq_id, name, target_backup_vault_name, completion_window_minutes, copy_actions, enable_continuous_backup, delete_after_days, move_to_cold_storage_after_days, recovery_point_tags, id, schedule_expression, start_window_minutes) FROM stdin;
\.


--
-- Data for Name: aws_backup_plan_selections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_plan_selections (cq_id, cq_meta, plan_cq_id, creation_date, creator_request_id, iam_role_arn, selection_id, selection_name, conditions, list_of_tags, not_resources, resources) FROM stdin;
\.


--
-- Data for Name: aws_backup_plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_plans (cq_id, cq_meta, account_id, region, arn, id, name, creation_date, creator_request_id, last_execution_date, version_id, advanced_backup_settings, tags) FROM stdin;
\.


--
-- Data for Name: aws_backup_region_settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_region_settings (cq_id, cq_meta, account_id, region, resource_type_management_preference, resource_type_opt_in_preference) FROM stdin;
\.


--
-- Data for Name: aws_backup_vault_recovery_points; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_vault_recovery_points (cq_id, cq_meta, vault_cq_id, backup_size, calculated_delete_at, calculated_move_to_cold_storage_at, completion_date, created_by, creation_date, encryption_key_arn, iam_role_arn, is_encrypted, last_restore_time, delete_after, move_to_cold_storage_after, arn, resource_arn, resource_type, source_backup_vault_arn, status, status_message, tags) FROM stdin;
\.


--
-- Data for Name: aws_backup_vaults; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_backup_vaults (cq_id, cq_meta, account_id, region, arn, name, creation_date, creator_request_id, encryption_key_arn, lock_date, locked, max_retention_days, min_retention_days, number_of_recovery_points, access_policy, notification_events, notification_sns_topic_arn, tags) FROM stdin;
\.


--
-- Data for Name: aws_cloudformation_stack_outputs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudformation_stack_outputs (cq_id, cq_meta, stack_cq_id, description, export_name, output_key, output_value) FROM stdin;
\.


--
-- Data for Name: aws_cloudformation_stack_resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudformation_stack_resources (cq_id, cq_meta, stack_cq_id, last_updated_timestamp, logical_resource_id, resource_status, resource_type, stack_resource_drift_status, drift_last_check_timestamp, module_info_logical_id_hierarchy, module_info_type_hierarchy, physical_resource_id, resource_status_reason) FROM stdin;
\.


--
-- Data for Name: aws_cloudformation_stacks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudformation_stacks (cq_id, cq_meta, account_id, region, arn, creation_time, stack, status, capabilities, change_set_id, deletion_time, description, disable_rollback, stack_drift_status, drift_last_check_timestamp, enable_termination_protection, last_updated_time, notification_arns, parameters, parent_id, role_arn, rollback_configuration_monitoring_time_in_minutes, rollback_configuration_rollback_triggers, root_id, id, stack_status_reason, tags, timeout_in_minutes) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_cache_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_cache_policies (cq_id, cq_meta, account_id, arn, min_ttl, name, comment, default_ttl, max_ttl, cookies_behavior, cookies_quantity, cookies, enable_accept_encoding_gzip, headers_behavior, headers_quantity, headers, query_strings_behavior, query_strings_quantity, query_strings, enable_accept_encoding_brotli, id, last_modified_time, type) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_cache_behavior_lambda_functions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_cache_behavior_lambda_functions (cq_id, cq_meta, distribution_cache_behavior_cq_id, event_type, lambda_function_arn, include_body) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_cache_behaviors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_cache_behaviors (cq_id, cq_meta, distribution_cq_id, path_pattern, target_origin_id, viewer_protocol_policy, allowed_methods, cached_methods, cache_policy_id, compress, default_ttl, field_level_encryption_id, forwarded_values_cookies_forward, forwarded_values_cookies_whitelisted_names, forwarded_values_query_string, forwarded_values_headers, forwarded_values_query_string_cache_keys, max_ttl, min_ttl, origin_request_policy_id, realtime_log_config_arn, smooth_streaming, trusted_key_groups_enabled, trusted_key_groups, trusted_signers_enabled, trusted_signers) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_custom_error_responses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_custom_error_responses (cq_id, cq_meta, distribution_cq_id, error_code, error_caching_min_ttl, response_code, response_page_path) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_default_cache_behavior_functions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_default_cache_behavior_functions (cq_id, cq_meta, distribution_cq_id, event_type, lambda_function_arn, include_body) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_origin_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_origin_groups (cq_id, cq_meta, distribution_cq_id, failover_criteria_status_codes, id, members_origin_ids) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distribution_origins; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distribution_origins (cq_id, cq_meta, distribution_cq_id, domain_name, id, connection_attempts, connection_timeout, custom_headers, custom_origin_config_http_port, custom_origin_config_https_port, custom_origin_config_protocol_policy, custom_origin_config_keepalive_timeout, custom_origin_config_read_timeout, custom_origin_config_ssl_protocols, origin_path, origin_shield_enabled, origin_shield_region, s3_origin_config_origin_access_identity) FROM stdin;
\.


--
-- Data for Name: aws_cloudfront_distributions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudfront_distributions (cq_id, cq_meta, account_id, tags, arn, caller_reference, comment, cache_behavior_target_origin_id, cache_behavior_viewer_protocol_policy, cache_behavior_allowed_methods, cache_behavior_allowed_methods_cached_methods, cache_behavior_cache_policy_id, cache_behavior_compress, cache_behavior_default_ttl, cache_behavior_field_level_encryption_id, cache_behavior_forwarded_values_cookies_forward, cache_behavior_forwarded_values_cookies_whitelisted_names, cache_behavior_forwarded_values_query_string, cache_behavior_forwarded_values_headers, cache_behavior_forwarded_values_query_string_cache_keys, cache_behavior_max_ttl, cache_behavior_min_ttl, cache_behavior_origin_request_policy_id, cache_behavior_realtime_log_config_arn, cache_behavior_smooth_streaming, cache_behavior_trusted_key_groups_enabled, cache_behavior_trusted_key_groups, cache_behavior_trusted_signers_enabled, cache_behavior_trusted_signers, enabled, aliases, default_root_object, http_version, ipv6_enabled, logging_bucket, logging_enabled, logging_include_cookies, logging_prefix, price_class, geo_restriction_type, geo_restrictions, viewer_certificate_acm_certificate_arn, viewer_certificate, viewer_certificate_source, viewer_certificate_cloudfront_default_certificate, viewer_certificate_iam_certificate_id, viewer_certificate_minimum_protocol_version, viewer_certificate_ssl_support_method, web_acl_id, domain_name, id, in_progress_invalidation_batches, last_modified_time, status, active_trusted_key_groups_enabled, active_trusted_key_groups, active_trusted_signers_enabled, active_trusted_signers, alias_icp_recordals) FROM stdin;
\.


--
-- Data for Name: aws_cloudtrail_trail_event_selectors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudtrail_trail_event_selectors (cq_id, cq_meta, trail_cq_id, trail_arn, exclude_management_event_sources, include_management_events, read_write_type) FROM stdin;
\.


--
-- Data for Name: aws_cloudtrail_trails; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudtrail_trails (cq_id, cq_meta, account_id, tags, cloudwatch_logs_log_group_name, is_logging, latest_cloud_watch_logs_delivery_error, latest_cloud_watch_logs_delivery_time, latest_delivery_error, latest_delivery_time, latest_digest_delivery_error, latest_digest_delivery_time, latest_notification_error, latest_notification_time, start_logging_time, stop_logging_time, cloud_watch_logs_log_group_arn, cloud_watch_logs_role_arn, has_custom_event_selectors, has_insight_selectors, region, include_global_service_events, is_multi_region_trail, is_organization_trail, kms_key_id, log_file_validation_enabled, name, s3_bucket_name, s3_key_prefix, sns_topic_arn, sns_topic_name, arn) FROM stdin;
\.


--
-- Data for Name: aws_cloudwatch_alarm_metrics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudwatch_alarm_metrics (cq_id, cq_meta, alarm_cq_id, alarm_arn, alarm_name, id, expression, label, metric_stat_metric_dimensions, metric_stat_metric_name, metric_stat_metric_namespace, metric_stat_period, metric_stat, metric_stat_unit, period, return_data) FROM stdin;
\.


--
-- Data for Name: aws_cloudwatch_alarms; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudwatch_alarms (cq_id, cq_meta, account_id, region, actions_enabled, actions, arn, configuration_updated_timestamp, description, name, comparison_operator, datapoints_to_alarm, dimensions, evaluate_low_sample_count_percentile, evaluation_periods, extended_statistic, insufficient_data_actions, metric_name, namespace, ok_actions, period, state_reason, state_reason_data, state_updated_timestamp, state_value, statistic, threshold, threshold_metric_id, treat_missing_data, unit) FROM stdin;
\.


--
-- Data for Name: aws_cloudwatchlogs_filter_metric_transformations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudwatchlogs_filter_metric_transformations (cq_id, cq_meta, filter_cq_id, metric_name, metric_namespace, metric_value, default_value) FROM stdin;
\.


--
-- Data for Name: aws_cloudwatchlogs_filters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cloudwatchlogs_filters (cq_id, cq_meta, account_id, region, creation_time, name, pattern, log_group_name) FROM stdin;
\.


--
-- Data for Name: aws_codebuild_project_environment_variables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codebuild_project_environment_variables (cq_id, cq_meta, project_cq_id, name, value, type) FROM stdin;
\.


--
-- Data for Name: aws_codebuild_project_file_system_locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codebuild_project_file_system_locations (cq_id, cq_meta, project_cq_id, identifier, location, mount_options, mount_point, type) FROM stdin;
\.


--
-- Data for Name: aws_codebuild_project_secondary_artifacts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codebuild_project_secondary_artifacts (cq_id, cq_meta, project_cq_id, type, artifact_identifier, bucket_owner_access, encryption_disabled, location, name, namespace_type, override_artifact_name, packaging, path) FROM stdin;
\.


--
-- Data for Name: aws_codebuild_project_secondary_sources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codebuild_project_secondary_sources (cq_id, cq_meta, project_cq_id, type, auth_type, auth_resource, build_status_config_context, build_status_config_target_url, buildspec, git_clone_depth, git_submodules_config_fetch_submodules, insecure_ssl, location, report_build_status, source_identifier) FROM stdin;
\.


--
-- Data for Name: aws_codebuild_projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codebuild_projects (cq_id, cq_meta, account_id, region, arn, artifacts_type, artifacts_artifact_identifier, artifacts_bucket_owner_access, artifacts_encryption_disabled, artifacts_location, artifacts_name, artifacts_namespace_type, artifacts_override_artifact_name, artifacts_packaging, artifacts_path, badge_enabled, badge_request_url, build_batch_config_batch_report_mode, build_batch_config_combine_artifacts, build_batch_config_restrictions_compute_types_allowed, build_batch_config_restrictions_maximum_builds_allowed, build_batch_config_service_role, build_batch_config_timeout_in_mins, cache_type, cache_location, cache_modes, concurrent_build_limit, created, description, encryption_key, environment_compute_type, environment_image, environment_type, environment_certificate, environment_image_pull_credentials_type, environment_privileged_mode, environment_registry_credential, environment_registry_credential_credential_provider, last_modified, logs_config_cloud_watch_logs_status, logs_config_cloud_watch_logs_group_name, logs_config_cloud_watch_logs_stream_name, logs_config_s3_logs_status, logs_config_s3_logs_bucket_owner_access, logs_config_s3_logs_encryption_disabled, logs_config_s3_logs_location, name, project_visibility, public_project_alias, queued_timeout_in_minutes, resource_access_role, secondary_source_versions, service_role, source_type, source_auth_type, source_auth_resource, source_build_status_config_context, source_build_status_config_target_url, source_buildspec, source_git_clone_depth, source_git_submodules_config_fetch_submodules, source_insecure_ssl, source_location, source_report_build_status, source_identifier, source_version, tags, timeout_in_minutes, vpc_config_security_group_ids, vpc_config_subnets, vpc_config_vpc_id, webhook_branch_filter, webhook_build_type, webhook_filter_groups, webhook_last_modified_secret, webhook_payload_url, webhook_secret, webhook_url) FROM stdin;
\.


--
-- Data for Name: aws_codepipeline_pipeline_stage_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codepipeline_pipeline_stage_actions (cq_id, cq_meta, pipeline_stage_cq_id, category, owner, provider, version, name, configuration, input_artifacts, namespace, output_artifacts, region, role_arn, run_order) FROM stdin;
\.


--
-- Data for Name: aws_codepipeline_pipeline_stages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codepipeline_pipeline_stages (cq_id, cq_meta, pipeline_cq_id, stage_order, name, blockers) FROM stdin;
\.


--
-- Data for Name: aws_codepipeline_pipelines; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codepipeline_pipelines (cq_id, cq_meta, account_id, region, tags, created, arn, updated, name, role_arn, artifact_store_location, artifact_store_type, artifact_store_encryption_key_id, artifact_store_encryption_key_type, artifact_stores, version) FROM stdin;
\.


--
-- Data for Name: aws_codepipeline_webhook_filters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codepipeline_webhook_filters (cq_id, cq_meta, webhook_cq_id, json_path, match_equals) FROM stdin;
\.


--
-- Data for Name: aws_codepipeline_webhooks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_codepipeline_webhooks (cq_id, cq_meta, account_id, region, authentication, authentication_allowed_ip_range, authentication_secret_token, name, target_action, target_pipeline, url, arn, error_code, error_message, last_triggered, tags) FROM stdin;
\.


--
-- Data for Name: aws_cognito_identity_pool_cognito_identity_providers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cognito_identity_pool_cognito_identity_providers (cq_id, cq_meta, identity_pool_cq_id, identity_pool_id, client_id, provider_name, server_side_token_check) FROM stdin;
\.


--
-- Data for Name: aws_cognito_identity_pools; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cognito_identity_pools (cq_id, cq_meta, account_id, region, arn, allow_unauthenticated_identities, id, identity_pool_name, allow_classic_flow, developer_provider_name, identity_pool_tags, open_id_connect_provider_arns, saml_provider_arns, supported_login_providers) FROM stdin;
\.


--
-- Data for Name: aws_cognito_user_pool_identity_providers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cognito_user_pool_identity_providers (cq_id, cq_meta, user_pool_cq_id, user_pool_id, account_id, region, attribute_mapping, creation_date, idp_identifiers, last_modified_date, provider_details, provider_name, provider_type) FROM stdin;
\.


--
-- Data for Name: aws_cognito_user_pool_schema_attributes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cognito_user_pool_schema_attributes (cq_id, cq_meta, user_pool_cq_id, user_pool_id, attribute_data_type, developer_only_attribute, mutable, name, number_attribute_constraints_max_value, number_attribute_constraints_min_value, required, string_attribute_constraints_max_length, string_attribute_constraints_min_length) FROM stdin;
\.


--
-- Data for Name: aws_cognito_user_pools; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_cognito_user_pools (cq_id, cq_meta, account_id, region, account_recovery_setting, admin_create_user_admin_only, admin_create_user_invite_email_message, admin_create_user_invite_email_subject, admin_create_user_invite_sms, admin_create_user_config_unused_account_validity_days, alias_attributes, arn, auto_verified_attributes, creation_date, custom_domain, challenge_required_on_new_device, device_only_remembered_on_user_prompt, domain, email_configuration_set, email_configuration_sending_account, email_configuration_from, email_configuration_reply_to_address, email_configuration_source_arn, email_configuration_failure, email_verification_message, email_verification_subject, estimated_number_of_users, id, lambda_config_create_auth_challenge, lambda_config_custom_email_sender_lambda_arn, lambda_config_custom_email_sender_lambda_version, lambda_config_custom_message, lambda_config_custom_sms_sender_lambda_arn, lambda_config_custom_sms_sender_lambda_version, lambda_config_define_auth_challenge, lambda_config_kms_key_id, lambda_config_post_authentication, lambda_config_post_confirmation, lambda_config_pre_authentication, lambda_config_pre_sign_up, lambda_config_pre_token_generation, lambda_config_user_migration, lambda_config_verify_auth_challenge_response, last_modified_date, mfa_configuration, name, policies_password_policy_minimum_length, policies_password_policy_require_lowercase, policies_password_policy_require_numbers, policies_password_policy_require_symbols, policies_password_policy_require_uppercase, policies_password_policy_temporary_password_validity_days, sms_authentication_message, sms_configuration_sns_caller_arn, sms_configuration_external_id, sms_configuration_failure, sms_verification_message, status, user_pool_add_ons_advanced_security_mode, user_pool_tags, username_attributes, username_configuration_case_sensitive, verification_message_template_default_email_option, verification_message_template_email_message, verification_message_template_email_message_by_link, verification_message_template_email_subject, verification_message_template_email_subject_by_link, verification_message_template_sms_message) FROM stdin;
\.


--
-- Data for Name: aws_config_configuration_recorders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_config_configuration_recorders (cq_id, cq_meta, account_id, region, arn, name, recording_group_all_supported, recording_group_include_global_resource_types, recording_group_resource_types, role_arn, status_last_error_code, status_last_error_message, status_last_start_time, status_last_status, status_last_status_change_time, status_last_stop_time, status_recording) FROM stdin;
\.


--
-- Data for Name: aws_config_conformance_pack_rule_compliances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_config_conformance_pack_rule_compliances (cq_id, cq_meta, conformance_pack_cq_id, compliance_type, config_rule_name, controls, config_rule_invoked_time, resource_id, resource_type, ordering_timestamp, result_recorded_time, annotation) FROM stdin;
\.


--
-- Data for Name: aws_config_conformance_packs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_config_conformance_packs (cq_id, cq_meta, account_id, region, arn, conformance_pack_id, conformance_pack_name, conformance_pack_input_parameters, created_by, delivery_s3_bucket, delivery_s3_key_prefix, last_update_requested_time) FROM stdin;
\.


--
-- Data for Name: aws_dax_cluster_nodes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dax_cluster_nodes (cq_id, cq_meta, cluster_cq_id, availability_zone, endpoint_address, endpoint_port, endpoint_url, node_create_time, node_id, node_status, parameter_group_status) FROM stdin;
\.


--
-- Data for Name: aws_dax_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dax_clusters (cq_id, cq_meta, account_id, region, tags, active_nodes, arn, cluster_discovery_endpoint_address, cluster_discovery_endpoint_port, cluster_discovery_endpoint_url, cluster_endpoint_encryption_type, name, description, iam_role_arn, node_ids_to_remove, node_type, notification_configuration_topic_arn, notification_configuration_topic_status, node_ids_to_reboot, parameter_apply_status, parameter_group_name, preferred_maintenance_window, sse_description_status, security_groups, status, subnet_group, total_nodes) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_connection_mac_sec_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_connection_mac_sec_keys (cq_id, cq_meta, connection_cq_id, connection_id, ckn, secret_arn, start_on, state) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_connections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_connections (cq_id, cq_meta, account_id, region, arn, aws_device_v2, bandwidth, id, name, connection_state, encryption_mode, has_logical_redundancy, jumbo_frame_capable, lag_id, loa_issue_time, location, mac_sec_capable, owner_account, partner_name, port_encryption_status, provider_name, tags, vlan) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_gateway_associations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_gateway_associations (cq_id, cq_meta, gateway_cq_id, gateway_id, allowed_prefixes_to_direct_connect_gateway, associated_gateway_id, associated_gateway_owner_account, associated_gateway_region, associated_gateway_type, association_id, association_state, direct_connect_gateway_owner_account, state_change_error, virtual_gateway_id, virtual_gateway_owner_account, resource_id) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_gateway_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_gateway_attachments (cq_id, cq_meta, gateway_cq_id, gateway_id, attachment_state, attachment_type, state_change_error, virtual_interface_id, virtual_interface_owner_account, virtual_interface_region) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_gateways (cq_id, cq_meta, account_id, arn, amazon_side_asn, id, name, state, owner_account, state_change_error) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_lag_mac_sec_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_lag_mac_sec_keys (cq_id, cq_meta, lag_cq_id, lag_id, ckn, secret_arn, start_on, state) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_lags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_lags (cq_id, cq_meta, account_id, region, arn, allows_hosted_connections, aws_device_v2, connection_ids, connections_bandwidth, encryption_mode, has_logical_redundancy, jumbo_frame_capable, id, name, state, location, mac_sec_capable, minimum_links, number_of_connections, owner_account, provider_name, tags) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_virtual_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_virtual_gateways (cq_id, cq_meta, account_id, region, id, state) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_virtual_interface_bgp_peers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_virtual_interface_bgp_peers (cq_id, cq_meta, virtual_interface_cq_id, virtual_interface_id, address_family, amazon_address, asn, auth_key, aws_device_v2, bgp_peer_id, bgp_peer_state, bgp_status, customer_address) FROM stdin;
\.


--
-- Data for Name: aws_directconnect_virtual_interfaces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_directconnect_virtual_interfaces (cq_id, cq_meta, account_id, arn, address_family, amazon_address, amazon_side_asn, asn, auth_key, aws_device_v2, connection_id, customer_address, customer_router_config, direct_connect_gateway_id, jumbo_frame_capable, location, mtu, owner_account, region, route_filter_prefixes, tags, virtual_gateway_id, id, virtual_interface_name, virtual_interface_state, virtual_interface_type, vlan) FROM stdin;
\.


--
-- Data for Name: aws_dms_replication_instance_replication_subnet_group_subnets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dms_replication_instance_replication_subnet_group_subnets (cq_id, cq_meta, replication_instance_cq_id, subnet_availability_zone_name, subnet_identifier, subnet_status) FROM stdin;
\.


--
-- Data for Name: aws_dms_replication_instance_vpc_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dms_replication_instance_vpc_security_groups (cq_id, cq_meta, replication_instance_cq_id, status, vpc_security_group_id) FROM stdin;
\.


--
-- Data for Name: aws_dms_replication_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dms_replication_instances (cq_id, cq_meta, account_id, region, tags, allocated_storage, auto_minor_version_upgrade, availability_zone, dns_name_servers, engine_version, free_until, instance_create_time, kms_key_id, multi_az, pending_modified_values_allocated_storage, pending_modified_values_engine_version, pending_modified_values_multi_az, pending_modified_values_class, preferred_maintenance_window, publicly_accessible, arn, class, identifier, private_ip_address, private_ip_addresses, public_ip_address, public_ip_addresses, status, replication_subnet_group_description, replication_subnet_group_identifier, replication_subnet_group_subnet_group_status, replication_subnet_group_vpc_id, secondary_availability_zone) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_table_continuous_backups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_table_continuous_backups (cq_id, cq_meta, table_cq_id, continuous_backups_status, earliest_restorable_date_time, latest_restorable_date_time, point_in_time_recovery_status) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_table_global_secondary_indexes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_table_global_secondary_indexes (cq_id, cq_meta, table_cq_id, backfilling, arn, name, index_size_bytes, status, item_count, key_schema, projection_non_key_attributes, projection_type, provisioned_throughput_last_decrease_date_time, provisioned_throughput_last_increase_date_time, provisioned_throughput_number_of_decreases_today, provisioned_throughput_read_capacity_units, provisioned_throughput_write_capacity_units) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_table_local_secondary_indexes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_table_local_secondary_indexes (cq_id, cq_meta, table_cq_id, arn, name, index_size_bytes, item_count, key_schema, projection_non_key_attributes, projection_type) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_table_replica_auto_scalings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_table_replica_auto_scalings (cq_id, cq_meta, table_cq_id, global_secondary_indexes, region_name, read_capacity, write_capacity, replica_status) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_table_replicas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_table_replicas (cq_id, cq_meta, table_cq_id, global_secondary_indexes, kms_master_key_id, provisioned_throughput_override_read_capacity_units, region_name, replica_inaccessible_date_time, replica_status, replica_status_description, replica_status_percent_progress, summary_last_update_date_time, summary_table_class) FROM stdin;
\.


--
-- Data for Name: aws_dynamodb_tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_dynamodb_tables (cq_id, cq_meta, account_id, region, tags, archival_summary, attribute_definitions, billing_mode_summary, creation_date_time, global_table_version, item_count, key_schema, latest_stream_arn, latest_stream_label, provisioned_throughput_last_decrease_date_time, provisioned_throughput_last_increase_date_time, provisioned_throughput_number_of_decreases_today, provisioned_throughput_read_capacity_units, provisioned_throughput_write_capacity_units, restore_summary, inaccessible_encryption_date_time, kms_master_key_arn, sse_type, sse_status, stream_specification, arn, table_class_last_update, table_class, id, name, size_bytes, status) FROM stdin;
\.


--
-- Data for Name: aws_ec2_byoip_cidrs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_byoip_cidrs (cq_id, cq_meta, account_id, region, cidr, description, state, status_message) FROM stdin;
\.


--
-- Data for Name: aws_ec2_customer_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_customer_gateways (cq_id, cq_meta, account_id, region, id, bgp_asn, certificate_arn, arn, device_name, ip_address, state, tags, type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_ebs_snapshots; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_ebs_snapshots (cq_id, cq_meta, account_id, region, create_volume_permissions, data_encryption_key_id, description, encrypted, kms_key_id, outpost_arn, owner_alias, owner_id, progress, snapshot_id, start_time, state, state_message, tags, volume_id, volume_size) FROM stdin;
\.


--
-- Data for Name: aws_ec2_ebs_volume_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_ebs_volume_attachments (cq_id, cq_meta, ebs_volume_cq_id, attach_time, delete_on_termination, device, instance_id, state, volume_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_ebs_volumes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_ebs_volumes (cq_id, cq_meta, account_id, region, id, arn, availability_zone, create_time, encrypted, fast_restored, iops, kms_key_id, multi_attach_enabled, outpost_arn, size, snapshot_id, state, tags, throughput, volume_type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_egress_only_internet_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_egress_only_internet_gateways (cq_id, cq_meta, account_id, region, arn, attachments, id, tags) FROM stdin;
\.


--
-- Data for Name: aws_ec2_eips; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_eips (cq_id, cq_meta, account_id, region, allocation_id, association_id, carrier_ip, customer_owned_ip, customer_owned_ipv4_pool, domain, instance_id, network_border_group, network_interface_id, network_interface_owner_id, private_ip_address, public_ip, public_ipv4_pool, tags) FROM stdin;
\.


--
-- Data for Name: aws_ec2_flow_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_flow_logs (cq_id, cq_meta, account_id, region, arn, id, creation_time, deliver_logs_error_message, deliver_logs_permission_arn, deliver_logs_status, flow_log_id, flow_log_status, log_destination, log_destination_type, log_format, log_group_name, max_aggregation_interval, resource_id, tags, traffic_type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_host_available_instance_capacity; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_host_available_instance_capacity (cq_id, cq_meta, host_cq_id, available_capacity, instance_type, total_capacity) FROM stdin;
\.


--
-- Data for Name: aws_ec2_host_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_host_instances (cq_id, cq_meta, host_cq_id, instance_id, instance_type, owner_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_hosts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_hosts (cq_id, cq_meta, account_id, region, arn, allocation_time, allows_multiple_instance_types, auto_placement, availability_zone, availability_zone_id, available_vcpus, client_token, id, cores, instance_family, instance_type, sockets, total_vcpus, host_recovery, reservation_id, member_of_service_linked_resource_group, owner_id, release_time, state, tags) FROM stdin;
\.


--
-- Data for Name: aws_ec2_image_block_device_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_image_block_device_mappings (cq_id, cq_meta, image_cq_id, device_name, ebs_delete_on_termination, ebs_encrypted, ebs_iops, ebs_kms_key_id, ebs_outpost_arn, ebs_snapshot_id, ebs_throughput, ebs_volume_size, ebs_volume_type, no_device, virtual_name) FROM stdin;
\.


--
-- Data for Name: aws_ec2_images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_images (cq_id, cq_meta, account_id, region, arn, id, architecture, creation_date, description, ena_support, hypervisor, image_location, image_owner_alias, image_type, kernel_id, name, owner_id, platform, platform_details, product_codes, public, ramdisk_id, root_device_name, root_device_type, sriov_net_support, state, state_reason_code, state_reason_message, tags, usage_operation, virtualization_type, last_launched_time) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_block_device_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_block_device_mappings (cq_id, cq_meta, instance_cq_id, device_name, ebs_attach_time, ebs_delete_on_termination, ebs_status, ebs_volume_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_elastic_gpu_associations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_elastic_gpu_associations (cq_id, cq_meta, instance_cq_id, elastic_gpu_association_id, elastic_gpu_association_state, elastic_gpu_association_time, elastic_gpu_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_elastic_inference_accelerator_associations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_elastic_inference_accelerator_associations (cq_id, cq_meta, instance_cq_id, elastic_inference_accelerator_arn, elastic_inference_accelerator_association_id, elastic_inference_accelerator_association_state, elastic_inference_accelerator_association_time) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_network_interface_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_network_interface_groups (cq_id, cq_meta, instance_network_interface_cq_id, network_interface_id, group_id, group_name) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_network_interface_ipv6_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_network_interface_ipv6_addresses (cq_id, cq_meta, instance_network_interface_cq_id, ipv6_address) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_network_interface_private_ip_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_network_interface_private_ip_addresses (cq_id, cq_meta, instance_network_interface_cq_id, association_carrier_ip, association_ip_owner_id, association_public_dns_name, association_public_ip, is_primary, private_dns_name, private_ip_address) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_network_interfaces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_network_interfaces (cq_id, cq_meta, instance_cq_id, arn, association_carrier_ip, association_ip_owner_id, association_public_dns_name, association_public_ip, attachment_attach_time, attachment_id, attachment_delete_on_termination, attachment_device_index, attachment_network_card_index, attachment_status, description, interface_type, ipv4_prefixes, ipv6_prefixes, mac_address, network_interface_id, owner_id, private_dns_name, private_ip_address, source_dest_check, status, subnet_id, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_product_codes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_product_codes (cq_id, cq_meta, instance_cq_id, product_code_id, product_code_type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_security_groups (cq_id, cq_meta, instance_cq_id, group_id, group_name) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_status_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_status_events (cq_id, cq_meta, instance_status_cq_id, code, description, id, not_after, not_before, not_before_deadline) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instance_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instance_statuses (cq_id, cq_meta, account_id, region, arn, availability_zone, instance_id, instance_state_code, instance_state_name, details, status, outpost_arn, system_status, system_status_details) FROM stdin;
\.


--
-- Data for Name: aws_ec2_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_instances (cq_id, cq_meta, account_id, region, arn, state_transition_reason_time, ami_launch_index, architecture, boot_mode, capacity_reservation_id, cap_reservation_preference, cap_reservation_target_capacity_reservation_id, cap_reservation_target_capacity_reservation_rg_arn, client_token, cpu_options_core_count, cpu_options_threads_per_core, ebs_optimized, ena_support, enclave_options_enabled, hibernation_options_configured, hypervisor, iam_instance_profile_arn, iam_instance_profile_id, image_id, id, instance_lifecycle, instance_type, kernel_id, key_name, launch_time, licenses, metadata_options_http_endpoint, metadata_options_http_protocol_ipv6, metadata_options_http_put_response_hop_limit, metadata_options_http_tokens, metadata_options_state, monitoring_state, outpost_arn, placement_affinity, placement_availability_zone, placement_group_name, placement_host_id, placement_host_resource_group_arn, placement_partition_number, placement_spread_domain, placement_tenancy, platform, private_dns_name, private_ip_address, public_dns_name, public_ip_address, ramdisk_id, root_device_name, root_device_type, source_dest_check, spot_instance_request_id, sriov_net_support, state_code, state_name, state_reason_code, state_reason_message, state_transition_reason, subnet_id, tags, virtualization_type, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_internet_gateway_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_internet_gateway_attachments (cq_id, cq_meta, internet_gateway_cq_id, state, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_internet_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_internet_gateways (cq_id, cq_meta, account_id, region, arn, id, owner_id, tags) FROM stdin;
\.


--
-- Data for Name: aws_ec2_nat_gateway_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_nat_gateway_addresses (cq_id, cq_meta, nat_gateway_cq_id, allocation_id, network_interface_id, private_ip, public_ip) FROM stdin;
\.


--
-- Data for Name: aws_ec2_nat_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_nat_gateways (cq_id, cq_meta, account_id, region, arn, id, create_time, delete_time, failure_code, failure_message, provisioned_bandwidth_provision_time, provisioned_bandwidth_provisioned, provisioned_bandwidth_request_time, provisioned_bandwidth_requested, provisioned_bandwidth_status, state, subnet_id, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_network_acl_associations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_network_acl_associations (cq_id, cq_meta, network_acl_cq_id, network_acl_association_id, subnet_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_network_acl_entries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_network_acl_entries (cq_id, cq_meta, network_acl_cq_id, cidr_block, egress, icmp_type_code, icmp_type_code_type, ipv6_cidr_block, port_range_from, port_range_to, protocol, rule_action, rule_number) FROM stdin;
\.


--
-- Data for Name: aws_ec2_network_acls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_network_acls (cq_id, cq_meta, account_id, region, arn, is_default, id, owner_id, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_network_interface_private_ip_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_network_interface_private_ip_addresses (cq_id, cq_meta, network_interface_cq_id, association_allocation_id, association_id, association_carrier_ip, association_customer_owned_ip, association_ip_owner_id, association_public_dns_name, association_public_ip, "primary", private_dns_name, private_ip_address) FROM stdin;
\.


--
-- Data for Name: aws_ec2_network_interfaces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_network_interfaces (cq_id, cq_meta, account_id, region, arn, tags, association_allocation_id, association_id, association_carrier_ip, association_customer_owned_ip, association_ip_owner_id, association_public_dns_name, association_public_ip, attachment_attach_time, attachment_id, attachment_delete_on_termination, attachment_device_index, attachment_instance_id, attachment_instance_owner_id, attachment_network_card_index, attachment_status, availability_zone, deny_all_igw_traffic, description, groups, interface_type, ipv4_prefixes, ipv6_address, ipv6_addresses, ipv6_native, ipv6_prefixes, mac_address, id, outpost_arn, owner_id, private_dns_name, private_ip_address, requester_id, requester_managed, source_dest_check, status, subnet_id, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_regional_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_regional_config (cq_id, cq_meta, account_id, region, ebs_encryption_enabled_by_default, ebs_default_kms_key_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_route_table_associations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_route_table_associations (cq_id, cq_meta, route_table_cq_id, id, association_state, association_state_status_message, gateway_id, main, subnet_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_route_table_propagating_vgws; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_route_table_propagating_vgws (cq_id, cq_meta, route_table_cq_id, gateway_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_route_table_routes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_route_table_routes (cq_id, cq_meta, route_table_cq_id, carrier_gateway_id, destination_cidr_block, destination_ipv6_cidr_block, destination_prefix_list_id, egress_only_internet_gateway_id, gateway_id, instance_id, instance_owner_id, local_gateway_id, nat_gateway_id, network_interface_id, origin, state, transit_gateway_id, vpc_peering_connection_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_route_tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_route_tables (cq_id, cq_meta, account_id, region, arn, owner_id, id, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_security_group_ip_permission_ip_ranges; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_security_group_ip_permission_ip_ranges (cq_id, cq_meta, security_group_ip_permission_cq_id, cidr, description, cidr_type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_security_group_ip_permission_prefix_list_ids; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_security_group_ip_permission_prefix_list_ids (cq_id, cq_meta, security_group_ip_permission_cq_id, description, prefix_list_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_security_group_ip_permission_user_id_group_pairs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_security_group_ip_permission_user_id_group_pairs (cq_id, cq_meta, security_group_ip_permission_cq_id, description, group_id, group_name, peering_status, user_id, vpc_id, vpc_peering_connection_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_security_group_ip_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_security_group_ip_permissions (cq_id, cq_meta, security_group_cq_id, from_port, ip_protocol, to_port, permission_type) FROM stdin;
\.


--
-- Data for Name: aws_ec2_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_security_groups (cq_id, cq_meta, account_id, region, arn, description, id, group_name, owner_id, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_subnet_ipv6_cidr_block_association_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_subnet_ipv6_cidr_block_association_sets (cq_id, cq_meta, subnet_cq_id, association_id, ipv6_cidr_block, ipv6_cidr_block_state, ipv6_cidr_block_state_status_message) FROM stdin;
\.


--
-- Data for Name: aws_ec2_subnets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_subnets (cq_id, cq_meta, account_id, region, assign_ipv6_address_on_creation, availability_zone, availability_zone_id, available_ip_address_count, cidr_block, customer_owned_ipv4_pool, default_for_az, map_customer_owned_ip_on_launch, map_public_ip_on_launch, outpost_arn, owner_id, state, arn, id, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateway_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateway_attachments (cq_id, cq_meta, transit_gateway_cq_id, association_state, association_route_table_id, creation_time, resource_id, resource_owner_id, resource_type, state, tags, transit_gateway_owner_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateway_multicast_domains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateway_multicast_domains (cq_id, cq_meta, transit_gateway_cq_id, creation_time, auto_accept_shared_associations, igmpv2_support, static_sources_support, owner_id, state, tags, transit_gateway_multicast_domain_arn, transit_gateway_multicast_domain_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateway_peering_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateway_peering_attachments (cq_id, cq_meta, transit_gateway_cq_id, accepter_owner_id, accepter_region, accepter_transit_gateway_id, creation_time, requester_owner_id, requester_region, requester_transit_gateway_id, state, status_code, status_message, tags, transit_gateway_attachment_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateway_route_tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateway_route_tables (cq_id, cq_meta, transit_gateway_cq_id, creation_time, default_association_route_table, default_propagation_route_table, state, tags, transit_gateway_route_table_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateway_vpc_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateway_vpc_attachments (cq_id, cq_meta, transit_gateway_cq_id, creation_time, appliance_mode_support, dns_support, ipv6_support, state, tags, transit_gateway_attachment_id, vpc_id, vpc_owner_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_transit_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_transit_gateways (cq_id, cq_meta, account_id, region, amazon_side_asn, association_default_route_table_id, auto_accept_shared_attachments, creation_time, default_route_table_association, default_route_table_propagation, description, dns_support, multicast_support, owner_id, propagation_default_route_table_id, state, tags, arn, transit_gateway_cidr_blocks, id, vpn_ecmp_support) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_attachment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_attachment (cq_id, cq_meta, vpn_gateway_cq_id, state, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_cidr_block_association_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_cidr_block_association_sets (cq_id, cq_meta, vpc_cq_id, association_id, cidr_block, cidr_block_state, cidr_block_state_status_message) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_endpoint_dns_entries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_endpoint_dns_entries (cq_id, cq_meta, vpc_endpoint_cq_id, dns_name, hosted_zone_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_endpoint_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_endpoint_groups (cq_id, cq_meta, vpc_endpoint_cq_id, group_id, group_name) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_endpoints; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_endpoints (cq_id, cq_meta, account_id, region, arn, creation_timestamp, last_error_code, last_error_message, network_interface_ids, owner_id, policy_document, private_dns_enabled, requester_managed, route_table_ids, service_name, state, subnet_ids, tags, id, vpc_endpoint_type, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_ipv6_cidr_block_association_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_ipv6_cidr_block_association_sets (cq_id, cq_meta, vpc_cq_id, association_id, ipv6_cidr_block, ipv6_cidr_block_state, ipv6_cidr_block_state_status_message, ipv6_pool, network_border_group) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpc_peering_connections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpc_peering_connections (cq_id, cq_meta, account_id, region, arn, accepter_cidr_block, accepter_cidr_block_set, accepter_ipv6_cidr_block_set, accepter_owner_id, accepter_allow_dns_resolution_from_remote_vpc, accepter_allow_egress_local_classic_link_to_remote_vpc, accepter_allow_egress_local_vpc_to_remote_classic_link, accepter_vpc_region, accepter_vpc_id, expiration_time, requester_cidr_block, requester_cidr_block_set, requester_ipv6_cidr_block_set, requester_owner_id, requester_allow_dns_resolution_from_remote_vpc, requester_allow_egress_local_classic_link_to_remote_vpc, requester_allow_egress_local_vpc_to_remote_classic_link, requester_vpc_region, requester_vpc_id, status_code, status_message, tags, id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpcs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpcs (cq_id, cq_meta, account_id, region, arn, cidr_block, dhcp_options_id, instance_tenancy, is_default, owner_id, state, tags, id) FROM stdin;
\.


--
-- Data for Name: aws_ec2_vpn_gateways; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ec2_vpn_gateways (cq_id, cq_meta, account_id, region, arn, amazon_side_asn, availability_zone, state, tags, type, id) FROM stdin;
\.


--
-- Data for Name: aws_ecr_repositories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecr_repositories (cq_id, cq_meta, account_id, region, created_at, encryption_configuration_encryption_type, encryption_configuration_kms_key, image_scanning_configuration_scan_on_push, image_tag_mutability, registry_id, arn, name, uri) FROM stdin;
\.


--
-- Data for Name: aws_ecr_repository_images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecr_repository_images (cq_id, cq_meta, repository_cq_id, account_id, region, artifact_media_type, image_digest, image_manifest_media_type, image_pushed_at, image_scan_findings_summary_finding_severity_counts, image_scan_findings_summary_image_scan_completed_at, image_scan_findings_summary_vulnerability_source_updated_at, image_scan_status_description, image_scan_status, image_size_in_bytes, image_tags, registry_id, repository_name) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_attachments (cq_id, cq_meta, cluster_cq_id, details, id, status, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instance_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instance_attachments (cq_id, cq_meta, cluster_container_instance_cq_id, details, id, status, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instance_attributes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instance_attributes (cq_id, cq_meta, cluster_container_instance_cq_id, name, target_id, target_type, value) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instance_health_status_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instance_health_status_details (cq_id, cq_meta, cluster_container_instance_cq_id, last_status_change, last_updated, status, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instance_registered_resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instance_registered_resources (cq_id, cq_meta, cluster_container_instance_cq_id, double_value, integer_value, long_value, name, string_set_value, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instance_remaining_resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instance_remaining_resources (cq_id, cq_meta, cluster_container_instance_cq_id, double_value, integer_value, long_value, name, string_set_value, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_container_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_container_instances (cq_id, cq_meta, cluster_cq_id, agent_connected, agent_update_status, capacity_provider_name, container_instance_arn, ec2_instance_id, health_status_overall_status, pending_tasks_count, registered_at, running_tasks_count, status, status_reason, tags, version, version_info_agent_hash, version_info_agent_version, version_info_docker_version) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_deployments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_deployments (cq_id, cq_meta, cluster_service_cq_id, capacity_provider_strategy, created_at, desired_count, failed_tasks, id, launch_type, network_configuration_awsvpc_configuration_subnets, network_configuration_awsvpc_configuration_assign_public_ip, network_configuration_awsvpc_configuration_security_groups, pending_count, platform_family, platform_version, rollout_state, rollout_state_reason, running_count, status, task_definition, updated_at) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_events (cq_id, cq_meta, cluster_service_cq_id, created_at, id, message) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_load_balancers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_load_balancers (cq_id, cq_meta, cluster_service_cq_id, container_name, container_port, load_balancer_name, target_group_arn) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_service_registries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_service_registries (cq_id, cq_meta, cluster_service_cq_id, container_name, container_port, port, registry_arn) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_task_set_load_balancers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_task_set_load_balancers (cq_id, cq_meta, cluster_service_task_set_cq_id, container_name, container_port, load_balancer_name, target_group_arn) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_task_set_service_registries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_task_set_service_registries (cq_id, cq_meta, cluster_service_task_set_cq_id, container_name, container_port, port, arn) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_service_task_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_service_task_sets (cq_id, cq_meta, cluster_service_cq_id, capacity_provider_strategy, cluster_arn, computed_desired_count, created_at, external_id, id, launch_type, network_configuration_awsvpc_configuration_subnets, network_configuration_awsvpc_configuration_assign_public_ip, network_configuration_awsvpc_configuration_security_groups, pending_count, platform_family, platform_version, running_count, scale_unit, scale_value, service_arn, stability_status, stability_status_at, started_by, status, tags, task_definition, arn, updated_at) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_services; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_services (cq_id, cq_meta, cluster_cq_id, capacity_provider_strategy, cluster_arn, created_at, created_by, deployment_configuration_deployment_circuit_breaker_enable, deployment_configuration_deployment_circuit_breaker_rollback, deployment_configuration_maximum_percent, deployment_configuration_minimum_healthy_percent, deployment_controller_type, desired_count, enable_ecs_managed_tags, enable_execute_command, health_check_grace_period_seconds, launch_type, network_configuration_awsvpc_configuration_subnets, network_configuration_awsvpc_configuration_assign_public_ip, network_configuration_awsvpc_configuration_security_groups, pending_count, placement_constraints, placement_strategy, platform_family, platform_version, propagate_tags, role_arn, running_count, scheduling_strategy, arn, name, status, tags, task_definition) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_task_attachments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_task_attachments (cq_id, cq_meta, cluster_task_cq_id, details, id, status, type) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_task_containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_task_containers (cq_id, cq_meta, cluster_task_cq_id, container_arn, cpu, exit_code, gpu_ids, health_status, image, image_digest, last_status, managed_agents, memory, memory_reservation, name, network_bindings, network_interfaces, reason, runtime_id, task_arn) FROM stdin;
\.


--
-- Data for Name: aws_ecs_cluster_tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_cluster_tasks (cq_id, cq_meta, cluster_cq_id, attributes, availability_zone, capacity_provider_name, cluster_arn, connectivity, connectivity_at, container_instance_arn, cpu, created_at, desired_status, enable_execute_command, ephemeral_storage_size_in_gib, execution_stopped_at, "group", health_status, inference_accelerators, last_status, launch_type, memory, overrides, platform_family, platform_version, pull_started_at, pull_stopped_at, started_at, started_by, stop_code, stopped_at, stopped_reason, stopping_at, tags, arn, task_definition_arn, version) FROM stdin;
\.


--
-- Data for Name: aws_ecs_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_clusters (cq_id, cq_meta, account_id, region, active_services_count, attachments_status, capacity_providers, arn, name, execute_config_kms_key_id, execute_config_logs_cloud_watch_encryption_enabled, execute_config_log_cloud_watch_log_group_name, execute_config_log_s3_bucket_name, execute_config_log_s3_encryption_enabled, execute_config_log_s3_key_prefix, execute_config_logging, default_capacity_provider_strategy, pending_tasks_count, registered_container_instances_count, running_tasks_count, settings, statistics, status, tags) FROM stdin;
\.


--
-- Data for Name: aws_ecs_task_definition_container_definitions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_task_definition_container_definitions (cq_id, cq_meta, task_definition_cq_id, command, cpu, depends_on, disable_networking, dns_search_domains, dns_servers, docker_labels, docker_security_options, entry_point, environment, environment_files, essential, extra_hosts, firelens_configuration_type, firelens_configuration_options, health_check_command, health_check_interval, health_check_retries, health_check_start_period, health_check_timeout, hostname, image, interactive, links, linux_parameters_capabilities_add, linux_parameters_capabilities_drop, linux_parameters_devices, linux_parameters_init_process_enabled, linux_parameters_max_swap, linux_parameters_shared_memory_size, linux_parameters_swappiness, linux_parameters_tmpfs, log_configuration_log_driver, log_configuration_options, log_configuration_secret_options, memory, memory_reservation, mount_points, name, port_mappings, privileged, pseudo_terminal, readonly_root_filesystem, repository_credentials_parameter, resource_requirements, secrets, start_timeout, stop_timeout, system_controls, ulimits, "user", volumes_from, working_directory) FROM stdin;
\.


--
-- Data for Name: aws_ecs_task_definition_volumes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_task_definition_volumes (cq_id, cq_meta, task_definition_cq_id, docker_autoprovision, docker_driver, docker_driver_opts, docker_labels, docker_scope, efs_file_system_id, efs_authorization_config_access_point_id, efs_authorization_config_iam, efs_root_directory, efs_volume_configuration_transit_encryption, efs_transit_encryption_port, fsx_wfs_authorization_config_credentials_parameter, fsx_wfs_authorization_config_domain, fsx_wfs_file_system_id, fsx_wfs_root_directory, host_source_path, name) FROM stdin;
\.


--
-- Data for Name: aws_ecs_task_definitions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ecs_task_definitions (cq_id, cq_meta, account_id, region, tags, compatibilities, cpu, deregistered_at, ephemeral_storage_size, execution_role_arn, family, inference_accelerators, ipc_mode, memory, network_mode, pid_mode, placement_constraints, proxy_configuration_container_name, proxy_configuration_properties, proxy_configuration_type, registered_at, registered_by, requires_attributes, requires_compatibilities, revision, runtime_platform_cpu_architecture, runtime_platform_os_family, status, arn, task_role_arn) FROM stdin;
\.


--
-- Data for Name: aws_efs_filesystems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_efs_filesystems (cq_id, cq_meta, account_id, region, backup_policy_status, creation_time, creation_token, id, life_cycle_state, number_of_mount_targets, owner_id, performance_mode, size_in_bytes_value, size_in_bytes_timestamp, size_in_bytes_value_in_ia, size_in_bytes_value_in_standard, tags, availability_zone_id, availability_zone_name, encrypted, arn, kms_key_id, name, provisioned_throughput_in_mibps, throughput_mode) FROM stdin;
\.


--
-- Data for Name: aws_eks_cluster_encryption_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_eks_cluster_encryption_configs (cq_id, cq_meta, cluster_cq_id, provider_key_arn, resources) FROM stdin;
\.


--
-- Data for Name: aws_eks_cluster_loggings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_eks_cluster_loggings (cq_id, cq_meta, cluster_cq_id, enabled, types) FROM stdin;
\.


--
-- Data for Name: aws_eks_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_eks_clusters (cq_id, cq_meta, account_id, region, arn, certificate_authority_data, client_request_token, created_at, endpoint, identity_oidc_issuer, kubernetes_network_config_service_ipv4_cidr, name, platform_version, resources_vpc_config_cluster_security_group_id, resources_vpc_config_endpoint_private_access, resources_vpc_config_endpoint_public_access, resources_vpc_config_public_access_cidrs, resources_vpc_config_security_group_ids, resources_vpc_config_subnet_ids, resources_vpc_config_vpc_id, role_arn, status, tags, version) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_application_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_application_versions (cq_id, cq_meta, account_id, region, application_name, arn, build_arn, date_created, date_updated, description, source_location, source_repository, source_type, source_bundle_s3_bucket, source_bundle_s3_key, status, version_label) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_applications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_applications (cq_id, cq_meta, account_id, region, arn, name, configuration_templates, date_created, date_updated, description, resource_lifecycle_config_service_role, max_age_rule_enabled, max_age_rule_delete_source_from_s3, max_age_rule_max_age_in_days, max_count_rule_enabled, max_count_rule_delete_source_from_s3, max_count_rule_max_count, versions) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_configuration_options; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_configuration_options (cq_id, cq_meta, environment_cq_id, application_arn, name, namespace, change_severity, default_value, max_length, max_value, min_value, regex_label, regex_pattern, user_defined, value_options, value_type) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_configuration_setting_options; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_configuration_setting_options (cq_id, cq_meta, configuration_setting_cq_id, namespace, option_name, resource_name, value) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_configuration_settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_configuration_settings (cq_id, cq_meta, environment_cq_id, application_name, application_arn, date_created, date_updated, deployment_status, description, environment_name, platform_arn, solution_stack_name, template_name) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_environment_links; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_environment_links (cq_id, cq_meta, environment_cq_id, environment_name, link_name) FROM stdin;
\.


--
-- Data for Name: aws_elasticbeanstalk_environments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticbeanstalk_environments (cq_id, cq_meta, account_id, region, tags, abortable_operation_in_progress, application_name, cname, date_created, date_updated, description, endpoint_url, arn, id, name, health, health_status, operations_role, platform_arn, load_balancer_domain, listeners, load_balancer_name, solution_stack_name, status, template_name, tier_name, tier_type, tier_version, version_label) FROM stdin;
\.


--
-- Data for Name: aws_elasticsearch_domains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elasticsearch_domains (cq_id, cq_meta, account_id, region, tags, arn, id, name, cluster_cold_storage_options_enabled, cluster_dedicated_master_count, cluster_dedicated_master_enabled, cluster_dedicated_master_type, cluster_instance_count, cluster_instance_type, cluster_warm_count, cluster_warm_enabled, cluster_warm_type, cluster_zone_awareness_config_availability_zone_count, cluster_zone_awareness_enabled, access_policies, advanced_options, advanced_security_enabled, advanced_security_internal_user_database_enabled, advanced_security_saml_enabled, advanced_security_saml_idp_entity_id, advanced_security_saml_roles_key, advanced_security_options_saml_options_roles_key, advanced_security_saml_session_timeout_minutes, advanced_security_saml_subject_key, auto_tune_error_message, auto_tune_options_state, cognito_enabled, cognito_identity_pool_id, cognito_role_arn, cognito_user_pool_id, created, deleted, domain_endpoint_custom, domain_endpoint_custom_certificate_arn, domain_endpoint_custom_enabled, domain_endpoint_enforce_https, domain_endpoint_tls_security_policy, ebs_enabled, ebs_iops, ebs_volume_size, ebs_volume_type, elasticsearch_version, encryption_at_rest_enabled, encryption_at_rest_kms_key_id, endpoint, endpoints, log_publishing_options, node_to_node_encryption_enabled, processing, service_software_automated_update_date, service_software_cancellable, service_software_current_version, service_software_description, service_software_new_version, service_software_optional_deployment, service_software_update_available, service_software_update_status, snapshot_options_automated_snapshot_start_hour, upgrade_processing, vpc_availability_zones, vpc_security_group_ids, vpc_subnet_ids, vpc_vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancer_backend_server_descriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancer_backend_server_descriptions (cq_id, cq_meta, load_balancer_cq_id, name, instance_port, policy_names) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancer_listeners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancer_listeners (cq_id, cq_meta, load_balancer_cq_id, load_balance_name, listener_instance_port, listener_load_balancer_port, listener_protocol, listener_instance_protocol, listener_ssl_certificate_id, policy_names) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancer_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancer_policies (cq_id, cq_meta, load_balancer_cq_id, load_balance_name, policy_attribute_descriptions, policy_name, policy_type_name) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancer_policies_app_cookie_stickiness; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancer_policies_app_cookie_stickiness (cq_id, cq_meta, load_balancer_cq_id, load_balance_name, cookie_name, policy_name) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancer_policies_lb_cookie_stickiness; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancer_policies_lb_cookie_stickiness (cq_id, cq_meta, load_balancer_cq_id, load_balance_name, cookie_expiration_period, policy_name) FROM stdin;
\.


--
-- Data for Name: aws_elbv1_load_balancers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv1_load_balancers (cq_id, cq_meta, account_id, region, arn, attributes_access_log_enabled, attributes_access_log_s3_bucket_name, attributes_access_log_s3_bucket_prefix, attributes_access_log_emit_interval, attributes_connection_settings_idle_timeout, attributes_cross_zone_load_balancing_enabled, attributes_connection_draining_enabled, attributes_connection_draining_timeout, attributes_additional_attributes, tags, availability_zones, canonical_hosted_zone_name, canonical_hosted_zone_name_id, created_time, dns_name, health_check_healthy_threshold, health_check_interval, health_check_target, health_check_timeout, health_check_unhealthy_threshold, instances, name, other_policies, scheme, security_groups, source_security_group_name, source_security_group_owner_alias, subnets, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_listener_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_listener_certificates (cq_id, cq_meta, listener_cq_id, certificate_arn, is_default) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_listener_default_action_forward_config_target_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_listener_default_action_forward_config_target_groups (cq_id, cq_meta, listener_default_action_cq_id, target_group_arn, weight) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_listener_default_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_listener_default_actions (cq_id, cq_meta, listener_cq_id, type, auth_cognito_user_pool_arn, auth_cognito_user_pool_client_id, auth_cognito_user_pool_domain, auth_cognito_authentication_request_extra_params, auth_cognito_on_unauthenticated_request, auth_cognito_scope, auth_cognito_session_cookie_name, auth_cognito_session_timeout, auth_oidc_authorization_endpoint, auth_oidc_client_id, auth_oidc_issuer, auth_oidc_token_endpoint, auth_oidc_user_info_endpoint, auth_oidc_authentication_request_extra_params, auth_oidc_client_secret, auth_oidc_on_unauthenticated_request, auth_oidc_scope, auth_oidc_session_cookie_name, auth_oidc_session_timeout, auth_oidc_use_existing_client_secret, fixed_response_config_status_code, fixed_response_config_content_type, fixed_response_config_message_body, forward_config_target_group_stickiness_config_duration_seconds, forward_config_target_group_stickiness_config_enabled, "order", redirect_config_status_code, redirect_config_host, redirect_config_path, redirect_config_port, redirect_config_protocol, redirect_config_query, target_group_arn) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_listeners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_listeners (cq_id, cq_meta, account_id, region, load_balancer_cq_id, tags, alpn_policy, arn, load_balancer_arn, port, protocol, ssl_policy) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_load_balancer_attributes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_load_balancer_attributes (cq_id, cq_meta, load_balancer_cq_id, access_logs_s3_enabled, access_logs_s3_bucket, access_logs_s3_prefix, deletion_protection, idle_timeout, routing_http_desync_mitigation_mode, routing_http_drop_invalid_header_fields, routing_http_xamzntls_enabled, routing_http_xff_client_port, routing_http2, waf_fail_open, load_balancing_cross_zone) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_load_balancer_availability_zone_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_load_balancer_availability_zone_addresses (cq_id, cq_meta, load_balancer_availability_zone_cq_id, zone_name, allocation_id, ipv6_address, ip_address, private_ipv4_address) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_load_balancer_availability_zones; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_load_balancer_availability_zones (cq_id, cq_meta, load_balancer_cq_id, load_balance_name, outpost_id, subnet_id, zone_name) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_load_balancers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_load_balancers (cq_id, cq_meta, account_id, region, web_acl_arn, tags, canonical_hosted_zone_id, created_time, customer_owned_ipv4_pool, dns_name, ip_address_type, arn, name, scheme, security_groups, state_code, state_reason, type, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_elbv2_target_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_elbv2_target_groups (cq_id, cq_meta, account_id, region, tags, health_check_enabled, health_check_interval_seconds, health_check_path, health_check_port, health_check_protocol, health_check_timeout_seconds, healthy_threshold_count, load_balancer_arns, matcher_grpc_code, matcher_http_code, port, protocol, protocol_version, arn, name, target_type, unhealthy_threshold_count, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_emr_block_public_access_config_port_ranges; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_emr_block_public_access_config_port_ranges (cq_id, cq_meta, block_public_access_config_cq_id, min_range, max_range) FROM stdin;
\.


--
-- Data for Name: aws_emr_block_public_access_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_emr_block_public_access_configs (cq_id, cq_meta, account_id, region, block_public_security_group_rules, classification, configurations, properties, created_by_arn, creation_date_time) FROM stdin;
\.


--
-- Data for Name: aws_emr_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_emr_clusters (cq_id, cq_meta, account_id, region, applications, auto_scaling_role, auto_terminate, arn, configurations, custom_ami_id, ebs_root_volume_size, ec2_instance_attribute_additional_master_security_groups, ec2_instance_attribute_additional_slave_security_groups, ec2_instance_attribute_availability_zone, ec2_instance_attribute_key_name, ec2_instance_attribute_subnet_id, ec2_instance_attribute_emr_managed_master_security_group, ec2_instance_attribute_emr_managed_slave_security_group, ec2_instance_attribute_iam_instance_profile, ec2_instance_attribute_requested_availability_zones, ec2_instance_attribute_requested_subnet_ids, ec2_instance_attribute_service_access_security_group, id, instance_collection_type, kerberos_kdc_admin_password, kerberos_realm, kerberos_ad_domain_join_password, kerberos_ad_domain_join_user, kerberos_cross_realm_trust_principal_password, log_encryption_kms_key_id, log_uri, master_public_dns_name, name, normalized_instance_hours, outpost_arn, placement_groups, release_label, repo_upgrade_on_boot, requested_ami_version, running_ami_version, scale_down_behavior, security_configuration, service_role, state, state_change_reason_code, state_change_reason_message, creation_date_time, end_date_time, ready_date_time, step_concurrency_level, tags, termination_protected, visible_to_all_users) FROM stdin;
\.


--
-- Data for Name: aws_fsx_backups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_fsx_backups (cq_id, cq_meta, account_id, region, id, creation_time, lifecycle, type, directory_information_active_directory_id, directory_information_domain_name, failure_details_message, kms_key_id, progress_percent, arn, tags) FROM stdin;
\.


--
-- Data for Name: aws_guardduty_detector_members; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_guardduty_detector_members (cq_id, cq_meta, detector_cq_id, account_id, email, master_id, relationship_status, updated_at, detector_id, invited_at) FROM stdin;
\.


--
-- Data for Name: aws_guardduty_detectors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_guardduty_detectors (cq_id, cq_meta, account_id, region, arn, id, service_role, status, created_at, data_sources_cloud_trail_status, data_sources_dns_logs_status, data_sources_flow_logs_status, data_sources_s3_logs_status, finding_publishing_frequency, tags, updated_at) FROM stdin;
\.


--
-- Data for Name: aws_iam_group_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_group_policies (cq_id, cq_meta, account_id, group_cq_id, group_id, group_name, policy_document, policy_name) FROM stdin;
\.


--
-- Data for Name: aws_iam_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_groups (cq_id, cq_meta, account_id, policies, arn, create_date, id, name, path) FROM stdin;
\.


--
-- Data for Name: aws_iam_openid_connect_identity_providers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_openid_connect_identity_providers (cq_id, cq_meta, account_id, arn, client_id_list, create_date, tags, thumbprint_list, url) FROM stdin;
\.


--
-- Data for Name: aws_iam_password_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_password_policies (cq_id, cq_meta, account_id, allow_users_to_change_password, expire_passwords, hard_expiry, max_password_age, minimum_password_length, password_reuse_prevention, require_lowercase_characters, require_numbers, require_symbols, require_uppercase_characters, policy_exists) FROM stdin;
\.


--
-- Data for Name: aws_iam_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_policies (cq_id, cq_meta, account_id, arn, attachment_count, create_date, default_version_id, description, is_attachable, path, permissions_boundary_usage_count, id, name, update_date) FROM stdin;
\.


--
-- Data for Name: aws_iam_policy_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_policy_versions (cq_id, cq_meta, policy_cq_id, policy_id, create_date, document, is_default_version, version_id) FROM stdin;
\.


--
-- Data for Name: aws_iam_role_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_role_policies (cq_id, cq_meta, role_cq_id, role_id, account_id, policy_document, policy_name, role_name) FROM stdin;
\.


--
-- Data for Name: aws_iam_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_roles (cq_id, cq_meta, account_id, policies, arn, create_date, path, id, name, assume_role_policy_document, description, max_session_duration, permissions_boundary_arn, permissions_boundary_type, role_last_used_last_used_date, role_last_used_region, tags) FROM stdin;
\.


--
-- Data for Name: aws_iam_saml_identity_providers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_saml_identity_providers (cq_id, cq_meta, account_id, arn, create_date, saml_metadata_document, tags, valid_until) FROM stdin;
\.


--
-- Data for Name: aws_iam_server_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_server_certificates (cq_id, cq_meta, account_id, id, arn, path, name, expiration, upload_date) FROM stdin;
\.


--
-- Data for Name: aws_iam_user_access_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_user_access_keys (cq_id, cq_meta, user_cq_id, user_id, access_key_id, create_date, status, last_used, last_rotated, last_used_service_name) FROM stdin;
\.


--
-- Data for Name: aws_iam_user_attached_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_user_attached_policies (cq_id, cq_meta, user_cq_id, user_id, policy_arn, policy_name) FROM stdin;
\.


--
-- Data for Name: aws_iam_user_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_user_groups (cq_id, cq_meta, user_cq_id, user_id, group_arn, create_date, group_id, group_name, path) FROM stdin;
\.


--
-- Data for Name: aws_iam_user_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_user_policies (cq_id, cq_meta, user_cq_id, account_id, user_id, policy_document, policy_name, user_name) FROM stdin;
\.


--
-- Data for Name: aws_iam_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_users (cq_id, cq_meta, account_id, id, password_last_used, arn, password_enabled, password_status, password_last_changed, password_next_rotation, mfa_active, create_date, path, permissions_boundary_arn, permissions_boundary_type, tags, user_id, user_name, access_key_1_active, access_key_1_last_rotated, access_key_2_active, access_key_2_last_rotated, cert_1_active, cert_1_last_rotated, cert_2_active, cert_2_last_rotated) FROM stdin;
\.


--
-- Data for Name: aws_iam_virtual_mfa_devices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iam_virtual_mfa_devices (cq_id, cq_meta, account_id, serial_number, base32_string_seed, enable_date, qr_code_png, tags, user_arn, user_create_date, user_path, user_id, user_name, user_password_last_used, user_permissions_boundary_permissions_boundary_arn, user_permissions_boundary_permissions_boundary_type, user_tags) FROM stdin;
\.


--
-- Data for Name: aws_iot_billing_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_billing_groups (cq_id, cq_meta, account_id, region, things_in_group, tags, arn, id, creation_date, name, description, version) FROM stdin;
\.


--
-- Data for Name: aws_iot_ca_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_ca_certificates (cq_id, cq_meta, account_id, region, certificates, auto_registration_status, arn, id, pem, creation_date, customer_version, generation_id, last_modified_date, owned_by, status, validity_not_after, validity_not_before) FROM stdin;
\.


--
-- Data for Name: aws_iot_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_certificates (cq_id, cq_meta, account_id, region, policies, ca_certificate_id, arn, id, mode, pem, creation_date, customer_version, generation_id, last_modified_date, owned_by, previous_owned_by, status, transfer_data_accept_date, transfer_data_reject_date, transfer_data_reject_reason, transfer_data_transfer_date, transfer_data_transfer_message, validity_not_after, validity_not_before) FROM stdin;
\.


--
-- Data for Name: aws_iot_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_policies (cq_id, cq_meta, account_id, region, tags, creation_date, default_version_id, generation_id, last_modified_date, arn, document, name) FROM stdin;
\.


--
-- Data for Name: aws_iot_stream_files; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_stream_files (cq_id, cq_meta, stream_cq_id, file_id, s3_location_bucket, s3_location_key, s3_location_version) FROM stdin;
\.


--
-- Data for Name: aws_iot_streams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_streams (cq_id, cq_meta, account_id, region, created_at, description, last_updated_at, role_arn, arn, id, version) FROM stdin;
\.


--
-- Data for Name: aws_iot_thing_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_thing_groups (cq_id, cq_meta, account_id, region, things_in_group, policies, tags, index_name, query_string, query_version, status, arn, id, creation_date, parent_group_name, root_to_parent_thing_groups, name, attribute_payload_attributes, attribute_payload_merge, thing_group_description, version) FROM stdin;
\.


--
-- Data for Name: aws_iot_thing_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_thing_types (cq_id, cq_meta, account_id, region, tags, arn, creation_date, deprecated, deprecation_date, name, searchable_attributes, description) FROM stdin;
\.


--
-- Data for Name: aws_iot_things; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_things (cq_id, cq_meta, account_id, region, principals, attributes, arn, name, type_name, version) FROM stdin;
\.


--
-- Data for Name: aws_iot_topic_rule_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_topic_rule_actions (cq_id, cq_meta, topic_rule_cq_id, cloudwatch_alarm_alarm_name, cloudwatch_alarm_role_arn, cloudwatch_alarm_state_reason, cloudwatch_alarm_state_value, cloudwatch_logs_log_group_name, cloudwatch_logs_role_arn, cloudwatch_metric_metric_name, cloudwatch_metric_metric_namespace, cloudwatch_metric_metric_unit, cloudwatch_metric_metric_value, cloudwatch_metric_role_arn, cloudwatch_metric_metric_timestamp, dynamo_db_hash_key_field, dynamo_db_hash_key_value, dynamo_db_role_arn, dynamo_db_table_name, dynamo_db_hash_key_type, dynamo_db_operation, dynamo_db_payload_field, dynamo_db_range_key_field, dynamo_db_range_key_type, dynamo_db_range_key_value, dynamo_db_v2_put_item_table_name, dynamo_db_v2_role_arn, elasticsearch_endpoint, elasticsearch_id, elasticsearch_index, elasticsearch_role_arn, elasticsearch_type, firehose_delivery_stream_name, firehose_role_arn, firehose_batch_mode, firehose_separator, http_url, http_auth_sigv4_role_arn, http_auth_sigv4_service_name, http_auth_sigv4_signing_region, http_confirmation_url, http_headers, iot_analytics_batch_mode, iot_analytics_channel_arn, iot_analytics_channel_name, iot_analytics_role_arn, iot_events_input_name, iot_events_role_arn, iot_events_batch_mode, iot_events_message_id, iot_site_wise, kafka_client_properties, kafka_destination_arn, kafka_topic, kafka_key, kafka_partition, kinesis_role_arn, kinesis_stream_name, kinesis_partition_key, lambda_function_arn, open_search_endpoint, open_search_id, open_search_index, open_search_role_arn, open_search_type, republish_role_arn, republish_topic, republish_qos, s3_bucket_name, s3_key, s3_role_arn, s3_canned_acl, salesforce_token, salesforce_url, sns_role_arn, sns_target_arn, sns_message_format, sqs_queue_url, sqs_role_arn, sqs_use_base64, step_functions_role_arn, step_functions_state_machine_name, step_functions_execution_name_prefix, timestream_database_name, timestream_dimensions, timestream_role_arn, timestream_table_name, timestream_timestamp_unit, timestream_timestamp_value) FROM stdin;
\.


--
-- Data for Name: aws_iot_topic_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_iot_topic_rules (cq_id, cq_meta, account_id, region, tags, aws_iot_sql_version, created_at, description, error_action_cloudwatch_alarm_name, error_action_cloudwatch_alarm_role_arn, error_action_cloudwatch_alarm_state_reason, error_action_cloudwatch_alarm_state_value, error_action_cloudwatch_logs_log_group_name, error_action_cloudwatch_logs_role_arn, error_action_cloudwatch_metric_metric_name, error_action_cloudwatch_metric_metric_namespace, error_action_cloudwatch_metric_unit, error_action_cloudwatch_metric_value, error_action_cloudwatch_metric_role_arn, error_action_cloudwatch_metric_timestamp, error_action_dynamo_db_hash_key_field, error_action_dynamo_db_hash_key_value, error_action_dynamo_db_role_arn, error_action_dynamo_db_table_name, error_action_dynamo_db_hash_key_type, error_action_dynamo_db_operation, error_action_dynamo_db_payload_field, error_action_dynamo_db_range_key_field, error_action_dynamo_db_range_key_type, error_action_dynamo_db_range_key_value, error_action_dynamo_db_v2_put_item_table_name, error_action_dynamo_db_v2_role_arn, error_action_elasticsearch_endpoint, error_action_elasticsearch_id, error_action_elasticsearch_index, error_action_elasticsearch_role_arn, error_action_elasticsearch_type, error_action_firehose_delivery_stream_name, error_action_firehose_role_arn, error_action_firehose_batch_mode, error_action_firehose_separator, error_action_http_url, error_action_http_auth_sigv4_role_arn, error_action_http_auth_sigv4_service_name, error_action_http_auth_sigv4_signing_region, error_action_http_confirmation_url, error_action_http_headers, error_action_iot_analytics_batch_mode, error_action_iot_analytics_channel_arn, error_action_iot_analytics_channel_name, error_action_iot_analytics_role_arn, error_action_iot_events_input_name, error_action_iot_events_role_arn, error_action_iot_events_batch_mode, error_action_iot_events_message_id, error_action_iot_site_wise, error_action_kafka_client_properties, error_action_kafka_destination_arn, error_action_kafka_topic, error_action_kafka_key, error_action_kafka_partition, error_action_kinesis_role_arn, error_action_kinesis_stream_name, error_action_kinesis_partition_key, error_action_lambda_function_arn, error_action_open_search_endpoint, error_action_open_search_id, error_action_open_search_index, error_action_open_search_role_arn, error_action_open_search_type, error_action_republish_role_arn, error_action_republish_topic, error_action_republish_qos, error_action_s3_bucket_name, error_action_s3_key, error_action_s3_role_arn, error_action_s3_canned_acl, error_action_salesforce_token, error_action_salesforce_url, error_action_sns_role_arn, error_action_sns_target_arn, error_action_sns_message_format, error_action_sqs_queue_url, error_action_sqs_role_arn, error_action_sqs_use_base64, error_action_step_functions_role_arn, error_action_step_functions_state_machine_name, error_action_step_functions_execution_name_prefix, error_action_timestream_database_name, error_action_timestream_dimensions, error_action_timestream_role_arn, error_action_timestream_table_name, error_action_timestream_timestamp_unit, error_action_timestream_timestamp_value, rule_disabled, rule_name, sql, arn) FROM stdin;
\.


--
-- Data for Name: aws_kms_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_kms_keys (cq_id, cq_meta, account_id, region, rotation_enabled, tags, id, aws_account_id, arn, cloud_hsm_cluster_id, creation_date, custom_key_store_id, deletion_date, description, enabled, encryption_algorithms, expiration_model, manager, key_spec, key_state, key_usage, mac_algorithms, multi_region, multi_region_key_type, primary_key_arn, primary_key_region, replica_keys, origin, pending_deletion_window_in_days, signing_algorithms, valid_to) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_aliases; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_aliases (cq_id, cq_meta, function_cq_id, function_arn, arn, description, function_version, name, revision_id, routing_config_additional_version_weights, url_config_auth_type, url_config_creation_time, url_config_function_arn, url_config_function_url, url_config_last_modified_time, url_config_cors) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_concurrency_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_concurrency_configs (cq_id, cq_meta, function_cq_id, allocated_provisioned_concurrent_executions, available_provisioned_concurrent_executions, function_arn, last_modified, requested_provisioned_concurrent_executions, status, status_reason) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_event_invoke_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_event_invoke_configs (cq_id, cq_meta, function_cq_id, on_failure_destination, on_success_destination, function_arn, last_modified, maximum_event_age_in_seconds, maximum_retry_attempts) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_event_source_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_event_source_mappings (cq_id, cq_meta, function_cq_id, batch_size, bisect_batch_on_function_error, on_failure_destination, on_success_destination, event_source_arn, criteria_filters, function_arn, function_response_types, last_modified, last_processing_result, maximum_batching_window_in_seconds, maximum_record_age_in_seconds, maximum_retry_attempts, parallelization_factor, queues, self_managed_event_source_endpoints, source_access_configurations, starting_position, starting_position_timestamp, state, state_transition_reason, topics, tumbling_window_in_seconds, uuid) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_file_system_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_file_system_configs (cq_id, cq_meta, function_cq_id, function_arn, arn, local_mount_path) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_layers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_layers (cq_id, cq_meta, function_cq_id, function_arn, arn, code_size, signing_job_arn, signing_profile_version_arn) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_version_file_system_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_version_file_system_configs (cq_id, cq_meta, function_version_cq_id, arn, local_mount_path) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_version_layers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_version_layers (cq_id, cq_meta, function_version_cq_id, arn, code_size, signing_job_arn, signing_profile_version_arn) FROM stdin;
\.


--
-- Data for Name: aws_lambda_function_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_function_versions (cq_id, cq_meta, function_cq_id, architectures, code_sha256, code_size, dead_letter_config_target_arn, description, environment_error_error_code, environment_error_message, environment_variables, ephemeral_storage_size, function_arn, function_name, handler, error_code, error_message, image_config_command, image_config_entry_point, image_config_working_directory, kms_key_arn, last_modified, last_update_status, last_update_status_reason, last_update_status_reason_code, master_arn, memory_size, package_type, revision_id, role, runtime, signing_job_arn, signing_profile_version_arn, state, state_reason, state_reason_code, timeout, tracing_config_mode, version, vpc_config_security_group_ids, vpc_config_subnet_ids, vpc_config_vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_lambda_functions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_functions (cq_id, cq_meta, account_id, region, policy_document, policy_revision_id, code_signing_allowed_publishers_version_arns, code_signing_config_arn, code_signing_config_id, code_signing_policies_untrusted_artifact_on_deployment, code_signing_description, code_signing_last_modified, code_image_uri, code_location, code_repository_type, code_resolved_image_uri, concurrency_reserved_concurrent_executions, architectures, code_sha256, code_size, dead_letter_config_target_arn, description, environment_error_code, environment_error_message, environment_variables, ephemeral_storage_size, arn, name, handler, error_code, error_message, image_config_command, image_config_entry_point, image_config_working_directory, kms_key_arn, last_modified, last_update_status, last_update_status_reason, last_update_status_reason_code, master_arn, memory_size, package_type, revision_id, role, runtime, signing_job_arn, signing_profile_version_arn, state, state_reason, state_reason_code, timeout, tracing_config_mode, version, vpc_config_security_group_ids, vpc_config_subnet_ids, vpc_config_vpc_id, tags) FROM stdin;
\.


--
-- Data for Name: aws_lambda_layer_version_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_layer_version_policies (cq_id, cq_meta, layer_version_cq_id, layer_version, policy, revision_id) FROM stdin;
\.


--
-- Data for Name: aws_lambda_layer_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_layer_versions (cq_id, cq_meta, layer_cq_id, compatible_runtimes, created_date, description, layer_version_arn, license_info, version) FROM stdin;
\.


--
-- Data for Name: aws_lambda_layers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_layers (cq_id, cq_meta, account_id, region, latest_matching_version_compatible_runtimes, latest_matching_version_created_date, latest_matching_version_description, latest_matching_version_layer_version_arn, latest_matching_version_license_info, latest_matching_version, arn, name) FROM stdin;
\.


--
-- Data for Name: aws_lambda_runtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_lambda_runtimes (cq_id, cq_meta, name) FROM stdin;
\.


--
-- Data for Name: aws_mq_broker_configuration_revisions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_mq_broker_configuration_revisions (cq_id, cq_meta, broker_configuration_cq_id, configuration_id, created, data, description) FROM stdin;
\.


--
-- Data for Name: aws_mq_broker_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_mq_broker_configurations (cq_id, cq_meta, broker_cq_id, account_id, region, arn, authentication_strategy, created, description, engine_type, engine_version, id, latest_revision_created, latest_revision, latest_revision_description, name, tags) FROM stdin;
\.


--
-- Data for Name: aws_mq_broker_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_mq_broker_users (cq_id, cq_meta, broker_cq_id, account_id, region, console_access, groups, pending, username) FROM stdin;
\.


--
-- Data for Name: aws_mq_brokers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_mq_brokers (cq_id, cq_meta, account_id, region, authentication_strategy, auto_minor_version_upgrade, arn, id, broker_instances, broker_name, broker_state, created, deployment_mode, encryption_options_use_aws_owned_key, encryption_options_kms_key_id, engine_type, engine_version, host_instance_type, ldap_server_metadata, logs, maintenance_window_start_time, pending_authentication_strategy, pending_engine_version, pending_host_instance_type, pending_ldap_server_metadata, pending_security_groups, publicly_accessible, security_groups, storage_type, subnet_ids, tags) FROM stdin;
\.


--
-- Data for Name: aws_organizations_accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_organizations_accounts (cq_id, cq_meta, account_id, tags, arn, email, id, joined_method, joined_timestamp, name, status) FROM stdin;
\.


--
-- Data for Name: aws_qldb_ledger_journal_kinesis_streams; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_qldb_ledger_journal_kinesis_streams (cq_id, cq_meta, ledger_cq_id, stream_arn, aggregation_enabled, ledger_name, role_arn, status, stream_id, stream_name, arn, creation_time, error_cause, exclusive_end_time, inclusive_start_time) FROM stdin;
\.


--
-- Data for Name: aws_qldb_ledger_journal_s3_exports; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_qldb_ledger_journal_s3_exports (cq_id, cq_meta, ledger_cq_id, exclusive_end_time, export_creation_time, export_id, inclusive_start_time, ledger_name, role_arn, bucket, object_encryption_type, kms_key_arn, prefix, status, output_format) FROM stdin;
\.


--
-- Data for Name: aws_qldb_ledgers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_qldb_ledgers (cq_id, cq_meta, account_id, region, tags, arn, creation_date_time, deletion_protection, encryption_status, kms_key_arn, inaccessible_kms_key_date_time, name, permissions_mode, state) FROM stdin;
\.


--
-- Data for Name: aws_rds_certificates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_certificates (cq_id, cq_meta, account_id, region, arn, certificate_identifier, certificate_type, customer_override, customer_override_valid_till, thumbprint, valid_from, valid_till) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_associated_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_associated_roles (cq_id, cq_meta, cluster_cq_id, feature_name, role_arn, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_db_cluster_members; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_db_cluster_members (cq_id, cq_meta, cluster_cq_id, db_cluster_parameter_group_status, db_instance_identifier, is_cluster_writer, promotion_tier) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_domain_memberships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_domain_memberships (cq_id, cq_meta, cluster_cq_id, domain, fqdn, iam_role_name, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_parameter_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_parameter_groups (cq_id, cq_meta, account_id, region, arn, name, family, description, tags) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_parameters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_parameters (cq_id, cq_meta, cluster_parameter_group_cq_id, allowed_values, apply_method, apply_type, data_type, description, is_modifiable, minimum_engine_version, parameter_name, parameter_value, source, supported_engine_modes) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_snapshots; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_snapshots (cq_id, cq_meta, account_id, region, allocated_storage, availability_zones, cluster_create_time, db_cluster_identifier, arn, db_cluster_snapshot_identifier, engine, engine_mode, engine_version, iam_database_authentication_enabled, kms_key_id, license_model, master_username, percent_progress, port, snapshot_create_time, snapshot_type, source_db_cluster_snapshot_arn, status, storage_encrypted, vpc_id, tags, attributes) FROM stdin;
\.


--
-- Data for Name: aws_rds_cluster_vpc_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_cluster_vpc_security_groups (cq_id, cq_meta, cluster_cq_id, status, vpc_security_group_id) FROM stdin;
\.


--
-- Data for Name: aws_rds_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_clusters (cq_id, cq_meta, account_id, region, activity_stream_kinesis_stream_name, activity_stream_kms_key_id, activity_stream_mode, activity_stream_status, allocated_storage, availability_zones, backtrack_consumed_change_records, backtrack_window, backup_retention_period, capacity, character_set_name, clone_group_id, cluster_create_time, copy_tags_to_snapshot, cross_account_clone, custom_endpoints, arn, db_cluster_identifier, db_cluster_parameter_group, db_cluster_option_group_memberships, db_subnet_group, database_name, id, deletion_protection, earliest_backtrack_time, earliest_restorable_time, enabled_cloudwatch_logs_exports, endpoint, engine, engine_mode, engine_version, global_write_forwarding_requested, global_write_forwarding_status, hosted_zone_id, http_endpoint_enabled, iam_database_authentication_enabled, kms_key_id, latest_restorable_time, master_username, multi_az, pending_modified_values_db_cluster_identifier, pending_modified_values_engine_version, pending_modified_values_iam_database_authentication_enabled, pending_modified_values_master_user_password, pending_cloudwatch_logs_types_to_disable, pending_cloudwatch_logs_types_to_enable, percent_progress, port, preferred_backup_window, preferred_maintenance_window, read_replica_identifiers, reader_endpoint, replication_source_identifier, scaling_configuration_info_auto_pause, scaling_configuration_info_max_capacity, scaling_configuration_info_min_capacity, scaling_configuration_info_seconds_until_auto_pause, scaling_configuration_info_timeout_action, status, storage_encrypted, tags) FROM stdin;
\.


--
-- Data for Name: aws_rds_db_parameter_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_db_parameter_groups (cq_id, cq_meta, account_id, region, arn, family, name, description, tags) FROM stdin;
\.


--
-- Data for Name: aws_rds_db_parameters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_db_parameters (cq_id, cq_meta, db_parameter_group_cq_id, allowed_values, apply_method, apply_type, data_type, description, is_modifiable, minimum_engine_version, parameter_name, parameter_value, source, supported_engine_modes) FROM stdin;
\.


--
-- Data for Name: aws_rds_db_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_db_security_groups (cq_id, cq_meta, account_id, region, arn, description, name, ec2_security_groups, ip_ranges, owner_id, vpc_id, tags) FROM stdin;
\.


--
-- Data for Name: aws_rds_db_snapshots; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_db_snapshots (cq_id, cq_meta, account_id, region, allocated_storage, availability_zone, db_instance_identifier, arn, db_snapshot_identifier, dbi_resource_id, encrypted, engine, engine_version, iam_database_authentication_enabled, instance_create_time, iops, kms_key_id, license_model, master_username, option_group_name, percent_progress, port, processor_features, snapshot_create_time, snapshot_type, source_db_snapshot_identifier, source_region, status, storage_type, tde_credential_arn, timezone, vpc_id, tags, attributes) FROM stdin;
\.


--
-- Data for Name: aws_rds_event_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_event_subscriptions (cq_id, cq_meta, account_id, region, cust_subscription_id, customer_aws_id, enabled, event_categories_list, arn, sns_topic_arn, source_ids_list, source_type, status, subscription_creation_time, tags) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_associated_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_associated_roles (cq_id, cq_meta, instance_cq_id, instance_id, feature_name, role_arn, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_db_instance_automated_backups_replications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_db_instance_automated_backups_replications (cq_id, cq_meta, instance_cq_id, instance_id, db_instance_automated_backups_arn) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_db_parameter_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_db_parameter_groups (cq_id, cq_meta, instance_cq_id, instance_id, db_parameter_group_name, parameter_apply_status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_db_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_db_security_groups (cq_id, cq_meta, instance_cq_id, instance_id, db_security_group_name, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_db_subnet_group_subnets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_db_subnet_group_subnets (cq_id, cq_meta, instance_cq_id, instance_id, subnet_availability_zone_name, subnet_identifier, subnet_outpost_arn, subnet_status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_domain_memberships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_domain_memberships (cq_id, cq_meta, instance_cq_id, instance_id, domain, fqdn, iam_role_name, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_option_group_memberships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_option_group_memberships (cq_id, cq_meta, instance_cq_id, option_group_name, status) FROM stdin;
\.


--
-- Data for Name: aws_rds_instance_vpc_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instance_vpc_security_groups (cq_id, cq_meta, instance_cq_id, instance_id, status, vpc_security_group_id) FROM stdin;
\.


--
-- Data for Name: aws_rds_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_instances (cq_id, cq_meta, account_id, region, allocated_storage, auto_minor_version_upgrade, availability_zone, aws_backup_recovery_point_arn, backup_retention_period, ca_certificate_identifier, character_set_name, copy_tags_to_snapshot, customer_owned_ip_enabled, cluster_identifier, arn, db_instance_class, user_instance_id, db_instance_status, db_name, subnet_group_arn, subnet_group_description, subnet_group_name, subnet_group_subnet_group_status, subnet_group_vpc_id, instance_port, id, deletion_protection, enabled_cloudwatch_logs_exports, endpoint_address, endpoint_hosted_zone_id, endpoint_port, engine, engine_version, enhanced_monitoring_resource_arn, iam_database_authentication_enabled, instance_create_time, iops, kms_key_id, latest_restorable_time, license_model, listener_endpoint_address, listener_endpoint_hosted_zone_id, listener_endpoint_port, master_username, max_allocated_storage, monitoring_interval, monitoring_role_arn, multi_az, nchar_character_set_name, pending_modified_values_allocated_storage, pending_modified_values_backup_retention_period, pending_modified_values_ca_certificate_identifier, pending_modified_values_db_instance_class, pending_modified_values_db_instance_identifier, pending_modified_values_db_subnet_group_name, pending_modified_values_engine_version, pending_modified_values_iam_database_authentication_enabled, pending_modified_values_iops, pending_modified_values_license_model, pending_modified_values_master_user_password, pending_modified_values_multi_az, pending_cloudwatch_logs_types_to_disable, pending_cloudwatch_logs_types_to_enable, pending_modified_values_port, pending_modified_values_processor_features, pending_modified_values_storage_type, performance_insights_enabled, performance_insights_kms_key_id, performance_insights_retention_period, preferred_backup_window, preferred_maintenance_window, processor_features, promotion_tier, publicly_accessible, read_replica_db_cluster_identifiers, read_replica_db_instance_identifiers, read_replica_source_db_instance_identifier, replica_mode, secondary_availability_zone, storage_encrypted, storage_type, tags, tde_credential_arn, timezone, status_infos) FROM stdin;
\.


--
-- Data for Name: aws_rds_subnet_group_subnets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_subnet_group_subnets (cq_id, cq_meta, subnet_group_cq_id, subnet_availability_zone_name, subnet_identifier, subnet_outpost_arn, subnet_status) FROM stdin;
\.


--
-- Data for Name: aws_rds_subnet_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_rds_subnet_groups (cq_id, cq_meta, account_id, region, arn, description, name, status, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_deferred_maintenance_windows; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_deferred_maintenance_windows (cq_id, cq_meta, cluster_cq_id, defer_maintenance_end_time, defer_maintenance_identifier, defer_maintenance_start_time) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces (cq_id, cq_meta, cluster_endpoint_vpc_endpoint_cq_id, availability_zone, network_interface_id, private_ip_address, subnet_id) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_endpoint_vpc_endpoints; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_endpoint_vpc_endpoints (cq_id, cq_meta, cluster_cq_id, vpc_endpoint_id, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_iam_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_iam_roles (cq_id, cq_meta, cluster_cq_id, apply_status, iam_role_arn) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_nodes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_nodes (cq_id, cq_meta, cluster_cq_id, node_role, private_ip_address, public_ip_address) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_parameter_group_status_lists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_parameter_group_status_lists (cq_id, cq_meta, cluster_parameter_group_cq_id, parameter_apply_error_description, parameter_apply_status, parameter_name) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_parameter_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_parameter_groups (cq_id, cq_meta, cluster_cq_id, parameter_apply_status, parameter_group_name) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_parameters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_parameters (cq_id, cq_meta, cluster_parameter_group_cq_id, allowed_values, apply_type, data_type, description, is_modifiable, minimum_engine_version, parameter_name, parameter_value, source) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_security_groups (cq_id, cq_meta, cluster_cq_id, cluster_security_group_name, status) FROM stdin;
\.


--
-- Data for Name: aws_redshift_cluster_vpc_security_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_cluster_vpc_security_groups (cq_id, cq_meta, cluster_cq_id, status, vpc_security_group_id) FROM stdin;
\.


--
-- Data for Name: aws_redshift_clusters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_clusters (cq_id, cq_meta, account_id, region, arn, allow_version_upgrade, automated_snapshot_retention_period, availability_zone, availability_zone_relocation_status, cluster_availability_status, cluster_create_time, id, cluster_namespace_arn, cluster_public_key, cluster_revision_number, cluster_snapshot_copy_status_destination_region, cluster_snapshot_copy_status_manual_snapshot_retention_period, cluster_snapshot_copy_status_retention_period, cluster_snapshot_copy_status_snapshot_copy_grant_name, cluster_status, cluster_subnet_group_name, cluster_version, db_name, data_transfer_progress_current_rate_in_mega_bytes_per_second, data_transfer_progress_data_transferred_in_mega_bytes, data_transfer_progress_elapsed_time_in_seconds, data_transfer_progress_estimated_time_to_completion_in_seconds, data_transfer_progress_status, data_transfer_progress_total_data_in_mega_bytes, elastic_ip_status_elastic_ip, elastic_ip_status, elastic_resize_number_of_node_options, encrypted, endpoint_address, endpoint_port, enhanced_vpc_routing, expected_next_snapshot_schedule_time, expected_next_snapshot_schedule_time_status, hsm_status_hsm_client_certificate_identifier, hsm_status_hsm_configuration_identifier, hsm_status, kms_key_id, maintenance_track_name, manual_snapshot_retention_period, master_username, modify_status, next_maintenance_window_start_time, node_type, number_of_nodes, pending_actions, pending_modified_values_automated_snapshot_retention_period, pending_modified_values_cluster_identifier, pending_modified_values_cluster_type, pending_modified_values_cluster_version, pending_modified_values_encryption_type, pending_modified_values_enhanced_vpc_routing, pending_modified_values_maintenance_track_name, pending_modified_values_master_user_password, pending_modified_values_node_type, pending_modified_values_number_of_nodes, pending_modified_values_publicly_accessible, preferred_maintenance_window, publicly_accessible, resize_info_allow_cancel_resize, resize_info_resize_type, restore_status_current_restore_rate_in_mega_bytes_per_second, restore_status_elapsed_time_in_seconds, restore_status_estimated_time_to_completion_in_seconds, restore_status_progress_in_mega_bytes, restore_status_snapshot_size_in_mega_bytes, restore_status, snapshot_schedule_identifier, snapshot_schedule_state, tags, total_storage_capacity_in_mega_bytes, vpc_id, logging_status) FROM stdin;
\.


--
-- Data for Name: aws_redshift_event_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_event_subscriptions (cq_id, cq_meta, arn, account_id, region, id, customer_aws_id, enabled, event_categories_list, severity, sns_topic_arn, source_ids_list, source_type, status, subscription_creation_time, tags) FROM stdin;
\.


--
-- Data for Name: aws_redshift_snapshot_accounts_with_restore_access; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_snapshot_accounts_with_restore_access (cq_id, cq_meta, snapshot_cq_id, account_alias, account_id) FROM stdin;
\.


--
-- Data for Name: aws_redshift_snapshots; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_snapshots (cq_id, cq_meta, cluster_cq_id, arn, actual_incremental_backup_size, availability_zone, backup_progress, cluster_create_time, cluster_identifier, cluster_version, current_backup_rate, db_name, elapsed_time, encrypted, encrypted_with_hsm, engine_full_version, enhanced_vpc_routing, estimated_seconds_to_completion, kms_key_id, maintenance_track_name, manual_snapshot_remaining_days, manual_snapshot_retention_period, master_username, node_type, number_of_nodes, owner_account, port, restorable_node_types, snapshot_create_time, snapshot_identifier, snapshot_retention_start_time, snapshot_type, source_region, status, total_backup_size_in_mega_bytes, vpc_id, tags) FROM stdin;
\.


--
-- Data for Name: aws_redshift_subnet_group_subnets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_subnet_group_subnets (cq_id, cq_meta, subnet_group_cq_id, subnet_availability_zone_name, subnet_availability_zone_supported_platforms, subnet_identifier, subnet_status) FROM stdin;
\.


--
-- Data for Name: aws_redshift_subnet_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_redshift_subnet_groups (cq_id, cq_meta, account_id, region, arn, cluster_subnet_group_name, description, subnet_group_status, tags, vpc_id) FROM stdin;
\.


--
-- Data for Name: aws_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_regions (cq_id, cq_meta, account_id, enabled, endpoint, opt_in_status, region, partition) FROM stdin;
\.


--
-- Data for Name: aws_route53_domain_nameservers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_domain_nameservers (cq_id, cq_meta, domain_cq_id, name, glue_ips) FROM stdin;
\.


--
-- Data for Name: aws_route53_domains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_domains (cq_id, cq_meta, account_id, admin_contact_address_line1, admin_contact_address_line2, admin_contact_city, admin_contact_type, admin_contact_country_code, admin_contact_email, admin_contact_fax, admin_contact_first_name, admin_contact_last_name, admin_contact_organization_name, admin_contact_phone_number, admin_contact_state, admin_contact_zip_code, admin_contact_extra_params, domain_name, registrant_contact_address_line1, registrant_contact_address_line2, registrant_contact_city, registrant_contact_type, registrant_contact_country_code, registrant_contact_email, registrant_contact_fax, registrant_contact_first_name, registrant_contact_last_name, registrant_contact_organization_name, registrant_contact_phone_number, registrant_contact_state, registrant_contact_zip_code, registrant_contact_extra_params, tech_contact_address_line1, tech_contact_address_line2, tech_contact_city, tech_contact_type, tech_contact_country_code, tech_contact_email, tech_contact_fax, tech_contact_first_name, tech_contact_last_name, tech_contact_organization_name, tech_contact_phone_number, tech_contact_state, tech_contact_zip_code, tech_contact_extra_params, abuse_contact_email, abuse_contact_phone, admin_privacy, auto_renew, creation_date, dns_sec, expiration_date, registrant_privacy, registrar_name, registrar_url, registry_domain_id, reseller, status_list, tech_privacy, updated_date, who_is_server, tags) FROM stdin;
\.


--
-- Data for Name: aws_route53_health_checks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_health_checks (cq_id, cq_meta, account_id, cloud_watch_alarm_configuration_dimensions, tags, caller_reference, type, alarm_identifier_name, alarm_identifier_region, child_health_checks, disabled, enable_sni, failure_threshold, fully_qualified_domain_name, health_threshold, ip_address, insufficient_data_health_status, inverted, measure_latency, port, regions, request_interval, resource_path, search_string, health_check_version, id, cloud_watch_alarm_config_comparison_operator, cloud_watch_alarm_config_evaluation_periods, cloud_watch_alarm_config_metric_name, cloud_watch_alarm_config_namespace, cloud_watch_alarm_config_period, cloud_watch_alarm_config_statistic, cloud_watch_alarm_config_threshold, linked_service_description, linked_service_service_principal, arn) FROM stdin;
\.


--
-- Data for Name: aws_route53_hosted_zone_query_logging_configs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_hosted_zone_query_logging_configs (cq_id, cq_meta, hosted_zone_cq_id, cloud_watch_logs_log_group_arn, id, arn) FROM stdin;
\.


--
-- Data for Name: aws_route53_hosted_zone_resource_record_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_hosted_zone_resource_record_sets (cq_id, cq_meta, hosted_zone_cq_id, resource_records, name, type, dns_name, evaluate_target_health, failover, geo_location_continent_code, geo_location_country_code, geo_location_subdivision_code, health_check_id, multi_value_answer, region, set_identifier, ttl, traffic_policy_instance_id, weight) FROM stdin;
\.


--
-- Data for Name: aws_route53_hosted_zone_traffic_policy_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_hosted_zone_traffic_policy_instances (cq_id, cq_meta, hosted_zone_cq_id, id, message, name, state, ttl, traffic_policy_id, traffic_policy_type, traffic_policy_version, arn) FROM stdin;
\.


--
-- Data for Name: aws_route53_hosted_zone_vpc_association_authorizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_hosted_zone_vpc_association_authorizations (cq_id, cq_meta, hosted_zone_cq_id, vpc_id, vpc_region, vpc_arn) FROM stdin;
\.


--
-- Data for Name: aws_route53_hosted_zones; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_hosted_zones (cq_id, cq_meta, account_id, tags, arn, delegation_set_id, caller_reference, id, name, config_comment, config_private_zone, linked_service_description, linked_service_principal, resource_record_set_count) FROM stdin;
\.


--
-- Data for Name: aws_route53_reusable_delegation_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_reusable_delegation_sets (cq_id, cq_meta, account_id, arn, name_servers, caller_reference, id) FROM stdin;
\.


--
-- Data for Name: aws_route53_traffic_policies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_traffic_policies (cq_id, cq_meta, account_id, id, latest_version, name, traffic_policy_count, type, arn) FROM stdin;
\.


--
-- Data for Name: aws_route53_traffic_policy_versions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_route53_traffic_policy_versions (cq_id, cq_meta, traffic_policy_cq_id, document, id, name, type, version, comment) FROM stdin;
\.


--
-- Data for Name: aws_s3_account_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_account_config (cq_id, cq_meta, account_id, config_exists, block_public_acls, block_public_policy, ignore_public_acls, restrict_public_buckets) FROM stdin;
59dd4900-c27e-5e41-841a-c629fde27d5b	{"last_updated": "2022-06-06T11:55:36.567892Z"}	123456123456	f	f	f	f	f
\.


--
-- Data for Name: aws_s3_bucket_cors_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_bucket_cors_rules (cq_id, cq_meta, bucket_cq_id, allowed_methods, allowed_origins, allowed_headers, expose_headers, id, max_age_seconds) FROM stdin;
\.


--
-- Data for Name: aws_s3_bucket_encryption_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_bucket_encryption_rules (cq_id, cq_meta, bucket_cq_id, sse_algorithm, kms_master_key_id, bucket_key_enabled) FROM stdin;
\.


--
-- Data for Name: aws_s3_bucket_grants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_bucket_grants (cq_id, cq_meta, bucket_cq_id, type, display_name, email_address, grantee_id, uri, permission) FROM stdin;
8a927cab-48c8-4f66-8f74-eb8252736cc7	{"last_updated": "2022-06-06T11:55:40.545998Z"}	ada7fe55-19f4-5deb-8c47-1a47a7464dce	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
4db63ef5-e731-4e7e-8d9e-3400f52df54e	{"last_updated": "2022-06-06T11:55:41.741402Z"}	8976c778-204d-54b0-9800-05216d728e34	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
e29129de-6795-44aa-b4eb-9d4b5f9bc109	{"last_updated": "2022-06-06T11:55:42.56965Z"}	e82b6da0-cbed-5906-9af4-527f918181fc	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
8f89536f-da4d-4c3b-8039-1d3ccae2fff7	{"last_updated": "2022-06-06T11:55:43.354431Z"}	73ae8aff-4487-5eab-b961-c04bde3b62aa	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
9ea2a94b-4746-47ef-8797-ab464fc6959f	{"last_updated": "2022-06-06T11:55:44.167305Z"}	d4c47aff-7e38-5a51-a17f-a40574365ad2	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
00a99694-9fc2-4560-85ab-9a9236829278	{"last_updated": "2022-06-06T11:55:45.027239Z"}	8e7a9a85-0f48-5a1c-ba3a-bf952d16870f	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
1e248de4-49eb-439c-bed7-b29a5a2a47d8	{"last_updated": "2022-06-06T11:55:45.842778Z"}	faf13aa0-eb1e-5ac0-af6f-d82981d44d60	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
e1c6bc41-b36c-40ef-92b9-75f980ee9b36	{"last_updated": "2022-06-06T11:55:48.490434Z"}	8368daae-6227-5ce6-8a78-cb7e2ac2bd63	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
df1049aa-0cb0-40c3-a75f-6a69136f22f3	{"last_updated": "2022-06-06T11:55:49.994383Z"}	067c0296-7df5-54bd-ac4f-b9002b457541	CanonicalUser	aws+cq-provider-aws	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
45f0e82f-d546-4132-8f64-1eaea0ea1d63	{"last_updated": "2022-06-06T11:55:51.88897Z"}	81a96b32-cd78-57f9-af43-5b50a2bd9478	CanonicalUser	\N	\N	be25ea540d744189366202a4b4fa23ef4c976afed1b094b472e4a2308f971016	\N	FULL_CONTROL
\.


--
-- Data for Name: aws_s3_bucket_lifecycles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_bucket_lifecycles (cq_id, cq_meta, bucket_cq_id, status, abort_incomplete_multipart_upload_days_after_initiation, expiration_date, expiration_days, expiration_expired_object_delete_marker, filter, id, noncurrent_version_expiration_days, noncurrent_version_transitions, prefix, transitions) FROM stdin;
\.


--
-- Data for Name: aws_s3_bucket_replication_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_bucket_replication_rules (cq_id, cq_meta, bucket_cq_id, destination_bucket, destination_access_control_translation_owner, destination_account, destination_encryption_configuration_replica_kms_key_id, destination_metrics_status, destination_metrics_event_threshold_minutes, destination_replication_time_status, destination_replication_time_minutes, destination_storage_class, status, delete_marker_replication_status, existing_object_replication_status, filter, id, prefix, priority, source_replica_modifications_status, source_sse_kms_encrypted_objects_status) FROM stdin;
\.


--
-- Data for Name: aws_s3_buckets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_s3_buckets (cq_id, cq_meta, account_id, region, logging_target_prefix, logging_target_bucket, versioning_status, versioning_mfa_delete, policy, tags, creation_date, name, block_public_acls, block_public_policy, ignore_public_acls, restrict_public_buckets, replication_role, arn, ownership_controls) FROM stdin;
ada7fe55-19f4-5deb-8c47-1a47a7464dce	{"last_updated": "2022-06-06T11:55:40.270801Z"}	123456123456	us-east-1	\N	\N			\N	\N	2022-04-24 07:51:57	cq-test-bucket-1	t	t	t	t	\N	arn:aws:s3:::cq-test-bucket-1	\N
8976c778-204d-54b0-9800-05216d728e34	{"last_updated": "2022-06-06T11:55:41.485358Z"}	123456123456	us-east-1	\N	\N	Enabled		\N	\N	2022-03-10 18:48:53	cq-test-bucket-2	t	t	t	t	\N	arn:aws:s3:::cq-test-bucket-2	{BucketOwnerEnforced}
e82b6da0-cbed-5906-9af4-527f918181fc	{"last_updated": "2022-06-06T11:55:42.303096Z"}	123456123456	us-east-1	\N	\N			\N	\N	2022-03-11 12:49:53	cq-test-bucket-3	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-3	\N
73ae8aff-4487-5eab-b961-c04bde3b62aa	{"last_updated": "2022-06-06T11:55:43.148075Z"}	123456123456	us-east-1	\N	\N			\N	{"Name": "cq-test-bucket-4", "Environment": "cq-provider-aws"}	2022-05-03 12:05:41	cq-test-bucket-4	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-4	\N
d4c47aff-7e38-5a51-a17f-a40574365ad2	{"last_updated": "2022-06-06T11:55:43.95286Z"}	123456123456	us-east-1	\N	\N			\N	\N	2022-03-11 16:20:59	cq-test-bucket-5	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-5	\N
8e7a9a85-0f48-5a1c-ba3a-bf952d16870f	{"last_updated": "2022-06-06T11:55:44.760743Z"}	123456123456	us-east-1	\N	\N			{"Id": "AWSLogDeliveryWrite20150319", "Version": "2012-10-17", "Statement": [{"Sid": "AWSLogDeliveryWrite", "Action": "s3:PutObject", "Effect": "Allow", "Resource": "arn:aws:s3:::cq-test-bucket-6/AWSLogs/123456123456/*", "Condition": {"ArnLike": {"aws:SourceArn": "arn:aws:logs:us-east-1:123456123456:*"}, "StringEquals": {"s3:x-amz-acl": "bucket-owner-full-control", "aws:SourceAccount": "123456123456"}}, "Principal": {"Service": "delivery.logs.amazonaws.com"}}, {"Sid": "AWSLogDeliveryAclCheck", "Action": "s3:GetBucketAcl", "Effect": "Allow", "Resource": "arn:aws:s3:::cq-test-bucket-6", "Condition": {"ArnLike": {"aws:SourceArn": "arn:aws:logs:us-east-1:123456123456:*"}, "StringEquals": {"aws:SourceAccount": "123456123456"}}, "Principal": {"Service": "delivery.logs.amazonaws.com"}}]}	\N	2022-05-03 12:06:58	cq-test-bucket-6	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-6	\N
faf13aa0-eb1e-5ac0-af6f-d82981d44d60	{"last_updated": "2022-06-06T11:55:45.595111Z"}	123456123456	us-east-1	\N	\N			\N	\N	2022-05-04 16:56:51	cq-test-bucket-7	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-7	\N
8368daae-6227-5ce6-8a78-cb7e2ac2bd63	{"last_updated": "2022-06-06T11:55:46.551328Z"}	123456123456	us-east-1	\N	\N			{"Version": "2008-10-17", "Statement": [{"Sid": "eb-af163bf3-d27b-4712-b795-d1e33e331ca4", "Action": ["s3:ListBucket", "s3:ListBucketVersions", "s3:GetObject", "s3:GetObjectVersion"], "Effect": "Allow", "Resource": ["arn:aws:s3:::cq-test-bucket-8", "arn:aws:s3:::cq-test-bucket-8/resources/environments/*"], "Principal": {"AWS": "arn:aws:iam::123456123456:role/cq-elastic-beanstalk-role"}}, {"Sid": "eb-58950a8c-feb6-11e2-89e0-0800277d041b", "Action": "s3:DeleteBucket", "Effect": "Deny", "Resource": "arn:aws:s3:::cq-test-bucket-8", "Principal": {"AWS": "*"}}]}	\N	2022-04-24 07:52:07	cq-test-bucket-8	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-8	\N
067c0296-7df5-54bd-ac4f-b9002b457541	{"last_updated": "2022-06-06T11:55:49.129881Z"}	123456123456	us-east-1	\N	\N			\N	{"Name": "cq-test-bucket-9", "Environment": "cq-provider-aws"}	2022-04-11 16:15:41	cq-test-bucket-9	f	f	f	f	\N	arn:aws:s3:::cq-test-bucket-9	\N
81a96b32-cd78-57f9-af43-5b50a2bd9478	{"last_updated": "2022-06-06T11:55:50.559798Z"}	123456123456	eu-central-1	\N	\N	Enabled		\N	\N	2021-12-27 11:30:44	cq-test-bucket-10	t	t	t	t	\N	arn:aws:s3:::cq-test-bucket-10	{BucketOwnerEnforced}
\.


--
-- Data for Name: aws_sagemaker_endpoint_configuration_production_variants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_endpoint_configuration_production_variants (cq_id, cq_meta, endpoint_configuration_cq_id, initial_instance_count, instance_type, model_name, variant_name, accelerator_type, core_dump_config_destination_s3_uri, core_dump_config_kms_key_id, initial_variant_weight) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_endpoint_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_endpoint_configurations (cq_id, cq_meta, account_id, region, kms_key_id, data_capture_config, tags, creation_time, arn, name) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_model_containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_model_containers (cq_id, cq_meta, model_cq_id, container_hostname, environment, image, image_config_repository_access_mode, image_config_repository_auth_config_repo_creds_provider_arn, mode, model_data_url, model_package_name, multi_model_config_model_cache_setting) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_model_vpc_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_model_vpc_config (cq_id, cq_meta, model_cq_id, security_group_ids, subnets) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_models; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_models (cq_id, cq_meta, account_id, region, enable_network_isolation, execution_role_arn, inference_execution_config, primary_container, tags, creation_time, arn, name) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_notebook_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_notebook_instances (cq_id, cq_meta, account_id, region, network_interface_id, kms_key_id, subnet_id, volume_size_in_gb, accelerator_types, security_groups, direct_internet_access, tags, arn, name, additional_code_repositories, creation_time, default_code_repository, instance_type, last_modified_time, notebook_instance_lifecycle_config_name, notebook_instance_status, url) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_algorithm_specification; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_algorithm_specification (cq_id, cq_meta, training_job_cq_id, training_input_mode, algorithm_name, enable_sage_maker_metrics_time_series, metric_definitions, training_image) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_debug_hook_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_debug_hook_config (cq_id, cq_meta, training_job_cq_id, s3_output_path, collection_configurations, hook_parameters, local_path) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_debug_rule_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_debug_rule_configurations (cq_id, cq_meta, training_job_cq_id, rule_configuration_name, rule_evaluator_image, instance_type, local_path, rule_parameters, s3_output_path, volume_size_in_gb) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_debug_rule_evaluation_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_debug_rule_evaluation_statuses (cq_id, cq_meta, training_job_cq_id, last_modified_time, rule_configuration_name, rule_evaluation_job_arn, rule_evaluation_status, status_details) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_input_data_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_input_data_config (cq_id, cq_meta, training_job_cq_id, channel_name, data_source_file_directory_path, data_source_file_system_access_mode, data_source_file_system_id, data_source_file_system_type, data_source_s3_data_type, data_source_s3_uri, data_source_attribute_names, data_source_s3_data_distribution_type, compression_type, content_type, input_mode, record_wrapper_type, shuffle_config_seed) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_profiler_rule_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_profiler_rule_configurations (cq_id, cq_meta, training_job_cq_id, rule_configuration_name, rule_evaluator_image, instance_type, local_path, rule_parameters, s3_output_path, volume_size_in_gb) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_job_profiler_rule_evaluation_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_job_profiler_rule_evaluation_statuses (cq_id, cq_meta, training_job_cq_id, last_modified_time, rule_configuration_name, rule_evaluation_job_arn, rule_evaluation_status, status_details) FROM stdin;
\.


--
-- Data for Name: aws_sagemaker_training_jobs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sagemaker_training_jobs (cq_id, cq_meta, account_id, region, auto_ml_job_arn, billable_time_in_seconds, enable_managed_spot_training, enable_network_isolation, enable_inter_container_traffic_encryption, failure_reason, labeling_job_arn, last_modified_time, profiling_status, role_arn, secondary_status, training_end_time, training_start_time, training_time_in_seconds, tuning_job_arn, checkpoint_config, environment, experiment_config, hyper_parameters, model_artifacts, output_data_config, profiler_config, resource_config, stopping_condition, tensor_board_output_config, vpc_config, tags, creation_time, arn, name, training_job_status, secondary_status_transitions, final_metric_data_list) FROM stdin;
\.


--
-- Data for Name: aws_secretsmanager_secrets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_secretsmanager_secrets (cq_id, cq_meta, account_id, region, policy, replication_status, arn, created_date, deleted_date, description, kms_key_id, last_accessed_date, last_changed_date, last_rotated_date, name, owning_service, primary_region, rotation_enabled, rotation_lambda_arn, rotation_rules_automatically_after_days, secret_versions_to_stages, tags) FROM stdin;
\.


--
-- Data for Name: aws_shield_attack_properties; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_attack_properties (cq_id, cq_meta, attack_cq_id, attack_layer, attack_property_identifier, top_contributors, total, unit) FROM stdin;
\.


--
-- Data for Name: aws_shield_attack_sub_resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_attack_sub_resources (cq_id, cq_meta, attack_cq_id, attack_vectors, counters, id, type) FROM stdin;
\.


--
-- Data for Name: aws_shield_attacks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_attacks (cq_id, cq_meta, account_id, attack_counters, id, end_time, mitigations, resource_arn, start_time) FROM stdin;
\.


--
-- Data for Name: aws_shield_protection_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_protection_groups (cq_id, cq_meta, account_id, tags, aggregation, members, pattern, id, arn, resource_type) FROM stdin;
\.


--
-- Data for Name: aws_shield_protections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_protections (cq_id, cq_meta, account_id, region, tags, application_automatic_response_configuration_status, health_check_ids, id, name, arn, resource_arn) FROM stdin;
\.


--
-- Data for Name: aws_shield_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_shield_subscriptions (cq_id, cq_meta, account_id, protection_group_limits_max_protection_groups, protection_group_limits_arbitrary_pattern_limits_max_members, protected_resource_type_limits, auto_renew, end_time, limits, proactive_engagement_status, start_time, arn, time_commitment_in_seconds) FROM stdin;
\.


--
-- Data for Name: aws_sns_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sns_subscriptions (cq_id, cq_meta, account_id, region, endpoint, owner, protocol, arn, topic_arn) FROM stdin;
\.


--
-- Data for Name: aws_sns_topics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sns_topics (cq_id, cq_meta, account_id, region, owner, policy, delivery_policy, display_name, subscriptions_confirmed, subscriptions_deleted, subscriptions_pending, effective_delivery_policy, fifo_topic, content_based_deduplication, kms_master_key_id, arn, tags) FROM stdin;
\.


--
-- Data for Name: aws_sqs_queues; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_sqs_queues (cq_id, cq_meta, account_id, region, url, policy, visibility_timeout, maximum_message_size, message_retention_period, approximate_number_of_messages, approximate_number_of_messages_not_visible, created_timestamp, last_modified_timestamp, arn, approximate_number_of_messages_delayed, delay_seconds, receive_message_wait_time_seconds, redrive_policy, fifo_queue, content_based_deduplication, kms_master_key_id, kms_data_key_reuse_period_seconds, deduplication_scope, fifo_throughput_limit, redrive_allow_policy, tags, unknown_fields) FROM stdin;
\.


--
-- Data for Name: aws_ssm_documents; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ssm_documents (cq_id, cq_meta, account_id, region, arn, approved_version, attachments_information, author, created_date, default_version, description, display_name, document_format, document_type, document_version, hash, hash_type, latest_version, name, owner, parameters, pending_review_version, platform_types, requires, review_status, schema_version, sha1, status, status_information, target_type, version_name, review_information, tags, account_ids, account_sharing_info_list) FROM stdin;
\.


--
-- Data for Name: aws_ssm_instance_compliance_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ssm_instance_compliance_items (cq_id, cq_meta, instance_cq_id, compliance_type, details, execution_summary_execution_time, execution_summary_execution_id, execution_summary_execution_type, id, resource_id, resource_type, severity, status, title) FROM stdin;
\.


--
-- Data for Name: aws_ssm_instances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_ssm_instances (cq_id, cq_meta, account_id, region, arn, activation_id, agent_version, association_overview_detailed_status, association_instance_status_aggregated_count, association_status, computer_name, ip_address, iam_role, instance_id, is_latest_version, last_association_execution_date, last_ping_date_time, last_successful_association_execution_date, name, ping_status, platform_name, platform_type, platform_version, registration_date, resource_type) FROM stdin;
\.


--
-- Data for Name: aws_waf_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_rule_groups (cq_id, cq_meta, account_id, arn, rule_ids, tags, id, metric_name, name) FROM stdin;
\.


--
-- Data for Name: aws_waf_rule_predicates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_rule_predicates (cq_id, cq_meta, rule_cq_id, data_id, negated, type) FROM stdin;
\.


--
-- Data for Name: aws_waf_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_rules (cq_id, cq_meta, account_id, arn, tags, id, metric_name, name) FROM stdin;
\.


--
-- Data for Name: aws_waf_subscribed_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_subscribed_rule_groups (cq_id, cq_meta, account_id, metric_name, name, rule_group_id) FROM stdin;
\.


--
-- Data for Name: aws_waf_web_acl_logging_configuration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_web_acl_logging_configuration (cq_id, cq_meta, web_acl_cq_id, log_destination_configs, resource_arn, redacted_fields) FROM stdin;
\.


--
-- Data for Name: aws_waf_web_acl_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_web_acl_rules (cq_id, cq_meta, web_acl_cq_id, priority, rule_id, action_type, excluded_rules, override_action_type, type) FROM stdin;
\.


--
-- Data for Name: aws_waf_web_acls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_waf_web_acls (cq_id, cq_meta, account_id, tags, default_action_type, id, metric_name, name, arn, logging_configuration) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_rate_based_rule_match_predicates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_rate_based_rule_match_predicates (cq_id, cq_meta, rate_based_rule_cq_id, data_id, negated, type) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_rate_based_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_rate_based_rules (cq_id, cq_meta, account_id, region, arn, tags, rate_key, rate_limit, id, metric_name, name) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_rule_groups (cq_id, cq_meta, account_id, region, arn, tags, id, metric_name, name) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_rule_predicates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_rule_predicates (cq_id, cq_meta, rule_cq_id, data_id, negated, type) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_rules (cq_id, cq_meta, account_id, region, arn, tags, id, metric_name, name) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_web_acl_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_web_acl_rules (cq_id, cq_meta, web_acl_cq_id, priority, rule_id, action, excluded_rules, override_action, type) FROM stdin;
\.


--
-- Data for Name: aws_wafregional_web_acls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafregional_web_acls (cq_id, cq_meta, account_id, region, tags, default_action, id, metric_name, name, arn) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_ipsets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_ipsets (cq_id, cq_meta, account_id, region, scope, arn, addresses, ip_address_version, id, name, description, tags) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_managed_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_managed_rule_groups (cq_id, cq_meta, account_id, region, scope, available_labels, consumed_labels, capacity, label_namespace, rules, description, name, vendor_name) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_regex_pattern_sets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_regex_pattern_sets (cq_id, cq_meta, account_id, region, scope, arn, description, id, name, regular_expression_list, tags) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_rule_groups (cq_id, cq_meta, account_id, region, scope, tags, policy, arn, capacity, id, name, visibility_config_cloud_watch_metrics_enabled, visibility_config_metric_name, visibility_config_sampled_requests_enabled, custom_response_bodies, description, label_namespace, rules, available_labels, consumed_labels) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_web_acl_logging_configuration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_web_acl_logging_configuration (cq_id, cq_meta, web_acl_cq_id, log_destination_configs, resource_arn, logging_filter, managed_by_firewall_manager, redacted_fields) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_web_acl_post_process_firewall_manager_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_web_acl_post_process_firewall_manager_rule_groups (cq_id, cq_meta, web_acl_cq_id, statement, name, override_action, priority, visibility_config_cloud_watch_metrics_enabled, visibility_config_metric_name, visibility_config_sampled_requests_enabled) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups (cq_id, cq_meta, web_acl_cq_id, statement, name, override_action, priority, visibility_config_cloud_watch_metrics_enabled, visibility_config_metric_name, visibility_config_sampled_requests_enabled) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_web_acl_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_web_acl_rules (cq_id, cq_meta, web_acl_cq_id, name, priority, statement, visibility_config_cloud_watch_metrics_enabled, visibility_config_metric_name, visibility_config_sampled_requests_enabled, action, override_action, labels) FROM stdin;
\.


--
-- Data for Name: aws_wafv2_web_acls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_wafv2_web_acls (cq_id, cq_meta, account_id, region, scope, resources_for_web_acl, tags, arn, default_action, id, name, visibility_config_cloud_watch_metrics_enabled, visibility_config_metric_name, visibility_config_sampled_requests_enabled, capacity, custom_response_bodies, description, label_namespace, managed_by_firewall_manager, logging_configuration) FROM stdin;
\.


--
-- Data for Name: aws_workspaces_directories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_workspaces_directories (cq_id, cq_meta, account_id, region, arn, alias, customer_user_name, id, name, type, dns_ip_addresses, iam_role_id, ip_group_ids, registration_code, change_compute_type, increase_volume_size, rebuild_workspace, restart_workspace, switch_running_mode, state, subnet_ids, tenancy, device_type_android, device_type_chrome_os, device_type_ios, device_type_linux, device_type_osx, device_type_web, device_type_windows, device_type_zero_client, custom_security_group_id, default_ou, enable_internet_access, enable_maintenance_mode, enable_work_docs, user_enabled_as_local_administrator, workspace_security_group_id) FROM stdin;
\.


--
-- Data for Name: aws_workspaces_workspaces; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_workspaces_workspaces (cq_id, cq_meta, account_id, region, arn, bundle_id, computer_name, directory_id, error_code, error_message, ip_address, modification_states, root_volume_encryption_enabled, state, subnet_id, user_name, user_volume_encryption_enabled, volume_encryption_key, id, compute_type_name, root_volume_size_gib, running_mode, running_mode_auto_stop_timeout_in_minutes, user_volume_size_gib) FROM stdin;
\.


--
-- Data for Name: aws_xray_encryption_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_xray_encryption_config (cq_id, cq_meta, account_id, region, key_id, status, type) FROM stdin;
\.


--
-- Data for Name: aws_xray_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_xray_groups (cq_id, cq_meta, account_id, region, tags, filter_expression, arn, group_name, insights_enabled, notifications_enabled) FROM stdin;
\.


--
-- Data for Name: aws_xray_sampling_rules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aws_xray_sampling_rules (cq_id, cq_meta, account_id, region, tags, created_at, modified_at, fixed_rate, http_method, host, priority, reservoir_size, resource_arn, service_name, service_type, url_path, version, attributes, arn, rule_name) FROM stdin;
\.


--
-- Name: check_results check_results_pkey; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.check_results
    ADD CONSTRAINT check_results_pkey PRIMARY KEY (execution_id, selector);


--
-- Name: cloudquery_core_schema_migrations cloudquery_core_schema_migrations_pkey; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.cloudquery_core_schema_migrations
    ADD CONSTRAINT cloudquery_core_schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: fetches fetches_id; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.fetches
    ADD CONSTRAINT fetches_id PRIMARY KEY (id);


--
-- Name: fetches fetches_pk; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.fetches
    ADD CONSTRAINT fetches_pk UNIQUE (fetch_id, provider_name, provider_alias);


--
-- Name: policy_executions policy_executions_pkey; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.policy_executions
    ADD CONSTRAINT policy_executions_pkey PRIMARY KEY (id);


--
-- Name: providers providers_id; Type: CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.providers
    ADD CONSTRAINT providers_id PRIMARY KEY (source, name);


--
-- Name: aws_access_analyzer_analyzer_archive_rules aws_access_analyzer_analyzer_archive_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_archive_rules
    ADD CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_access_analyzer_analyzer_finding_sources aws_access_analyzer_analyzer_finding_sources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_finding_sources
    ADD CONSTRAINT aws_access_analyzer_analyzer_finding_sources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_access_analyzer_analyzer_findings aws_access_analyzer_analyzer_findings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_findings
    ADD CONSTRAINT aws_access_analyzer_analyzer_findings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_access_analyzer_analyzers aws_access_analyzer_analyzers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzers
    ADD CONSTRAINT aws_access_analyzer_analyzers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_access_analyzer_analyzers aws_access_analyzer_analyzers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzers
    ADD CONSTRAINT aws_access_analyzer_analyzers_pk PRIMARY KEY (arn);


--
-- Name: aws_accounts aws_accounts_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_accounts
    ADD CONSTRAINT aws_accounts_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_accounts aws_accounts_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_accounts
    ADD CONSTRAINT aws_accounts_pk PRIMARY KEY (account_id);


--
-- Name: aws_acm_certificates aws_acm_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_acm_certificates
    ADD CONSTRAINT aws_acm_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_acm_certificates aws_acm_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_acm_certificates
    ADD CONSTRAINT aws_acm_certificates_pk PRIMARY KEY (arn);


--
-- Name: aws_apigateway_api_keys aws_apigateway_api_keys_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_api_keys
    ADD CONSTRAINT aws_apigateway_api_keys_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_api_keys aws_apigateway_api_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_api_keys
    ADD CONSTRAINT aws_apigateway_api_keys_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigateway_client_certificates aws_apigateway_client_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_client_certificates
    ADD CONSTRAINT aws_apigateway_client_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_client_certificates aws_apigateway_client_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_client_certificates
    ADD CONSTRAINT aws_apigateway_client_certificates_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigateway_domain_name_base_path_mappings aws_apigateway_domain_name_base_path_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_domain_name_base_path_mappings
    ADD CONSTRAINT aws_apigateway_domain_name_base_path_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_domain_names aws_apigateway_domain_names_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_domain_names
    ADD CONSTRAINT aws_apigateway_domain_names_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_domain_names aws_apigateway_domain_names_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_domain_names
    ADD CONSTRAINT aws_apigateway_domain_names_pk PRIMARY KEY (account_id, region, domain_name);


--
-- Name: aws_apigateway_rest_api_authorizers aws_apigateway_rest_api_authorizers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_authorizers
    ADD CONSTRAINT aws_apigateway_rest_api_authorizers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_deployments aws_apigateway_rest_api_deployments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_deployments
    ADD CONSTRAINT aws_apigateway_rest_api_deployments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_documentation_parts aws_apigateway_rest_api_documentation_parts_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_documentation_parts
    ADD CONSTRAINT aws_apigateway_rest_api_documentation_parts_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_documentation_versions aws_apigateway_rest_api_documentation_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_documentation_versions
    ADD CONSTRAINT aws_apigateway_rest_api_documentation_versions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_gateway_responses aws_apigateway_rest_api_gateway_responses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_gateway_responses
    ADD CONSTRAINT aws_apigateway_rest_api_gateway_responses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_models aws_apigateway_rest_api_models_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_models
    ADD CONSTRAINT aws_apigateway_rest_api_models_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_request_validators aws_apigateway_rest_api_request_validators_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_request_validators
    ADD CONSTRAINT aws_apigateway_rest_api_request_validators_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_resources aws_apigateway_rest_api_resources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_resources
    ADD CONSTRAINT aws_apigateway_rest_api_resources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_api_stages aws_apigateway_rest_api_stages_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_stages
    ADD CONSTRAINT aws_apigateway_rest_api_stages_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_rest_apis aws_apigateway_rest_apis_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_apis
    ADD CONSTRAINT aws_apigateway_rest_apis_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_rest_apis aws_apigateway_rest_apis_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_apis
    ADD CONSTRAINT aws_apigateway_rest_apis_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigateway_usage_plan_api_stages aws_apigateway_usage_plan_api_stages_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plan_api_stages
    ADD CONSTRAINT aws_apigateway_usage_plan_api_stages_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_usage_plan_keys aws_apigateway_usage_plan_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plan_keys
    ADD CONSTRAINT aws_apigateway_usage_plan_keys_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigateway_usage_plans aws_apigateway_usage_plans_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plans
    ADD CONSTRAINT aws_apigateway_usage_plans_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_usage_plans aws_apigateway_usage_plans_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plans
    ADD CONSTRAINT aws_apigateway_usage_plans_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigateway_vpc_links aws_apigateway_vpc_links_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_vpc_links
    ADD CONSTRAINT aws_apigateway_vpc_links_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigateway_vpc_links aws_apigateway_vpc_links_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_vpc_links
    ADD CONSTRAINT aws_apigateway_vpc_links_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigatewayv2_api_authorizers aws_apigatewayv2_api_authorizers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_authorizers
    ADD CONSTRAINT aws_apigatewayv2_api_authorizers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_deployments aws_apigatewayv2_api_deployments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_deployments
    ADD CONSTRAINT aws_apigatewayv2_api_deployments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_integration_responses aws_apigatewayv2_api_integration_responses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_integration_responses
    ADD CONSTRAINT aws_apigatewayv2_api_integration_responses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_integrations aws_apigatewayv2_api_integrations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_integrations
    ADD CONSTRAINT aws_apigatewayv2_api_integrations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_models aws_apigatewayv2_api_models_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_models
    ADD CONSTRAINT aws_apigatewayv2_api_models_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_route_responses aws_apigatewayv2_api_route_responses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_route_responses
    ADD CONSTRAINT aws_apigatewayv2_api_route_responses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_routes aws_apigatewayv2_api_routes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_routes
    ADD CONSTRAINT aws_apigatewayv2_api_routes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_api_stages aws_apigatewayv2_api_stages_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_stages
    ADD CONSTRAINT aws_apigatewayv2_api_stages_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_apis aws_apigatewayv2_apis_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_apis
    ADD CONSTRAINT aws_apigatewayv2_apis_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigatewayv2_apis aws_apigatewayv2_apis_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_apis
    ADD CONSTRAINT aws_apigatewayv2_apis_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_apigatewayv2_domain_name_configurations aws_apigatewayv2_domain_name_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_name_configurations
    ADD CONSTRAINT aws_apigatewayv2_domain_name_configurations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_domain_name_rest_api_mappings aws_apigatewayv2_domain_name_rest_api_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_name_rest_api_mappings
    ADD CONSTRAINT aws_apigatewayv2_domain_name_rest_api_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_apigatewayv2_domain_names aws_apigatewayv2_domain_names_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_names
    ADD CONSTRAINT aws_apigatewayv2_domain_names_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigatewayv2_domain_names aws_apigatewayv2_domain_names_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_names
    ADD CONSTRAINT aws_apigatewayv2_domain_names_pk PRIMARY KEY (account_id, region, domain_name);


--
-- Name: aws_apigatewayv2_vpc_links aws_apigatewayv2_vpc_links_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_vpc_links
    ADD CONSTRAINT aws_apigatewayv2_vpc_links_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_apigatewayv2_vpc_links aws_apigatewayv2_vpc_links_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_vpc_links
    ADD CONSTRAINT aws_apigatewayv2_vpc_links_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_applicationautoscaling_policies aws_applicationautoscaling_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_applicationautoscaling_policies
    ADD CONSTRAINT aws_applicationautoscaling_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_applicationautoscaling_policies aws_applicationautoscaling_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_applicationautoscaling_policies
    ADD CONSTRAINT aws_applicationautoscaling_policies_pk PRIMARY KEY (arn);


--
-- Name: aws_athena_data_catalog_database_table_columns aws_athena_data_catalog_database_table_columns_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_table_columns
    ADD CONSTRAINT aws_athena_data_catalog_database_table_columns_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_data_catalog_database_table_partition_keys aws_athena_data_catalog_database_table_partition_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_table_partition_keys
    ADD CONSTRAINT aws_athena_data_catalog_database_table_partition_keys_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_data_catalog_database_tables aws_athena_data_catalog_database_tables_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_tables
    ADD CONSTRAINT aws_athena_data_catalog_database_tables_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_data_catalog_databases aws_athena_data_catalog_databases_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_databases
    ADD CONSTRAINT aws_athena_data_catalog_databases_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_data_catalogs aws_athena_data_catalogs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalogs
    ADD CONSTRAINT aws_athena_data_catalogs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_athena_data_catalogs aws_athena_data_catalogs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalogs
    ADD CONSTRAINT aws_athena_data_catalogs_pk PRIMARY KEY (arn);


--
-- Name: aws_athena_work_group_named_queries aws_athena_work_group_named_queries_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_named_queries
    ADD CONSTRAINT aws_athena_work_group_named_queries_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_work_group_prepared_statements aws_athena_work_group_prepared_statements_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_prepared_statements
    ADD CONSTRAINT aws_athena_work_group_prepared_statements_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_work_group_query_executions aws_athena_work_group_query_executions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_query_executions
    ADD CONSTRAINT aws_athena_work_group_query_executions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_athena_work_groups aws_athena_work_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_groups
    ADD CONSTRAINT aws_athena_work_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_athena_work_groups aws_athena_work_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_groups
    ADD CONSTRAINT aws_athena_work_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_autoscaling_group_instances aws_autoscaling_group_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_instances
    ADD CONSTRAINT aws_autoscaling_group_instances_pk PRIMARY KEY (cq_id);


--
-- Name: aws_autoscaling_group_lifecycle_hooks aws_autoscaling_group_lifecycle_hooks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_lifecycle_hooks
    ADD CONSTRAINT aws_autoscaling_group_lifecycle_hooks_pk PRIMARY KEY (cq_id);


--
-- Name: aws_autoscaling_group_scaling_policies aws_autoscaling_group_scaling_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_scaling_policies
    ADD CONSTRAINT aws_autoscaling_group_scaling_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_autoscaling_group_tags aws_autoscaling_group_tags_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_tags
    ADD CONSTRAINT aws_autoscaling_group_tags_pk PRIMARY KEY (cq_id);


--
-- Name: aws_autoscaling_groups aws_autoscaling_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_groups
    ADD CONSTRAINT aws_autoscaling_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_autoscaling_groups aws_autoscaling_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_groups
    ADD CONSTRAINT aws_autoscaling_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_autoscaling_launch_configuration_block_device_mappings aws_autoscaling_launch_configuration_block_device_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_launch_configuration_block_device_mappings
    ADD CONSTRAINT aws_autoscaling_launch_configuration_block_device_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_autoscaling_launch_configurations aws_autoscaling_launch_configurations_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_launch_configurations
    ADD CONSTRAINT aws_autoscaling_launch_configurations_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_autoscaling_launch_configurations aws_autoscaling_launch_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_launch_configurations
    ADD CONSTRAINT aws_autoscaling_launch_configurations_pk PRIMARY KEY (arn);


--
-- Name: aws_autoscaling_scheduled_actions aws_autoscaling_scheduled_actions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_scheduled_actions
    ADD CONSTRAINT aws_autoscaling_scheduled_actions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_autoscaling_scheduled_actions aws_autoscaling_scheduled_actions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_scheduled_actions
    ADD CONSTRAINT aws_autoscaling_scheduled_actions_pk PRIMARY KEY (arn);


--
-- Name: aws_backup_global_settings aws_backup_global_settings_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_global_settings
    ADD CONSTRAINT aws_backup_global_settings_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_backup_global_settings aws_backup_global_settings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_global_settings
    ADD CONSTRAINT aws_backup_global_settings_pk PRIMARY KEY (account_id);


--
-- Name: aws_backup_plan_rules aws_backup_plan_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plan_rules
    ADD CONSTRAINT aws_backup_plan_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_backup_plan_selections aws_backup_plan_selections_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plan_selections
    ADD CONSTRAINT aws_backup_plan_selections_pk PRIMARY KEY (cq_id);


--
-- Name: aws_backup_plans aws_backup_plans_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plans
    ADD CONSTRAINT aws_backup_plans_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_backup_plans aws_backup_plans_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plans
    ADD CONSTRAINT aws_backup_plans_pk PRIMARY KEY (arn);


--
-- Name: aws_backup_region_settings aws_backup_region_settings_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_region_settings
    ADD CONSTRAINT aws_backup_region_settings_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_backup_region_settings aws_backup_region_settings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_region_settings
    ADD CONSTRAINT aws_backup_region_settings_pk PRIMARY KEY (account_id, region);


--
-- Name: aws_backup_vault_recovery_points aws_backup_vault_recovery_points_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_vault_recovery_points
    ADD CONSTRAINT aws_backup_vault_recovery_points_pk PRIMARY KEY (cq_id);


--
-- Name: aws_backup_vaults aws_backup_vaults_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_vaults
    ADD CONSTRAINT aws_backup_vaults_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_backup_vaults aws_backup_vaults_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_vaults
    ADD CONSTRAINT aws_backup_vaults_pk PRIMARY KEY (arn);


--
-- Name: aws_cloudformation_stack_outputs aws_cloudformation_stack_outputs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stack_outputs
    ADD CONSTRAINT aws_cloudformation_stack_outputs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudformation_stack_resources aws_cloudformation_stack_resources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stack_resources
    ADD CONSTRAINT aws_cloudformation_stack_resources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudformation_stacks aws_cloudformation_stacks_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stacks
    ADD CONSTRAINT aws_cloudformation_stacks_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudformation_stacks aws_cloudformation_stacks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stacks
    ADD CONSTRAINT aws_cloudformation_stacks_pk PRIMARY KEY (id);


--
-- Name: aws_cloudfront_cache_policies aws_cloudfront_cache_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_cache_policies
    ADD CONSTRAINT aws_cloudfront_cache_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudfront_cache_policies aws_cloudfront_cache_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_cache_policies
    ADD CONSTRAINT aws_cloudfront_cache_policies_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_cloudfront_distribution_cache_behavior_lambda_functions aws_cloudfront_distribution_cache_behavior_lambda_functions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_cache_behavior_lambda_functions
    ADD CONSTRAINT aws_cloudfront_distribution_cache_behavior_lambda_functions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distribution_cache_behaviors aws_cloudfront_distribution_cache_behaviors_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_cache_behaviors
    ADD CONSTRAINT aws_cloudfront_distribution_cache_behaviors_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distribution_custom_error_responses aws_cloudfront_distribution_custom_error_responses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_custom_error_responses
    ADD CONSTRAINT aws_cloudfront_distribution_custom_error_responses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distribution_default_cache_behavior_functions aws_cloudfront_distribution_default_cache_behavior_functions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_default_cache_behavior_functions
    ADD CONSTRAINT aws_cloudfront_distribution_default_cache_behavior_functions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distribution_origin_groups aws_cloudfront_distribution_origin_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_origin_groups
    ADD CONSTRAINT aws_cloudfront_distribution_origin_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distribution_origins aws_cloudfront_distribution_origins_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_origins
    ADD CONSTRAINT aws_cloudfront_distribution_origins_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudfront_distributions aws_cloudfront_distributions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distributions
    ADD CONSTRAINT aws_cloudfront_distributions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudfront_distributions aws_cloudfront_distributions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distributions
    ADD CONSTRAINT aws_cloudfront_distributions_pk PRIMARY KEY (arn);


--
-- Name: aws_cloudtrail_trail_event_selectors aws_cloudtrail_trail_event_selectors_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudtrail_trail_event_selectors
    ADD CONSTRAINT aws_cloudtrail_trail_event_selectors_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudtrail_trails aws_cloudtrail_trails_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudtrail_trails
    ADD CONSTRAINT aws_cloudtrail_trails_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudtrail_trails aws_cloudtrail_trails_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudtrail_trails
    ADD CONSTRAINT aws_cloudtrail_trails_pk PRIMARY KEY (account_id, arn);


--
-- Name: aws_cloudwatch_alarm_metrics aws_cloudwatch_alarm_metrics_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatch_alarm_metrics
    ADD CONSTRAINT aws_cloudwatch_alarm_metrics_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudwatch_alarms aws_cloudwatch_alarms_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatch_alarms
    ADD CONSTRAINT aws_cloudwatch_alarms_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudwatch_alarms aws_cloudwatch_alarms_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatch_alarms
    ADD CONSTRAINT aws_cloudwatch_alarms_pk PRIMARY KEY (arn);


--
-- Name: aws_cloudwatchlogs_filter_metric_transformations aws_cloudwatchlogs_filter_metric_transformations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatchlogs_filter_metric_transformations
    ADD CONSTRAINT aws_cloudwatchlogs_filter_metric_transformations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cloudwatchlogs_filters aws_cloudwatchlogs_filters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatchlogs_filters
    ADD CONSTRAINT aws_cloudwatchlogs_filters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cloudwatchlogs_filters aws_cloudwatchlogs_filters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatchlogs_filters
    ADD CONSTRAINT aws_cloudwatchlogs_filters_pk PRIMARY KEY (account_id, region, name, log_group_name);


--
-- Name: aws_codebuild_project_environment_variables aws_codebuild_project_environment_variables_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_environment_variables
    ADD CONSTRAINT aws_codebuild_project_environment_variables_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codebuild_project_file_system_locations aws_codebuild_project_file_system_locations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_file_system_locations
    ADD CONSTRAINT aws_codebuild_project_file_system_locations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codebuild_project_secondary_artifacts aws_codebuild_project_secondary_artifacts_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_secondary_artifacts
    ADD CONSTRAINT aws_codebuild_project_secondary_artifacts_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codebuild_project_secondary_sources aws_codebuild_project_secondary_sources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_secondary_sources
    ADD CONSTRAINT aws_codebuild_project_secondary_sources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codebuild_projects aws_codebuild_projects_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_projects
    ADD CONSTRAINT aws_codebuild_projects_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_codebuild_projects aws_codebuild_projects_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_projects
    ADD CONSTRAINT aws_codebuild_projects_pk PRIMARY KEY (arn);


--
-- Name: aws_codepipeline_pipeline_stage_actions aws_codepipeline_pipeline_stage_actions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipeline_stage_actions
    ADD CONSTRAINT aws_codepipeline_pipeline_stage_actions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codepipeline_pipeline_stages aws_codepipeline_pipeline_stages_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipeline_stages
    ADD CONSTRAINT aws_codepipeline_pipeline_stages_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codepipeline_pipelines aws_codepipeline_pipelines_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipelines
    ADD CONSTRAINT aws_codepipeline_pipelines_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_codepipeline_pipelines aws_codepipeline_pipelines_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipelines
    ADD CONSTRAINT aws_codepipeline_pipelines_pk PRIMARY KEY (arn);


--
-- Name: aws_codepipeline_webhook_filters aws_codepipeline_webhook_filters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_webhook_filters
    ADD CONSTRAINT aws_codepipeline_webhook_filters_pk PRIMARY KEY (cq_id);


--
-- Name: aws_codepipeline_webhooks aws_codepipeline_webhooks_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_webhooks
    ADD CONSTRAINT aws_codepipeline_webhooks_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_codepipeline_webhooks aws_codepipeline_webhooks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_webhooks
    ADD CONSTRAINT aws_codepipeline_webhooks_pk PRIMARY KEY (arn);


--
-- Name: aws_cognito_identity_pool_cognito_identity_providers aws_cognito_identity_pool_cognito_identity_providers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_identity_pool_cognito_identity_providers
    ADD CONSTRAINT aws_cognito_identity_pool_cognito_identity_providers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cognito_identity_pools aws_cognito_identity_pools_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_identity_pools
    ADD CONSTRAINT aws_cognito_identity_pools_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cognito_identity_pools aws_cognito_identity_pools_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_identity_pools
    ADD CONSTRAINT aws_cognito_identity_pools_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_cognito_user_pool_identity_providers aws_cognito_user_pool_identity_providers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pool_identity_providers
    ADD CONSTRAINT aws_cognito_user_pool_identity_providers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cognito_user_pool_schema_attributes aws_cognito_user_pool_schema_attributes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pool_schema_attributes
    ADD CONSTRAINT aws_cognito_user_pool_schema_attributes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_cognito_user_pools aws_cognito_user_pools_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pools
    ADD CONSTRAINT aws_cognito_user_pools_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_cognito_user_pools aws_cognito_user_pools_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pools
    ADD CONSTRAINT aws_cognito_user_pools_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_config_configuration_recorders aws_config_configuration_recorders_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_configuration_recorders
    ADD CONSTRAINT aws_config_configuration_recorders_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_config_configuration_recorders aws_config_configuration_recorders_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_configuration_recorders
    ADD CONSTRAINT aws_config_configuration_recorders_pk PRIMARY KEY (arn);


--
-- Name: aws_config_conformance_pack_rule_compliances aws_config_conformance_pack_rule_compliances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_conformance_pack_rule_compliances
    ADD CONSTRAINT aws_config_conformance_pack_rule_compliances_pk PRIMARY KEY (cq_id);


--
-- Name: aws_config_conformance_packs aws_config_conformance_packs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_conformance_packs
    ADD CONSTRAINT aws_config_conformance_packs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_config_conformance_packs aws_config_conformance_packs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_conformance_packs
    ADD CONSTRAINT aws_config_conformance_packs_pk PRIMARY KEY (arn);


--
-- Name: aws_dax_cluster_nodes aws_dax_cluster_nodes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dax_cluster_nodes
    ADD CONSTRAINT aws_dax_cluster_nodes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dax_clusters aws_dax_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dax_clusters
    ADD CONSTRAINT aws_dax_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_dax_clusters aws_dax_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dax_clusters
    ADD CONSTRAINT aws_dax_clusters_pk PRIMARY KEY (arn);


--
-- Name: aws_directconnect_connection_mac_sec_keys aws_directconnect_connection_mac_sec_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_connection_mac_sec_keys
    ADD CONSTRAINT aws_directconnect_connection_mac_sec_keys_pk PRIMARY KEY (cq_id);


--
-- Name: aws_directconnect_connections aws_directconnect_connections_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_connections
    ADD CONSTRAINT aws_directconnect_connections_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_directconnect_connections aws_directconnect_connections_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_connections
    ADD CONSTRAINT aws_directconnect_connections_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_directconnect_gateway_associations aws_directconnect_gateway_associations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateway_associations
    ADD CONSTRAINT aws_directconnect_gateway_associations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_directconnect_gateway_attachments aws_directconnect_gateway_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateway_attachments
    ADD CONSTRAINT aws_directconnect_gateway_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_directconnect_gateways aws_directconnect_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateways
    ADD CONSTRAINT aws_directconnect_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_directconnect_gateways aws_directconnect_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateways
    ADD CONSTRAINT aws_directconnect_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_directconnect_lag_mac_sec_keys aws_directconnect_lag_mac_sec_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_lag_mac_sec_keys
    ADD CONSTRAINT aws_directconnect_lag_mac_sec_keys_pk PRIMARY KEY (cq_id);


--
-- Name: aws_directconnect_lags aws_directconnect_lags_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_lags
    ADD CONSTRAINT aws_directconnect_lags_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_directconnect_lags aws_directconnect_lags_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_lags
    ADD CONSTRAINT aws_directconnect_lags_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_directconnect_virtual_gateways aws_directconnect_virtual_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_gateways
    ADD CONSTRAINT aws_directconnect_virtual_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_directconnect_virtual_gateways aws_directconnect_virtual_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_gateways
    ADD CONSTRAINT aws_directconnect_virtual_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_directconnect_virtual_interface_bgp_peers aws_directconnect_virtual_interface_bgp_peers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_interface_bgp_peers
    ADD CONSTRAINT aws_directconnect_virtual_interface_bgp_peers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_directconnect_virtual_interfaces aws_directconnect_virtual_interfaces_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_interfaces
    ADD CONSTRAINT aws_directconnect_virtual_interfaces_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_directconnect_virtual_interfaces aws_directconnect_virtual_interfaces_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_interfaces
    ADD CONSTRAINT aws_directconnect_virtual_interfaces_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_dms_replication_instance_replication_subnet_group_subnets aws_dms_replication_instance_replication_subnet_group_subnet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instance_replication_subnet_group_subnets
    ADD CONSTRAINT aws_dms_replication_instance_replication_subnet_group_subnet_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dms_replication_instance_vpc_security_groups aws_dms_replication_instance_vpc_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instance_vpc_security_groups
    ADD CONSTRAINT aws_dms_replication_instance_vpc_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dms_replication_instances aws_dms_replication_instances_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instances
    ADD CONSTRAINT aws_dms_replication_instances_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_dms_replication_instances aws_dms_replication_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instances
    ADD CONSTRAINT aws_dms_replication_instances_pk PRIMARY KEY (account_id, arn);


--
-- Name: aws_dynamodb_table_continuous_backups aws_dynamodb_table_continuous_backups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_continuous_backups
    ADD CONSTRAINT aws_dynamodb_table_continuous_backups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dynamodb_table_global_secondary_indexes aws_dynamodb_table_global_secondary_indexes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_global_secondary_indexes
    ADD CONSTRAINT aws_dynamodb_table_global_secondary_indexes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dynamodb_table_local_secondary_indexes aws_dynamodb_table_local_secondary_indexes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_local_secondary_indexes
    ADD CONSTRAINT aws_dynamodb_table_local_secondary_indexes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dynamodb_table_replica_auto_scalings aws_dynamodb_table_replica_auto_scalings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_replica_auto_scalings
    ADD CONSTRAINT aws_dynamodb_table_replica_auto_scalings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dynamodb_table_replicas aws_dynamodb_table_replicas_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_replicas
    ADD CONSTRAINT aws_dynamodb_table_replicas_pk PRIMARY KEY (cq_id);


--
-- Name: aws_dynamodb_tables aws_dynamodb_tables_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_tables
    ADD CONSTRAINT aws_dynamodb_tables_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_dynamodb_tables aws_dynamodb_tables_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_tables
    ADD CONSTRAINT aws_dynamodb_tables_pk PRIMARY KEY (arn);


--
-- Name: aws_ec2_byoip_cidrs aws_ec2_byoip_cidrs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_byoip_cidrs
    ADD CONSTRAINT aws_ec2_byoip_cidrs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_byoip_cidrs aws_ec2_byoip_cidrs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_byoip_cidrs
    ADD CONSTRAINT aws_ec2_byoip_cidrs_pk PRIMARY KEY (account_id, region, cidr);


--
-- Name: aws_ec2_customer_gateways aws_ec2_customer_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_customer_gateways
    ADD CONSTRAINT aws_ec2_customer_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_customer_gateways aws_ec2_customer_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_customer_gateways
    ADD CONSTRAINT aws_ec2_customer_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_ebs_snapshots aws_ec2_ebs_snapshots_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_snapshots
    ADD CONSTRAINT aws_ec2_ebs_snapshots_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_ebs_snapshots aws_ec2_ebs_snapshots_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_snapshots
    ADD CONSTRAINT aws_ec2_ebs_snapshots_pk PRIMARY KEY (account_id, snapshot_id);


--
-- Name: aws_ec2_ebs_volume_attachments aws_ec2_ebs_volume_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_volume_attachments
    ADD CONSTRAINT aws_ec2_ebs_volume_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_ebs_volumes aws_ec2_ebs_volumes_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_volumes
    ADD CONSTRAINT aws_ec2_ebs_volumes_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_ebs_volumes aws_ec2_ebs_volumes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_volumes
    ADD CONSTRAINT aws_ec2_ebs_volumes_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_egress_only_internet_gateways aws_ec2_egress_only_internet_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_egress_only_internet_gateways
    ADD CONSTRAINT aws_ec2_egress_only_internet_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_egress_only_internet_gateways aws_ec2_egress_only_internet_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_egress_only_internet_gateways
    ADD CONSTRAINT aws_ec2_egress_only_internet_gateways_pk PRIMARY KEY (arn);


--
-- Name: aws_ec2_eips aws_ec2_eips_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_eips aws_ec2_eips_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (account_id, allocation_id);


--
-- Name: aws_ec2_flow_logs aws_ec2_flow_logs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_flow_logs
    ADD CONSTRAINT aws_ec2_flow_logs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_flow_logs aws_ec2_flow_logs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_flow_logs
    ADD CONSTRAINT aws_ec2_flow_logs_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_host_available_instance_capacity aws_ec2_host_available_instance_capacity_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_host_available_instance_capacity
    ADD CONSTRAINT aws_ec2_host_available_instance_capacity_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_host_instances aws_ec2_host_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_host_instances
    ADD CONSTRAINT aws_ec2_host_instances_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_hosts aws_ec2_hosts_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_hosts
    ADD CONSTRAINT aws_ec2_hosts_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_hosts aws_ec2_hosts_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_hosts
    ADD CONSTRAINT aws_ec2_hosts_pk PRIMARY KEY (arn);


--
-- Name: aws_ec2_image_block_device_mappings aws_ec2_image_block_device_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_image_block_device_mappings
    ADD CONSTRAINT aws_ec2_image_block_device_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_images aws_ec2_images_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_images
    ADD CONSTRAINT aws_ec2_images_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_images aws_ec2_images_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_images
    ADD CONSTRAINT aws_ec2_images_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_instance_block_device_mappings aws_ec2_instance_block_device_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_block_device_mappings
    ADD CONSTRAINT aws_ec2_instance_block_device_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_elastic_gpu_associations aws_ec2_instance_elastic_gpu_associations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_elastic_gpu_associations
    ADD CONSTRAINT aws_ec2_instance_elastic_gpu_associations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_elastic_inference_accelerator_associations aws_ec2_instance_elastic_inference_accelerator_associations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_elastic_inference_accelerator_associations
    ADD CONSTRAINT aws_ec2_instance_elastic_inference_accelerator_associations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_network_interface_groups aws_ec2_instance_network_interface_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_groups
    ADD CONSTRAINT aws_ec2_instance_network_interface_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_network_interface_ipv6_addresses aws_ec2_instance_network_interface_ipv6_addresses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_ipv6_addresses
    ADD CONSTRAINT aws_ec2_instance_network_interface_ipv6_addresses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_network_interface_private_ip_addresses aws_ec2_instance_network_interface_private_ip_addresses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_private_ip_addresses
    ADD CONSTRAINT aws_ec2_instance_network_interface_private_ip_addresses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_network_interfaces aws_ec2_instance_network_interfaces_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interfaces
    ADD CONSTRAINT aws_ec2_instance_network_interfaces_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_product_codes aws_ec2_instance_product_codes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_product_codes
    ADD CONSTRAINT aws_ec2_instance_product_codes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_security_groups aws_ec2_instance_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_security_groups
    ADD CONSTRAINT aws_ec2_instance_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_status_events aws_ec2_instance_status_events_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_status_events
    ADD CONSTRAINT aws_ec2_instance_status_events_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_instance_statuses aws_ec2_instance_statuses_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_statuses
    ADD CONSTRAINT aws_ec2_instance_statuses_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_instance_statuses aws_ec2_instance_statuses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_statuses
    ADD CONSTRAINT aws_ec2_instance_statuses_pk PRIMARY KEY (arn);


--
-- Name: aws_ec2_instances aws_ec2_instances_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instances
    ADD CONSTRAINT aws_ec2_instances_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_instances aws_ec2_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instances
    ADD CONSTRAINT aws_ec2_instances_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_internet_gateway_attachments aws_ec2_internet_gateway_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_internet_gateway_attachments
    ADD CONSTRAINT aws_ec2_internet_gateway_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_internet_gateways aws_ec2_internet_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_internet_gateways
    ADD CONSTRAINT aws_ec2_internet_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_internet_gateways aws_ec2_internet_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_internet_gateways
    ADD CONSTRAINT aws_ec2_internet_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_nat_gateway_addresses aws_ec2_nat_gateway_addresses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_nat_gateway_addresses
    ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_nat_gateways aws_ec2_nat_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_nat_gateways
    ADD CONSTRAINT aws_ec2_nat_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_nat_gateways aws_ec2_nat_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_nat_gateways
    ADD CONSTRAINT aws_ec2_nat_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_network_acl_associations aws_ec2_network_acl_associations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acl_associations
    ADD CONSTRAINT aws_ec2_network_acl_associations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_network_acl_entries aws_ec2_network_acl_entries_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acl_entries
    ADD CONSTRAINT aws_ec2_network_acl_entries_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_network_acls aws_ec2_network_acls_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acls
    ADD CONSTRAINT aws_ec2_network_acls_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_network_acls aws_ec2_network_acls_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acls
    ADD CONSTRAINT aws_ec2_network_acls_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_network_interface_private_ip_addresses aws_ec2_network_interface_private_ip_addresses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_interface_private_ip_addresses
    ADD CONSTRAINT aws_ec2_network_interface_private_ip_addresses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_network_interfaces aws_ec2_network_interfaces_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_interfaces
    ADD CONSTRAINT aws_ec2_network_interfaces_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_network_interfaces aws_ec2_network_interfaces_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_interfaces
    ADD CONSTRAINT aws_ec2_network_interfaces_pk PRIMARY KEY (arn);


--
-- Name: aws_ec2_regional_config aws_ec2_regional_config_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_regional_config
    ADD CONSTRAINT aws_ec2_regional_config_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_regional_config aws_ec2_regional_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_regional_config
    ADD CONSTRAINT aws_ec2_regional_config_pk PRIMARY KEY (account_id, region);


--
-- Name: aws_ec2_route_table_associations aws_ec2_route_table_associations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_associations
    ADD CONSTRAINT aws_ec2_route_table_associations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_route_table_propagating_vgws aws_ec2_route_table_propagating_vgws_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_propagating_vgws
    ADD CONSTRAINT aws_ec2_route_table_propagating_vgws_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_route_table_routes aws_ec2_route_table_routes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_routes
    ADD CONSTRAINT aws_ec2_route_table_routes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_route_tables aws_ec2_route_tables_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_tables
    ADD CONSTRAINT aws_ec2_route_tables_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_route_tables aws_ec2_route_tables_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_tables
    ADD CONSTRAINT aws_ec2_route_tables_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_security_group_ip_permission_ip_ranges aws_ec2_security_group_ip_permission_ip_ranges_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_ip_ranges
    ADD CONSTRAINT aws_ec2_security_group_ip_permission_ip_ranges_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_security_group_ip_permission_prefix_list_ids aws_ec2_security_group_ip_permission_prefix_list_ids_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_prefix_list_ids
    ADD CONSTRAINT aws_ec2_security_group_ip_permission_prefix_list_ids_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_security_group_ip_permission_user_id_group_pairs aws_ec2_security_group_ip_permission_user_id_group_pairs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_user_id_group_pairs
    ADD CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_security_group_ip_permissions aws_ec2_security_group_ip_permissions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permissions
    ADD CONSTRAINT aws_ec2_security_group_ip_permissions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_security_groups aws_ec2_security_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_groups
    ADD CONSTRAINT aws_ec2_security_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_security_groups aws_ec2_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_groups
    ADD CONSTRAINT aws_ec2_security_groups_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_subnet_ipv6_cidr_block_association_sets aws_ec2_subnet_ipv6_cidr_block_association_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_subnet_ipv6_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_subnet_ipv6_cidr_block_association_sets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_subnets aws_ec2_subnets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_subnets
    ADD CONSTRAINT aws_ec2_subnets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_subnets aws_ec2_subnets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_subnets
    ADD CONSTRAINT aws_ec2_subnets_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_transit_gateway_attachments aws_ec2_transit_gateway_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_transit_gateway_multicast_domains aws_ec2_transit_gateway_multicast_domains_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_multicast_domains
    ADD CONSTRAINT aws_ec2_transit_gateway_multicast_domains_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_transit_gateway_peering_attachments aws_ec2_transit_gateway_peering_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_peering_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_peering_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_transit_gateway_route_tables aws_ec2_transit_gateway_route_tables_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_route_tables
    ADD CONSTRAINT aws_ec2_transit_gateway_route_tables_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_transit_gateway_vpc_attachments aws_ec2_transit_gateway_vpc_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_vpc_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_vpc_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_transit_gateways aws_ec2_transit_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateways
    ADD CONSTRAINT aws_ec2_transit_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_transit_gateways aws_ec2_transit_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateways
    ADD CONSTRAINT aws_ec2_transit_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_vpc_attachment aws_ec2_vpc_attachment_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_attachment
    ADD CONSTRAINT aws_ec2_vpc_attachment_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_vpc_cidr_block_association_sets aws_ec2_vpc_cidr_block_association_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_vpc_cidr_block_association_sets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_vpc_endpoint_dns_entries aws_ec2_vpc_endpoint_dns_entries_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoint_dns_entries
    ADD CONSTRAINT aws_ec2_vpc_endpoint_dns_entries_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_vpc_endpoint_groups aws_ec2_vpc_endpoint_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoint_groups
    ADD CONSTRAINT aws_ec2_vpc_endpoint_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_vpc_endpoints aws_ec2_vpc_endpoints_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoints
    ADD CONSTRAINT aws_ec2_vpc_endpoints_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_vpc_endpoints aws_ec2_vpc_endpoints_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoints
    ADD CONSTRAINT aws_ec2_vpc_endpoints_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_vpc_ipv6_cidr_block_association_sets aws_ec2_vpc_ipv6_cidr_block_association_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_ipv6_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_vpc_ipv6_cidr_block_association_sets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ec2_vpc_peering_connections aws_ec2_vpc_peering_connections_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_peering_connections
    ADD CONSTRAINT aws_ec2_vpc_peering_connections_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_vpc_peering_connections aws_ec2_vpc_peering_connections_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_peering_connections
    ADD CONSTRAINT aws_ec2_vpc_peering_connections_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_ec2_vpcs aws_ec2_vpcs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpcs
    ADD CONSTRAINT aws_ec2_vpcs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_vpcs aws_ec2_vpcs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpcs
    ADD CONSTRAINT aws_ec2_vpcs_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ec2_vpn_gateways aws_ec2_vpn_gateways_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpn_gateways
    ADD CONSTRAINT aws_ec2_vpn_gateways_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ec2_vpn_gateways aws_ec2_vpn_gateways_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpn_gateways
    ADD CONSTRAINT aws_ec2_vpn_gateways_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_ecr_repositories aws_ecr_repositories_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecr_repositories
    ADD CONSTRAINT aws_ecr_repositories_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ecr_repositories aws_ecr_repositories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecr_repositories
    ADD CONSTRAINT aws_ecr_repositories_pk PRIMARY KEY (account_id, arn);


--
-- Name: aws_ecr_repository_images aws_ecr_repository_images_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecr_repository_images
    ADD CONSTRAINT aws_ecr_repository_images_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_attachments aws_ecs_cluster_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_attachments
    ADD CONSTRAINT aws_ecs_cluster_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instance_attachments aws_ecs_cluster_container_instance_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_attachments
    ADD CONSTRAINT aws_ecs_cluster_container_instance_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instance_attributes aws_ecs_cluster_container_instance_attributes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_attributes
    ADD CONSTRAINT aws_ecs_cluster_container_instance_attributes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instance_health_status_details aws_ecs_cluster_container_instance_health_status_details_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_health_status_details
    ADD CONSTRAINT aws_ecs_cluster_container_instance_health_status_details_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instance_registered_resources aws_ecs_cluster_container_instance_registered_resources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_registered_resources
    ADD CONSTRAINT aws_ecs_cluster_container_instance_registered_resources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instance_remaining_resources aws_ecs_cluster_container_instance_remaining_resources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_remaining_resources
    ADD CONSTRAINT aws_ecs_cluster_container_instance_remaining_resources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_container_instances aws_ecs_cluster_container_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instances
    ADD CONSTRAINT aws_ecs_cluster_container_instances_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_deployments aws_ecs_cluster_service_deployments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_deployments
    ADD CONSTRAINT aws_ecs_cluster_service_deployments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_events aws_ecs_cluster_service_events_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_events
    ADD CONSTRAINT aws_ecs_cluster_service_events_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_load_balancers aws_ecs_cluster_service_load_balancers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_load_balancers
    ADD CONSTRAINT aws_ecs_cluster_service_load_balancers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_service_registries aws_ecs_cluster_service_service_registries_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_service_registries
    ADD CONSTRAINT aws_ecs_cluster_service_service_registries_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_task_set_load_balancers aws_ecs_cluster_service_task_set_load_balancers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_set_load_balancers
    ADD CONSTRAINT aws_ecs_cluster_service_task_set_load_balancers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_task_set_service_registries aws_ecs_cluster_service_task_set_service_registries_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_set_service_registries
    ADD CONSTRAINT aws_ecs_cluster_service_task_set_service_registries_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_service_task_sets aws_ecs_cluster_service_task_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_sets
    ADD CONSTRAINT aws_ecs_cluster_service_task_sets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_services aws_ecs_cluster_services_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_services
    ADD CONSTRAINT aws_ecs_cluster_services_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_task_attachments aws_ecs_cluster_task_attachments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_task_attachments
    ADD CONSTRAINT aws_ecs_cluster_task_attachments_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_task_containers aws_ecs_cluster_task_containers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_task_containers
    ADD CONSTRAINT aws_ecs_cluster_task_containers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_cluster_tasks aws_ecs_cluster_tasks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_tasks
    ADD CONSTRAINT aws_ecs_cluster_tasks_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_clusters aws_ecs_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_clusters
    ADD CONSTRAINT aws_ecs_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ecs_clusters aws_ecs_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_clusters
    ADD CONSTRAINT aws_ecs_clusters_pk PRIMARY KEY (arn);


--
-- Name: aws_ecs_task_definition_container_definitions aws_ecs_task_definition_container_definitions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definition_container_definitions
    ADD CONSTRAINT aws_ecs_task_definition_container_definitions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_task_definition_volumes aws_ecs_task_definition_volumes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definition_volumes
    ADD CONSTRAINT aws_ecs_task_definition_volumes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ecs_task_definitions aws_ecs_task_definitions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definitions
    ADD CONSTRAINT aws_ecs_task_definitions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ecs_task_definitions aws_ecs_task_definitions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definitions
    ADD CONSTRAINT aws_ecs_task_definitions_pk PRIMARY KEY (arn);


--
-- Name: aws_efs_filesystems aws_efs_filesystems_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_efs_filesystems
    ADD CONSTRAINT aws_efs_filesystems_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_efs_filesystems aws_efs_filesystems_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_efs_filesystems
    ADD CONSTRAINT aws_efs_filesystems_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_eks_cluster_encryption_configs aws_eks_cluster_encryption_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_cluster_encryption_configs
    ADD CONSTRAINT aws_eks_cluster_encryption_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_eks_cluster_loggings aws_eks_cluster_loggings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_cluster_loggings
    ADD CONSTRAINT aws_eks_cluster_loggings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_eks_clusters aws_eks_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_clusters
    ADD CONSTRAINT aws_eks_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_eks_clusters aws_eks_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_clusters
    ADD CONSTRAINT aws_eks_clusters_pk PRIMARY KEY (arn);


--
-- Name: aws_elasticbeanstalk_application_versions aws_elasticbeanstalk_application_versions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_application_versions
    ADD CONSTRAINT aws_elasticbeanstalk_application_versions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elasticbeanstalk_application_versions aws_elasticbeanstalk_application_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_application_versions
    ADD CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY (arn);


--
-- Name: aws_elasticbeanstalk_applications aws_elasticbeanstalk_applications_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_applications
    ADD CONSTRAINT aws_elasticbeanstalk_applications_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elasticbeanstalk_applications aws_elasticbeanstalk_applications_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_applications
    ADD CONSTRAINT aws_elasticbeanstalk_applications_pk PRIMARY KEY (arn, date_created);


--
-- Name: aws_elasticbeanstalk_configuration_options aws_elasticbeanstalk_configuration_options_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_options
    ADD CONSTRAINT aws_elasticbeanstalk_configuration_options_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elasticbeanstalk_configuration_setting_options aws_elasticbeanstalk_configuration_setting_options_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_setting_options
    ADD CONSTRAINT aws_elasticbeanstalk_configuration_setting_options_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elasticbeanstalk_configuration_settings aws_elasticbeanstalk_configuration_settings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_settings
    ADD CONSTRAINT aws_elasticbeanstalk_configuration_settings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elasticbeanstalk_environment_links aws_elasticbeanstalk_environment_links_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_environment_links
    ADD CONSTRAINT aws_elasticbeanstalk_environment_links_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elasticbeanstalk_environments aws_elasticbeanstalk_environments_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_environments
    ADD CONSTRAINT aws_elasticbeanstalk_environments_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elasticbeanstalk_environments aws_elasticbeanstalk_environments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_environments
    ADD CONSTRAINT aws_elasticbeanstalk_environments_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_elasticsearch_domains aws_elasticsearch_domains_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticsearch_domains
    ADD CONSTRAINT aws_elasticsearch_domains_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elasticsearch_domains aws_elasticsearch_domains_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticsearch_domains
    ADD CONSTRAINT aws_elasticsearch_domains_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_elbv1_load_balancer_backend_server_descriptions aws_elbv1_load_balancer_backend_server_descriptions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_backend_server_descriptions
    ADD CONSTRAINT aws_elbv1_load_balancer_backend_server_descriptions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv1_load_balancer_listeners aws_elbv1_load_balancer_listeners_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_listeners
    ADD CONSTRAINT aws_elbv1_load_balancer_listeners_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv1_load_balancer_policies_app_cookie_stickiness aws_elbv1_load_balancer_policies_app_cookie_stickiness_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies_app_cookie_stickiness
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_app_cookie_stickiness_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv1_load_balancer_policies_lb_cookie_stickiness aws_elbv1_load_balancer_policies_lb_cookie_stickiness_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies_lb_cookie_stickiness
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_lb_cookie_stickiness_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv1_load_balancer_policies aws_elbv1_load_balancer_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv1_load_balancers aws_elbv1_load_balancers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancers
    ADD CONSTRAINT aws_elbv1_load_balancers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elbv1_load_balancers aws_elbv1_load_balancers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancers
    ADD CONSTRAINT aws_elbv1_load_balancers_pk PRIMARY KEY (account_id, region, name);


--
-- Name: aws_elbv2_listener_certificates aws_elbv2_listener_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_certificates
    ADD CONSTRAINT aws_elbv2_listener_certificates_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_listener_default_action_forward_config_target_groups aws_elbv2_listener_default_action_forward_config_target_grou_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_default_action_forward_config_target_groups
    ADD CONSTRAINT aws_elbv2_listener_default_action_forward_config_target_grou_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_listener_default_actions aws_elbv2_listener_default_actions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_default_actions
    ADD CONSTRAINT aws_elbv2_listener_default_actions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_listeners aws_elbv2_listeners_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listeners
    ADD CONSTRAINT aws_elbv2_listeners_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_load_balancer_attributes aws_elbv2_load_balancer_attributes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_attributes
    ADD CONSTRAINT aws_elbv2_load_balancer_attributes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_load_balancer_availability_zone_addresses aws_elbv2_load_balancer_availability_zone_addresses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_availability_zone_addresses
    ADD CONSTRAINT aws_elbv2_load_balancer_availability_zone_addresses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_load_balancer_availability_zones aws_elbv2_load_balancer_availability_zones_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_availability_zones
    ADD CONSTRAINT aws_elbv2_load_balancer_availability_zones_pk PRIMARY KEY (cq_id);


--
-- Name: aws_elbv2_load_balancers aws_elbv2_load_balancers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancers
    ADD CONSTRAINT aws_elbv2_load_balancers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elbv2_load_balancers aws_elbv2_load_balancers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancers
    ADD CONSTRAINT aws_elbv2_load_balancers_pk PRIMARY KEY (arn);


--
-- Name: aws_elbv2_target_groups aws_elbv2_target_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_target_groups
    ADD CONSTRAINT aws_elbv2_target_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_elbv2_target_groups aws_elbv2_target_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_target_groups
    ADD CONSTRAINT aws_elbv2_target_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_emr_block_public_access_config_port_ranges aws_emr_block_public_access_config_port_ranges_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_block_public_access_config_port_ranges
    ADD CONSTRAINT aws_emr_block_public_access_config_port_ranges_pk PRIMARY KEY (cq_id);


--
-- Name: aws_emr_block_public_access_configs aws_emr_block_public_access_configs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_block_public_access_configs
    ADD CONSTRAINT aws_emr_block_public_access_configs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_emr_block_public_access_configs aws_emr_block_public_access_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_block_public_access_configs
    ADD CONSTRAINT aws_emr_block_public_access_configs_pk PRIMARY KEY (account_id, region);


--
-- Name: aws_emr_clusters aws_emr_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_emr_clusters aws_emr_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_pk PRIMARY KEY (arn);


--
-- Name: aws_fsx_backups aws_fsx_backups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_fsx_backups
    ADD CONSTRAINT aws_fsx_backups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_fsx_backups aws_fsx_backups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_fsx_backups
    ADD CONSTRAINT aws_fsx_backups_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_guardduty_detector_members aws_guardduty_detector_members_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_guardduty_detector_members
    ADD CONSTRAINT aws_guardduty_detector_members_pk PRIMARY KEY (cq_id);


--
-- Name: aws_guardduty_detectors aws_guardduty_detectors_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_guardduty_detectors
    ADD CONSTRAINT aws_guardduty_detectors_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_guardduty_detectors aws_guardduty_detectors_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_guardduty_detectors
    ADD CONSTRAINT aws_guardduty_detectors_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_iam_group_policies aws_iam_group_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_group_policies
    ADD CONSTRAINT aws_iam_group_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_groups aws_iam_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_groups
    ADD CONSTRAINT aws_iam_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_groups aws_iam_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_groups
    ADD CONSTRAINT aws_iam_groups_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_iam_openid_connect_identity_providers aws_iam_openid_connect_identity_providers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_openid_connect_identity_providers
    ADD CONSTRAINT aws_iam_openid_connect_identity_providers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_openid_connect_identity_providers aws_iam_openid_connect_identity_providers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_openid_connect_identity_providers
    ADD CONSTRAINT aws_iam_openid_connect_identity_providers_pk PRIMARY KEY (arn);


--
-- Name: aws_iam_password_policies aws_iam_password_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_password_policies
    ADD CONSTRAINT aws_iam_password_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_password_policies aws_iam_password_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_password_policies
    ADD CONSTRAINT aws_iam_password_policies_pk PRIMARY KEY (account_id);


--
-- Name: aws_iam_policies aws_iam_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_policies
    ADD CONSTRAINT aws_iam_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_policies aws_iam_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_policies
    ADD CONSTRAINT aws_iam_policies_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_iam_policy_versions aws_iam_policy_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_policy_versions
    ADD CONSTRAINT aws_iam_policy_versions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_role_policies aws_iam_role_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_role_policies
    ADD CONSTRAINT aws_iam_role_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_roles aws_iam_roles_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_roles
    ADD CONSTRAINT aws_iam_roles_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_roles aws_iam_roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_roles
    ADD CONSTRAINT aws_iam_roles_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_iam_saml_identity_providers aws_iam_saml_identity_providers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_saml_identity_providers
    ADD CONSTRAINT aws_iam_saml_identity_providers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_saml_identity_providers aws_iam_saml_identity_providers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_saml_identity_providers
    ADD CONSTRAINT aws_iam_saml_identity_providers_pk PRIMARY KEY (arn);


--
-- Name: aws_iam_server_certificates aws_iam_server_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_server_certificates
    ADD CONSTRAINT aws_iam_server_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_server_certificates aws_iam_server_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_server_certificates
    ADD CONSTRAINT aws_iam_server_certificates_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_iam_user_access_keys aws_iam_user_access_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_access_keys
    ADD CONSTRAINT aws_iam_user_access_keys_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_user_attached_policies aws_iam_user_attached_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_attached_policies
    ADD CONSTRAINT aws_iam_user_attached_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_user_groups aws_iam_user_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_groups
    ADD CONSTRAINT aws_iam_user_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_user_policies aws_iam_user_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_policies
    ADD CONSTRAINT aws_iam_user_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iam_users aws_iam_users_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_users
    ADD CONSTRAINT aws_iam_users_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_users aws_iam_users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_users
    ADD CONSTRAINT aws_iam_users_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_iam_virtual_mfa_devices aws_iam_virtual_mfa_devices_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_virtual_mfa_devices
    ADD CONSTRAINT aws_iam_virtual_mfa_devices_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iam_virtual_mfa_devices aws_iam_virtual_mfa_devices_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_virtual_mfa_devices
    ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (serial_number);


--
-- Name: aws_iot_billing_groups aws_iot_billing_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_billing_groups
    ADD CONSTRAINT aws_iot_billing_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_billing_groups aws_iot_billing_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_billing_groups
    ADD CONSTRAINT aws_iot_billing_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_ca_certificates aws_iot_ca_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_ca_certificates
    ADD CONSTRAINT aws_iot_ca_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_ca_certificates aws_iot_ca_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_ca_certificates
    ADD CONSTRAINT aws_iot_ca_certificates_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_certificates aws_iot_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_certificates
    ADD CONSTRAINT aws_iot_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_certificates aws_iot_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_certificates
    ADD CONSTRAINT aws_iot_certificates_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_policies aws_iot_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_policies
    ADD CONSTRAINT aws_iot_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_policies aws_iot_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_policies
    ADD CONSTRAINT aws_iot_policies_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_stream_files aws_iot_stream_files_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_stream_files
    ADD CONSTRAINT aws_iot_stream_files_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iot_streams aws_iot_streams_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_streams
    ADD CONSTRAINT aws_iot_streams_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_streams aws_iot_streams_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_streams
    ADD CONSTRAINT aws_iot_streams_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_thing_groups aws_iot_thing_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_thing_groups
    ADD CONSTRAINT aws_iot_thing_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_thing_groups aws_iot_thing_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_thing_groups
    ADD CONSTRAINT aws_iot_thing_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_thing_types aws_iot_thing_types_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_thing_types
    ADD CONSTRAINT aws_iot_thing_types_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_thing_types aws_iot_thing_types_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_thing_types
    ADD CONSTRAINT aws_iot_thing_types_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_things aws_iot_things_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_things
    ADD CONSTRAINT aws_iot_things_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_things aws_iot_things_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_things
    ADD CONSTRAINT aws_iot_things_pk PRIMARY KEY (arn);


--
-- Name: aws_iot_topic_rule_actions aws_iot_topic_rule_actions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_topic_rule_actions
    ADD CONSTRAINT aws_iot_topic_rule_actions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_iot_topic_rules aws_iot_topic_rules_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_topic_rules
    ADD CONSTRAINT aws_iot_topic_rules_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_iot_topic_rules aws_iot_topic_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_topic_rules
    ADD CONSTRAINT aws_iot_topic_rules_pk PRIMARY KEY (arn);


--
-- Name: aws_kms_keys aws_kms_keys_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_kms_keys
    ADD CONSTRAINT aws_kms_keys_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_kms_keys aws_kms_keys_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_kms_keys
    ADD CONSTRAINT aws_kms_keys_pk PRIMARY KEY (arn);


--
-- Name: aws_lambda_function_aliases aws_lambda_function_aliases_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_aliases
    ADD CONSTRAINT aws_lambda_function_aliases_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_concurrency_configs aws_lambda_function_concurrency_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_concurrency_configs
    ADD CONSTRAINT aws_lambda_function_concurrency_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_event_invoke_configs aws_lambda_function_event_invoke_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_event_invoke_configs
    ADD CONSTRAINT aws_lambda_function_event_invoke_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_event_source_mappings aws_lambda_function_event_source_mappings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_event_source_mappings
    ADD CONSTRAINT aws_lambda_function_event_source_mappings_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_file_system_configs aws_lambda_function_file_system_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_file_system_configs
    ADD CONSTRAINT aws_lambda_function_file_system_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_layers aws_lambda_function_layers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_layers
    ADD CONSTRAINT aws_lambda_function_layers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_version_file_system_configs aws_lambda_function_version_file_system_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_version_file_system_configs
    ADD CONSTRAINT aws_lambda_function_version_file_system_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_version_layers aws_lambda_function_version_layers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_version_layers
    ADD CONSTRAINT aws_lambda_function_version_layers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_function_versions aws_lambda_function_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_versions
    ADD CONSTRAINT aws_lambda_function_versions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_functions aws_lambda_functions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_functions
    ADD CONSTRAINT aws_lambda_functions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_lambda_functions aws_lambda_functions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_functions
    ADD CONSTRAINT aws_lambda_functions_pk PRIMARY KEY (arn);


--
-- Name: aws_lambda_layer_version_policies aws_lambda_layer_version_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layer_version_policies
    ADD CONSTRAINT aws_lambda_layer_version_policies_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_layer_versions aws_lambda_layer_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layer_versions
    ADD CONSTRAINT aws_lambda_layer_versions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_lambda_layers aws_lambda_layers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layers
    ADD CONSTRAINT aws_lambda_layers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_lambda_layers aws_lambda_layers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layers
    ADD CONSTRAINT aws_lambda_layers_pk PRIMARY KEY (arn);


--
-- Name: aws_lambda_runtimes aws_lambda_runtimes_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_runtimes
    ADD CONSTRAINT aws_lambda_runtimes_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_lambda_runtimes aws_lambda_runtimes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_runtimes
    ADD CONSTRAINT aws_lambda_runtimes_pk PRIMARY KEY (name);


--
-- Name: aws_mq_broker_configuration_revisions aws_mq_broker_configuration_revisions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_configuration_revisions
    ADD CONSTRAINT aws_mq_broker_configuration_revisions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_mq_broker_configurations aws_mq_broker_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_configurations
    ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_mq_broker_users aws_mq_broker_users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (cq_id);


--
-- Name: aws_mq_brokers aws_mq_brokers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_mq_brokers aws_mq_brokers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_organizations_accounts aws_organizations_accounts_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_organizations_accounts
    ADD CONSTRAINT aws_organizations_accounts_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_organizations_accounts aws_organizations_accounts_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_organizations_accounts
    ADD CONSTRAINT aws_organizations_accounts_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_qldb_ledger_journal_kinesis_streams aws_qldb_ledger_journal_kinesis_streams_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledger_journal_kinesis_streams
    ADD CONSTRAINT aws_qldb_ledger_journal_kinesis_streams_pk PRIMARY KEY (cq_id);


--
-- Name: aws_qldb_ledger_journal_s3_exports aws_qldb_ledger_journal_s3_exports_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledger_journal_s3_exports
    ADD CONSTRAINT aws_qldb_ledger_journal_s3_exports_pk PRIMARY KEY (cq_id);


--
-- Name: aws_qldb_ledgers aws_qldb_ledgers_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledgers
    ADD CONSTRAINT aws_qldb_ledgers_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_qldb_ledgers aws_qldb_ledgers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledgers
    ADD CONSTRAINT aws_qldb_ledgers_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_certificates aws_rds_certificates_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_certificates
    ADD CONSTRAINT aws_rds_certificates_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_certificates aws_rds_certificates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_certificates
    ADD CONSTRAINT aws_rds_certificates_pk PRIMARY KEY (account_id, arn);


--
-- Name: aws_rds_cluster_associated_roles aws_rds_cluster_associated_roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_associated_roles
    ADD CONSTRAINT aws_rds_cluster_associated_roles_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_cluster_db_cluster_members aws_rds_cluster_db_cluster_members_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_db_cluster_members
    ADD CONSTRAINT aws_rds_cluster_db_cluster_members_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_cluster_domain_memberships aws_rds_cluster_domain_memberships_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_domain_memberships
    ADD CONSTRAINT aws_rds_cluster_domain_memberships_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_cluster_parameter_groups aws_rds_cluster_parameter_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_parameter_groups
    ADD CONSTRAINT aws_rds_cluster_parameter_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_cluster_parameter_groups aws_rds_cluster_parameter_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_parameter_groups
    ADD CONSTRAINT aws_rds_cluster_parameter_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_cluster_parameters aws_rds_cluster_parameters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_parameters
    ADD CONSTRAINT aws_rds_cluster_parameters_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_cluster_snapshots aws_rds_cluster_snapshots_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_snapshots
    ADD CONSTRAINT aws_rds_cluster_snapshots_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_cluster_snapshots aws_rds_cluster_snapshots_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_snapshots
    ADD CONSTRAINT aws_rds_cluster_snapshots_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_cluster_vpc_security_groups aws_rds_cluster_vpc_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_vpc_security_groups
    ADD CONSTRAINT aws_rds_cluster_vpc_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_clusters aws_rds_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_clusters
    ADD CONSTRAINT aws_rds_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_clusters aws_rds_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_clusters
    ADD CONSTRAINT aws_rds_clusters_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_rds_db_parameter_groups aws_rds_db_parameter_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_parameter_groups
    ADD CONSTRAINT aws_rds_db_parameter_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_db_parameter_groups aws_rds_db_parameter_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_parameter_groups
    ADD CONSTRAINT aws_rds_db_parameter_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_db_parameters aws_rds_db_parameters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_parameters
    ADD CONSTRAINT aws_rds_db_parameters_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_db_security_groups aws_rds_db_security_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_security_groups
    ADD CONSTRAINT aws_rds_db_security_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_db_security_groups aws_rds_db_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_security_groups
    ADD CONSTRAINT aws_rds_db_security_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_db_snapshots aws_rds_db_snapshots_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_snapshots
    ADD CONSTRAINT aws_rds_db_snapshots_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_db_snapshots aws_rds_db_snapshots_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_snapshots
    ADD CONSTRAINT aws_rds_db_snapshots_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_event_subscriptions aws_rds_event_subscriptions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_event_subscriptions
    ADD CONSTRAINT aws_rds_event_subscriptions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_event_subscriptions aws_rds_event_subscriptions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_event_subscriptions
    ADD CONSTRAINT aws_rds_event_subscriptions_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_instance_associated_roles aws_rds_instance_associated_roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_associated_roles
    ADD CONSTRAINT aws_rds_instance_associated_roles_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_db_instance_automated_backups_replications aws_rds_instance_db_instance_automated_backups_replications_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_instance_automated_backups_replications
    ADD CONSTRAINT aws_rds_instance_db_instance_automated_backups_replications_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_db_parameter_groups aws_rds_instance_db_parameter_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_parameter_groups
    ADD CONSTRAINT aws_rds_instance_db_parameter_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_db_security_groups aws_rds_instance_db_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_security_groups
    ADD CONSTRAINT aws_rds_instance_db_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_db_subnet_group_subnets aws_rds_instance_db_subnet_group_subnets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_subnet_group_subnets
    ADD CONSTRAINT aws_rds_instance_db_subnet_group_subnets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_domain_memberships aws_rds_instance_domain_memberships_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_domain_memberships
    ADD CONSTRAINT aws_rds_instance_domain_memberships_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_option_group_memberships aws_rds_instance_option_group_memberships_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_option_group_memberships
    ADD CONSTRAINT aws_rds_instance_option_group_memberships_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instance_vpc_security_groups aws_rds_instance_vpc_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_vpc_security_groups
    ADD CONSTRAINT aws_rds_instance_vpc_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_instances aws_rds_instances_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instances
    ADD CONSTRAINT aws_rds_instances_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_instances aws_rds_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instances
    ADD CONSTRAINT aws_rds_instances_pk PRIMARY KEY (arn);


--
-- Name: aws_rds_subnet_group_subnets aws_rds_subnet_group_subnets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_subnet_group_subnets
    ADD CONSTRAINT aws_rds_subnet_group_subnets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_rds_subnet_groups aws_rds_subnet_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_subnet_groups
    ADD CONSTRAINT aws_rds_subnet_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_rds_subnet_groups aws_rds_subnet_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_subnet_groups
    ADD CONSTRAINT aws_rds_subnet_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_redshift_cluster_deferred_maintenance_windows aws_redshift_cluster_deferred_maintenance_windows_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_deferred_maintenance_windows
    ADD CONSTRAINT aws_redshift_cluster_deferred_maintenance_windows_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces aws_redshift_cluster_endpoint_vpc_endpoint_network_interface_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces
    ADD CONSTRAINT aws_redshift_cluster_endpoint_vpc_endpoint_network_interface_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoints aws_redshift_cluster_endpoint_vpc_endpoints_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_endpoint_vpc_endpoints
    ADD CONSTRAINT aws_redshift_cluster_endpoint_vpc_endpoints_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_iam_roles aws_redshift_cluster_iam_roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_iam_roles
    ADD CONSTRAINT aws_redshift_cluster_iam_roles_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_nodes aws_redshift_cluster_nodes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_nodes
    ADD CONSTRAINT aws_redshift_cluster_nodes_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_parameter_group_status_lists aws_redshift_cluster_parameter_group_status_lists_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameter_group_status_lists
    ADD CONSTRAINT aws_redshift_cluster_parameter_group_status_lists_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_parameter_groups aws_redshift_cluster_parameter_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameter_groups
    ADD CONSTRAINT aws_redshift_cluster_parameter_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_parameters aws_redshift_cluster_parameters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameters
    ADD CONSTRAINT aws_redshift_cluster_parameters_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_security_groups aws_redshift_cluster_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_security_groups
    ADD CONSTRAINT aws_redshift_cluster_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_cluster_vpc_security_groups aws_redshift_cluster_vpc_security_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_vpc_security_groups
    ADD CONSTRAINT aws_redshift_cluster_vpc_security_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_clusters aws_redshift_clusters_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_clusters
    ADD CONSTRAINT aws_redshift_clusters_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_redshift_clusters aws_redshift_clusters_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_clusters
    ADD CONSTRAINT aws_redshift_clusters_pk PRIMARY KEY (arn);


--
-- Name: aws_redshift_event_subscriptions aws_redshift_event_subscriptions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_event_subscriptions
    ADD CONSTRAINT aws_redshift_event_subscriptions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_redshift_event_subscriptions aws_redshift_event_subscriptions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_event_subscriptions
    ADD CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY (arn);


--
-- Name: aws_redshift_snapshot_accounts_with_restore_access aws_redshift_snapshot_accounts_with_restore_access_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_snapshot_accounts_with_restore_access
    ADD CONSTRAINT aws_redshift_snapshot_accounts_with_restore_access_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_snapshots aws_redshift_snapshots_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_subnet_group_subnets aws_redshift_subnet_group_subnets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_subnet_group_subnets
    ADD CONSTRAINT aws_redshift_subnet_group_subnets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_redshift_subnet_groups aws_redshift_subnet_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_subnet_groups
    ADD CONSTRAINT aws_redshift_subnet_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_redshift_subnet_groups aws_redshift_subnet_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_subnet_groups
    ADD CONSTRAINT aws_redshift_subnet_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_regions aws_regions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_regions
    ADD CONSTRAINT aws_regions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_domain_nameservers aws_route53_domain_nameservers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_domain_nameservers
    ADD CONSTRAINT aws_route53_domain_nameservers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_domains aws_route53_domains_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_domains
    ADD CONSTRAINT aws_route53_domains_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_route53_domains aws_route53_domains_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_domains
    ADD CONSTRAINT aws_route53_domains_pk PRIMARY KEY (account_id, domain_name);


--
-- Name: aws_route53_health_checks aws_route53_health_checks_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_health_checks
    ADD CONSTRAINT aws_route53_health_checks_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_route53_health_checks aws_route53_health_checks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_health_checks
    ADD CONSTRAINT aws_route53_health_checks_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_route53_hosted_zone_query_logging_configs aws_route53_hosted_zone_query_logging_configs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_query_logging_configs
    ADD CONSTRAINT aws_route53_hosted_zone_query_logging_configs_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_hosted_zone_resource_record_sets aws_route53_hosted_zone_resource_record_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_resource_record_sets
    ADD CONSTRAINT aws_route53_hosted_zone_resource_record_sets_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_hosted_zone_traffic_policy_instances aws_route53_hosted_zone_traffic_policy_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_traffic_policy_instances
    ADD CONSTRAINT aws_route53_hosted_zone_traffic_policy_instances_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_hosted_zone_vpc_association_authorizations aws_route53_hosted_zone_vpc_association_authorizations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_vpc_association_authorizations
    ADD CONSTRAINT aws_route53_hosted_zone_vpc_association_authorizations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_route53_hosted_zones aws_route53_hosted_zones_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zones
    ADD CONSTRAINT aws_route53_hosted_zones_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_route53_hosted_zones aws_route53_hosted_zones_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zones
    ADD CONSTRAINT aws_route53_hosted_zones_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_route53_reusable_delegation_sets aws_route53_reusable_delegation_sets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_reusable_delegation_sets
    ADD CONSTRAINT aws_route53_reusable_delegation_sets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_route53_reusable_delegation_sets aws_route53_reusable_delegation_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_reusable_delegation_sets
    ADD CONSTRAINT aws_route53_reusable_delegation_sets_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_route53_traffic_policies aws_route53_traffic_policies_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_traffic_policies
    ADD CONSTRAINT aws_route53_traffic_policies_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_route53_traffic_policies aws_route53_traffic_policies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_traffic_policies
    ADD CONSTRAINT aws_route53_traffic_policies_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_route53_traffic_policy_versions aws_route53_traffic_policy_versions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_traffic_policy_versions
    ADD CONSTRAINT aws_route53_traffic_policy_versions_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_account_config aws_s3_account_config_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_account_config
    ADD CONSTRAINT aws_s3_account_config_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_s3_account_config aws_s3_account_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_account_config
    ADD CONSTRAINT aws_s3_account_config_pk PRIMARY KEY (account_id);


--
-- Name: aws_s3_bucket_cors_rules aws_s3_bucket_cors_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_cors_rules
    ADD CONSTRAINT aws_s3_bucket_cors_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_bucket_encryption_rules aws_s3_bucket_encryption_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_encryption_rules
    ADD CONSTRAINT aws_s3_bucket_encryption_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_bucket_grants aws_s3_bucket_grants_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_grants
    ADD CONSTRAINT aws_s3_bucket_grants_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_bucket_lifecycles aws_s3_bucket_lifecycles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_lifecycles
    ADD CONSTRAINT aws_s3_bucket_lifecycles_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_bucket_replication_rules aws_s3_bucket_replication_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_replication_rules
    ADD CONSTRAINT aws_s3_bucket_replication_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_s3_buckets aws_s3_buckets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_buckets
    ADD CONSTRAINT aws_s3_buckets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_s3_buckets aws_s3_buckets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_buckets
    ADD CONSTRAINT aws_s3_buckets_pk PRIMARY KEY (account_id, name);


--
-- Name: aws_sagemaker_endpoint_configuration_production_variants aws_sagemaker_endpoint_configuration_production_variants_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_endpoint_configuration_production_variants
    ADD CONSTRAINT aws_sagemaker_endpoint_configuration_production_variants_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_endpoint_configurations aws_sagemaker_endpoint_configurations_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_endpoint_configurations
    ADD CONSTRAINT aws_sagemaker_endpoint_configurations_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sagemaker_endpoint_configurations aws_sagemaker_endpoint_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_endpoint_configurations
    ADD CONSTRAINT aws_sagemaker_endpoint_configurations_pk PRIMARY KEY (arn);


--
-- Name: aws_sagemaker_model_containers aws_sagemaker_model_containers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_model_containers
    ADD CONSTRAINT aws_sagemaker_model_containers_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_model_vpc_config aws_sagemaker_model_vpc_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_model_vpc_config
    ADD CONSTRAINT aws_sagemaker_model_vpc_config_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_models aws_sagemaker_models_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_models
    ADD CONSTRAINT aws_sagemaker_models_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sagemaker_models aws_sagemaker_models_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_models
    ADD CONSTRAINT aws_sagemaker_models_pk PRIMARY KEY (arn);


--
-- Name: aws_sagemaker_notebook_instances aws_sagemaker_notebook_instances_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_notebook_instances
    ADD CONSTRAINT aws_sagemaker_notebook_instances_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sagemaker_notebook_instances aws_sagemaker_notebook_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_notebook_instances
    ADD CONSTRAINT aws_sagemaker_notebook_instances_pk PRIMARY KEY (arn);


--
-- Name: aws_sagemaker_training_job_algorithm_specification aws_sagemaker_training_job_algorithm_specification_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_algorithm_specification
    ADD CONSTRAINT aws_sagemaker_training_job_algorithm_specification_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_debug_hook_config aws_sagemaker_training_job_debug_hook_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_hook_config
    ADD CONSTRAINT aws_sagemaker_training_job_debug_hook_config_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_debug_rule_configurations aws_sagemaker_training_job_debug_rule_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_rule_configurations
    ADD CONSTRAINT aws_sagemaker_training_job_debug_rule_configurations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_debug_rule_evaluation_statuses aws_sagemaker_training_job_debug_rule_evaluation_statuses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_rule_evaluation_statuses
    ADD CONSTRAINT aws_sagemaker_training_job_debug_rule_evaluation_statuses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_input_data_config aws_sagemaker_training_job_input_data_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_input_data_config
    ADD CONSTRAINT aws_sagemaker_training_job_input_data_config_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_profiler_rule_configurations aws_sagemaker_training_job_profiler_rule_configurations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_profiler_rule_configurations
    ADD CONSTRAINT aws_sagemaker_training_job_profiler_rule_configurations_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_job_profiler_rule_evaluation_statuses aws_sagemaker_training_job_profiler_rule_evaluation_statuses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_profiler_rule_evaluation_statuses
    ADD CONSTRAINT aws_sagemaker_training_job_profiler_rule_evaluation_statuses_pk PRIMARY KEY (cq_id);


--
-- Name: aws_sagemaker_training_jobs aws_sagemaker_training_jobs_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_jobs
    ADD CONSTRAINT aws_sagemaker_training_jobs_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sagemaker_training_jobs aws_sagemaker_training_jobs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_jobs
    ADD CONSTRAINT aws_sagemaker_training_jobs_pk PRIMARY KEY (arn);


--
-- Name: aws_secretsmanager_secrets aws_secretsmanager_secrets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_secretsmanager_secrets
    ADD CONSTRAINT aws_secretsmanager_secrets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_secretsmanager_secrets aws_secretsmanager_secrets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_secretsmanager_secrets
    ADD CONSTRAINT aws_secretsmanager_secrets_pk PRIMARY KEY (arn);


--
-- Name: aws_shield_attack_properties aws_shield_attack_properties_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attack_properties
    ADD CONSTRAINT aws_shield_attack_properties_pk PRIMARY KEY (cq_id);


--
-- Name: aws_shield_attack_sub_resources aws_shield_attack_sub_resources_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attack_sub_resources
    ADD CONSTRAINT aws_shield_attack_sub_resources_pk PRIMARY KEY (cq_id);


--
-- Name: aws_shield_attacks aws_shield_attacks_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attacks
    ADD CONSTRAINT aws_shield_attacks_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_shield_attacks aws_shield_attacks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attacks
    ADD CONSTRAINT aws_shield_attacks_pk PRIMARY KEY (id);


--
-- Name: aws_shield_protection_groups aws_shield_protection_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_protection_groups
    ADD CONSTRAINT aws_shield_protection_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_shield_protection_groups aws_shield_protection_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_protection_groups
    ADD CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_shield_protections aws_shield_protections_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_protections
    ADD CONSTRAINT aws_shield_protections_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_shield_protections aws_shield_protections_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_protections
    ADD CONSTRAINT aws_shield_protections_pk PRIMARY KEY (arn);


--
-- Name: aws_shield_subscriptions aws_shield_subscriptions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_subscriptions
    ADD CONSTRAINT aws_shield_subscriptions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_shield_subscriptions aws_shield_subscriptions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_subscriptions
    ADD CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY (arn);


--
-- Name: aws_sns_subscriptions aws_sns_subscriptions_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sns_subscriptions
    ADD CONSTRAINT aws_sns_subscriptions_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sns_subscriptions aws_sns_subscriptions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sns_subscriptions
    ADD CONSTRAINT aws_sns_subscriptions_pk PRIMARY KEY (endpoint, owner, protocol, arn, topic_arn);


--
-- Name: aws_sns_topics aws_sns_topics_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sns_topics
    ADD CONSTRAINT aws_sns_topics_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sns_topics aws_sns_topics_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sns_topics
    ADD CONSTRAINT aws_sns_topics_pk PRIMARY KEY (arn);


--
-- Name: aws_sqs_queues aws_sqs_queues_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sqs_queues
    ADD CONSTRAINT aws_sqs_queues_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_sqs_queues aws_sqs_queues_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sqs_queues
    ADD CONSTRAINT aws_sqs_queues_pk PRIMARY KEY (arn);


--
-- Name: aws_ssm_documents aws_ssm_documents_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_documents
    ADD CONSTRAINT aws_ssm_documents_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ssm_documents aws_ssm_documents_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_documents
    ADD CONSTRAINT aws_ssm_documents_pk PRIMARY KEY (arn);


--
-- Name: aws_ssm_instance_compliance_items aws_ssm_instance_compliance_items_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_instance_compliance_items
    ADD CONSTRAINT aws_ssm_instance_compliance_items_pk PRIMARY KEY (cq_id);


--
-- Name: aws_ssm_instances aws_ssm_instances_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_instances
    ADD CONSTRAINT aws_ssm_instances_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_ssm_instances aws_ssm_instances_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_instances
    ADD CONSTRAINT aws_ssm_instances_pk PRIMARY KEY (arn);


--
-- Name: aws_waf_rule_groups aws_waf_rule_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rule_groups
    ADD CONSTRAINT aws_waf_rule_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_waf_rule_groups aws_waf_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rule_groups
    ADD CONSTRAINT aws_waf_rule_groups_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_waf_rule_predicates aws_waf_rule_predicates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rule_predicates
    ADD CONSTRAINT aws_waf_rule_predicates_pk PRIMARY KEY (cq_id);


--
-- Name: aws_waf_rules aws_waf_rules_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rules
    ADD CONSTRAINT aws_waf_rules_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_waf_rules aws_waf_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rules
    ADD CONSTRAINT aws_waf_rules_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_waf_subscribed_rule_groups aws_waf_subscribed_rule_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_subscribed_rule_groups
    ADD CONSTRAINT aws_waf_subscribed_rule_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_waf_subscribed_rule_groups aws_waf_subscribed_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_subscribed_rule_groups
    ADD CONSTRAINT aws_waf_subscribed_rule_groups_pk PRIMARY KEY (account_id, rule_group_id);


--
-- Name: aws_waf_web_acl_logging_configuration aws_waf_web_acl_logging_configuration_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acl_logging_configuration
    ADD CONSTRAINT aws_waf_web_acl_logging_configuration_pk PRIMARY KEY (cq_id);


--
-- Name: aws_waf_web_acl_rules aws_waf_web_acl_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acl_rules
    ADD CONSTRAINT aws_waf_web_acl_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_waf_web_acls aws_waf_web_acls_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acls
    ADD CONSTRAINT aws_waf_web_acls_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_waf_web_acls aws_waf_web_acls_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acls
    ADD CONSTRAINT aws_waf_web_acls_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_wafregional_rate_based_rule_match_predicates aws_wafregional_rate_based_rule_match_predicates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rate_based_rule_match_predicates
    ADD CONSTRAINT aws_wafregional_rate_based_rule_match_predicates_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafregional_rate_based_rules aws_wafregional_rate_based_rules_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rate_based_rules
    ADD CONSTRAINT aws_wafregional_rate_based_rules_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafregional_rate_based_rules aws_wafregional_rate_based_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rate_based_rules
    ADD CONSTRAINT aws_wafregional_rate_based_rules_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_wafregional_rule_groups aws_wafregional_rule_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rule_groups
    ADD CONSTRAINT aws_wafregional_rule_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafregional_rule_groups aws_wafregional_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rule_groups
    ADD CONSTRAINT aws_wafregional_rule_groups_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_wafregional_rule_predicates aws_wafregional_rule_predicates_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rule_predicates
    ADD CONSTRAINT aws_wafregional_rule_predicates_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafregional_rules aws_wafregional_rules_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rules
    ADD CONSTRAINT aws_wafregional_rules_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafregional_rules aws_wafregional_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rules
    ADD CONSTRAINT aws_wafregional_rules_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_wafregional_web_acl_rules aws_wafregional_web_acl_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_web_acl_rules
    ADD CONSTRAINT aws_wafregional_web_acl_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafregional_web_acls aws_wafregional_web_acls_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_web_acls
    ADD CONSTRAINT aws_wafregional_web_acls_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafregional_web_acls aws_wafregional_web_acls_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_web_acls
    ADD CONSTRAINT aws_wafregional_web_acls_pk PRIMARY KEY (account_id, region, id);


--
-- Name: aws_wafv2_ipsets aws_wafv2_ipsets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_ipsets
    ADD CONSTRAINT aws_wafv2_ipsets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafv2_ipsets aws_wafv2_ipsets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_ipsets
    ADD CONSTRAINT aws_wafv2_ipsets_pk PRIMARY KEY (arn);


--
-- Name: aws_wafv2_managed_rule_groups aws_wafv2_managed_rule_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_managed_rule_groups
    ADD CONSTRAINT aws_wafv2_managed_rule_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafv2_managed_rule_groups aws_wafv2_managed_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_managed_rule_groups
    ADD CONSTRAINT aws_wafv2_managed_rule_groups_pk PRIMARY KEY (account_id, region, scope, vendor_name, name);


--
-- Name: aws_wafv2_regex_pattern_sets aws_wafv2_regex_pattern_sets_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_regex_pattern_sets
    ADD CONSTRAINT aws_wafv2_regex_pattern_sets_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafv2_regex_pattern_sets aws_wafv2_regex_pattern_sets_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_regex_pattern_sets
    ADD CONSTRAINT aws_wafv2_regex_pattern_sets_pk PRIMARY KEY (arn);


--
-- Name: aws_wafv2_rule_groups aws_wafv2_rule_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_rule_groups
    ADD CONSTRAINT aws_wafv2_rule_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafv2_rule_groups aws_wafv2_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_rule_groups
    ADD CONSTRAINT aws_wafv2_rule_groups_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_wafv2_web_acl_logging_configuration aws_wafv2_web_acl_logging_configuration_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_logging_configuration
    ADD CONSTRAINT aws_wafv2_web_acl_logging_configuration_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafv2_web_acl_post_process_firewall_manager_rule_groups aws_wafv2_web_acl_post_process_firewall_manager_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_post_process_firewall_manager_rule_groups
    ADD CONSTRAINT aws_wafv2_web_acl_post_process_firewall_manager_rule_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups
    ADD CONSTRAINT aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafv2_web_acl_rules aws_wafv2_web_acl_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_rules
    ADD CONSTRAINT aws_wafv2_web_acl_rules_pk PRIMARY KEY (cq_id);


--
-- Name: aws_wafv2_web_acls aws_wafv2_web_acls_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acls
    ADD CONSTRAINT aws_wafv2_web_acls_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_wafv2_web_acls aws_wafv2_web_acls_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acls
    ADD CONSTRAINT aws_wafv2_web_acls_pk PRIMARY KEY (account_id, id);


--
-- Name: aws_workspaces_directories aws_workspaces_directories_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_workspaces_directories
    ADD CONSTRAINT aws_workspaces_directories_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_workspaces_directories aws_workspaces_directories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_workspaces_directories
    ADD CONSTRAINT aws_workspaces_directories_pk PRIMARY KEY (id);


--
-- Name: aws_workspaces_workspaces aws_workspaces_workspaces_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_workspaces_workspaces
    ADD CONSTRAINT aws_workspaces_workspaces_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_workspaces_workspaces aws_workspaces_workspaces_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_workspaces_workspaces
    ADD CONSTRAINT aws_workspaces_workspaces_pk PRIMARY KEY (id);


--
-- Name: aws_xray_encryption_config aws_xray_encryption_config_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_encryption_config
    ADD CONSTRAINT aws_xray_encryption_config_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_xray_encryption_config aws_xray_encryption_config_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_encryption_config
    ADD CONSTRAINT aws_xray_encryption_config_pk PRIMARY KEY (account_id, region);


--
-- Name: aws_xray_groups aws_xray_groups_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_groups
    ADD CONSTRAINT aws_xray_groups_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_xray_groups aws_xray_groups_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_groups
    ADD CONSTRAINT aws_xray_groups_pk PRIMARY KEY (arn);


--
-- Name: aws_xray_sampling_rules aws_xray_sampling_rules_cq_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_sampling_rules
    ADD CONSTRAINT aws_xray_sampling_rules_cq_id_key UNIQUE (cq_id);


--
-- Name: aws_xray_sampling_rules aws_xray_sampling_rules_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_xray_sampling_rules
    ADD CONSTRAINT aws_xray_sampling_rules_pk PRIMARY KEY (arn);


--
-- Name: check_results_execution_timestamp_idx; Type: INDEX; Schema: cloudquery; Owner: postgres
--

CREATE INDEX check_results_execution_timestamp_idx ON cloudquery.check_results USING btree (execution_timestamp);


--
-- Name: check_results_selector_idx; Type: INDEX; Schema: cloudquery; Owner: postgres
--

CREATE INDEX check_results_selector_idx ON cloudquery.check_results USING gin (selector cloudquery.gin_trgm_ops);


--
-- Name: policy_executions_timestamp_idx; Type: INDEX; Schema: cloudquery; Owner: postgres
--

CREATE INDEX policy_executions_timestamp_idx ON cloudquery.policy_executions USING btree ("timestamp");


--
-- Name: policy_executions_version_idx; Type: INDEX; Schema: cloudquery; Owner: postgres
--

CREATE INDEX policy_executions_version_idx ON cloudquery.policy_executions USING brin (version);


--
-- Name: check_results CalculatePolicyExecutionsStats; Type: TRIGGER; Schema: cloudquery; Owner: postgres
--

CREATE TRIGGER "CalculatePolicyExecutionsStats" AFTER INSERT ON cloudquery.check_results FOR EACH ROW EXECUTE PROCEDURE cloudquery.calculate_policy_executions_stats();


--
-- Name: check_results check_results_execution_id_fkey; Type: FK CONSTRAINT; Schema: cloudquery; Owner: postgres
--

ALTER TABLE ONLY cloudquery.check_results
    ADD CONSTRAINT check_results_execution_id_fkey FOREIGN KEY (execution_id) REFERENCES cloudquery.policy_executions(id) ON DELETE CASCADE;


--
-- Name: aws_access_analyzer_analyzer_archive_rules aws_access_analyzer_analyzer_archive_rules_analyzer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_archive_rules
    ADD CONSTRAINT aws_access_analyzer_analyzer_archive_rules_analyzer_cq_id_fkey FOREIGN KEY (analyzer_cq_id) REFERENCES public.aws_access_analyzer_analyzers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_access_analyzer_analyzer_finding_sources aws_access_analyzer_analyzer_findin_analyzer_finding_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_finding_sources
    ADD CONSTRAINT aws_access_analyzer_analyzer_findin_analyzer_finding_cq_id_fkey FOREIGN KEY (analyzer_finding_cq_id) REFERENCES public.aws_access_analyzer_analyzer_findings(cq_id) ON DELETE CASCADE;


--
-- Name: aws_access_analyzer_analyzer_findings aws_access_analyzer_analyzer_findings_analyzer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_access_analyzer_analyzer_findings
    ADD CONSTRAINT aws_access_analyzer_analyzer_findings_analyzer_cq_id_fkey FOREIGN KEY (analyzer_cq_id) REFERENCES public.aws_access_analyzer_analyzers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_domain_name_base_path_mappings aws_apigateway_domain_name_base_path_map_domain_name_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_domain_name_base_path_mappings
    ADD CONSTRAINT aws_apigateway_domain_name_base_path_map_domain_name_cq_id_fkey FOREIGN KEY (domain_name_cq_id) REFERENCES public.aws_apigateway_domain_names(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_authorizers aws_apigateway_rest_api_authorizers_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_authorizers
    ADD CONSTRAINT aws_apigateway_rest_api_authorizers_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_deployments aws_apigateway_rest_api_deployments_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_deployments
    ADD CONSTRAINT aws_apigateway_rest_api_deployments_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_documentation_parts aws_apigateway_rest_api_documentation_parts_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_documentation_parts
    ADD CONSTRAINT aws_apigateway_rest_api_documentation_parts_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_documentation_versions aws_apigateway_rest_api_documentation_versi_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_documentation_versions
    ADD CONSTRAINT aws_apigateway_rest_api_documentation_versi_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_gateway_responses aws_apigateway_rest_api_gateway_responses_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_gateway_responses
    ADD CONSTRAINT aws_apigateway_rest_api_gateway_responses_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_models aws_apigateway_rest_api_models_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_models
    ADD CONSTRAINT aws_apigateway_rest_api_models_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_request_validators aws_apigateway_rest_api_request_validators_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_request_validators
    ADD CONSTRAINT aws_apigateway_rest_api_request_validators_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_resources aws_apigateway_rest_api_resources_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_resources
    ADD CONSTRAINT aws_apigateway_rest_api_resources_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_rest_api_stages aws_apigateway_rest_api_stages_rest_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_rest_api_stages
    ADD CONSTRAINT aws_apigateway_rest_api_stages_rest_api_cq_id_fkey FOREIGN KEY (rest_api_cq_id) REFERENCES public.aws_apigateway_rest_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_usage_plan_api_stages aws_apigateway_usage_plan_api_stages_usage_plan_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plan_api_stages
    ADD CONSTRAINT aws_apigateway_usage_plan_api_stages_usage_plan_cq_id_fkey FOREIGN KEY (usage_plan_cq_id) REFERENCES public.aws_apigateway_usage_plans(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigateway_usage_plan_keys aws_apigateway_usage_plan_keys_usage_plan_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigateway_usage_plan_keys
    ADD CONSTRAINT aws_apigateway_usage_plan_keys_usage_plan_cq_id_fkey FOREIGN KEY (usage_plan_cq_id) REFERENCES public.aws_apigateway_usage_plans(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_authorizers aws_apigatewayv2_api_authorizers_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_authorizers
    ADD CONSTRAINT aws_apigatewayv2_api_authorizers_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_deployments aws_apigatewayv2_api_deployments_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_deployments
    ADD CONSTRAINT aws_apigatewayv2_api_deployments_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_integration_responses aws_apigatewayv2_api_integration_res_api_integration_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_integration_responses
    ADD CONSTRAINT aws_apigatewayv2_api_integration_res_api_integration_cq_id_fkey FOREIGN KEY (api_integration_cq_id) REFERENCES public.aws_apigatewayv2_api_integrations(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_integrations aws_apigatewayv2_api_integrations_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_integrations
    ADD CONSTRAINT aws_apigatewayv2_api_integrations_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_models aws_apigatewayv2_api_models_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_models
    ADD CONSTRAINT aws_apigatewayv2_api_models_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_route_responses aws_apigatewayv2_api_route_responses_api_route_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_route_responses
    ADD CONSTRAINT aws_apigatewayv2_api_route_responses_api_route_cq_id_fkey FOREIGN KEY (api_route_cq_id) REFERENCES public.aws_apigatewayv2_api_routes(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_routes aws_apigatewayv2_api_routes_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_routes
    ADD CONSTRAINT aws_apigatewayv2_api_routes_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_api_stages aws_apigatewayv2_api_stages_api_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_api_stages
    ADD CONSTRAINT aws_apigatewayv2_api_stages_api_cq_id_fkey FOREIGN KEY (api_cq_id) REFERENCES public.aws_apigatewayv2_apis(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_domain_name_configurations aws_apigatewayv2_domain_name_configurati_domain_name_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_name_configurations
    ADD CONSTRAINT aws_apigatewayv2_domain_name_configurati_domain_name_cq_id_fkey FOREIGN KEY (domain_name_cq_id) REFERENCES public.aws_apigatewayv2_domain_names(cq_id) ON DELETE CASCADE;


--
-- Name: aws_apigatewayv2_domain_name_rest_api_mappings aws_apigatewayv2_domain_name_rest_api_ma_domain_name_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_apigatewayv2_domain_name_rest_api_mappings
    ADD CONSTRAINT aws_apigatewayv2_domain_name_rest_api_ma_domain_name_cq_id_fkey FOREIGN KEY (domain_name_cq_id) REFERENCES public.aws_apigatewayv2_domain_names(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_data_catalog_database_table_partition_keys aws_athena_data_catalog_data_data_catalog_database_table__fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_table_partition_keys
    ADD CONSTRAINT aws_athena_data_catalog_data_data_catalog_database_table__fkey1 FOREIGN KEY (data_catalog_database_table_cq_id) REFERENCES public.aws_athena_data_catalog_database_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_data_catalog_database_table_columns aws_athena_data_catalog_datab_data_catalog_database_table__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_table_columns
    ADD CONSTRAINT aws_athena_data_catalog_datab_data_catalog_database_table__fkey FOREIGN KEY (data_catalog_database_table_cq_id) REFERENCES public.aws_athena_data_catalog_database_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_data_catalog_database_tables aws_athena_data_catalog_databa_data_catalog_database_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_database_tables
    ADD CONSTRAINT aws_athena_data_catalog_databa_data_catalog_database_cq_id_fkey FOREIGN KEY (data_catalog_database_cq_id) REFERENCES public.aws_athena_data_catalog_databases(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_data_catalog_databases aws_athena_data_catalog_databases_data_catalog_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_data_catalog_databases
    ADD CONSTRAINT aws_athena_data_catalog_databases_data_catalog_cq_id_fkey FOREIGN KEY (data_catalog_cq_id) REFERENCES public.aws_athena_data_catalogs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_work_group_named_queries aws_athena_work_group_named_queries_work_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_named_queries
    ADD CONSTRAINT aws_athena_work_group_named_queries_work_group_cq_id_fkey FOREIGN KEY (work_group_cq_id) REFERENCES public.aws_athena_work_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_work_group_prepared_statements aws_athena_work_group_prepared_statements_work_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_prepared_statements
    ADD CONSTRAINT aws_athena_work_group_prepared_statements_work_group_cq_id_fkey FOREIGN KEY (work_group_cq_id) REFERENCES public.aws_athena_work_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_athena_work_group_query_executions aws_athena_work_group_query_executions_work_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_athena_work_group_query_executions
    ADD CONSTRAINT aws_athena_work_group_query_executions_work_group_cq_id_fkey FOREIGN KEY (work_group_cq_id) REFERENCES public.aws_athena_work_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_autoscaling_group_instances aws_autoscaling_group_instances_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_instances
    ADD CONSTRAINT aws_autoscaling_group_instances_group_cq_id_fkey FOREIGN KEY (group_cq_id) REFERENCES public.aws_autoscaling_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_autoscaling_group_lifecycle_hooks aws_autoscaling_group_lifecycle_hooks_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_lifecycle_hooks
    ADD CONSTRAINT aws_autoscaling_group_lifecycle_hooks_group_cq_id_fkey FOREIGN KEY (group_cq_id) REFERENCES public.aws_autoscaling_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_autoscaling_group_scaling_policies aws_autoscaling_group_scaling_policies_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_scaling_policies
    ADD CONSTRAINT aws_autoscaling_group_scaling_policies_group_cq_id_fkey FOREIGN KEY (group_cq_id) REFERENCES public.aws_autoscaling_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_autoscaling_group_tags aws_autoscaling_group_tags_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_group_tags
    ADD CONSTRAINT aws_autoscaling_group_tags_group_cq_id_fkey FOREIGN KEY (group_cq_id) REFERENCES public.aws_autoscaling_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_autoscaling_launch_configuration_block_device_mappings aws_autoscaling_launch_configur_launch_configuration_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_autoscaling_launch_configuration_block_device_mappings
    ADD CONSTRAINT aws_autoscaling_launch_configur_launch_configuration_cq_id_fkey FOREIGN KEY (launch_configuration_cq_id) REFERENCES public.aws_autoscaling_launch_configurations(cq_id) ON DELETE CASCADE;


--
-- Name: aws_backup_plan_rules aws_backup_plan_rules_plan_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plan_rules
    ADD CONSTRAINT aws_backup_plan_rules_plan_cq_id_fkey FOREIGN KEY (plan_cq_id) REFERENCES public.aws_backup_plans(cq_id) ON DELETE CASCADE;


--
-- Name: aws_backup_plan_selections aws_backup_plan_selections_plan_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_plan_selections
    ADD CONSTRAINT aws_backup_plan_selections_plan_cq_id_fkey FOREIGN KEY (plan_cq_id) REFERENCES public.aws_backup_plans(cq_id) ON DELETE CASCADE;


--
-- Name: aws_backup_vault_recovery_points aws_backup_vault_recovery_points_vault_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_backup_vault_recovery_points
    ADD CONSTRAINT aws_backup_vault_recovery_points_vault_cq_id_fkey FOREIGN KEY (vault_cq_id) REFERENCES public.aws_backup_vaults(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudformation_stack_outputs aws_cloudformation_stack_outputs_stack_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stack_outputs
    ADD CONSTRAINT aws_cloudformation_stack_outputs_stack_cq_id_fkey FOREIGN KEY (stack_cq_id) REFERENCES public.aws_cloudformation_stacks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudformation_stack_resources aws_cloudformation_stack_resources_stack_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudformation_stack_resources
    ADD CONSTRAINT aws_cloudformation_stack_resources_stack_cq_id_fkey FOREIGN KEY (stack_cq_id) REFERENCES public.aws_cloudformation_stacks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_cache_behavior_lambda_functions aws_cloudfront_distribution_c_distribution_cache_behavior__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_cache_behavior_lambda_functions
    ADD CONSTRAINT aws_cloudfront_distribution_c_distribution_cache_behavior__fkey FOREIGN KEY (distribution_cache_behavior_cq_id) REFERENCES public.aws_cloudfront_distribution_cache_behaviors(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_cache_behaviors aws_cloudfront_distribution_cache_behav_distribution_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_cache_behaviors
    ADD CONSTRAINT aws_cloudfront_distribution_cache_behav_distribution_cq_id_fkey FOREIGN KEY (distribution_cq_id) REFERENCES public.aws_cloudfront_distributions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_custom_error_responses aws_cloudfront_distribution_custom_erro_distribution_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_custom_error_responses
    ADD CONSTRAINT aws_cloudfront_distribution_custom_erro_distribution_cq_id_fkey FOREIGN KEY (distribution_cq_id) REFERENCES public.aws_cloudfront_distributions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_default_cache_behavior_functions aws_cloudfront_distribution_default_cac_distribution_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_default_cache_behavior_functions
    ADD CONSTRAINT aws_cloudfront_distribution_default_cac_distribution_cq_id_fkey FOREIGN KEY (distribution_cq_id) REFERENCES public.aws_cloudfront_distributions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_origin_groups aws_cloudfront_distribution_origin_grou_distribution_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_origin_groups
    ADD CONSTRAINT aws_cloudfront_distribution_origin_grou_distribution_cq_id_fkey FOREIGN KEY (distribution_cq_id) REFERENCES public.aws_cloudfront_distributions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudfront_distribution_origins aws_cloudfront_distribution_origins_distribution_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudfront_distribution_origins
    ADD CONSTRAINT aws_cloudfront_distribution_origins_distribution_cq_id_fkey FOREIGN KEY (distribution_cq_id) REFERENCES public.aws_cloudfront_distributions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudtrail_trail_event_selectors aws_cloudtrail_trail_event_selectors_trail_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudtrail_trail_event_selectors
    ADD CONSTRAINT aws_cloudtrail_trail_event_selectors_trail_cq_id_fkey FOREIGN KEY (trail_cq_id) REFERENCES public.aws_cloudtrail_trails(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudwatch_alarm_metrics aws_cloudwatch_alarm_metrics_alarm_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatch_alarm_metrics
    ADD CONSTRAINT aws_cloudwatch_alarm_metrics_alarm_cq_id_fkey FOREIGN KEY (alarm_cq_id) REFERENCES public.aws_cloudwatch_alarms(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cloudwatchlogs_filter_metric_transformations aws_cloudwatchlogs_filter_metric_transformati_filter_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cloudwatchlogs_filter_metric_transformations
    ADD CONSTRAINT aws_cloudwatchlogs_filter_metric_transformati_filter_cq_id_fkey FOREIGN KEY (filter_cq_id) REFERENCES public.aws_cloudwatchlogs_filters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codebuild_project_environment_variables aws_codebuild_project_environment_variables_project_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_environment_variables
    ADD CONSTRAINT aws_codebuild_project_environment_variables_project_cq_id_fkey FOREIGN KEY (project_cq_id) REFERENCES public.aws_codebuild_projects(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codebuild_project_file_system_locations aws_codebuild_project_file_system_locations_project_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_file_system_locations
    ADD CONSTRAINT aws_codebuild_project_file_system_locations_project_cq_id_fkey FOREIGN KEY (project_cq_id) REFERENCES public.aws_codebuild_projects(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codebuild_project_secondary_artifacts aws_codebuild_project_secondary_artifacts_project_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_secondary_artifacts
    ADD CONSTRAINT aws_codebuild_project_secondary_artifacts_project_cq_id_fkey FOREIGN KEY (project_cq_id) REFERENCES public.aws_codebuild_projects(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codebuild_project_secondary_sources aws_codebuild_project_secondary_sources_project_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codebuild_project_secondary_sources
    ADD CONSTRAINT aws_codebuild_project_secondary_sources_project_cq_id_fkey FOREIGN KEY (project_cq_id) REFERENCES public.aws_codebuild_projects(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codepipeline_pipeline_stage_actions aws_codepipeline_pipeline_stage_actio_pipeline_stage_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipeline_stage_actions
    ADD CONSTRAINT aws_codepipeline_pipeline_stage_actio_pipeline_stage_cq_id_fkey FOREIGN KEY (pipeline_stage_cq_id) REFERENCES public.aws_codepipeline_pipeline_stages(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codepipeline_pipeline_stages aws_codepipeline_pipeline_stages_pipeline_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_pipeline_stages
    ADD CONSTRAINT aws_codepipeline_pipeline_stages_pipeline_cq_id_fkey FOREIGN KEY (pipeline_cq_id) REFERENCES public.aws_codepipeline_pipelines(cq_id) ON DELETE CASCADE;


--
-- Name: aws_codepipeline_webhook_filters aws_codepipeline_webhook_filters_webhook_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_codepipeline_webhook_filters
    ADD CONSTRAINT aws_codepipeline_webhook_filters_webhook_cq_id_fkey FOREIGN KEY (webhook_cq_id) REFERENCES public.aws_codepipeline_webhooks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cognito_identity_pool_cognito_identity_providers aws_cognito_identity_pool_cognito_iden_identity_pool_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_identity_pool_cognito_identity_providers
    ADD CONSTRAINT aws_cognito_identity_pool_cognito_iden_identity_pool_cq_id_fkey FOREIGN KEY (identity_pool_cq_id) REFERENCES public.aws_cognito_identity_pools(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cognito_user_pool_identity_providers aws_cognito_user_pool_identity_providers_user_pool_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pool_identity_providers
    ADD CONSTRAINT aws_cognito_user_pool_identity_providers_user_pool_cq_id_fkey FOREIGN KEY (user_pool_cq_id) REFERENCES public.aws_cognito_user_pools(cq_id) ON DELETE CASCADE;


--
-- Name: aws_cognito_user_pool_schema_attributes aws_cognito_user_pool_schema_attributes_user_pool_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_cognito_user_pool_schema_attributes
    ADD CONSTRAINT aws_cognito_user_pool_schema_attributes_user_pool_cq_id_fkey FOREIGN KEY (user_pool_cq_id) REFERENCES public.aws_cognito_user_pools(cq_id) ON DELETE CASCADE;


--
-- Name: aws_config_conformance_pack_rule_compliances aws_config_conformance_pack_rule_co_conformance_pack_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_config_conformance_pack_rule_compliances
    ADD CONSTRAINT aws_config_conformance_pack_rule_co_conformance_pack_cq_id_fkey FOREIGN KEY (conformance_pack_cq_id) REFERENCES public.aws_config_conformance_packs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dax_cluster_nodes aws_dax_cluster_nodes_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dax_cluster_nodes
    ADD CONSTRAINT aws_dax_cluster_nodes_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_dax_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_directconnect_connection_mac_sec_keys aws_directconnect_connection_mac_sec_keys_connection_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_connection_mac_sec_keys
    ADD CONSTRAINT aws_directconnect_connection_mac_sec_keys_connection_cq_id_fkey FOREIGN KEY (connection_cq_id) REFERENCES public.aws_directconnect_connections(cq_id) ON DELETE CASCADE;


--
-- Name: aws_directconnect_gateway_associations aws_directconnect_gateway_associations_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateway_associations
    ADD CONSTRAINT aws_directconnect_gateway_associations_gateway_cq_id_fkey FOREIGN KEY (gateway_cq_id) REFERENCES public.aws_directconnect_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_directconnect_gateway_attachments aws_directconnect_gateway_attachments_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_gateway_attachments
    ADD CONSTRAINT aws_directconnect_gateway_attachments_gateway_cq_id_fkey FOREIGN KEY (gateway_cq_id) REFERENCES public.aws_directconnect_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_directconnect_lag_mac_sec_keys aws_directconnect_lag_mac_sec_keys_lag_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_lag_mac_sec_keys
    ADD CONSTRAINT aws_directconnect_lag_mac_sec_keys_lag_cq_id_fkey FOREIGN KEY (lag_cq_id) REFERENCES public.aws_directconnect_lags(cq_id) ON DELETE CASCADE;


--
-- Name: aws_directconnect_virtual_interface_bgp_peers aws_directconnect_virtual_interfac_virtual_interface_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_directconnect_virtual_interface_bgp_peers
    ADD CONSTRAINT aws_directconnect_virtual_interfac_virtual_interface_cq_id_fkey FOREIGN KEY (virtual_interface_cq_id) REFERENCES public.aws_directconnect_virtual_interfaces(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dms_replication_instance_replication_subnet_group_subnets aws_dms_replication_instance_re_replication_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instance_replication_subnet_group_subnets
    ADD CONSTRAINT aws_dms_replication_instance_re_replication_instance_cq_id_fkey FOREIGN KEY (replication_instance_cq_id) REFERENCES public.aws_dms_replication_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dms_replication_instance_vpc_security_groups aws_dms_replication_instance_vp_replication_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dms_replication_instance_vpc_security_groups
    ADD CONSTRAINT aws_dms_replication_instance_vp_replication_instance_cq_id_fkey FOREIGN KEY (replication_instance_cq_id) REFERENCES public.aws_dms_replication_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dynamodb_table_continuous_backups aws_dynamodb_table_continuous_backups_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_continuous_backups
    ADD CONSTRAINT aws_dynamodb_table_continuous_backups_table_cq_id_fkey FOREIGN KEY (table_cq_id) REFERENCES public.aws_dynamodb_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dynamodb_table_global_secondary_indexes aws_dynamodb_table_global_secondary_indexes_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_global_secondary_indexes
    ADD CONSTRAINT aws_dynamodb_table_global_secondary_indexes_table_cq_id_fkey FOREIGN KEY (table_cq_id) REFERENCES public.aws_dynamodb_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dynamodb_table_local_secondary_indexes aws_dynamodb_table_local_secondary_indexes_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_local_secondary_indexes
    ADD CONSTRAINT aws_dynamodb_table_local_secondary_indexes_table_cq_id_fkey FOREIGN KEY (table_cq_id) REFERENCES public.aws_dynamodb_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dynamodb_table_replica_auto_scalings aws_dynamodb_table_replica_auto_scalings_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_replica_auto_scalings
    ADD CONSTRAINT aws_dynamodb_table_replica_auto_scalings_table_cq_id_fkey FOREIGN KEY (table_cq_id) REFERENCES public.aws_dynamodb_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_dynamodb_table_replicas aws_dynamodb_table_replicas_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_dynamodb_table_replicas
    ADD CONSTRAINT aws_dynamodb_table_replicas_table_cq_id_fkey FOREIGN KEY (table_cq_id) REFERENCES public.aws_dynamodb_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_ebs_volume_attachments aws_ec2_ebs_volume_attachments_ebs_volume_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_ebs_volume_attachments
    ADD CONSTRAINT aws_ec2_ebs_volume_attachments_ebs_volume_cq_id_fkey FOREIGN KEY (ebs_volume_cq_id) REFERENCES public.aws_ec2_ebs_volumes(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_host_available_instance_capacity aws_ec2_host_available_instance_capacity_host_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_host_available_instance_capacity
    ADD CONSTRAINT aws_ec2_host_available_instance_capacity_host_cq_id_fkey FOREIGN KEY (host_cq_id) REFERENCES public.aws_ec2_hosts(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_host_instances aws_ec2_host_instances_host_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_host_instances
    ADD CONSTRAINT aws_ec2_host_instances_host_cq_id_fkey FOREIGN KEY (host_cq_id) REFERENCES public.aws_ec2_hosts(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_image_block_device_mappings aws_ec2_image_block_device_mappings_image_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_image_block_device_mappings
    ADD CONSTRAINT aws_ec2_image_block_device_mappings_image_cq_id_fkey FOREIGN KEY (image_cq_id) REFERENCES public.aws_ec2_images(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_block_device_mappings aws_ec2_instance_block_device_mappings_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_block_device_mappings
    ADD CONSTRAINT aws_ec2_instance_block_device_mappings_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_elastic_gpu_associations aws_ec2_instance_elastic_gpu_associations_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_elastic_gpu_associations
    ADD CONSTRAINT aws_ec2_instance_elastic_gpu_associations_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_elastic_inference_accelerator_associations aws_ec2_instance_elastic_inference_accelera_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_elastic_inference_accelerator_associations
    ADD CONSTRAINT aws_ec2_instance_elastic_inference_accelera_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_network_interface_ipv6_addresses aws_ec2_instance_network_int_instance_network_interface_c_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_ipv6_addresses
    ADD CONSTRAINT aws_ec2_instance_network_int_instance_network_interface_c_fkey1 FOREIGN KEY (instance_network_interface_cq_id) REFERENCES public.aws_ec2_instance_network_interfaces(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_network_interface_private_ip_addresses aws_ec2_instance_network_int_instance_network_interface_c_fkey2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_private_ip_addresses
    ADD CONSTRAINT aws_ec2_instance_network_int_instance_network_interface_c_fkey2 FOREIGN KEY (instance_network_interface_cq_id) REFERENCES public.aws_ec2_instance_network_interfaces(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_network_interface_groups aws_ec2_instance_network_inte_instance_network_interface_c_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interface_groups
    ADD CONSTRAINT aws_ec2_instance_network_inte_instance_network_interface_c_fkey FOREIGN KEY (instance_network_interface_cq_id) REFERENCES public.aws_ec2_instance_network_interfaces(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_network_interfaces aws_ec2_instance_network_interfaces_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_network_interfaces
    ADD CONSTRAINT aws_ec2_instance_network_interfaces_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_product_codes aws_ec2_instance_product_codes_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_product_codes
    ADD CONSTRAINT aws_ec2_instance_product_codes_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_security_groups aws_ec2_instance_security_groups_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_security_groups
    ADD CONSTRAINT aws_ec2_instance_security_groups_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ec2_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_instance_status_events aws_ec2_instance_status_events_instance_status_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_instance_status_events
    ADD CONSTRAINT aws_ec2_instance_status_events_instance_status_cq_id_fkey FOREIGN KEY (instance_status_cq_id) REFERENCES public.aws_ec2_instance_statuses(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_internet_gateway_attachments aws_ec2_internet_gateway_attachment_internet_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_internet_gateway_attachments
    ADD CONSTRAINT aws_ec2_internet_gateway_attachment_internet_gateway_cq_id_fkey FOREIGN KEY (internet_gateway_cq_id) REFERENCES public.aws_ec2_internet_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_nat_gateway_addresses aws_ec2_nat_gateway_addresses_nat_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_nat_gateway_addresses
    ADD CONSTRAINT aws_ec2_nat_gateway_addresses_nat_gateway_cq_id_fkey FOREIGN KEY (nat_gateway_cq_id) REFERENCES public.aws_ec2_nat_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_network_acl_associations aws_ec2_network_acl_associations_network_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acl_associations
    ADD CONSTRAINT aws_ec2_network_acl_associations_network_acl_cq_id_fkey FOREIGN KEY (network_acl_cq_id) REFERENCES public.aws_ec2_network_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_network_acl_entries aws_ec2_network_acl_entries_network_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_acl_entries
    ADD CONSTRAINT aws_ec2_network_acl_entries_network_acl_cq_id_fkey FOREIGN KEY (network_acl_cq_id) REFERENCES public.aws_ec2_network_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_network_interface_private_ip_addresses aws_ec2_network_interface_private__network_interface_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_network_interface_private_ip_addresses
    ADD CONSTRAINT aws_ec2_network_interface_private__network_interface_cq_id_fkey FOREIGN KEY (network_interface_cq_id) REFERENCES public.aws_ec2_network_interfaces(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_route_table_associations aws_ec2_route_table_associations_route_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_associations
    ADD CONSTRAINT aws_ec2_route_table_associations_route_table_cq_id_fkey FOREIGN KEY (route_table_cq_id) REFERENCES public.aws_ec2_route_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_route_table_propagating_vgws aws_ec2_route_table_propagating_vgws_route_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_propagating_vgws
    ADD CONSTRAINT aws_ec2_route_table_propagating_vgws_route_table_cq_id_fkey FOREIGN KEY (route_table_cq_id) REFERENCES public.aws_ec2_route_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_route_table_routes aws_ec2_route_table_routes_route_table_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_route_table_routes
    ADD CONSTRAINT aws_ec2_route_table_routes_route_table_cq_id_fkey FOREIGN KEY (route_table_cq_id) REFERENCES public.aws_ec2_route_tables(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_security_group_ip_permission_prefix_list_ids aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_prefix_list_ids
    ADD CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey1 FOREIGN KEY (security_group_ip_permission_cq_id) REFERENCES public.aws_ec2_security_group_ip_permissions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_security_group_ip_permission_user_id_group_pairs aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_user_id_group_pairs
    ADD CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey2 FOREIGN KEY (security_group_ip_permission_cq_id) REFERENCES public.aws_ec2_security_group_ip_permissions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_security_group_ip_permission_ip_ranges aws_ec2_security_group_ip_per_security_group_ip_permission_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permission_ip_ranges
    ADD CONSTRAINT aws_ec2_security_group_ip_per_security_group_ip_permission_fkey FOREIGN KEY (security_group_ip_permission_cq_id) REFERENCES public.aws_ec2_security_group_ip_permissions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_security_group_ip_permissions aws_ec2_security_group_ip_permissions_security_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_security_group_ip_permissions
    ADD CONSTRAINT aws_ec2_security_group_ip_permissions_security_group_cq_id_fkey FOREIGN KEY (security_group_cq_id) REFERENCES public.aws_ec2_security_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_subnet_ipv6_cidr_block_association_sets aws_ec2_subnet_ipv6_cidr_block_association_se_subnet_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_subnet_ipv6_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_subnet_ipv6_cidr_block_association_se_subnet_cq_id_fkey FOREIGN KEY (subnet_cq_id) REFERENCES public.aws_ec2_subnets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_transit_gateway_attachments aws_ec2_transit_gateway_attachments_transit_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_attachments_transit_gateway_cq_id_fkey FOREIGN KEY (transit_gateway_cq_id) REFERENCES public.aws_ec2_transit_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_transit_gateway_multicast_domains aws_ec2_transit_gateway_multicast_do_transit_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_multicast_domains
    ADD CONSTRAINT aws_ec2_transit_gateway_multicast_do_transit_gateway_cq_id_fkey FOREIGN KEY (transit_gateway_cq_id) REFERENCES public.aws_ec2_transit_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_transit_gateway_peering_attachments aws_ec2_transit_gateway_peering_atta_transit_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_peering_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_peering_atta_transit_gateway_cq_id_fkey FOREIGN KEY (transit_gateway_cq_id) REFERENCES public.aws_ec2_transit_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_transit_gateway_route_tables aws_ec2_transit_gateway_route_tables_transit_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_route_tables
    ADD CONSTRAINT aws_ec2_transit_gateway_route_tables_transit_gateway_cq_id_fkey FOREIGN KEY (transit_gateway_cq_id) REFERENCES public.aws_ec2_transit_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_transit_gateway_vpc_attachments aws_ec2_transit_gateway_vpc_attachme_transit_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_transit_gateway_vpc_attachments
    ADD CONSTRAINT aws_ec2_transit_gateway_vpc_attachme_transit_gateway_cq_id_fkey FOREIGN KEY (transit_gateway_cq_id) REFERENCES public.aws_ec2_transit_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_vpc_attachment aws_ec2_vpc_attachment_vpn_gateway_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_attachment
    ADD CONSTRAINT aws_ec2_vpc_attachment_vpn_gateway_cq_id_fkey FOREIGN KEY (vpn_gateway_cq_id) REFERENCES public.aws_ec2_vpn_gateways(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_vpc_cidr_block_association_sets aws_ec2_vpc_cidr_block_association_sets_vpc_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_vpc_cidr_block_association_sets_vpc_cq_id_fkey FOREIGN KEY (vpc_cq_id) REFERENCES public.aws_ec2_vpcs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_vpc_endpoint_dns_entries aws_ec2_vpc_endpoint_dns_entries_vpc_endpoint_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoint_dns_entries
    ADD CONSTRAINT aws_ec2_vpc_endpoint_dns_entries_vpc_endpoint_cq_id_fkey FOREIGN KEY (vpc_endpoint_cq_id) REFERENCES public.aws_ec2_vpc_endpoints(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_vpc_endpoint_groups aws_ec2_vpc_endpoint_groups_vpc_endpoint_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_endpoint_groups
    ADD CONSTRAINT aws_ec2_vpc_endpoint_groups_vpc_endpoint_cq_id_fkey FOREIGN KEY (vpc_endpoint_cq_id) REFERENCES public.aws_ec2_vpc_endpoints(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ec2_vpc_ipv6_cidr_block_association_sets aws_ec2_vpc_ipv6_cidr_block_association_sets_vpc_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ec2_vpc_ipv6_cidr_block_association_sets
    ADD CONSTRAINT aws_ec2_vpc_ipv6_cidr_block_association_sets_vpc_cq_id_fkey FOREIGN KEY (vpc_cq_id) REFERENCES public.aws_ec2_vpcs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecr_repository_images aws_ecr_repository_images_repository_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecr_repository_images
    ADD CONSTRAINT aws_ecr_repository_images_repository_cq_id_fkey FOREIGN KEY (repository_cq_id) REFERENCES public.aws_ecr_repositories(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_attachments aws_ecs_cluster_attachments_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_attachments
    ADD CONSTRAINT aws_ecs_cluster_attachments_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_ecs_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instance_attributes aws_ecs_cluster_container_in_cluster_container_instance_c_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_attributes
    ADD CONSTRAINT aws_ecs_cluster_container_in_cluster_container_instance_c_fkey1 FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES public.aws_ecs_cluster_container_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instance_health_status_details aws_ecs_cluster_container_in_cluster_container_instance_c_fkey2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_health_status_details
    ADD CONSTRAINT aws_ecs_cluster_container_in_cluster_container_instance_c_fkey2 FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES public.aws_ecs_cluster_container_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instance_registered_resources aws_ecs_cluster_container_in_cluster_container_instance_c_fkey3; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_registered_resources
    ADD CONSTRAINT aws_ecs_cluster_container_in_cluster_container_instance_c_fkey3 FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES public.aws_ecs_cluster_container_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instance_remaining_resources aws_ecs_cluster_container_in_cluster_container_instance_c_fkey4; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_remaining_resources
    ADD CONSTRAINT aws_ecs_cluster_container_in_cluster_container_instance_c_fkey4 FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES public.aws_ecs_cluster_container_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instance_attachments aws_ecs_cluster_container_ins_cluster_container_instance_c_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instance_attachments
    ADD CONSTRAINT aws_ecs_cluster_container_ins_cluster_container_instance_c_fkey FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES public.aws_ecs_cluster_container_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_container_instances aws_ecs_cluster_container_instances_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_container_instances
    ADD CONSTRAINT aws_ecs_cluster_container_instances_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_ecs_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_deployments aws_ecs_cluster_service_deployments_cluster_service_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_deployments
    ADD CONSTRAINT aws_ecs_cluster_service_deployments_cluster_service_cq_id_fkey FOREIGN KEY (cluster_service_cq_id) REFERENCES public.aws_ecs_cluster_services(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_events aws_ecs_cluster_service_events_cluster_service_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_events
    ADD CONSTRAINT aws_ecs_cluster_service_events_cluster_service_cq_id_fkey FOREIGN KEY (cluster_service_cq_id) REFERENCES public.aws_ecs_cluster_services(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_load_balancers aws_ecs_cluster_service_load_balance_cluster_service_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_load_balancers
    ADD CONSTRAINT aws_ecs_cluster_service_load_balance_cluster_service_cq_id_fkey FOREIGN KEY (cluster_service_cq_id) REFERENCES public.aws_ecs_cluster_services(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_service_registries aws_ecs_cluster_service_service_regi_cluster_service_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_service_registries
    ADD CONSTRAINT aws_ecs_cluster_service_service_regi_cluster_service_cq_id_fkey FOREIGN KEY (cluster_service_cq_id) REFERENCES public.aws_ecs_cluster_services(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_task_set_load_balancers aws_ecs_cluster_service_task__cluster_service_task_set_cq__fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_set_load_balancers
    ADD CONSTRAINT aws_ecs_cluster_service_task__cluster_service_task_set_cq__fkey FOREIGN KEY (cluster_service_task_set_cq_id) REFERENCES public.aws_ecs_cluster_service_task_sets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_task_set_service_registries aws_ecs_cluster_service_task_cluster_service_task_set_cq__fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_set_service_registries
    ADD CONSTRAINT aws_ecs_cluster_service_task_cluster_service_task_set_cq__fkey1 FOREIGN KEY (cluster_service_task_set_cq_id) REFERENCES public.aws_ecs_cluster_service_task_sets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_service_task_sets aws_ecs_cluster_service_task_sets_cluster_service_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_service_task_sets
    ADD CONSTRAINT aws_ecs_cluster_service_task_sets_cluster_service_cq_id_fkey FOREIGN KEY (cluster_service_cq_id) REFERENCES public.aws_ecs_cluster_services(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_services aws_ecs_cluster_services_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_services
    ADD CONSTRAINT aws_ecs_cluster_services_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_ecs_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_task_attachments aws_ecs_cluster_task_attachments_cluster_task_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_task_attachments
    ADD CONSTRAINT aws_ecs_cluster_task_attachments_cluster_task_cq_id_fkey FOREIGN KEY (cluster_task_cq_id) REFERENCES public.aws_ecs_cluster_tasks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_task_containers aws_ecs_cluster_task_containers_cluster_task_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_task_containers
    ADD CONSTRAINT aws_ecs_cluster_task_containers_cluster_task_cq_id_fkey FOREIGN KEY (cluster_task_cq_id) REFERENCES public.aws_ecs_cluster_tasks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_cluster_tasks aws_ecs_cluster_tasks_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_cluster_tasks
    ADD CONSTRAINT aws_ecs_cluster_tasks_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_ecs_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_task_definition_container_definitions aws_ecs_task_definition_container_de_task_definition_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definition_container_definitions
    ADD CONSTRAINT aws_ecs_task_definition_container_de_task_definition_cq_id_fkey FOREIGN KEY (task_definition_cq_id) REFERENCES public.aws_ecs_task_definitions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ecs_task_definition_volumes aws_ecs_task_definition_volumes_task_definition_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ecs_task_definition_volumes
    ADD CONSTRAINT aws_ecs_task_definition_volumes_task_definition_cq_id_fkey FOREIGN KEY (task_definition_cq_id) REFERENCES public.aws_ecs_task_definitions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_eks_cluster_encryption_configs aws_eks_cluster_encryption_configs_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_cluster_encryption_configs
    ADD CONSTRAINT aws_eks_cluster_encryption_configs_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_eks_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_eks_cluster_loggings aws_eks_cluster_loggings_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_eks_cluster_loggings
    ADD CONSTRAINT aws_eks_cluster_loggings_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_eks_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elasticbeanstalk_configuration_setting_options aws_elasticbeanstalk_configura_configuration_setting_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_setting_options
    ADD CONSTRAINT aws_elasticbeanstalk_configura_configuration_setting_cq_id_fkey FOREIGN KEY (configuration_setting_cq_id) REFERENCES public.aws_elasticbeanstalk_configuration_settings(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elasticbeanstalk_configuration_options aws_elasticbeanstalk_configuration_optio_environment_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_options
    ADD CONSTRAINT aws_elasticbeanstalk_configuration_optio_environment_cq_id_fkey FOREIGN KEY (environment_cq_id) REFERENCES public.aws_elasticbeanstalk_environments(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elasticbeanstalk_configuration_settings aws_elasticbeanstalk_configuration_setti_environment_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_configuration_settings
    ADD CONSTRAINT aws_elasticbeanstalk_configuration_setti_environment_cq_id_fkey FOREIGN KEY (environment_cq_id) REFERENCES public.aws_elasticbeanstalk_environments(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elasticbeanstalk_environment_links aws_elasticbeanstalk_environment_links_environment_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elasticbeanstalk_environment_links
    ADD CONSTRAINT aws_elasticbeanstalk_environment_links_environment_cq_id_fkey FOREIGN KEY (environment_cq_id) REFERENCES public.aws_elasticbeanstalk_environments(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv1_load_balancer_backend_server_descriptions aws_elbv1_load_balancer_backend_server_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_backend_server_descriptions
    ADD CONSTRAINT aws_elbv1_load_balancer_backend_server_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv1_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv1_load_balancer_listeners aws_elbv1_load_balancer_listeners_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_listeners
    ADD CONSTRAINT aws_elbv1_load_balancer_listeners_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv1_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv1_load_balancer_policies_app_cookie_stickiness aws_elbv1_load_balancer_policies_app_c_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies_app_cookie_stickiness
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_app_c_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv1_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv1_load_balancer_policies_lb_cookie_stickiness aws_elbv1_load_balancer_policies_lb_co_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies_lb_cookie_stickiness
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_lb_co_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv1_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv1_load_balancer_policies aws_elbv1_load_balancer_policies_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv1_load_balancer_policies
    ADD CONSTRAINT aws_elbv1_load_balancer_policies_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv1_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_listener_certificates aws_elbv2_listener_certificates_listener_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_certificates
    ADD CONSTRAINT aws_elbv2_listener_certificates_listener_cq_id_fkey FOREIGN KEY (listener_cq_id) REFERENCES public.aws_elbv2_listeners(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_listener_default_action_forward_config_target_groups aws_elbv2_listener_default_ac_listener_default_action_cq_i_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_default_action_forward_config_target_groups
    ADD CONSTRAINT aws_elbv2_listener_default_ac_listener_default_action_cq_i_fkey FOREIGN KEY (listener_default_action_cq_id) REFERENCES public.aws_elbv2_listener_default_actions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_listener_default_actions aws_elbv2_listener_default_actions_listener_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listener_default_actions
    ADD CONSTRAINT aws_elbv2_listener_default_actions_listener_cq_id_fkey FOREIGN KEY (listener_cq_id) REFERENCES public.aws_elbv2_listeners(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_listeners aws_elbv2_listeners_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_listeners
    ADD CONSTRAINT aws_elbv2_listeners_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv2_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_load_balancer_attributes aws_elbv2_load_balancer_attributes_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_attributes
    ADD CONSTRAINT aws_elbv2_load_balancer_attributes_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv2_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_load_balancer_availability_zone_addresses aws_elbv2_load_balancer_avail_load_balancer_availability_z_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_availability_zone_addresses
    ADD CONSTRAINT aws_elbv2_load_balancer_avail_load_balancer_availability_z_fkey FOREIGN KEY (load_balancer_availability_zone_cq_id) REFERENCES public.aws_elbv2_load_balancer_availability_zones(cq_id) ON DELETE CASCADE;


--
-- Name: aws_elbv2_load_balancer_availability_zones aws_elbv2_load_balancer_availability_z_load_balancer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_elbv2_load_balancer_availability_zones
    ADD CONSTRAINT aws_elbv2_load_balancer_availability_z_load_balancer_cq_id_fkey FOREIGN KEY (load_balancer_cq_id) REFERENCES public.aws_elbv2_load_balancers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_emr_block_public_access_config_port_ranges aws_emr_block_public_access_c_block_public_access_config_c_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_emr_block_public_access_config_port_ranges
    ADD CONSTRAINT aws_emr_block_public_access_c_block_public_access_config_c_fkey FOREIGN KEY (block_public_access_config_cq_id) REFERENCES public.aws_emr_block_public_access_configs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_guardduty_detector_members aws_guardduty_detector_members_detector_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_guardduty_detector_members
    ADD CONSTRAINT aws_guardduty_detector_members_detector_cq_id_fkey FOREIGN KEY (detector_cq_id) REFERENCES public.aws_guardduty_detectors(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_group_policies aws_iam_group_policies_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_group_policies
    ADD CONSTRAINT aws_iam_group_policies_group_cq_id_fkey FOREIGN KEY (group_cq_id) REFERENCES public.aws_iam_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_policy_versions aws_iam_policy_versions_policy_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_policy_versions
    ADD CONSTRAINT aws_iam_policy_versions_policy_cq_id_fkey FOREIGN KEY (policy_cq_id) REFERENCES public.aws_iam_policies(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_role_policies aws_iam_role_policies_role_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_role_policies
    ADD CONSTRAINT aws_iam_role_policies_role_cq_id_fkey FOREIGN KEY (role_cq_id) REFERENCES public.aws_iam_roles(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_user_access_keys aws_iam_user_access_keys_user_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_access_keys
    ADD CONSTRAINT aws_iam_user_access_keys_user_cq_id_fkey FOREIGN KEY (user_cq_id) REFERENCES public.aws_iam_users(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_user_attached_policies aws_iam_user_attached_policies_user_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_attached_policies
    ADD CONSTRAINT aws_iam_user_attached_policies_user_cq_id_fkey FOREIGN KEY (user_cq_id) REFERENCES public.aws_iam_users(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_user_groups aws_iam_user_groups_user_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_groups
    ADD CONSTRAINT aws_iam_user_groups_user_cq_id_fkey FOREIGN KEY (user_cq_id) REFERENCES public.aws_iam_users(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iam_user_policies aws_iam_user_policies_user_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iam_user_policies
    ADD CONSTRAINT aws_iam_user_policies_user_cq_id_fkey FOREIGN KEY (user_cq_id) REFERENCES public.aws_iam_users(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iot_stream_files aws_iot_stream_files_stream_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_stream_files
    ADD CONSTRAINT aws_iot_stream_files_stream_cq_id_fkey FOREIGN KEY (stream_cq_id) REFERENCES public.aws_iot_streams(cq_id) ON DELETE CASCADE;


--
-- Name: aws_iot_topic_rule_actions aws_iot_topic_rule_actions_topic_rule_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_iot_topic_rule_actions
    ADD CONSTRAINT aws_iot_topic_rule_actions_topic_rule_cq_id_fkey FOREIGN KEY (topic_rule_cq_id) REFERENCES public.aws_iot_topic_rules(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_aliases aws_lambda_function_aliases_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_aliases
    ADD CONSTRAINT aws_lambda_function_aliases_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_concurrency_configs aws_lambda_function_concurrency_configs_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_concurrency_configs
    ADD CONSTRAINT aws_lambda_function_concurrency_configs_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_event_invoke_configs aws_lambda_function_event_invoke_configs_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_event_invoke_configs
    ADD CONSTRAINT aws_lambda_function_event_invoke_configs_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_event_source_mappings aws_lambda_function_event_source_mappings_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_event_source_mappings
    ADD CONSTRAINT aws_lambda_function_event_source_mappings_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_file_system_configs aws_lambda_function_file_system_configs_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_file_system_configs
    ADD CONSTRAINT aws_lambda_function_file_system_configs_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_layers aws_lambda_function_layers_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_layers
    ADD CONSTRAINT aws_lambda_function_layers_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_version_file_system_configs aws_lambda_function_version_file_sy_function_version_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_version_file_system_configs
    ADD CONSTRAINT aws_lambda_function_version_file_sy_function_version_cq_id_fkey FOREIGN KEY (function_version_cq_id) REFERENCES public.aws_lambda_function_versions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_version_layers aws_lambda_function_version_layers_function_version_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_version_layers
    ADD CONSTRAINT aws_lambda_function_version_layers_function_version_cq_id_fkey FOREIGN KEY (function_version_cq_id) REFERENCES public.aws_lambda_function_versions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_function_versions aws_lambda_function_versions_function_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_function_versions
    ADD CONSTRAINT aws_lambda_function_versions_function_cq_id_fkey FOREIGN KEY (function_cq_id) REFERENCES public.aws_lambda_functions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_layer_version_policies aws_lambda_layer_version_policies_layer_version_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layer_version_policies
    ADD CONSTRAINT aws_lambda_layer_version_policies_layer_version_cq_id_fkey FOREIGN KEY (layer_version_cq_id) REFERENCES public.aws_lambda_layer_versions(cq_id) ON DELETE CASCADE;


--
-- Name: aws_lambda_layer_versions aws_lambda_layer_versions_layer_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_lambda_layer_versions
    ADD CONSTRAINT aws_lambda_layer_versions_layer_cq_id_fkey FOREIGN KEY (layer_cq_id) REFERENCES public.aws_lambda_layers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_mq_broker_configuration_revisions aws_mq_broker_configuration_rev_broker_configuration_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_configuration_revisions
    ADD CONSTRAINT aws_mq_broker_configuration_rev_broker_configuration_cq_id_fkey FOREIGN KEY (broker_configuration_cq_id) REFERENCES public.aws_mq_broker_configurations(cq_id) ON DELETE CASCADE;


--
-- Name: aws_mq_broker_configurations aws_mq_broker_configurations_broker_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_configurations
    ADD CONSTRAINT aws_mq_broker_configurations_broker_cq_id_fkey FOREIGN KEY (broker_cq_id) REFERENCES public.aws_mq_brokers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_mq_broker_users aws_mq_broker_users_broker_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_broker_cq_id_fkey FOREIGN KEY (broker_cq_id) REFERENCES public.aws_mq_brokers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_qldb_ledger_journal_kinesis_streams aws_qldb_ledger_journal_kinesis_streams_ledger_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledger_journal_kinesis_streams
    ADD CONSTRAINT aws_qldb_ledger_journal_kinesis_streams_ledger_cq_id_fkey FOREIGN KEY (ledger_cq_id) REFERENCES public.aws_qldb_ledgers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_qldb_ledger_journal_s3_exports aws_qldb_ledger_journal_s3_exports_ledger_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_qldb_ledger_journal_s3_exports
    ADD CONSTRAINT aws_qldb_ledger_journal_s3_exports_ledger_cq_id_fkey FOREIGN KEY (ledger_cq_id) REFERENCES public.aws_qldb_ledgers(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_cluster_associated_roles aws_rds_cluster_associated_roles_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_associated_roles
    ADD CONSTRAINT aws_rds_cluster_associated_roles_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_rds_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_cluster_db_cluster_members aws_rds_cluster_db_cluster_members_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_db_cluster_members
    ADD CONSTRAINT aws_rds_cluster_db_cluster_members_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_rds_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_cluster_domain_memberships aws_rds_cluster_domain_memberships_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_domain_memberships
    ADD CONSTRAINT aws_rds_cluster_domain_memberships_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_rds_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_cluster_parameters aws_rds_cluster_parameters_cluster_parameter_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_parameters
    ADD CONSTRAINT aws_rds_cluster_parameters_cluster_parameter_group_cq_id_fkey FOREIGN KEY (cluster_parameter_group_cq_id) REFERENCES public.aws_rds_cluster_parameter_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_cluster_vpc_security_groups aws_rds_cluster_vpc_security_groups_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_cluster_vpc_security_groups
    ADD CONSTRAINT aws_rds_cluster_vpc_security_groups_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_rds_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_db_parameters aws_rds_db_parameters_db_parameter_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_db_parameters
    ADD CONSTRAINT aws_rds_db_parameters_db_parameter_group_cq_id_fkey FOREIGN KEY (db_parameter_group_cq_id) REFERENCES public.aws_rds_db_parameter_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_associated_roles aws_rds_instance_associated_roles_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_associated_roles
    ADD CONSTRAINT aws_rds_instance_associated_roles_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_db_instance_automated_backups_replications aws_rds_instance_db_instance_automated_back_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_instance_automated_backups_replications
    ADD CONSTRAINT aws_rds_instance_db_instance_automated_back_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_db_parameter_groups aws_rds_instance_db_parameter_groups_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_parameter_groups
    ADD CONSTRAINT aws_rds_instance_db_parameter_groups_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_db_security_groups aws_rds_instance_db_security_groups_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_security_groups
    ADD CONSTRAINT aws_rds_instance_db_security_groups_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_db_subnet_group_subnets aws_rds_instance_db_subnet_group_subnets_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_db_subnet_group_subnets
    ADD CONSTRAINT aws_rds_instance_db_subnet_group_subnets_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_domain_memberships aws_rds_instance_domain_memberships_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_domain_memberships
    ADD CONSTRAINT aws_rds_instance_domain_memberships_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_option_group_memberships aws_rds_instance_option_group_memberships_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_option_group_memberships
    ADD CONSTRAINT aws_rds_instance_option_group_memberships_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_instance_vpc_security_groups aws_rds_instance_vpc_security_groups_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_instance_vpc_security_groups
    ADD CONSTRAINT aws_rds_instance_vpc_security_groups_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_rds_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_rds_subnet_group_subnets aws_rds_subnet_group_subnets_subnet_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_rds_subnet_group_subnets
    ADD CONSTRAINT aws_rds_subnet_group_subnets_subnet_group_cq_id_fkey FOREIGN KEY (subnet_group_cq_id) REFERENCES public.aws_rds_subnet_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_deferred_maintenance_windows aws_redshift_cluster_deferred_maintenance_wi_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_deferred_maintenance_windows
    ADD CONSTRAINT aws_redshift_cluster_deferred_maintenance_wi_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces aws_redshift_cluster_endpoint_cluster_endpoint_vpc_endpoin_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces
    ADD CONSTRAINT aws_redshift_cluster_endpoint_cluster_endpoint_vpc_endpoin_fkey FOREIGN KEY (cluster_endpoint_vpc_endpoint_cq_id) REFERENCES public.aws_redshift_cluster_endpoint_vpc_endpoints(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_endpoint_vpc_endpoints aws_redshift_cluster_endpoint_vpc_endpoints_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_endpoint_vpc_endpoints
    ADD CONSTRAINT aws_redshift_cluster_endpoint_vpc_endpoints_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_iam_roles aws_redshift_cluster_iam_roles_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_iam_roles
    ADD CONSTRAINT aws_redshift_cluster_iam_roles_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_nodes aws_redshift_cluster_nodes_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_nodes
    ADD CONSTRAINT aws_redshift_cluster_nodes_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_parameter_group_status_lists aws_redshift_cluster_paramet_cluster_parameter_group_cq_i_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameter_group_status_lists
    ADD CONSTRAINT aws_redshift_cluster_paramet_cluster_parameter_group_cq_i_fkey1 FOREIGN KEY (cluster_parameter_group_cq_id) REFERENCES public.aws_redshift_cluster_parameter_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_parameters aws_redshift_cluster_paramete_cluster_parameter_group_cq_i_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameters
    ADD CONSTRAINT aws_redshift_cluster_paramete_cluster_parameter_group_cq_i_fkey FOREIGN KEY (cluster_parameter_group_cq_id) REFERENCES public.aws_redshift_cluster_parameter_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_parameter_groups aws_redshift_cluster_parameter_groups_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_parameter_groups
    ADD CONSTRAINT aws_redshift_cluster_parameter_groups_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_security_groups aws_redshift_cluster_security_groups_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_security_groups
    ADD CONSTRAINT aws_redshift_cluster_security_groups_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_cluster_vpc_security_groups aws_redshift_cluster_vpc_security_groups_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_cluster_vpc_security_groups
    ADD CONSTRAINT aws_redshift_cluster_vpc_security_groups_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_snapshot_accounts_with_restore_access aws_redshift_snapshot_accounts_with_restore_snapshot_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_snapshot_accounts_with_restore_access
    ADD CONSTRAINT aws_redshift_snapshot_accounts_with_restore_snapshot_cq_id_fkey FOREIGN KEY (snapshot_cq_id) REFERENCES public.aws_redshift_snapshots(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_snapshots aws_redshift_snapshots_cluster_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey FOREIGN KEY (cluster_cq_id) REFERENCES public.aws_redshift_clusters(cq_id) ON DELETE CASCADE;


--
-- Name: aws_redshift_subnet_group_subnets aws_redshift_subnet_group_subnets_subnet_group_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_redshift_subnet_group_subnets
    ADD CONSTRAINT aws_redshift_subnet_group_subnets_subnet_group_cq_id_fkey FOREIGN KEY (subnet_group_cq_id) REFERENCES public.aws_redshift_subnet_groups(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_domain_nameservers aws_route53_domain_nameservers_domain_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_domain_nameservers
    ADD CONSTRAINT aws_route53_domain_nameservers_domain_cq_id_fkey FOREIGN KEY (domain_cq_id) REFERENCES public.aws_route53_domains(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_hosted_zone_query_logging_configs aws_route53_hosted_zone_query_logging_co_hosted_zone_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_query_logging_configs
    ADD CONSTRAINT aws_route53_hosted_zone_query_logging_co_hosted_zone_cq_id_fkey FOREIGN KEY (hosted_zone_cq_id) REFERENCES public.aws_route53_hosted_zones(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_hosted_zone_resource_record_sets aws_route53_hosted_zone_resource_record__hosted_zone_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_resource_record_sets
    ADD CONSTRAINT aws_route53_hosted_zone_resource_record__hosted_zone_cq_id_fkey FOREIGN KEY (hosted_zone_cq_id) REFERENCES public.aws_route53_hosted_zones(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_hosted_zone_traffic_policy_instances aws_route53_hosted_zone_traffic_policy_i_hosted_zone_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_traffic_policy_instances
    ADD CONSTRAINT aws_route53_hosted_zone_traffic_policy_i_hosted_zone_cq_id_fkey FOREIGN KEY (hosted_zone_cq_id) REFERENCES public.aws_route53_hosted_zones(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_hosted_zone_vpc_association_authorizations aws_route53_hosted_zone_vpc_association__hosted_zone_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_hosted_zone_vpc_association_authorizations
    ADD CONSTRAINT aws_route53_hosted_zone_vpc_association__hosted_zone_cq_id_fkey FOREIGN KEY (hosted_zone_cq_id) REFERENCES public.aws_route53_hosted_zones(cq_id) ON DELETE CASCADE;


--
-- Name: aws_route53_traffic_policy_versions aws_route53_traffic_policy_versions_traffic_policy_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_route53_traffic_policy_versions
    ADD CONSTRAINT aws_route53_traffic_policy_versions_traffic_policy_cq_id_fkey FOREIGN KEY (traffic_policy_cq_id) REFERENCES public.aws_route53_traffic_policies(cq_id) ON DELETE CASCADE;


--
-- Name: aws_s3_bucket_cors_rules aws_s3_bucket_cors_rules_bucket_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_cors_rules
    ADD CONSTRAINT aws_s3_bucket_cors_rules_bucket_cq_id_fkey FOREIGN KEY (bucket_cq_id) REFERENCES public.aws_s3_buckets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_s3_bucket_encryption_rules aws_s3_bucket_encryption_rules_bucket_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_encryption_rules
    ADD CONSTRAINT aws_s3_bucket_encryption_rules_bucket_cq_id_fkey FOREIGN KEY (bucket_cq_id) REFERENCES public.aws_s3_buckets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_s3_bucket_grants aws_s3_bucket_grants_bucket_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_grants
    ADD CONSTRAINT aws_s3_bucket_grants_bucket_cq_id_fkey FOREIGN KEY (bucket_cq_id) REFERENCES public.aws_s3_buckets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_s3_bucket_lifecycles aws_s3_bucket_lifecycles_bucket_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_lifecycles
    ADD CONSTRAINT aws_s3_bucket_lifecycles_bucket_cq_id_fkey FOREIGN KEY (bucket_cq_id) REFERENCES public.aws_s3_buckets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_s3_bucket_replication_rules aws_s3_bucket_replication_rules_bucket_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_s3_bucket_replication_rules
    ADD CONSTRAINT aws_s3_bucket_replication_rules_bucket_cq_id_fkey FOREIGN KEY (bucket_cq_id) REFERENCES public.aws_s3_buckets(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_endpoint_configuration_production_variants aws_sagemaker_endpoint_config_endpoint_configuration_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_endpoint_configuration_production_variants
    ADD CONSTRAINT aws_sagemaker_endpoint_config_endpoint_configuration_cq_id_fkey FOREIGN KEY (endpoint_configuration_cq_id) REFERENCES public.aws_sagemaker_endpoint_configurations(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_model_containers aws_sagemaker_model_containers_model_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_model_containers
    ADD CONSTRAINT aws_sagemaker_model_containers_model_cq_id_fkey FOREIGN KEY (model_cq_id) REFERENCES public.aws_sagemaker_models(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_model_vpc_config aws_sagemaker_model_vpc_config_model_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_model_vpc_config
    ADD CONSTRAINT aws_sagemaker_model_vpc_config_model_cq_id_fkey FOREIGN KEY (model_cq_id) REFERENCES public.aws_sagemaker_models(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_algorithm_specification aws_sagemaker_training_job_algorithm_sp_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_algorithm_specification
    ADD CONSTRAINT aws_sagemaker_training_job_algorithm_sp_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_debug_hook_config aws_sagemaker_training_job_debug_hook_c_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_hook_config
    ADD CONSTRAINT aws_sagemaker_training_job_debug_hook_c_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_debug_rule_configurations aws_sagemaker_training_job_debug_rule_c_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_rule_configurations
    ADD CONSTRAINT aws_sagemaker_training_job_debug_rule_c_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_debug_rule_evaluation_statuses aws_sagemaker_training_job_debug_rule_e_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_debug_rule_evaluation_statuses
    ADD CONSTRAINT aws_sagemaker_training_job_debug_rule_e_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_input_data_config aws_sagemaker_training_job_input_data_c_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_input_data_config
    ADD CONSTRAINT aws_sagemaker_training_job_input_data_c_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_profiler_rule_evaluation_statuses aws_sagemaker_training_job_profiler_ru_training_job_cq_id_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_profiler_rule_evaluation_statuses
    ADD CONSTRAINT aws_sagemaker_training_job_profiler_ru_training_job_cq_id_fkey1 FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_sagemaker_training_job_profiler_rule_configurations aws_sagemaker_training_job_profiler_rul_training_job_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_sagemaker_training_job_profiler_rule_configurations
    ADD CONSTRAINT aws_sagemaker_training_job_profiler_rul_training_job_cq_id_fkey FOREIGN KEY (training_job_cq_id) REFERENCES public.aws_sagemaker_training_jobs(cq_id) ON DELETE CASCADE;


--
-- Name: aws_shield_attack_properties aws_shield_attack_properties_attack_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attack_properties
    ADD CONSTRAINT aws_shield_attack_properties_attack_cq_id_fkey FOREIGN KEY (attack_cq_id) REFERENCES public.aws_shield_attacks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_shield_attack_sub_resources aws_shield_attack_sub_resources_attack_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_shield_attack_sub_resources
    ADD CONSTRAINT aws_shield_attack_sub_resources_attack_cq_id_fkey FOREIGN KEY (attack_cq_id) REFERENCES public.aws_shield_attacks(cq_id) ON DELETE CASCADE;


--
-- Name: aws_ssm_instance_compliance_items aws_ssm_instance_compliance_items_instance_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_ssm_instance_compliance_items
    ADD CONSTRAINT aws_ssm_instance_compliance_items_instance_cq_id_fkey FOREIGN KEY (instance_cq_id) REFERENCES public.aws_ssm_instances(cq_id) ON DELETE CASCADE;


--
-- Name: aws_waf_rule_predicates aws_waf_rule_predicates_rule_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_rule_predicates
    ADD CONSTRAINT aws_waf_rule_predicates_rule_cq_id_fkey FOREIGN KEY (rule_cq_id) REFERENCES public.aws_waf_rules(cq_id) ON DELETE CASCADE;


--
-- Name: aws_waf_web_acl_logging_configuration aws_waf_web_acl_logging_configuration_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acl_logging_configuration
    ADD CONSTRAINT aws_waf_web_acl_logging_configuration_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_waf_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_waf_web_acl_rules aws_waf_web_acl_rules_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_waf_web_acl_rules
    ADD CONSTRAINT aws_waf_web_acl_rules_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_waf_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafregional_rate_based_rule_match_predicates aws_wafregional_rate_based_rule_matc_rate_based_rule_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rate_based_rule_match_predicates
    ADD CONSTRAINT aws_wafregional_rate_based_rule_matc_rate_based_rule_cq_id_fkey FOREIGN KEY (rate_based_rule_cq_id) REFERENCES public.aws_wafregional_rate_based_rules(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafregional_rule_predicates aws_wafregional_rule_predicates_rule_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_rule_predicates
    ADD CONSTRAINT aws_wafregional_rule_predicates_rule_cq_id_fkey FOREIGN KEY (rule_cq_id) REFERENCES public.aws_wafregional_rules(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafregional_web_acl_rules aws_wafregional_web_acl_rules_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafregional_web_acl_rules
    ADD CONSTRAINT aws_wafregional_web_acl_rules_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_wafregional_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafv2_web_acl_logging_configuration aws_wafv2_web_acl_logging_configuration_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_logging_configuration
    ADD CONSTRAINT aws_wafv2_web_acl_logging_configuration_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_wafv2_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafv2_web_acl_post_process_firewall_manager_rule_groups aws_wafv2_web_acl_post_process_firewall_mana_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_post_process_firewall_manager_rule_groups
    ADD CONSTRAINT aws_wafv2_web_acl_post_process_firewall_mana_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_wafv2_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups aws_wafv2_web_acl_pre_process_firewall_manag_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups
    ADD CONSTRAINT aws_wafv2_web_acl_pre_process_firewall_manag_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_wafv2_web_acls(cq_id) ON DELETE CASCADE;


--
-- Name: aws_wafv2_web_acl_rules aws_wafv2_web_acl_rules_web_acl_cq_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aws_wafv2_web_acl_rules
    ADD CONSTRAINT aws_wafv2_web_acl_rules_web_acl_cq_id_fkey FOREIGN KEY (web_acl_cq_id) REFERENCES public.aws_wafv2_web_acls(cq_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

