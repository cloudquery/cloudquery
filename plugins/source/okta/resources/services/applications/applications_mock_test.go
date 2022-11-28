package applications

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func buildApplications(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockApplicationService(ctrl)

	var a okta.Application
	if err := faker.FakeObject(&a); err != nil {
		t.Fatal(err)
	}
	a.Profile = map[string]string{"test": "value"}

	mock.EXPECT().ListApplications(
		gomock.Any(),
		gomock.Any(),
	).Return(
		[]okta.App{&a},
		nil,
		nil,
	)

	var aga okta.ApplicationGroupAssignment
	if err := faker.FakeObject(&aga); err != nil {
		t.Fatal(err)
	}
	aga.Profile = map[string]string{"test": "value"}

	mock.EXPECT().ListApplicationGroupAssignments(
		gomock.Any(),
		a.Id,
		gomock.Any(),
	).Return(
		[]*okta.ApplicationGroupAssignment{&aga},
		nil,
		nil,
	)

	var au okta.AppUser
	if err := faker.FakeObject(&au); err != nil {
		t.Fatal(err)
	}
	au.Profile = map[string]string{"test": "value"}

	mock.EXPECT().ListApplicationUsers(
		gomock.Any(),
		a.Id,
		gomock.Any(),
	).Return(
		[]*okta.AppUser{&au},
		nil,
		nil,
	)

	return client.Services{
		Applications: mock,
	}
}

func TestApplications(t *testing.T) {
	client.MockTestHelper(t, Applications(), buildApplications)
}
