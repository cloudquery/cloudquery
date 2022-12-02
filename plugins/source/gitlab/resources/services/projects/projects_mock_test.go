package projects

import (
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildProjects(t *testing.T, ctrl *gomock.Controller) client.Services {
	projectsMock := mocks.NewMockProjectsClient(ctrl)
	releaseMock := mocks.NewMockReleasesClient(ctrl)

	var release *gitlab.Release
	if err := faker.FakeObject(&release); err != nil {
		t.Fatal(err)
	}

	releaseMock.EXPECT().ListReleases(gomock.Any(), gomock.Any()).Return([]*gitlab.Release{release}, &gitlab.Response{}, nil)

	var project *gitlab.Project
	if err := faker.FakeObject(&project); err != nil {
		t.Fatal(err)
	}

	project.Permissions = &gitlab.Permissions{
		GroupAccess: &gitlab.GroupAccess{
			AccessLevel:       gitlab.GuestPermissions,
			NotificationLevel: gitlab.DisabledNotificationLevel,
		},
		ProjectAccess: &gitlab.ProjectAccess{
			AccessLevel:       gitlab.GuestPermissions,
			NotificationLevel: gitlab.DisabledNotificationLevel,
		},
	}

	isoTime := gitlab.ISOTime(time.Now())
	project.MarkedForDeletionAt = &isoTime

	projectsMock.EXPECT().ListProjects(gomock.Any(), gomock.Any()).Return([]*gitlab.Project{project}, &gitlab.Response{}, nil)

	return client.Services{
		Releases: releaseMock,
		Projects: projectsMock,
	}
}

func TestProjects(t *testing.T) {
	client.GitlabMockTestHelper(t, Projects(), buildProjects, client.TestOptions{})
}
