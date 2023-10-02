package dashboards

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildDashboardsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockDashboardsAPIClient(ctrl)
	services := client.DatadogServices{
		DashboardsAPI: m,
	}

	var d datadogV1.DashboardSummaryDefinition
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	desc := "test string"
	d.Description.Set(&desc)
	m.EXPECT().ListDashboardsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(d))

	return services
}

func TestDashboards(t *testing.T) {
	client.DatadogMockTestHelper(t, Dashboards(), buildDashboardsMock, client.TestOptions{})
}
