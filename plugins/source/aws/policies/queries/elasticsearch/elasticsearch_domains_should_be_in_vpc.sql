insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should be in a VPC' as title,
  account_id,
  arn as resource_id,
  case when
    vpc_options->>'VPCId' is null
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
