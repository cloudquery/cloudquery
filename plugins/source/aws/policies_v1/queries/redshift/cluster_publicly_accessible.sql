insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon Redshift clusters should prohibit public access' as title,
    account_id,
    arn AS resource_id,
    case when publicly_accessible is TRUE then 'fail' else 'pass' end as status
from aws_redshift_clusters
