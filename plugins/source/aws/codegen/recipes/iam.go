package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IAMResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct:     &models.Account{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "credential_reports",
			Struct:     &models.CredentialReportEntry{},
			SkipFields: []string{
				"Arn",
				"UserCreationTime",
				"PasswordLastChanged",
				"PasswordNextRotation",
				"AccessKey1LastRotated",
				"AccessKey2LastRotated",
				"Cert1LastRotated",
				"Cert2LastRotated",
				"AccessKey1LastUsedDate",
				"AccessKey2LastUsedDate",
				"PasswordLastUsed",
			},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "user_creation_time",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("UserCreationTime")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "password_last_changed",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("PasswordLastChanged")`,
				},
				{
					Name:     "password_next_rotation",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("PasswordNextRotation")`,
				},
				{
					Name:     "access_key_1_last_rotated",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("AccessKey1LastRotated")`,
				},
				{
					Name:     "access_key_2_last_rotated",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("AccessKey2LastRotated")`,
				},
				{
					Name:     "cert_1_last_rotated",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("Cert1LastRotated")`,
				},
				{
					Name:     "cert_2_last_rotated",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("Cert2LastRotated")`,
				},
				{
					Name:     "access_key_1_last_used_date",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("AccessKey1LastUsedDate")`,
				},
				{
					Name:     "access_key_2_last_used_date",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("AccessKey2LastUsedDate")`,
				},
				{
					Name:     "password_last_used",
					Type:     schema.TypeTimestamp,
					Resolver: `timestampPathResolver("PasswordLastUsed")`,
				},
				// password_enabled is an alias for the (now deprecated) password_status - https://github.com/cloudquery/cloudquery/issues/3145
				{
					Name:     "password_enabled",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("PasswordStatus")`,
				},
			},
			Relations: []string{},
		},
		{
			SubService:  "groups",
			Struct:      &types.Group{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html",
			SkipFields:  []string{"GroupId"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "policies",
					Type:     schema.TypeJSON,
					Resolver: `resolveIamGroupPolicies`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("GroupId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			Relations: []string{"GroupPolicies()"},
		},
		{
			SubService:          "group_policies",
			Struct:              &iam.GetGroupPolicyOutput{},
			SkipFields:          []string{"PolicyDocument"},
			PreResourceResolver: "getGroupPolicy",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "group_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveIamGroupPolicyPolicyDocument`,
					},
				}...),
		},
		{
			SubService:          "openid_connect_identity_providers",
			Struct:              &models.IamOpenIdIdentityProviderWrapper{},
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getOpenIdConnectIdentityProvider",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "password_policies",
			Struct:     &models.PasswordPolicyWrapper{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "policies",
			Struct:      &types.ManagedPolicyDetail{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html",
			SkipFields:  []string{"PolicyId", "Tags", "PolicyVersionList"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("PolicyId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveIamPolicyTags`,
				},
				{
					Name:     "policy_version_list",
					Type:     schema.TypeJSON,
					Resolver: `resolveIamPolicyVersionList`,
				},
			},
		},
		{
			SubService:          "roles",
			Struct:              &types.Role{},
			Description:         "https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html",
			SkipFields:          []string{"RoleId", "AssumeRolePolicyDocument"},
			PreResourceResolver: "getRole",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "policies",
					Type:     schema.TypeJSON,
					Resolver: `resolveIamRolePolicies`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("RoleId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "assume_role_policy_document",
					Type:     schema.TypeJSON,
					Resolver: `resolveRolesAssumeRolePolicyDocument`,
				},
			},
			Relations: []string{
				"RolePolicies()",
			},
		},
		{
			SubService:          "role_policies",
			Struct:              &iam.GetRolePolicyOutput{},
			SkipFields:          []string{"PolicyDocument"},
			PreResourceResolver: "getRolePolicy",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "role_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveRolePoliciesPolicyDocument`,
					},
				}...),
		},
		{
			SubService:          "saml_identity_providers",
			Struct:              &types.SAMLProviderListEntry{},
			Description:         "https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html",
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getSamlIdentityProvider",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "server_certificates",
			Struct:      &types.ServerCertificateMetadata{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html",
			SkipFields:  []string{"ServerCertificateId"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ServerCertificateId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:          "users",
			Struct:              &types.User{},
			Description:         "https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html",
			SkipFields:          []string{"Arn", "AccountId", "UserId"},
			PreResourceResolver: "getUser",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("UserId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			Relations: []string{
				"UserAccessKeys()",
				"UserGroups()",
				"UserAttachedPolicies()",
				"UserPolicies()",
			},
		},
		{
			SubService:           "user_access_keys",
			Struct:               &models.AccessKeyWrapper{},
			SkipFields:           []string{},
			PostResourceResolver: `postIamUserAccessKeyResolver`,
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name: "last_used",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_used_service_name",
						Type: schema.TypeString,
					},
				}...),
		},
		{
			SubService:  "user_groups",
			Struct:      &types.Group{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
				}...),
		},
		{
			SubService:  "user_attached_policies",
			Struct:      &types.AttachedPolicy{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
				}...),
		},
		{
			SubService:          "user_policies",
			Struct:              &iam.GetUserPolicyOutput{},
			SkipFields:          []string{"PolicyDocument"},
			PreResourceResolver: "getUserPolicy",
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveIamUserPolicyPolicyDocument`,
					},
				}...),
		},
		{
			SubService:  "virtual_mfa_devices",
			Struct:      &types.VirtualMFADevice{},
			Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html",
			SkipFields:  []string{"SerialNumber"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "serial_number",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "iam"
		r.Multiplex = `client.AccountMultiplex`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
