package rum

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildEventsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockRumAPIClient(ctrl)
	services := client.DatadogServices{
		RumAPI: m,
	}

	var i datadogV2.RUMEvent
	err := faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRUMEventsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(i))

	return services
}

func TestEvents(t *testing.T) {
	client.DatadogMockTestHelper(t, Events(), buildEventsMock, client.TestOptions{})
}
