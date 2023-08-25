insert into aws_policy_results

select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM policies should not allow full ''*'' administrative privileges' as title,
    account_id,
    arn AS resource_id,
    CASE SUM(attachment_count)
        WHEN 0 THEN 'pass'
        ELSE 'fail'
    END AS status
FROM view_aws_iam_policy_statements
WHERE effect = 'Allow' AND resources ? '*' AND actions ?| array['*', '*:*']
GROUP BY account_id, arn
