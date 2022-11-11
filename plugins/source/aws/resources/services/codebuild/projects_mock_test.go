package codebuild

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCodebuildProjects(t *testing.T, ctrl *gomock.Controller) client.Services {
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

	return client.Services{Codebuild: m}
}

func TestCodebuildProjects(t *testing.T) {
	client.AwsMockTestHelper(t, Projects(), buildCodebuildProjects, client.TestOptions{})
}
