package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	iamService "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IAMResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "accounts",
			Struct:     &iamService.Account{},
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
			Struct:     &iamService.CredentialReportEntry{},
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
			},
			Relations: []string{},
		},
		{
			SubService: "groups",
			Struct:     &types.Group{},
			SkipFields: []string{"GroupId"},
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
			SubService: "group_policies",
			Struct:     &iam.GetGroupPolicyOutput{},
			SkipFields: []string{"PolicyDocument"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "group_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveIamGroupPolicyPolicyDocument`,
					},
				}...),
		},
		{
			SubService: "openid_connect_identity_providers",
			Struct:     &types.OpenIDConnectProviderListEntry{},
			SkipFields: []string{"Arn", "Tags"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "password_policies",
			Struct:     &iamService.PasswordPolicyWrapper{},
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
			SubService: "policies",
			Struct:     &types.Policy{},
			SkipFields: []string{"PolicyId", "Tags"},
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
			},
		},
		{
			SubService: "roles",
			Struct:     &types.Role{},
			SkipFields: []string{"RoleId", "Tags", "AssumeRolePolicyDocument"},
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
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `client.ResolveTags`,
				},
			},
			Relations: []string{
				"RolePolicies()",
			},
		},
		{
			SubService: "role_policies",
			Struct:     &iam.GetRolePolicyOutput{},
			SkipFields: []string{"PolicyDocument"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "role_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveRolePoliciesPolicyDocument`,
					},
				}...),
		},
		{
			SubService: "saml_identity_providers",
			Struct:     &types.SAMLProviderListEntry{},
			SkipFields: []string{"Arn", "Tags"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "server_certificates",
			Struct:     &types.ServerCertificateMetadata{},
			SkipFields: []string{"ServerCertificateId"},
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
			SubService: "users",
			Struct:     &types.User{},
			SkipFields: []string{"Arn", "AccountId", "UserId", "Tags"},
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
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `client.ResolveTags`,
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
			Struct:               &iamService.AccessKeyWrapper{},
			SkipFields:           []string{},
			PostResourceResolver: `postIamUserAccessKeyResolver`,
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
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
			SubService: "user_groups",
			Struct:     &types.Group{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
				}...),
		},
		{
			SubService: "user_attached_policies",
			Struct:     &types.AttachedPolicy{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
				}...),
		},
		{
			SubService: "user_policies",
			Struct:     &iam.GetUserPolicyOutput{},
			SkipFields: []string{"PolicyDocument"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "policy_document",
						Type:     schema.TypeJSON,
						Resolver: `resolveIamUserPolicyPolicyDocument`,
					},
				}...),
		},
		{
			SubService: "virtual_mfa_devices",
			Struct:     &types.VirtualMFADevice{},
			SkipFields: []string{"SerialNumber", "Tags", "UserTags"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "serial_number",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:     "user_tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
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
