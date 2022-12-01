package users

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.GitlabServices {
	userMock := mocks.NewMockUsersClient(ctrl)

	var user *gitlab.User
	if err := faker.FakeObject(&user, faker.WithMaxDepth(25)); err != nil {
		t.Fatal(err)
	}

	userMock.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return([]*gitlab.User{user}, &gitlab.Response{}, nil)

	return client.GitlabServices{
		Users: userMock,
	}
}

func TestStorageBillings(t *testing.T) {
	client.GitlabMockTestHelper(t, Users(), buildUsers, client.TestOptions{})
}
