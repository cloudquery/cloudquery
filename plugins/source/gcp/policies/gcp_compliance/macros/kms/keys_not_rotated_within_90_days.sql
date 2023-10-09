{% macro kms_keys_not_rotated_within_90_days(framework, check_id) %}
    select
        "name" as resource_id,
        _cq_sync_time as sync_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure KMS encryption keys are rotated within a period of 90 days (Automated)'
        as title,
        project_id as project_id,
        case
            when
                (
                    make_interval(secs => rotation_period / 1000000000.0)
                    > make_interval(days => 90)
                )
                or next_rotation_time is null
                or date_part('day', current_date - next_rotation_time::timestamp) > 90
            then 'fail'
            else 'pass'
        end as status
    from gcp_kms_crypto_keys
{% endmacro %}
