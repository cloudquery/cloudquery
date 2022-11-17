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
			Struct: new(armsecurity.AutoProvisioningSetting),
			Resolver: &resource.FuncParams{
				Func: security.AutoProvisioningSettingsClient.NewListPager,
			},
		},
		{
			Struct: new(armsecurity.Contact),
			Resolver: &resource.FuncParams{
				Func: security.ContactsClient.NewListPager,
			},
		},
		{
			Struct: new(armsecurity.JitNetworkAccessPolicy),
			Resolver: &resource.FuncParams{
				Func: security.JitNetworkAccessPoliciesClient.NewListPager,
			},
		},
		{
			Struct: new(armsecurity.Pricing),
			Resolver: &resource.FuncParams{
				Func:   security.PricingsClient.List,
				Params: []string{"ctx"},
			},
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
			Resolver: &resource.FuncParams{
				Func: security.SettingsClient.NewListPager,
			},
			SkipMock: true,
		},
	}
}
