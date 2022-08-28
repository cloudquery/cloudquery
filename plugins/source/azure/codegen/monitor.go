package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
)

func Monitor() []Resource {
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
					azureStruct:        &insights.ActivityLogAlertResource{},
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
	}

	return generateResources(resourcesByTemplates)
}
