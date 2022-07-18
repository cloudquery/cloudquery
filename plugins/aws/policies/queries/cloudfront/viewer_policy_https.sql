insert into aws_policy_results
with data as (
    select distinct distribution_cq_id
    from aws_cloudfront_distribution_cache_behaviors
    where viewer_protocol_policy = 'allow-all'
)
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should require encryption in transit' as title,
    d.account_id,
    d.arn as resource_id,
    case
        when data.distribution_cq_id is not null
            or d.cache_behavior_viewer_protocol_policy = 'allow-all' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions d
left join data on data.distribution_cq_id = d.cq_id
