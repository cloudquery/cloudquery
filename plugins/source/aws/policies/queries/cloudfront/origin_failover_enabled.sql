with origin_groups as ( select acd.arn, distribution_config->'OriginGroups'->'Items' as ogs from aws_cloudfront_distributions acd),
     oids as (
select distinct
    account_id,
    acd.arn as resource_id,
    case
            when o.ogs = 'null' or
            o.ogs->'Members'->'Items' = 'null' or
            jsonb_array_length(o.ogs->'Members'->'Items') = 0  then 'fail'
    else 'pass'
    end as status
from aws_cloudfront_distributions acd
    left join origin_groups o on o.arn = acd.arn
)
insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have origin failover configured' as title,
    account_id,
    resource_id,
    status
from oids
