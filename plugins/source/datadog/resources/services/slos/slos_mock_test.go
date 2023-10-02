package slos

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildObjectivesMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockServiceLevelObjectivesAPIClient(ctrl)
	services := client.DatadogServices{
		ServiceLevelObjectivesAPI: m,
	}

	var d datadogV1.ServiceLevelObjective
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	str := "test"
	d.Description.Set(&str)

	m.EXPECT().ListSLOsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(d))

	return services
}

func TestObjectives(t *testing.T) {
	client.DatadogMockTestHelper(t, Objectives(), buildObjectivesMock, client.TestOptions{})
}
