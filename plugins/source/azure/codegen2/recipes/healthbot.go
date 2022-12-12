// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"

func init() {
	tables := []Table{
		{
			Service:        "armhealthbot",
			Name:           "bots",
			Struct:         &armhealthbot.HealthBot{},
			ResponseStruct: &armhealthbot.BotsClientListResponse{},
			Client:         &armhealthbot.BotsClient{},
			ListFunc:       (&armhealthbot.BotsClient{}).NewListPager,
			NewFunc:        armhealthbot.NewBotsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HealthBot/healthBots",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_HealthBot)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
