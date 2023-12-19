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

	var perms datadogV2.PermissionsResponse
	err := faker.FakeObject(&perms)
	if err != nil {
		t.Fatal(err)
	}
	perms.Data[0].AdditionalProperties = map[string]any{"key": "value"}
	perms.AdditionalProperties = map[string]any{"key": "value"}
	m.EXPECT().ListPermissions(gomock.Any()).Return(perms, nil, nil)

	return services
}

func TestPermissions(t *testing.T) {
	client.DatadogMockTestHelper(t, Permissions(), buildPermissionsMock, client.TestOptions{})
}
