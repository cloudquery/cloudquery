package dashboards

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildDashboardListsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockDashboardListsAPIClient(ctrl)
	services := client.DatadogServices{
		DashboardListsAPI: m,
	}

	var d datadogV1.DashboardListListResponse
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListDashboardLists(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestDashboardLists(t *testing.T) {
	client.DatadogMockTestHelper(t, Lists(), buildDashboardListsMock, client.TestOptions{})
}
