package authorization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RoleAssignments() *schema.Table {
	return &schema.Table{
		Name:        "azure_authorization_role_assignments",
		Resolver:    fetchRoleAssignments,
		Description: "https://learn.microsoft.com/en-us/rest/api/authorization/role-assignments/get?tabs=HTTP#roleassignment",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_authorization_role_assignments", client.Namespacemicrosoft_authorization),
		Transform:   transformers.TransformWithStruct(&armauthorization.RoleAssignment{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchRoleAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewRoleAssignmentsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListForSubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
