package roles

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildRolesMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockRolesAPIClient(ctrl)
	services := client.DatadogServices{
		RolesAPI: m,
	}

	var roles datadogV2.RolesResponse
	err := faker.FakeObject(&roles)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRoles(gomock.Any()).Return(roles, nil, nil)

	var permissions datadogV2.PermissionsResponse
	err = faker.FakeObject(&permissions)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRolePermissions(gomock.Any(), gomock.Any()).Return(permissions, nil, nil)

	var users datadogV2.UsersResponse
	err = faker.FakeObject(&users)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRoleUsers(gomock.Any(), gomock.Any()).Return(users, nil, nil)

	return services
}

func TestRoles(t *testing.T) {
	client.DatadogMockTestHelper(t, Roles(), buildRolesMock, client.TestOptions{})
}
