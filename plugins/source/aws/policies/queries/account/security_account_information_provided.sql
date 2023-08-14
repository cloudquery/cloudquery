insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Security contact information should be provided for an AWS account' as title,
  aws_iam_accounts.account_id,
  case when
    alternate_contact_type is null
    then 'fail'
    else 'pass'
  end as status
FROM aws_iam_accounts
LEFT JOIN (
	SELECT * from aws_account_alternate_contacts
WHERE alternate_contact_type = 'SECURITY'
) as account_security_contacts
ON aws_iam_accounts.account_id = account_security_contacts.account_id