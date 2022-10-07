insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should have logging enabled' as title,
    account_id,
    arn as resource_id,
    case
        when (distribution_config->'Logging'->>'Enabled')::boolean is distinct from true then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions
