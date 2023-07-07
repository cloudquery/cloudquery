package security

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Pricings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_pricings",
		Resolver:             fetchPricings,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/pricings/list?tabs=HTTP#pricing",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_pricings", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.Pricing{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
