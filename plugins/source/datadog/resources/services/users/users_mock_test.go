package users

import (
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildUsersMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockUsersAPIClient(ctrl)
	services := client.DatadogServices{
		UsersAPI: m,
	}
	var i datadogV2.User
	err := faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUsersWithPagination(gomock.Any()).Return(client.MockPaginatedResponse(i))

	var permissions datadogV2.PermissionsResponse
	err = faker.FakeObject(&permissions)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPermissions(gomock.Any(), gomock.Any()).Return(permissions, nil, nil)

	return services
}

func TestUsers(t *testing.T) {
	client.DatadogMockTestHelper(t, Users(), buildUsersMock, client.TestOptions{})
}
