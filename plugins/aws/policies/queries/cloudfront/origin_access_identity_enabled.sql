insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have origin access identity enabled' as title,
    d.account_id,
    d.arn as resource_id,
    case
        when o.domain_name like '%s3.amazonaws.com' and o.s3_origin_config_origin_access_identity = '' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distribution_origins o
inner join aws_cloudfront_distributions d on d.cq_id = o.distribution_cq_id
