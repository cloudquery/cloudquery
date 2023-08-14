package rulesets

import (
	"fmt"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListRulesetsResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.Rulesets[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/rulesets", response)

	rulesetruleResponse := pagerduty.ListRulesetRulesResponse{}
	if err := faker.FakeObject(&rulesetruleResponse); err != nil {
		panic(err)
	}
	rulesetruleResponse.More = false
	if err := client.FakeStringTimestamps(&rulesetruleResponse.Rules[0]); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		fmt.Sprintf("/rulesets/%s/rules",
			response.Rulesets[0].ID),
		rulesetruleResponse)

	return &mockHttpClient
}

func TestRulesets(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Rulesets(), buildMockHttpClient)
}
