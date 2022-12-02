package groups

import (
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	groupMock := mocks.NewMockGroupsClient(ctrl)

	var group *gitlab.Group
	if err := faker.FakeObject(&group, faker.WithMaxDepth(10)); err != nil {
		t.Fatal(err)
	}

	isoTime := gitlab.ISOTime(time.Now())
	group.MarkedForDeletionOn = &isoTime
	groupMock.EXPECT().ListGroups(gomock.Any(), gomock.Any()).Return([]*gitlab.Group{group}, &gitlab.Response{}, nil)

	var groupMember *gitlab.GroupMember
	if err := faker.FakeObject(&groupMember, faker.WithMaxDepth(12)); err != nil {
		t.Fatal(err)
	}

	groupMember.ExpiresAt = &isoTime
	groupMock.EXPECT().ListGroupMembers(gomock.Any(), gomock.Any()).Return([]*gitlab.GroupMember{groupMember}, &gitlab.Response{}, nil)
	return client.Services{
		Groups: groupMock,
	}
}

func TestGroups(t *testing.T) {
	client.GitlabMockTestHelper(t, Groups(), buildGroups, client.TestOptions{})
}
