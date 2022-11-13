insert into aws_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elastic Beanstalk managed platform updates should be enabled' as title,
  account_id,
  application_arn as resource_id,
  case when
    s->>'OptionName' = 'ManagedActionsEnabled' AND (s->>'Value')::boolean = true is distinct from true
    then 'fail'
    else 'pass'
  end as status
from aws_elasticbeanstalk_configuration_settings, jsonb_array_elements(aws_elasticbeanstalk_configuration_settings.option_settings) as s
