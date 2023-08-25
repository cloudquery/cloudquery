CREATE OR REPLACE VIEW view_aws_iam_policy_statements AS

WITH statement_array AS (
    SELECT
        aws_iam_policies.account_id,
        arn,
        attachment_count,
        CASE jsonb_typeof(document -> 'Statement')
            WHEN 'object' THEN jsonb_build_array(document -> 'Statement')
            WHEN 'array' THEN document -> 'Statement'
        END AS statements
    FROM aws_iam_policies JOIN aws_iam_policy_versions ON aws_iam_policies._cq_id = aws_iam_policy_versions._cq_parent_id
), statements AS (
    SELECT
        account_id,
        arn,
        attachment_count,
        statement ->> 'Effect' AS effect,
        CASE jsonb_typeof(statement -> 'Action')
            WHEN 'string' THEN jsonb_build_array(statement -> 'Action')
            WHEN 'array' THEN statement -> 'Action'
            ELSE jsonb_build_array()
        END AS actions,
        CASE jsonb_typeof(statement -> 'NotAction')
            WHEN 'string' THEN jsonb_build_array(statement -> 'NotAction')
            WHEN 'array' THEN statement -> 'NotAction'
            ELSE jsonb_build_array()
        END AS not_actions,
        CASE jsonb_typeof(statement -> 'Resource')
            WHEN 'string' THEN jsonb_build_array(statement -> 'Resource')
            WHEN 'array' THEN statement -> 'Resource'
            ELSE jsonb_build_array()
        END AS resources
    FROM statement_array s, jsonb_array_elements(s.statements) AS statement
)
SELECT * FROM statements;