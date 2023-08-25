insert into aws_policy_results

SELECT
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM customer managed policies that you create should not allow wildcard actions for services' as title,
    account_id,
    arn,
    CASE SUM(attachment_count)
        WHEN 0 THEN 'pass'
        ELSE 'fail'
    END AS status
FROM view_aws_iam_policy_statements s
WHERE effect = 'Allow' AND resources::text like '%"*"%' AND (actions::text LIKE '%:*%' OR not_actions::text LIKE '%:*%')
GROUP BY account_id, arn
