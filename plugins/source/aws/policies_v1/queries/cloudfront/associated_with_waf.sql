insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'API Gateway should be associated with an AWS WAF web ACL' as title,
    account_id,
    arn as resource_id,
    case
        when distribution_config->>'WebACLId' = '' then 'fail'
        else 'pass'
    end as status
from aws_cloudfront_distributions
