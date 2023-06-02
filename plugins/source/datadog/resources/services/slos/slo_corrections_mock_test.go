package slos

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildObjectiveCorrectionsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockServiceLevelObjectiveCorrectionsAPIClient(ctrl)
	services := client.DatadogServices{
		ServiceLevelObjectiveCorrectionsAPI: m,
	}

	var d datadogV1.SLOCorrectionListResponse
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSLOCorrection(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestObjectiveCorrections(t *testing.T) {
	client.DatadogMockTestHelper(t, ObjectiveCorrections(), buildObjectiveCorrectionsMock, client.TestOptions{})
}
