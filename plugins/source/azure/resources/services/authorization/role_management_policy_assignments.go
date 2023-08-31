package authorization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RoleManagementPolicyAssignments() *schema.Table {
	return &schema.Table{
		Name:                 "azure_role_management_policy_assignments",
		Resolver:             fetchRoleManagementPolicyAssignments,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/authorization/role-management-policy-assignments/list-for-scope?tabs=HTTP#rolemanagementpolicyassignment",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_role_management_policy_assignments", client.Namespacemicrosoft_authorization),
		Transform:            transformers.TransformWithStruct(&armauthorization.RoleManagementPolicyAssignment{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchRoleManagementPolicyAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewRoleManagementPolicyAssignmentsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	scope := "/subscriptions/" + cl.SubscriptionId
	pager := svc.NewListForScopePager(scope, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
