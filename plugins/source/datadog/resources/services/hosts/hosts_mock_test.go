package hosts

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildHostsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockHostsAPIClient(ctrl)
	services := client.DatadogServices{
		HostsAPI: m,
	}

	var h datadogV1.HostListResponse
	err := faker.FakeObject(&h)
	if err != nil {
		t.Fatal(err)
	}
	i64val := int64(123)
	h.HostList[0].MuteTimeout.Set(&i64val)

	m.EXPECT().ListHosts(gomock.Any()).Return(h, nil, nil)

	return services
}

func TestHosts(t *testing.T) {
	client.DatadogMockTestHelper(t, Hosts(), buildHostsMock, client.TestOptions{})
}
