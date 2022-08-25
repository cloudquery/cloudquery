insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure no root account access key exists (Scored)',
  account_id,
  arn,
  case when
    keys.count > 0 
    then 'fail'
    else 'pass'
  end
from aws_iam_users
left join (select count(*), user_cq_id from aws_iam_user_access_keys group by user_cq_id) keys on aws_iam_users.cq_id =  keys.user_cq_id
where aws_iam_users.arn = 'arn:aws:iam::'||account_id||':root' 