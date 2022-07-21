insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should have at least three data nodes' as title,
  account_id,
  arn as resource_id,
  case when
    not cluster_zone_awareness_enabled
    or cluster_instance_count is null
    or cluster_instance_count < 3
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
