{% macro kms_separation_of_duties(framework, check_id) %}
SELECT
    member                                                                                                AS resource_id,
    _cq_sync_time                                                                          AS sync_time,
    '{{framework}}'                                                                                          AS framework,
    '{{check_id}}'                                                                                           AS check_id,
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
FROM {{ ref('gcp_compliance__member_with_roles') }}
{% endmacro %}