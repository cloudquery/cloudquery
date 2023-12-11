package oncalls

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/v4/faker"

	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListOnCallsResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.OnCalls[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/oncalls", response)

	return &mockHttpClient
}

func TestOncalls(t *testing.T) {
	client.PagerdutyMockTestHelper(t, Oncalls(), buildMockHttpClient)
}
