// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"

func Armhealthbot() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhealthbot.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot",
		},
		{
			NewFunc: armhealthbot.NewBotsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhealthbot())
}