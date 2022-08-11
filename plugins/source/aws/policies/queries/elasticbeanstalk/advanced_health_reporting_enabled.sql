insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Elastic Beanstalk environments should have enhanced health reporting enabled' as title,
    account_id,
    arn as resource_id,
    case when
        health_status is null
        or health is null
        then 'fail'
        else 'pass'
    end as status
from aws_elasticbeanstalk_environments
