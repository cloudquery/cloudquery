insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    '' as title,
    account_id,
    arn as resource_id,
    case when
        kms_key_id is null
    then 'fail' else 'pass' end as status
from aws_cloudtrail_trails
