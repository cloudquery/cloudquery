create or replace view view_aws_apigateway_method_settings as
select
    s.arn,
    s.rest_api_arn,
    s.stage_name,
    s.tracing_enabled as stage_data_trace_enabled,
    s.cache_cluster_enabled as stage_caching_enabled,
    s.web_acl_arn as waf,
    s.client_certificate_id as cert,
    key as method,
    (
        value::JSON -> 'DataTraceEnabled'
    )::TEXT::BOOLEAN as data_trace_enabled,
    (value::JSON -> 'CachingEnabled')::TEXT::BOOLEAN as caching_enabled,
    (
        value::JSON -> 'CacheDataEncrypted'
    )::TEXT::BOOLEAN as cache_data_encrypted,
    (value::JSON -> 'LoggingLevel')::TEXT as logging_level,
    r.account_id
from aws_apigateway_rest_api_stages s, aws_apigateway_rest_apis r,
    JSONB_EACH_TEXT(s.method_settings)
where s.rest_api_arn=r.arn
