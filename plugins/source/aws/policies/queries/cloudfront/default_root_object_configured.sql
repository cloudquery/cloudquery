insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have a default root object configured' as title,
    account_id,
    arn as resource_id,
    case
        when distribution_config->>'DefaultRootObject' = '' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions
