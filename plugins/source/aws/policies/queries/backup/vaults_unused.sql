insert into aws_policy_results
with point as (select distinct vault_arn from aws_backup_vault_recovery_points)
select :'execution_time'                as execution_time,
       :'framework'                     as framework,
       :'check_id'                      as check_id,
       'Vaults with no recovery points' as title,
       vault.account_id,
       vault.arn                        as resource_id,
       'fail'                           as status
from aws_backup_vaults vault
         left join point on point.vault_arn = vault.arn
where point.vault_arn is null;