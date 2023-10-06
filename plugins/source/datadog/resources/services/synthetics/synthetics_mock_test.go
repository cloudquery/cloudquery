package synthetics

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildSyntheticsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockSyntheticsAPIClient(ctrl)
	services := client.DatadogServices{
		SyntheticsAPI: m,
	}

	var s datadogV1.SyntheticsTestDetails
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTestsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(s))

	return services
}

func TestSynthetics(t *testing.T) {
	client.DatadogMockTestHelper(t, Synthetics(), buildSyntheticsMock, client.TestOptions{})
}
