package codeartifact

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodeartifactClient(ctrl)

	repoSummary := types.RepositorySummary{}
	if err := faker.FakeObject(&repoSummary); err != nil {
		t.Fatal(err)
	}
	repo := types.RepositoryDescription{}
	if err := faker.FakeObject(&repo); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRepositories(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListRepositoriesOutput{
			Repositories: []types.RepositorySummary{repoSummary},
			NextToken:    nil,
		},
		nil,
	)

	m.EXPECT().DescribeRepository(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.DescribeRepositoryOutput{
			Repository: &repo,
		},
		nil,
	)

	tag := types.Tag{}
	if err := faker.FakeObject(&tag); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		},
		nil,
	)

	return client.Services{Codeartifact: m}
}

func TestRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}
