insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all CloudFront distributions' AS title,
    account_id,
    arn as resource_id,
    'fail' as status
from
    aws_cloudfront_distributions
