package codecommit

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodecommitClient(ctrl)

	repoMetadata := types.RepositoryMetadata{}
	require.NoError(t, faker.FakeObject(&repoMetadata))
	listOutput := &codecommit.ListRepositoriesOutput{Repositories: make([]types.RepositoryNameIdPair, 200)}
	batchGetInput := &codecommit.BatchGetRepositoriesInput{RepositoryNames: make([]string, 100)} // max batch get input len
	batchGetOutput := &codecommit.BatchGetRepositoriesOutput{Repositories: make([]types.RepositoryMetadata, 100)}
	for i := range listOutput.Repositories {
		listOutput.Repositories[i] = types.RepositoryNameIdPair{
			RepositoryName: repoMetadata.RepositoryName,
			RepositoryId:   repoMetadata.RepositoryId,
		}
	}
	for i := range batchGetInput.RepositoryNames {
		batchGetInput.RepositoryNames[i] = *repoMetadata.RepositoryName
		batchGetOutput.Repositories[i] = repoMetadata
	}

	m.EXPECT().ListRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(listOutput, nil)

	m.EXPECT().BatchGetRepositories(gomock.Any(), batchGetInput, gomock.Any()).
		Times(2).
		Return(batchGetOutput, nil)

	tags := map[string]string{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(200).
		Return(&codecommit.ListTagsForResourceOutput{Tags: tags}, nil)

	return client.Services{Codecommit: m}
}

func TestRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}
