package notebooks

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildNotebooksMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockNotebooksAPIClient(ctrl)
	services := client.DatadogServices{
		NotebooksAPI: m,
	}

	var n datadogV1.NotebooksResponse
	err := faker.FakeObject(&n)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListNotebooks(gomock.Any()).Return(n, nil, nil)

	return services
}

func TestNotebooks(t *testing.T) {
	client.DatadogMockTestHelper(t, Notebooks(), buildNotebooksMock, client.TestOptions{})
}
