insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager' as title,
  lb.account_id,
  lb.arn as resource_id,
  case when
    li->'Listener'->>'Protocol' = 'HTTPS' and aws_acm_certificates.arn is null
    then 'fail'
    else 'pass'
  end as status
from aws_elbv1_load_balancers lb, jsonb_array_elements(lb.listener_descriptions) as li
left join
    aws_acm_certificates on
        aws_acm_certificates.arn = li->'Listener'->>'SSLCertificateId'
