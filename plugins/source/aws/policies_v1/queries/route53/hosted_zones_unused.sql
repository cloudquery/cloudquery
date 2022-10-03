insert into aws_policy_results
select :'execution_time'              as execution_time,
       :'framework'                   as framework,
       :'check_id'                    as check_id,
       'Unused Route 53 hosted zones' as title,
       account_id,
       arn                            as resource_id,
       'fail'                         as status
from aws_route53_hosted_zones
where resource_record_set_count = 0