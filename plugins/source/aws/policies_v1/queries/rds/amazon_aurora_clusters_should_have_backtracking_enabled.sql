insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon Aurora clusters should have backtracking enabled' as title,
    account_id,
    arn AS resource_id,
    case when backtrack_window is null then 'fail' else 'pass' end as status
from aws_rds_clusters
where
    engine in ('aurora', 'aurora-mysql', 'mysql')
