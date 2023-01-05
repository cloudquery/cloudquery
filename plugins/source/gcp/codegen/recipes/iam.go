package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	iamadmin "cloud.google.com/go/iam/admin/apiv1"
	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	policiespb "google.golang.org/genproto/googleapis/iam/v2"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "roles",
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/roles#Role",

			Struct:         &iampb.Role{},
			RegisterServer: iampb.RegisterIAMServer,
			ProtobufImport: "cloud.google.com/go/iam/admin/apiv1/adminpb",
			ResponseStruct: &iampb.ListRolesResponse{},

			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			SkipFetch:   true,

			// These properties must be specified manually since they are only populated from reflection when SkipFetch is false
			ListFunctionName:  "ListRoles",
			RequestStructName: "ListRolesRequest",
		},
		{
			SubService:  "service_accounts",
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount",

			Struct:              &iampb.ServiceAccount{},
			NewFunction:         iamadmin.NewIamClient,
			MockImports:         []string{"cloud.google.com/go/iam/admin/apiv1"},
			ProtobufImport:      "cloud.google.com/go/iam/admin/apiv1/adminpb",
			ListFunction:        (&iamadmin.IamClient{}).ListServiceAccounts,
			RequestStructFields: `Name: "projects/" + c.ProjectId,`,

			NameTransformer:    CreateReplaceTransformer(map[string]string{"oauth_2": "oauth2"}),
			PrimaryKeys:        []string{"unique_id"},
			Relations:          []string{"ServiceAccountKeys()"},
			ServiceAPIOverride: "admin",
			SkipFields:         []string{"ProjectId"},
			SkipMock:           true,
		},
		{
			SubService:  "service_account_keys",
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey",

			Struct: &iampb.ServiceAccountKey{},

			ChildTable: true,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "service_account_unique_id",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.ParentColumnResolver("unique_id")`,
				},
			},
			PrimaryKeys: []string{"unique_id"},
			SkipFetch:   true,
			SkipFields:  []string{"ProjectId", "PrivateKeyData", "PrivateKeyType"},
			SkipMock:    true,
		},
		{
			SubService:  "deny_policies",
			Description: "https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy",

			Struct:         &policiespb.Policy{},
			RegisterServer: policiespb.RegisterPoliciesServer,
			ProtobufImport: "google.golang.org/genproto/googleapis/iam/v2",
			ResponseStruct: &policiespb.ListPoliciesResponse{},

			SkipFetch: true,

			// These properties must be specified manually since they are only populated from reflection when SkipFetch is false
			ListFunctionName:  "ListPolicies",
			RequestStructName: "ListPoliciesRequest",
		},
	}

	for _, resource := range resources {
		resource.Service = "iam"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
