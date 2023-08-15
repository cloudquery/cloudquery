package healthbot

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Bots() *schema.Table {
	return &schema.Table{
		Name:                 "azure_healthbot_bots",
		Resolver:             fetchBots,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot@v1.0.0#HealthBot",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_healthbot_bots", client.Namespacemicrosoft_healthbot),
		Transform:            transformers.TransformWithStruct(&armhealthbot.HealthBot{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
