insert into aws_policy_results
select :'execution_time',
       :'framework',
       :'check_id',
       'Ensure there is only one active access key available for any single IAM user (Automated)',
       account_id,
       user_arn,
       case
           when
                   count(*) > 1
               then 'fail'
           else 'pass'
           end
from aws_iam_user_access_keys
where status = 'Active'
group by account_id, user_arn