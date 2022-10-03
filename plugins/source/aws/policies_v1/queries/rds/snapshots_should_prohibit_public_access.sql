insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS snapshots should be private' as title,
    account_id,
    arn AS resource_id,
    case when
         (attrs ->> 'AttributeName' is not distinct from 'restore')
         and (attrs -> 'AttributeValues')::jsonb ? 'all'
    then 'fail' else 'pass' end as status
from aws_rds_cluster_snapshots, jsonb_array_elements(attributes) as attrs
