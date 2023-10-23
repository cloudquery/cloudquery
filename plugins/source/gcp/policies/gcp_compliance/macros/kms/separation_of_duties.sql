{% macro kms_separation_of_duties(framework, check_id) %}
    select
        member as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure that Separation of duties is enforced while assigning KMS related roles to users (Automated)'
        as title,
        project_id as project_id,
        case
            when
                member like 'user:%'
                and 'roles/cloudkms.admin' = any(roles)
                and roles && array[
                    'roles/cloudkms.cryptoKeyEncrypterDecrypter',
                    'roles/cloudkms.cryptoKeyEncrypter',
                    'roles/cloudkms.cryptoKeyDecrypter'
                ]
            then 'fail'
            else 'pass'
        end as status
    from {{ ref('gcp_compliance__member_with_roles') }}
{% endmacro %}
