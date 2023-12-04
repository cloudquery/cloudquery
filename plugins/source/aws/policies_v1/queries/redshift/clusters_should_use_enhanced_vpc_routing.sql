insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon Redshift clusters should use enhanced VPC routing' as title,
    account_id,
    arn as resource_id,
    case when
        enhanced_vpc_routing is FALSE or enhanced_vpc_routing is null
    then 'fail' else 'pass' end as status
from aws_redshift_clusters
