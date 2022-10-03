insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have origin access identity enabled' as title,
    account_id,
    arn as resource_id,
    case
        when o->>'DomainName' like '%s3.amazonaws.com' and o->'S3OriginConfig'->>'OriginAccessIdentity' = '' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions, JSONB_ARRAY_ELEMENTS(distribution_config->'Origins'->'Items') o
