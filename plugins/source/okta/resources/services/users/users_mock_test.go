package users

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockUserService(ctrl)

	var u okta.User
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	up := okta.UserProfile{}
	up["test"] = "value"
	u.Profile = &up

	mock.EXPECT().ListUsers(
		gomock.Any(),
		gomock.Any(),
	).Return(
		[]*okta.User{&u},
		nil,
		nil,
	)

	return client.Services{
		Users: mock,
	}
}

func TestUsers(t *testing.T) {
	client.MockTestHelper(t, Users(), buildUsers)
}
