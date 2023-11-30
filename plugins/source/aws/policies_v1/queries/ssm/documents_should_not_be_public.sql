insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'SSM documents should not be public' as title,
    account_id,
    arn as resource_id,
    case when
        'all' = ANY(ARRAY(SELECT JSONB_ARRAY_ELEMENTS_TEXT(p->'AccountIds'))) 
    then 'fail' else 'pass' end as status
from aws_ssm_documents, jsonb_array_elements(aws_ssm_documents.permissions) p
where owner in (select account_id from aws_iam_accounts)
