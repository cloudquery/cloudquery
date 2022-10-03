insert into aws_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elastic Beanstalk managed platform updates should be enabled' as title,
  aws_elasticbeanstalk_environments.account_id,
  aws_elasticbeanstalk_configuration_settings.application_arn as resource_id,
  case when
    option_name = 'ManagedActionsEnabled'
    and value::boolean is distinct
    from true
    then 'fail'
    else 'pass'
  end as status
from aws_elasticbeanstalk_configuration_setting_options
    left join
        aws_elasticbeanstalk_configuration_settings on
            aws_elasticbeanstalk_configuration_setting_options.configuration_setting_cq_id =
            aws_elasticbeanstalk_configuration_settings.cq_id
    left join
     aws_elasticbeanstalk_environments on
             aws_elasticbeanstalk_configuration_settings.environment_cq_id =
             aws_elasticbeanstalk_environments.cq_id
