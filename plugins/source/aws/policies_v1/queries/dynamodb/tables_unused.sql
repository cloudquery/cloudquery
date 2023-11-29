insert into aws_policy_results
select :'execution_time'              as execution_time,
       :'framework'                   as framework,
       :'check_id'                    as check_id,
       'DynamoDB table with no items' as title,
       account_id,
       arn                            as resource_id,
       'fail'                         as status
from aws_dynamodb_tables
where item_count = 0