package synthetics

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildGlobalVariablesMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockSyntheticsAPIClient(ctrl)
	services := client.DatadogServices{
		SyntheticsAPI: m,
	}

	var gv datadogV1.SyntheticsListGlobalVariablesResponse
	err := faker.FakeObject(&gv)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGlobalVariables(gomock.Any()).Return(gv, nil, nil)

	return services
}

func TestGlobalVariables(t *testing.T) {
	client.DatadogMockTestHelper(t, GlobalVariables(), buildGlobalVariablesMock, client.TestOptions{})
}
