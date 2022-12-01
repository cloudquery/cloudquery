package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Security() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &security.Assessment{},
					listFunction:         "List",
					listFunctionArgs:     []string{`fmt.Sprintf("/subscriptions/%s", meta.(*client.Client).SubscriptionId)`},
					mockListFunctionArgs: []string{`"/subscriptions/testSubscription"`},
					mockListResult:       "AssessmentList",
					mockFieldsToIgnore:   []string{"ResourceDetails"},
				},
				{
					azureStruct:    &security.AutoProvisioningSetting{},
					listFunction:   "List",
					mockListResult: "AutoProvisioningSettingList",
				},
				{
					azureStruct:    &security.Contact{},
					listFunction:   "List",
					mockListResult: "ContactList",
				},
				{
					azureStruct:    &security.JitNetworkAccessPolicy{},
					listFunction:   "List",
					mockListResult: "JitNetworkAccessPoliciesList",
				},
				{
					azureStruct:  &security.Setting{},
					listFunction: "List",
					listHandler: `if err != nil {
						return err
					}
					for response.NotDone() {
						for _, item := range response.Values() {
							if v, ok := item.AsSetting(); ok {
								res <- v
							} else if v, ok := item.AsDataExportSettings(); ok {
								res <- v
							} else if v, ok := item.AsAlertSyncSettings(); ok {
								res <- v
							} else {
								return errors.WithStack(fmt.Errorf("unexpected BasicSetting: %#v", item))
							}
						}
						if err := response.NextWithContext(ctx); err != nil {
							return err
						}
					}`,
					mockListResult: "SettingsList",
					mockValueType:  "BasicSetting",
					customColumns: []codegen.ColumnDefinition{{
						Name:     "enabled",
						Type:     schema.TypeBool,
						Resolver: "resolveEnabled",
					}},
					helpers: []string{`func resolveEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
						item := resource.Item.(security.BasicSetting)
						if v, ok := item.AsDataExportSettings(); ok {
							return errors.WithStack(resource.Set(c.Name, v.Enabled))
						}
						if v, ok := item.AsAlertSyncSettings(); ok {
							return errors.WithStack(resource.Set(c.Name, v.Enabled))
						}
						return errors.WithStack(resource.Set(c.Name, true))
					}
					`},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &security.Pricing{},
					listFunction:   "List",
					listHandler:    valueHandler,
					mockListResult: "PricingList",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
