insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have origin failover configured' as title,
    d.account_id,
    d.arn as resource_id,
    case
        when members_origin_ids is null then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distribution_origin_groups o
inner join aws_cloudfront_distributions d on d.cq_id = o.distribution_cq_id
