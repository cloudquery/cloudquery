insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure rotation for customer created CMKs is enabled (Scored)',
  account_id,
  arn,
  case when
    rotation_enabled is FALSE and key_manager = 'CUSTOMER'
    then 'fail'
    else 'pass'
  end
from aws_kms_keys
