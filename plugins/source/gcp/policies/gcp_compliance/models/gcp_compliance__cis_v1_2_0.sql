with
    aggregated as (
        {{ iam_managed_service_account_keys('cis_v1.2.0', '1.4') }}
        union
        {{ iam_service_account_admin_priv('cis_v1.2.0', '1.5') }}
        union
        {{ iam_users_with_service_account_token_creator_role('cis_v1.2.0', '1.6') }}
        union
        {{ iam_service_account_keys_not_rotated('cis_v1.2.0', '1.7') }}
        union
        {{ iam_separation_of_duties('cis_v1.2.0', '1.8') }}
        union
        {{ kms_publicly_accessible('cis_v1.2.0', '1.9') }}
        union
        {{ kms_keys_not_rotated_within_90_days('cis_v1.2.0', '1.10') }}
        union
        {{ kms_separation_of_duties('cis_v1.2.0', '1.11') }}

    )
select *
from aggregated
