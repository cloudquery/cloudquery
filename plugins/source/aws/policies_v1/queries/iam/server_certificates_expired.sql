insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Ensure that all the expired SSL/TLS certificates stored in AWS IAM are removed (Automated)' as title,
    account_id,
    arn AS resource_id,
    case when
                 expiration < NOW() AT TIME ZONE 'UTC'
             then 'fail'
         else 'pass'
        end as status
FROM aws_iam_server_certificates