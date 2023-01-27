package azuredata

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SqlServerRegistrations() *schema.Table {
	return &schema.Table{
		Name:        "azure_azuredata_sql_server_registrations",
		Resolver:    fetchSqlServerRegistrations,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata@v0.5.0#SQLServerRegistration",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_azuredata_sql_server_registrations", client.Namespacemicrosoft_azuredata),
		Transform:   transformers.TransformWithStruct(&armazuredata.SQLServerRegistration{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}
