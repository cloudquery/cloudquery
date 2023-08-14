package services

import (
	"fmt"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListServiceResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.Services[0]); err != nil {
		panic(err)
	}

	response.Services[0].Integrations[0].EmailFilterMode = pagerduty.EmailFilterModeAll
	response.Services[0].Integrations[0].EmailFilters[0].BodyMode = pagerduty.EmailFilterRuleModeAlways
	response.Services[0].Integrations[0].EmailFilters[0].FromEmailMode = pagerduty.EmailFilterRuleModeAlways
	response.Services[0].Integrations[0].EmailFilters[0].SubjectMode = pagerduty.EmailFilterRuleModeAlways

	response.More = false

	mockHttpClient.AddMockResponse("/services", response)

	serviceruleResponse := pagerduty.ListServiceRulesResponse{}
	if err := faker.FakeObject(&serviceruleResponse); err != nil {
		panic(err)
	}
	serviceruleResponse.More = false
	if err := client.FakeStringTimestamps(&serviceruleResponse.Rules[0]); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/services/%s/rules",
			response.Services[0].ID),
		serviceruleResponse)

	return &mockHttpClient
}

func TestServices(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Services(), buildMockHttpClient)
}
