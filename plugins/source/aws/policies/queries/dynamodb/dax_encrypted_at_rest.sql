insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'DynamoDB Accelerator (DAX) clusters should be encrypted at rest' as title,
    account_id,
    arn as resource_id,
  case when
    sse_description->>'Status' is distinct from 'ENABLED'
    then 'fail'
    else 'pass'
  end as status
from aws_dax_clusters
