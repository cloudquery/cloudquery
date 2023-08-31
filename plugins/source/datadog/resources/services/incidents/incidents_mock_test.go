package incidents

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildIncidentsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockIncidentsAPIClient(ctrl)
	services := client.DatadogServices{
		IncidentsAPI: m,
	}

	var i datadogV2.IncidentResponseData
	err := faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIncidentsWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(i))

	var ar datadogV2.IncidentAttachmentsResponse
	err = faker.FakeObject(&ar)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIncidentAttachments(gomock.Any(), gomock.Any()).Return(ar, nil, nil)

	return services
}

func TestIncidents(t *testing.T) {
	client.DatadogMockTestHelper(t, Incidents(), buildIncidentsMock, client.TestOptions{})
}
