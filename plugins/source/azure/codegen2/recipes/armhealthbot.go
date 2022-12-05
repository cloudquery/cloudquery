// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"

func Armhealthbot() []Table {
	tables := []Table{
		{
      Name: "health_bot",
      Struct: &armhealthbot.HealthBot{},
      ResponseStruct: &armhealthbot.BotsClientListResponse{},
      Client: &armhealthbot.BotsClient{},
      ListFunc: (&armhealthbot.BotsClient{}).NewListPager,
			NewFunc: armhealthbot.NewBotsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.HealthBot/healthBots",
		},
	}

	for i := range tables {
		tables[i].Service = "armhealthbot"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhealthbot()...)
}