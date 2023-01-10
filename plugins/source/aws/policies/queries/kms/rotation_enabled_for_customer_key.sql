insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure rotation for customer created custom master keys is enabled (Scored)' as title,
  account_id,
  arn,
  case when
    rotation_enabled is FALSE and key_manager = 'CUSTOMER'
    then 'fail'
    else 'pass'
  end
from aws_kms_keys
