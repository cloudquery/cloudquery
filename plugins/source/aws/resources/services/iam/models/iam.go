package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type AccessKeyWrapper struct {
	types.AccessKeyMetadata
	LastRotated time.Time
}

type Account struct {
	Users                             int32    `json:"users,omitempty"`
	UsersQuota                        int32    `json:"users_quota,omitempty"`
	Groups                            int32    `json:"groups,omitempty"`
	GroupsQuota                       int32    `json:"groups_quota,omitempty"`
	ServerCertificates                int32    `json:"server_certificates,omitempty"`
	ServerCertificatesQuota           int32    `json:"server_certificates_quota,omitempty"`
	UserPolicySizeQuota               int32    `json:"user_policy_size_quota,omitempty"`
	GroupPolicySizeQuota              int32    `json:"group_policy_size_quota,omitempty"`
	GroupsPerUserQuota                int32    `json:"groups_per_user_quota,omitempty"`
	SigningCertificatesPerUserQuota   int32    `json:"signing_certificates_per_user_quota,omitempty"`
	AccessKeysPerUserQuota            int32    `json:"access_keys_per_user_quota,omitempty"`
	MFADevices                        int32    `json:"mfa_devices,omitempty"`
	MFADevicesInUse                   int32    `json:"mfa_devices_in_use,omitempty"`
	AccountMFAEnabled                 bool     `json:"account_mfa_enabled,omitempty"`
	AccountAccessKeysPresent          bool     `json:"account_access_keys_present,omitempty"`
	AccountSigningCertificatesPresent bool     `json:"account_signing_certificates_present,omitempty"`
	AttachedPoliciesPerGroupQuota     int32    `json:"attached_policies_per_group_quota,omitempty"`
	AttachedPoliciesPerRoleQuota      int32    `json:"attached_policies_per_role_quota,omitempty"`
	AttachedPoliciesPerUserQuota      int32    `json:"attached_policies_per_user_quota,omitempty"`
	Policies                          int32    `json:"policies,omitempty"`
	PoliciesQuota                     int32    `json:"policies_quota,omitempty"`
	PolicySizeQuota                   int32    `json:"policy_size_quota,omitempty"`
	PolicyVersionsInUse               int32    `json:"policy_versions_in_use,omitempty"`
	PolicyVersionsInUseQuota          int32    `json:"policy_versions_in_use_quota,omitempty"`
	VersionsPerPolicyQuota            int32    `json:"versions_per_policy_quota,omitempty"`
	GlobalEndpointTokenVersion        int32    `json:"global_endpoint_token_version,omitempty"`
	Aliases                           []string `json:"aliases,omitempty"`
}

type DateTime struct {
	*time.Time
}

func (d *DateTime) UnmarshalCSV(val string) (err error) {
	switch val {
	case "N/A", "not_supported", "no_information":
		d.Time = nil
		return nil
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	d.Time = &t
	return nil
}

type CredentialReportEntry struct {
	User                      string   `csv:"user"`
	Arn                       string   `csv:"arn"`
	UserCreationTime          DateTime `csv:"user_creation_time"`
	PasswordStatus            string   `csv:"password_enabled"`
	PasswordLastChanged       DateTime `csv:"password_last_changed"`
	PasswordNextRotation      DateTime `csv:"password_next_rotation"`
	MfaActive                 bool     `csv:"mfa_active"`
	AccessKey1Active          bool     `csv:"access_key_1_active"`
	AccessKey2Active          bool     `csv:"access_key_2_active"`
	AccessKey1LastRotated     DateTime `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated     DateTime `csv:"access_key_2_last_rotated"`
	Cert1Active               bool     `csv:"cert_1_active"`
	Cert2Active               bool     `csv:"cert_2_active"`
	Cert1LastRotated          DateTime `csv:"cert_1_last_rotated"`
	Cert2LastRotated          DateTime `csv:"cert_2_last_rotated"`
	AccessKey1LastUsedDate    DateTime `csv:"access_key_1_last_used_date"`
	AccessKey1LastUsedRegion  string   `csv:"access_key_1_last_used_region"`
	AccessKey1LastUsedService string   `csv:"access_key_1_last_used_service"`
	AccessKey2LastUsedDate    DateTime `csv:"access_key_2_last_used_date"`
	AccessKey2LastUsedRegion  string   `csv:"access_key_2_last_used_region"`
	AccessKey2LastUsedService string   `csv:"access_key_2_last_used_service"`
	PasswordLastUsed          DateTime `csv:"password_last_used"`
}

type IamOpenIdIdentityProviderWrapper struct {
	*iam.GetOpenIDConnectProviderOutput
	Arn string
}

type IAMSAMLIdentityProviderWrapper struct {
	*iam.GetSAMLProviderOutput
	Arn string
}

type PasswordPolicyWrapper struct {
	types.PasswordPolicy
	PolicyExists bool
}
