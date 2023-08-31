INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)

WITH member_with_roles AS(
    SELECT
        project_id,
        member,
        array_agg(role) AS roles
    FROM gcp_project_policy_members
    GROUP BY member, project_id
)
SELECT
    member                                                                                                AS resource_id,
    :'execution_time'::timestamp                                                                          AS execution_time,
    :'framework'                                                                                          AS framework,
    :'check_id'                                                                                           AS check_id,
    'Ensure that Separation of duties is enforced while assigning KMS related roles to users (Automated)' AS title,
    project_id                                                                                            AS project_id,
    CASE
        WHEN
            member LIKE 'user:%'
            AND 'roles/cloudkms.admin' = ANY(roles)
            AND roles && ARRAY[
                'roles/cloudkms.cryptoKeyEncrypterDecrypter',
                'roles/cloudkms.cryptoKeyEncrypter',
                'roles/cloudkms.cryptoKeyDecrypter'
            ]
        THEN 'fail'
        ELSE 'pass'
    END                                                                                                   AS status
FROM member_with_roles;
