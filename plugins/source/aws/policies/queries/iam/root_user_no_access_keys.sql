insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure no root account access key exists (Scored)',
  account_id,
  user_arn AS resource_id,
  case when
    user_name IN ('<root>', '<root_account>')
    then 'fail'
    else 'pass'
  end
from aws_iam_user_access_keys;
