{% macro iam_service_account_keys_not_rotated(framework, check_id) %}
    select distinct
        gisa.name as resource_id,
        gisa._cq_sync_time as execution_time,
        '{{framework}}' as framework,
        '{{check_id}}' as check_id,
        'Ensure user-managed/external keys for service accounts are rotated every 90 days or less (Automated)'
        as title,
        gisa.project_id as project_id,
        case
            when
                gisa.email like '%iam.gserviceaccount.com'
                and gisak.valid_after_time::timestamp <= (now() - interval '90' day)
            then 'fail'
            else 'pass'
        end as status
    from gcp_iam_service_accounts gisa
    join
        gcp_iam_service_account_keys gisak
        on gisa.project_id = gisak.project_id
        and gisa.unique_id = gisak.service_account_unique_id
{% endmacro %}
