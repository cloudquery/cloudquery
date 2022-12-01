package groups

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.GitlabServices {
	userMock := mocks.NewMockUsersClient(ctrl)

	var user *gitlab.User
	if err := faker.FakeObject(&user, faker.WithMaxDepth(25)); err != nil {
		t.Fatal(err)
	}

	userMock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(user, &gitlab.Response{}, nil)
	// GetUser(user int, opt gitlab.GetUsersOptions, options ...gitlab.RequestOptionFunc) (*gitlab.User, *gitlab.Response, error)

	groupMock := mocks.NewMockGroupsClient(ctrl)

	var group *gitlab.Group
	if err := faker.FakeObject(&group, faker.WithMaxDepth(10)); err != nil {
		t.Fatal(err)
	}
	groupMock.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Return([]*gitlab.Group{group}, &gitlab.Response{}, nil)

	var groupMember *gitlab.GroupMember
	if err := faker.FakeObject(&groupMember, faker.WithMaxDepth(12)); err != nil {
		t.Fatal(err)
	}
	groupMock.EXPECT().ListGroupMembers(gomock.Any(), gomock.Any()).Return([]*gitlab.GroupMember{groupMember}, &gitlab.Response{}, nil)
	return client.GitlabServices{
		Users:  userMock,
		Groups: groupMock,
	}
}

func TestGroups(t *testing.T) {
	client.GitlabMockTestHelper(t, Groups(), buildGroups, client.TestOptions{})
}
