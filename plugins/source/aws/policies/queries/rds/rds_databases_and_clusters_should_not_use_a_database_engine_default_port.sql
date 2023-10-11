insert into aws_policy_results
(
    select
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS databases and clusters should not use a database engine default port' as title,
    account_id,
    arn AS resource_id,
    case when
        (engine in ('aurora', 'aurora-mysql', 'mysql') and port = 3306) or (engine like '%postgres%' and port = 5432)
    then 'fail' else 'pass' end as status
    from aws_rds_clusters
)
union
(
    select
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS databases and clusters should not use a database engine default port' as title,
    account_id,
    arn AS resource_id,
    case when
        (
            engine in ('aurora', 'aurora-mysql', 'mariadb', 'mysql')
            and (endpoint ->> 'Port')::integer = 3306
        )
        or (engine like '%postgres%' and (endpoint ->> 'Port')::integer = 5432)
        or (engine like '%oracle%' and (endpoint ->> 'Port')::integer = 1521)
        or (engine like '%sqlserver%' and (endpoint ->> 'Port')::integer = 1433)
    then 'fail' else 'pass' end as status
    from aws_rds_instances
)
