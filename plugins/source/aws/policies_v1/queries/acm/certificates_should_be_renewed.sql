insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'certificate has less than 30 days to be renewed' as title,
  account_id,
  arn AS resource_id,
  case when
    not_after < NOW() AT TIME ZONE 'UTC' + INTERVAL '30' DAY
    then 'fail'
    else 'pass'
  end as status
FROM aws_acm_certificates
