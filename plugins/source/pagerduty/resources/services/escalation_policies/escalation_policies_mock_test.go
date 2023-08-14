package escalation_policies

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListEscalationPoliciesResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.EscalationPolicies[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/escalation_policies", response)

	return &mockHttpClient
}

func TestEscalationPolicies(t *testing.T) {
	client.PagerdutyMockTestHelper(t, EscalationPolicies(), buildMockHttpClient)
}
