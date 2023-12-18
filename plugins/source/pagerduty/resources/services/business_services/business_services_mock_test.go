package business_services

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListBusinessServicesResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.BusinessServices[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/business_services", response)

	servicedependenciesResponse := pagerduty.ListServiceDependencies{}
	if err := faker.FakeObject(&servicedependenciesResponse); err != nil {
		panic(err)
	}
	mockHttpClient.AddMockResponse(
		"/service_dependencies/business_services/"+response.BusinessServices[0].ID,
		servicedependenciesResponse)

	return &mockHttpClient
}

func TestBusinessServices(t *testing.T) {
	client.PagerdutyMockTestHelper(t, BusinessServices(), buildMockHttpClient)
}
