package recipes

import (
	logAlerts "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Monitor() []Resource {
	var resourceRelation = []resourceDefinition{
		{
			azureStruct:        &insights.DiagnosticSettingsResource{},
			listFunction:       "List",
			subServiceOverride: "DiagnosticSettings",
			customColumns:      []codegen.ColumnDefinition{{Name: "resource_uri", Type: schema.TypeString, Resolver: `schema.PathResolver("ResourceURI")`}},
			helpers: []string{`
			func isResourceTypeNotSupported(err error) bool {
				var azureErr *azure.RequestError
				if errors.As(err, &azureErr) {
					return azureErr.ServiceError != nil && azureErr.ServiceError.Code == "ResourceTypeNotSupported"
				}
				return false
			}`, `// diagnosticSettingResource is a custom copy of insights.DiagnosticSettingsResource with extra ResourceURI field
			type diagnosticSettingResource struct {
				insights.DiagnosticSettingsResource
				ResourceURI string
			}`},
			listFunctionArgsInit: []string{`resource := parent.Item.(resources.GenericResourceExpanded)`},
			listFunctionArgs:     []string{"*resource.ID"},
			listHandler: `if err != nil {
				if isResourceTypeNotSupported(err) {
					return nil
				}
				return err
			}
			if response.Value == nil {
				return nil
			}
			for _, v := range *response.Value {
				res <- diagnosticSettingResource{
					DiagnosticSettingsResource: v,
					ResourceURI:                *resource.ID,
				}
			}`,
			mockListFunctionArgsInit: []string{`mockClient.EXPECT().List(gomock.Any(), "/subscriptions/testSubscription").Return(result, nil)`},
			mockListFunctionArgs:     []string{`"/subscriptions/test/resourceGroups/test/providers/test/test/test"`},
			mockListResult:           "DiagnosticSettingsResourceCollection",
		},
	}
	var resourcesByTemplates = []byTemplates{
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &logAlerts.ActivityLogAlertResource{},
					listFunction:       "ListBySubscriptionID",
					listHandler:        valueHandler,
					subServiceOverride: "ActivityLogAlerts",
					mockListResult:     "ActivityLogAlertList",
				},
			},
			serviceNameOverride: "Monitor",
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &insights.LogProfileResource{},
					listFunction:       "List",
					listHandler:        valueHandler,
					subServiceOverride: "LogProfiles",
					mockListResult:     "LogProfileCollection",
				},
			},
			serviceNameOverride: "Monitor",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"regexp", "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &insights.EventData{},
					listFunction:     "List",
					listFunctionArgs: []string{"filter", `""`},
					listFunctionArgsInit: []string{
						"const fetchWindow = 24 * time.Hour",
						"now := time.Now().UTC()",
						"past := now.Add(-fetchWindow)",
						`filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", past.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))`},
					subServiceOverride: "ActivityLogs",
					mockHelpers: []string{`
					type regexMatcher struct {
						re *regexp.Regexp
					}`, `
					func (m regexMatcher) Matches(x interface{}) bool {
						s, ok := x.(string)
						if !ok {
							return false
						}
						return m.re.MatchString(s)
					}`, `
					func (m regexMatcher) String() string {
						return m.re.String()
					}`},
					mockListResult:           "EventDataCollection",
					mockListFunctionArgs:     []string{"regexMatcher{filterRe}", `""`},
					mockListFunctionArgsInit: []string{"filterRe := regexp.MustCompile(`eventTimestamp ge '\\d{4}-\\d\\d-\\d\\dT\\d\\d:\\d\\d:\\d\\d(\\.\\d+)Z' and eventTimestamp le '\\d{4}-\\d\\d-\\d\\dT\\d\\d:\\d\\d:\\d\\d(\\.\\d+)Z'`)"},
				},
			},
			serviceNameOverride: "Monitor",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &resources.GenericResourceExpanded{},
					skipFields:       []string{"Properties"},
					includeColumns:   "^id$",
					listFunction:     "List",
					listFunctionArgs: []string{`""`, `""`, `nil`},
					listFunctionArgsInit: []string{`// Add subscription id as the first entry
					subscriptionId := "/" + client.ScopeSubscription(meta.(*client.Client).SubscriptionId)
					res <- resources.GenericResourceExpanded{ID: &subscriptionId}`},
					subServiceOverride:       "Resources",
					mockListResult:           "ListResult",
					mockListFunctionArgsInit: []string{``},
					relations:                resourceRelation,
				},
			},
			serviceNameOverride: "Monitor",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights",
						"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources",
					},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights",
						"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources",
					},
				},
			},
			definitions:         resourceRelation,
			serviceNameOverride: "Monitor",
		},
	}

	initParents(resourcesByTemplates)
	return generateResources(resourcesByTemplates)
}
