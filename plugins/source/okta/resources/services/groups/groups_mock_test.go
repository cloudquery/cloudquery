package groups

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockGroupService(ctrl)

	var g okta.Group
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListGroups(
		gomock.Any(),
		gomock.Any(),
	).Return(
		[]*okta.Group{&g},
		nil,
		nil,
	)

	var gu okta.User
	if err := faker.FakeObject(&gu); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListGroupUsers(
		gomock.Any(),
		g.Id,
		gomock.Any(),
	).Return(
		[]*okta.User{&gu},
		nil,
		nil,
	)

	return client.Services{
		Groups: mock,
	}
}

func TestGroups(t *testing.T) {
	client.MockTestHelper(t, Groups(), buildGroups)
}
