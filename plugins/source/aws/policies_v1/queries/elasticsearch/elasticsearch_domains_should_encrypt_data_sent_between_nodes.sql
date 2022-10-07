insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should encrypt data sent between nodes' as title,
  account_id,
  arn as resource_id,
  case when
        (node_to_node_encryption_options->>'Enabled')::boolean is not true
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
