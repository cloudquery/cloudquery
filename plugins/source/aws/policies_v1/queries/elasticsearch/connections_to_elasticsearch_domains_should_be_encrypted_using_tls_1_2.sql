insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Connections to Elasticsearch domains should be encrypted using TLS 1.2' as title,
  account_id,
  arn as resource_id,
  case when
    domain_endpoint_options->>'TLSSecurityPolicy' is distinct from 'Policy-Min-TLS-1-2-2019-07'
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
