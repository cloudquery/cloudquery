package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_servers",
		Resolver:    fetchServers,
		Description: "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/servers/list?tabs=HTTP#server",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_sql_servers", client.Namespacemicrosoft_sql),
		Transform:   transformers.TransformWithStruct(&armsql.Server{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			serverVulnerabilityAssessments(),
			server_blob_auditing_policies(),
			serverAdmins(),
			serverEncryptionProtectors(),
			serverDatabases(),
			virtualNetworkRules(),
			serverSecurityAlertPolicies(),
		},
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsql.NewServersClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
