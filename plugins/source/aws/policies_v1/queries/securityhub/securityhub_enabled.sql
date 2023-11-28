insert into aws_policy_results
with enabled_securityhub_regions as (select account_id, region from aws_securityhub_hubs)

select :'execution_time'::timestamp    as execution_time,
       :'framework'                    as framework,
       :'check_id'                     as check_id,
       'SecurityHub should be enabled' AS title,
       r.account_id,
       r.region                        AS resource_id,
       case
           when
               r.enabled = TRUE AND e.region is null
               then 'fail'
           else 'pass' end             AS status
from aws_regions r
         left join enabled_securityhub_regions e on e.region = r.region AND e.account_id = r.account_id
