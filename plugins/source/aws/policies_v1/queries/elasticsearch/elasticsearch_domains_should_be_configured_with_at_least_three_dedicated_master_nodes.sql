insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should be configured with at least three dedicated master nodes' as title,
  account_id,
  arn as resource_id,
  case when
    (elasticsearch_cluster_config->>'DedicatedMasterEnabled')::boolean is not TRUE
    or (elasticsearch_cluster_config->>'DedicatedMasterCount')::integer is null
    or (elasticsearch_cluster_config->>'DedicatedMasterCount')::integer < 3
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
