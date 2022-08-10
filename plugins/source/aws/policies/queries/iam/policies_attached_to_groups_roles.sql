insert into aws_policy_results
select distinct
    :'execution_time'::timestamp,
    :'framework',
    :'check_id',
    'IAM users should not have IAM policies attached',
    aws_iam_users.account_id,
    arn AS resource_id,
    case when
        aws_iam_user_attached_policies.user_cq_id is not null
    then 'fail' else 'pass' end as status
from aws_iam_users
left join aws_iam_user_attached_policies on aws_iam_users.cq_id = aws_iam_user_attached_policies.user_cq_id
