package codebuild

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildSourceCredentials(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodebuildClient(ctrl)

	projectsList := codebuild.ListProjectsOutput{}
	if err := faker.FakeObject(&projectsList); err != nil {
		t.Fatal(err)
	}
	projectsList.NextToken = nil
	m.EXPECT().ListProjects(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&projectsList,
		nil,
	)

	projects := codebuild.BatchGetProjectsOutput{}
	if err := faker.FakeObject(&projects); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().BatchGetProjects(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&projects,
		nil,
	)

	buildID := ""
	if err := faker.FakeObject(&buildID); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBuildsForProject(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codebuild.ListBuildsForProjectOutput{
			Ids: []string{buildID},
		},
		nil,
	).MinTimes(1)

	build := codebuild.BatchGetBuildsOutput{}
	if err := faker.FakeObject(&build); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().BatchGetBuilds(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&build,
		nil,
	).MinTimes(1)

	return client.Services{Codebuild: m}
}

func TestSourceCredentials(t *testing.T) {
	client.AwsMockTestHelper(t, Projects(), buildSourceCredentials, client.TestOptions{})
}
