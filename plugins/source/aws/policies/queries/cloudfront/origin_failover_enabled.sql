with oids as (
select distinct
    account_id,
    arn as resource_id,
    case
        when m->>'OriginId' is null then 'fail'
    else 'pass'
    end as status
from aws_cloudfront_distributions
  left join jsonb_array_elements(distribution_config->'OriginGroups'->'Items') o on true
  left join jsonb_array_elements(o->'Members'->'Items') m on true
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
