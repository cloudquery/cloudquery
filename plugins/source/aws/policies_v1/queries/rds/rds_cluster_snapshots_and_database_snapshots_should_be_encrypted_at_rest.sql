insert into aws_policy_results
(
select
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS cluster snapshots and database snapshots should be encrypted at rest' as title,
    account_id,
    arn AS resource_id,
    case when storage_encrypted is not TRUE then 'fail' else 'pass' end as status
from aws_rds_cluster_snapshots
)
union
(
    select
        :'execution_time'::timestamp as execution_time,
        :'framework' as framework,
        :'check_id' as check_id,
        'RDS cluster snapshots and database snapshots should be encrypted at rest' as title,
        account_id,
        arn AS resource_id,
        case when encrypted is not TRUE then 'fail' else 'pass' end as status
    from aws_rds_db_snapshots
)
