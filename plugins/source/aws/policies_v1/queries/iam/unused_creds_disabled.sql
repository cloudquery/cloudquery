insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure credentials unused for 90 days or greater are disabled (Scored)',
  split_part(r.arn, ':', 5) as account_id,
  r.arn,
  case when
      (r.password_status IN ('TRUE', 'true') and r.password_last_used < (now() - '90 days'::INTERVAL)
        or (k.last_used < (now() - '90 days'::INTERVAL)))
      then 'fail'
      else 'pass'
  end
from aws_iam_credential_reports r
left join aws_iam_user_access_keys k on k.user_arn = r.arn
