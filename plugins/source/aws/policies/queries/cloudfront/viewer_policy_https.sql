insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should require encryption in transit' as title,
    account_id,
    arn as resource_id,
    case
        when d->>'ViewerProtocolPolicy' is not null
            or d->>'ViewerProtocolPolicy' = 'allow-all' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions
left join jsonb_array_elements(distribution_config->'CacheBehaviors'->'Items') as d on d->>'ViewerProtocolPolicy' = 'allow-all'
