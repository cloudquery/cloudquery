package projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/xanzy/go-gitlab"
)

func buildProjects(mux *httprouter.Router) error {
	var project *gitlab.Project
	if err := faker.FakeObject(&project); err != nil {
		return err
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

	projectsResp, err := json.Marshal([]*gitlab.Project{project})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/projects", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(projectsResp))
	})

	var release *gitlab.Release
	if err := faker.FakeObject(&release); err != nil {
		return err
	}
	releaseResp, err := json.Marshal([]*gitlab.Release{release})
	if err != nil {
		return err
	}
	mux.GET("/api/v4/projects/:projectId/releases", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(releaseResp))
	})

	var branch *gitlab.Branch
	if err := faker.FakeObject(&branch); err != nil {
		return err
	}

	branchResp, err := json.Marshal([]*gitlab.Branch{branch})
	if err != nil {
		return err
	}

	mux.GET("/api/v4/projects/:projectId/repository/branches", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(branchResp))
	})
	return nil
}

func TestProjects(t *testing.T) {
	client.GitlabMockTestHelper(t, Projects(), buildProjects, client.TestOptions{})
}
