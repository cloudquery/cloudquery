package codebuild

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCodebuildProjects(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodebuildClient(ctrl)

	projectsList := codebuild.ListProjectsOutput{}
	require.NoError(t, faker.FakeObject(&projectsList))

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
	require.NoError(t, faker.FakeObject(&projects))

	m.EXPECT().BatchGetProjects(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&projects,
		nil,
	)

	buildID := ""
	require.NoError(t, faker.FakeObject(&buildID))

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
	require.NoError(t, faker.FakeObject(&build))

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

func TestCodebuildProjects(t *testing.T) {
	client.AwsMockTestHelper(t, Projects(), buildCodebuildProjects, client.TestOptions{})
}
