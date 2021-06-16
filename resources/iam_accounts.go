package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/mitchellh/mapstructure"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:         "aws_accounts",
		Description:  "Information about IAM entity usage and IAM quotas in the AWS account.",
		Resolver:     fetchAccountSummary,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "users",
				Description: "Current number of users.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "users_quota",
				Description: "Maximum allowed number of users.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "groups",
				Description: "Current number of groups.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "groups_quota",
				Description: "Maximum allowed number of groups.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "server_certificates",
				Description: "Current number of server certificates",
				Type:        schema.TypeInt,
			},
			{
				Name:        "server_certificates_quota",
				Description: "Maximum allowed number of server certificates.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "user_policy_size_quota",
				Description: "Maximum allowed policies per user",
				Type:        schema.TypeInt,
			},
			{
				Name:        "group_policy_size_quota",
				Description: "Maximum allowed policies per group",
				Type:        schema.TypeInt,
			},
			{
				Name:        "groups_per_user_quota",
				Description: "Maximum allowed groups per user",
				Type:        schema.TypeInt,
			},
			{
				Name:        "signing_certificates_per_user_quota",
				Description: "Maximum allowed server certificates per user",
				Type:        schema.TypeInt,
			},
			{
				Name:        "access_keys_per_user_quota",
				Description: "Maximum allowed quota of access keys per user.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "mfa_devices",
				Description: "Number of MFa devices",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MFADevices"),
			},
			{
				Name:        "mfa_devices_in_use",
				Description: "Number of MFA devices in use.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MFADevicesInUse"),
			},
			{
				Name:        "account_mfa_enabled",
				Description: "Whether MFA is enabled for the account.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountMFAEnabled"),
			},
			{
				Name:        "account_access_keys_present",
				Description: "Number of account level access keys present.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "account_signing_certificates_present",
				Description: "Number of account signing certificates present.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "attached_policies_per_group_quota",
				Description: "Maximum allowed attached policies per group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "policies",
				Description: "Current number of policies.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "policies_quota",
				Description: "Allowed number of policies.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "policy_size_quota",
				Description: "Allowed size of policies.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "policy_versions_in_use",
				Description: "Number of policy versions in use.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "policy_versions_in_use_quota",
				Description: "Allowed number of policy versions.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "versions_per_policy_quota",
				Description: " Allowed number of versions per policy.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "global_endpoint_token_version",
				Description: "Token version of the global endpoint.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "aliases",
				Description: "List of aliases associated with the account. AWS supports only one alias per account",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

type account struct {
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

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAccountSummary(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().IAM

	summary, err := svc.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
	if err != nil {
		return err
	}
	var accSummary account
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
