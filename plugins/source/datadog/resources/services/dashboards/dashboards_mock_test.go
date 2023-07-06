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

	var d datadogV1.DashboardSummary
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	desc := "test string"
	d.Dashboards[0].Description.Set(&desc)
	m.EXPECT().ListDashboards(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestDashboards(t *testing.T) {
	client.DatadogMockTestHelper(t, Dashboards(), buildDashboardsMock, client.TestOptions{})
}
