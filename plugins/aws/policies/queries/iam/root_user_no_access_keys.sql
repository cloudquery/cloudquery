insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure no root account access key exists (Scored)',
  account_id,
  arn,
  case when
    user_name = '<root>'
    then 'fail'
    else 'pass'
  end
from aws_iam_users
inner join
    aws_iam_user_access_keys on
        aws_iam_users.cq_id = aws_iam_user_access_keys.user_cq_id
