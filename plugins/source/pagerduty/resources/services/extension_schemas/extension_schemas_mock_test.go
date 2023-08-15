package extension_schemas

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListExtensionSchemaResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.ExtensionSchemas[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/extension_schemas", response)

	return &mockHttpClient
}

func TestExtensionSchemas(t *testing.T) {
	client.PagerdutyMockTestHelper(t, ExtensionSchemas(), buildMockHttpClient)
}
