insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon Redshift should have automatic upgrades to major versions enabled' as title,
    account_id,
    arn as resource_id,
    case when
        allow_version_upgrade is FALSE or allow_version_upgrade is null
    then 'fail' else 'pass' end as status
from aws_redshift_clusters
