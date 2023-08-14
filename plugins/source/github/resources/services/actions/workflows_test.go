package actions

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildWorkflows(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)
	workflowsMock := mocks.NewMockActionsService(ctrl)

	// create mock for repositories
	var repository github.Repository
	if err := faker.FakeObject(&repository); err != nil {
		t.Fatal(err)
	}

	var workflow github.Workflow
	if err := faker.FakeObject(&workflow); err != nil {
		t.Fatal(err)
	}
	workflow.HTMLURL = github.String("https://github.com/testorg/repo/blob/master/.github/workflows/161335")
	workflow.Path = github.String(".github/workflows/ci.yml")
	count := 1
	workflows := github.Workflows{Workflows: []*github.Workflow{&workflow}, TotalCount: &count}

	workflowContent := github.RepositoryContent{}
	if err := faker.FakeObject(&workflowContent); err != nil {
		t.Fatal(err)
	}
	workflowContent.Encoding = github.String("")
	opts := github.RepositoryContentGetOptions{Ref: "master"}

	repositoriesMock.EXPECT().GetContents(gomock.Any(), "testorg", "repo", *workflow.Path, &opts).Return(&workflowContent, nil, nil, nil)
	workflowsMock.EXPECT().ListWorkflows(gomock.Any(), *repository.Owner.Login, *repository.Name, gomock.Any()).Return(&workflows, &github.Response{}, nil)
	return client.GithubServices{Actions: workflowsMock, Repositories: repositoriesMock}
}

func TestActionBillings(t *testing.T) {
	client.GithubMockTestHelper(t, Workflows(), buildWorkflows, client.TestOptions{})
}
