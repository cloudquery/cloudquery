create or replace view view_aws_apigateway_method_settings as
with apigateway_rest_api_stages as(
  select
    *,
    method_settings::text = '{}'::text as is_empty_method_settings
  from aws_apigateway_rest_api_stages
)
select
     s.arn,
     s.rest_api_arn,
	 s.stage_name,
	 s.tracing_enabled as stage_data_trace_enabled,
     s.cache_cluster_enabled as stage_caching_enabled,
     s.web_acl_arn as waf,
     s.client_certificate_id as cert,
case when (s.is_empty_method_settings = false)
        then (select key from JSONB_EACH_TEXT(s.method_settings))
        else '/*/'
    end as method,
    case when (s.is_empty_method_settings = false)
        then (select (value::JSON -> 'DataTraceEnabled')::TEXT::BOOLEAN from JSONB_EACH_TEXT(s.method_settings))
        else false
    end as data_trace_enabled,
    case when (s.is_empty_method_settings = false)
        then (select (value::JSON -> 'CachingEnabled')::TEXT::BOOLEAN from JSONB_EACH_TEXT(s.method_settings))
        else false
    end as caching_enabled,
    case when (s.is_empty_method_settings = false)
        then (select (value::JSON -> 'CacheDataEncrypted')::TEXT::BOOLEAN from JSONB_EACH_TEXT(s.method_settings))
        else false
    end as cache_data_encrypted,
    case when (s.is_empty_method_settings = false)
        then (select (value::JSON -> 'LoggingLevel')::TEXT from JSONB_EACH_TEXT(s.method_settings))
        else '"OFF"'
   end as logging_level,
     r.account_id

from apigateway_rest_api_stages s, aws_apigateway_rest_apis r
 where s.rest_api_arn=r.arn
