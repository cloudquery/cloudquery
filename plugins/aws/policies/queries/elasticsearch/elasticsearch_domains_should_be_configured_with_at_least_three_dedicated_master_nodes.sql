insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should be configured with at least three dedicated master nodes' as title,
  account_id,
  arn as resource_id,
  case when
    cluster_dedicated_master_enabled is not TRUE
    or cluster_dedicated_master_count is null
    or cluster_dedicated_master_count < 3
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
