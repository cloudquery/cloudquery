insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Avoid the use of "root" account. Show used in last 30 days (Scored)' as title,
  account_id,
  arn as resource_id,
  case when
    user_name = '<root_account>' and password_last_used > (now() - '30 days'::INTERVAL)
    then 'fail'
    else 'pass'
  end as status
from aws_iam_users
