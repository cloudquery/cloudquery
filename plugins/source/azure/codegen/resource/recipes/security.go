package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Security() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armsecurity.Assessment),
			Resolver: &resource.FuncParams{
				Func:   security.AssessmentsClient.NewListPager,
				Params: []string{"c.ScopeSubscription()"},
			},
		},
		{
			Struct:   new(armsecurity.AutoProvisioningSetting),
			Resolver: security.AutoProvisioningSettingsClient.NewListPager,
		},
		{
			Struct:   new(armsecurity.Contact),
			Resolver: security.ContactsClient.NewListPager,
		},
		{
			Struct:   new(armsecurity.JitNetworkAccessPolicy),
			Resolver: security.JitNetworkAccessPoliciesClient.NewListPager,
		},
		{
			Struct:   new(armsecurity.Pricing),
			Resolver: security.PricingsClient.List,
		},
		{
			Struct: new(armsecurity.Setting),
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "enabled",
					Type:     schema.TypeBool,
					Resolver: "resolveEnabled",
				},
			},
			Resolver: security.SettingsClient.NewListPager,
			SkipMock: true,
		},
	}
}
