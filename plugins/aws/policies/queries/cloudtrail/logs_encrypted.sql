insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudTrail should have encryption at rest enabled' as title,
    account_id,
    arn as resource_id,
    case
        when kms_key_id is NULL then 'fail'
        else 'pass'
    end as status
FROM aws_cloudtrail_trails
