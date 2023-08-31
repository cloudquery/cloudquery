package peering

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ServiceProviders() *schema.Table {
	return &schema.Table{
		Name:                 "azure_peering_service_providers",
		Resolver:             fetchServiceProviders,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/peering/peering-service-providers/list?tabs=HTTP#peeringserviceprovider",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_peering_service_providers", client.Namespacemicrosoft_peering),
		Transform:            transformers.TransformWithStruct(&armpeering.ServiceProvider{}, transformers.WithPrimaryKeys("Name")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
	}
}
