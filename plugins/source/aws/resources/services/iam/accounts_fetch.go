package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

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

func fetchIamAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().IAM

	summary, err := svc.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
	if err != nil {
		return err
	}
	var accSummary Account
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", WeaklyTypedInput: true, Result: &accSummary})
	if err != nil {
		return err
	}
	if err := decoder.Decode(summary.SummaryMap); err != nil {
		return err
	}
	config := iam.ListAccountAliasesInput{}
	for {
		response, err := svc.ListAccountAliases(ctx, &config)
		if err != nil {
			return err
		}

		accSummary.Aliases = append(accSummary.Aliases, response.AccountAliases...)

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	res <- accSummary
	return nil
}
