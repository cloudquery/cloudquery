package roles

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildPermissionsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockRolesAPIClient(ctrl)
	services := client.DatadogServices{
		RolesAPI: m,
	}

	var roles datadogV2.PermissionsResponse
	err := faker.FakeObject(&roles)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListPermissions(gomock.Any()).Return(roles, nil, nil)

	return services
}

func TestPermissions(t *testing.T) {
	client.DatadogMockTestHelper(t, Permissions(), buildPermissionsMock, client.TestOptions{})
}
