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
	Users                             int32
	UsersQuota                        int32
	Groups                            int32
	GroupsQuota                       int32
	ServerCertificates                int32
	ServerCertificatesQuota           int32
	UserPolicySizeQuota               int32
	GroupPolicySizeQuota              int32
	GroupsPerUserQuota                int32
	SigningCertificatesPerUserQuota   int32
	AccessKeysPerUserQuota            int32
	MFADevices                        int32
	MFADevicesInUse                   int32
	AccountMFAEnabled                 bool
	AccountAccessKeysPresent          bool
	AccountSigningCertificatesPresent bool
	AttachedPoliciesPerGroupQuota     int32
	AttachedPoliciesPerRoleQuota      int32
	AttachedPoliciesPerUserQuota      int32
	Policies                          int32
	PoliciesQuota                     int32
	PolicySizeQuota                   int32
	PolicyVersionsInUse               int32
	PolicyVersionsInUseQuota          int32
	VersionsPerPolicyQuota            int32
	GlobalEndpointTokenVersion        int32
	Aliases                           []string
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
	Arn  string
	Tags map[string]string
}

type PasswordPolicyWrapper struct {
	types.PasswordPolicy
	PolicyExists bool
}
