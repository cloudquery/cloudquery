insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'AMIs should require IMDSv2' as title,
  account_id,
  arn              as resource_id,
  case when
    imds_support is distinct from 'v2.0'
    then 'fail'
    else 'pass'
  end as status
from aws_ec2_images
