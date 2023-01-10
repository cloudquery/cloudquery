insert into aws_policy_results
select :'execution_time' as execution_time,
       :'framework' as framework,
       :'check_id' as check_id,
       'Ensure credentials unused for 45 days or greater are disabled (Automated)' as title,
       split_part(r.arn, ':', 5) as account_id,
       r.arn,
       case
           when
                   (r.password_status in ('TRUE', 'true')
                        and r.password_last_used < (now() - '45 days'::interval)
                       or (r.password_status in ('TRUE', 'true')
                           and r.password_last_used is null
                           and r.password_last_changed < (now() - '45 days'::interval))
                       or (k.last_used < (now() - '45 days'::interval)))
                   or (r.access_key1_active
                   and r.access_key_1_last_used_date < (now() - '45 days'::interval))
                   or (r.access_key1_active
                   and r.access_key_1_last_used_date is null
                   and access_key_1_last_rotated < (now() - '45 days'::interval))
                   or (r.access_key2_active
                   and r.access_key_2_last_used_date < (now() - '45 days'::interval))
                   or (r.access_key2_active
                   and r.access_key_2_last_used_date is null
                   and access_key_2_last_rotated < (now() - '45 days'::interval))
               then 'fail'
           else 'pass'
           end
from aws_iam_credential_reports r
         left join aws_iam_user_access_keys k on
        k.user_arn = r.arn