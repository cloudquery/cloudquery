package actions

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildWorkflowRuns(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	actionsMock := mocks.NewMockActionsService(ctrl)

	// create mock for repositories
	var repository github.Repository
	require.NoError(t, faker.FakeObject(&repository))

	var runs github.WorkflowRuns
	require.NoError(t, faker.FakeObject(&runs))
	runs.TotalCount = github.Int(1)

	var runUsage github.WorkflowRunUsage
	require.NoError(t, faker.FakeObject(&runUsage))

	var workflowJobs github.Jobs
	require.NoError(t, faker.FakeObject(&workflowJobs))
	workflowJobs.TotalCount = github.Int(1)

	actionsMock.EXPECT().ListRepositoryWorkflowRuns(gomock.Any(), *repository.Owner.Login, *repository.Name, gomock.Any()).Return(&runs, &github.Response{}, nil)
	actionsMock.EXPECT().ListWorkflowJobs(gomock.Any(), *repository.Owner.Login, *repository.Name, *runs.WorkflowRuns[0].ID, gomock.Any()).Return(&workflowJobs, &github.Response{}, nil)
	actionsMock.EXPECT().GetWorkflowRunUsageByID(gomock.Any(), *repository.Owner.Login, *repository.Name, *runs.WorkflowRuns[0].ID).Return(&runUsage, &github.Response{}, nil)
	return client.GithubServices{Actions: actionsMock}
}

func TestWorkflowRuns(t *testing.T) {
	client.GithubMockTestHelper(t, WorkflowRuns(), buildWorkflowRuns, client.TestOptions{})
}
